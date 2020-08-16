package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	var url string //下载文件的地址
	url = "https://download.jetbrains.com/go/goland-2020.2.2.dmg"

	downloader := NewFileDownloader(url, "", "", 10)
	if err := downloader.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n 文件下载完成耗时：%f second\n", time.Now().Sub(startTime).Seconds())
}

// 文件下载器
type FileDownloader struct {
	fileSize       int
	url            string
	outputFileName string
	totalPart      int
	outputDir      string
	doneFilePart   []FilePart
}

// 分片文件
type FilePart struct {
	// 分片序号
	Index int
	// 起始byte
	From int
	// 结束byte
	To int
	// 文件内容
	Data []byte
}

// 初始化文件下载器
func NewFileDownloader(url, outputFileName, outputDir string, totalPart int) *FileDownloader {
	if outputDir == "" {
		// 当前工作路径
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		outputDir = wd
	}
	return &FileDownloader{
		fileSize:       0,
		url:            url,
		outputFileName: outputFileName,
		outputDir:      outputDir,
		totalPart:      totalPart,
		doneFilePart:   make([]FilePart, totalPart),
	}
}

func (d *FileDownloader) head() (int, error) {
	r, err := d.getNewRequest("HEAD")
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode > 299 {
		return 0, errors.New(fmt.Sprintf("Can't process, response is %v", resp.StatusCode))
	}
	// 检查是否支持断点续传
	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("服务器不支持文件断点续传")
	}

	d.outputFileName = parseFileInfoFrom(resp)
	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func parseFileInfoFrom(resp *http.Response) string {
	contentDisposition := resp.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		_, params, err := mime.ParseMediaType(contentDisposition)

		if err != nil {
			panic(err)
		}
		return params["filename"]
	}
	filename := filepath.Base(resp.Request.URL.Path)
	return filename
}

func (d *FileDownloader) Run() error {
	fileTotalSize, err := d.head()
	if err != nil {
		return err
	}
	d.fileSize = fileTotalSize

	jobs := make([]FilePart, d.totalPart)
	eachSize := fileTotalSize / d.totalPart

	for i := range jobs {
		jobs[i].Index = i
		if i == 0 {
			jobs[i].From = 0
		} else {
			jobs[i].From = jobs[i-1].To + 1
		}
		if i < d.totalPart-1 {
			jobs[i].To = jobs[i].From + eachSize
		} else {
			jobs[i].To = fileTotalSize - 1
		}
	}

	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		go func(job FilePart) {
			defer wg.Done()
			err := d.downloadPart(job)
			if err != nil {
				log.Println("下载文件失败：", err, job)
			}
		}(job)
	}
	wg.Wait()
	return d.mergeFileParts()
}

// 创建一个request
func (d *FileDownloader) getNewRequest(method string) (*http.Request, error) {
	r, err := http.NewRequest(
		method,
		d.url,
		nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "emily")
	return r, nil
}

// 下载分片
func (d *FileDownloader) downloadPart(job FilePart) error {
	r, err := d.getNewRequest("GET")
	if err != nil {
		return err
	}

	log.Printf("开始[%d]下载from:%d to %d\n", job.Index, job.From, job.To)
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", job.From, job.To))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("服务器状态码错误：%v", resp.StatusCode))
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(bs) != (job.To - job.From + 1) {
		return errors.New("下载文件分片长度错误")
	}
	job.Data = bs
	d.doneFilePart[job.Index] = job
	return nil
}

// 合并分片文件
func (d *FileDownloader) mergeFileParts() error {

	path := filepath.Join(d.outputDir, d.outputFileName)
	log.Printf("开始合并文件,下载地址：%s\n", path)
	mergedFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer mergedFile.Close()
	hash := sha256.New()
	totalSize := 0
	for _, filePart := range d.doneFilePart {
		mergedFile.Write(filePart.Data)
		hash.Write(filePart.Data)
		totalSize += len(filePart.Data)
	}
	if totalSize != d.fileSize {
		return errors.New("文件不完整")
	}

	if hex.EncodeToString(hash.Sum(nil)) != "3af4660ef22f805008e6773ac25f9edbc17c2014af18019b7374afbed63d4744" {
		return errors.New("文件损坏")
	} else {
		log.Println("文件SHA-256校验成功")
	}
	return nil
}
