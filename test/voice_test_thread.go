package main

import (
    "bytes"
    "context"
    "crypto/md5"
    "encoding/binary"
    "flag"
    "fmt"
    "github.com/gorilla/websocket"
    "io/ioutil"
    "log"
    "net/url"
    "os"
    "os/signal"
    "sync"
    "time"
)

// 结束标记
const AUX_BREAK_FLAG = "--end--"

// 每帧数据大小，单位Byte
const AUX_SLICE_SIZE = 2048

// 每帧数据发送间隔，单位ms
const AUX_INTERVAL = 40

// 音频文件地址
const AUX_FILE_PATH = "pcm/next.pcm"

var HEAT_BEAT_BYTES = []byte{0xa5, 0xa5, 0x06, 0x00, 0x88, 0x00}
var START_BYTES = []byte{0xa5, 0xa5, 0x50, 0x00, 0x02, 0x05, 0x74, 0x6f, 0x6b, 0x65,
    0x6e, 0x01, 0x01, 0x00, 0x01, 0x01, 0x00, 0x16, 0x76, 0x52, 0x34, 0x4a, 0x4c, 0x36,
    0x38, 0x62, 0x4b, 0x35, 0x51, 0x33, 0x41, 0x42, 0x54, 0x4b, 0x59, 0x42, 0x53, 0x4a,
    0x37, 0x48, 0x0a, 0x34, 0x32, 0x33, 0x36, 0x37, 0x39, 0x30, 0x31, 0x32, 0x33, 0x04,
    0x01, 0x48, 0x01, 0x02, 0x16, 0x00, 0x7b, 0x22, 0x6d, 0x61, 0x63, 0x22, 0x3a, 0x22,
    0x41, 0x43, 0x46, 0x32, 0x44, 0x43, 0x33, 0x42, 0x33, 0x35, 0x43, 0x38, 0x22, 0x7d}

var (
    // 接口地址
    AUX_WS_URL string
    // 线程数
    AUX_THREAD_NUM int
    // 每个线程循环次数
    AUX_REQ_NUM    int
    AUX_ORIGIN     string
    OPEN_HEARTBEAT bool
)

// 读取启动参数
func init() {
    flag.StringVar(&AUX_WS_URL, "AUX_WS_URL", "ws://smthomeav.aux-home.com/stressTest", "websocket url")
    flag.StringVar(&AUX_ORIGIN, "AUX_ORIGIN", "http://smthomeav.aux-home.com", "origin url")
    flag.IntVar(&AUX_THREAD_NUM, "AUX_THREAD_NUM", 1, "thread num")
    flag.IntVar(&AUX_REQ_NUM, "AUX_REQ_NUM", 1, "req num")
    flag.BoolVar(&OPEN_HEARTBEAT, "OPEN_HEARTBEAT", false, "heartbeat open")

}

// 主方法
func main() {
    // 暂停，获取参数，并打印
    flag.Parse()
    fmt.Printf("ws_url: %v\r\n", AUX_WS_URL)
    fmt.Printf("thread_num: %v\r\n", AUX_THREAD_NUM)
    fmt.Printf("req_num: %v\r\n", AUX_REQ_NUM)
    fmt.Printf("origin: %v\r\n", AUX_ORIGIN)
    fmt.Printf("heartbeat:%v\r\n", OPEN_HEARTBEAT)
    var addr = flag.String("addr", "192.168.1.105:18851", "http service address")
    u := url.URL{Scheme: "ws", Host: *addr, Path: "/voiceProxy"}

    // 读取待发送数据
    voiceData, _ := ioutil.ReadFile(AUX_FILE_PATH)

    // 起线程
    var wg sync.WaitGroup
    for i := 0; i < AUX_THREAD_NUM; i++ {
        wg.Add(1)
        go voiceClient(u, i, wg.Done, voiceData, OPEN_HEARTBEAT)

    }

    // 等待所有线程结束
    wg.Wait()

    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)
    <-interrupt

    os.Exit(0)

}

func voiceClient(u url.URL, threadIndex int, done func(), voiceData []byte, heartbeat bool) {
    defer done()
    log.Printf("thread-%v started!\r\n", threadIndex)
    conn, err := getConnection(u)
    if err != nil {
        log.Printf("t-%v, dial err: %v\r\n", threadIndex, err)
        return
    }
    defer func() {
        err2 := conn.Close()
        if err2 != nil {
            log.Printf("t-%v, close err: %v\r\n", threadIndex, err)
            return
        }
    }()
    addr := conn.RemoteAddr()
    log.Printf("t-%v,remote address:%s", threadIndex, addr)

    for i := 0; i < AUX_REQ_NUM; i++ {
        voiceRequest(conn, START_BYTES, threadIndex, i, voiceData, heartbeat)
    }

    return
}

func voiceRequest(conn *websocket.Conn, data []byte, threadIndex int, reqNum int, voiceData []byte, heartbeat bool) {
    sendChan := make(chan int, 1)
    receiveChan := make(chan int, 1)
    defer close(sendChan)
    defer close(receiveChan)
    // 发送数据
    go sendStartData(conn, data, threadIndex, reqNum, sendChan)
    // 接收数据
    go receiveData(conn, threadIndex, reqNum, receiveChan, voiceData, heartbeat)

    <-sendChan
    <-receiveChan
    return
}

func sendStartData(conn *websocket.Conn, data []byte, threadIndex int, reqNum int, sendChan chan int) {
    log.Printf("t-%v-r-%v, begin send start data\n", threadIndex, reqNum)
    // 上传start指令
    if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
        log.Printf("t-%v-r-%v, send start err: %v\r\n", threadIndex, reqNum, err)
        sendChan <- 1
        return
    }
    log.Printf("t-%v-r-%v, send start success\n", threadIndex, reqNum)
    sendChan <- 1
    return
}

func receiveData(conn *websocket.Conn, threadIndex int, reqNum int, receiveChan chan int, voiceData []byte, heartbeat bool) {
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            if err.Error() == "EOF" {
                log.Printf("t-%v-r-%v, receive msg end\r\n", threadIndex, reqNum)
            } else {
                log.Printf("t-%v-r-%v, receive msg error: %v\r\n", threadIndex, reqNum, err)

            }
            receiveChan <- 1
            log.Println("receive finish err")
            return
        }
        cmdTypeBytes := msg[4:5]
        cmdType := cmdTypeBytes[0] & 0xFF
        if cmdType == 130 {
            if heartbeat {
                if sendErr := conn.WriteMessage(websocket.BinaryMessage, HEAT_BEAT_BYTES); sendErr != nil {
                    log.Printf("t-%v-r-%v, send heartbeat msg err: %v\r\n", threadIndex, reqNum, sendErr)
                    return
                }
            } else {
                sessionId := msg[7:]
                log.Printf("t-%v-r-%v start sessionId:%s", threadIndex, reqNum, sessionId)
                voiceDataLen := len(voiceData)
                sliceNum := getSliceNumForAux(voiceDataLen, AUX_SLICE_SIZE)
                for i := 0; i < sliceNum; i++ {
                    time.Sleep(AUX_INTERVAL * time.Millisecond)
                    var sliceData []byte
                    var isFinish bool
                    if (i+1)*AUX_SLICE_SIZE < voiceDataLen {
                        sliceData = voiceData[i*AUX_SLICE_SIZE : (i+1)*AUX_SLICE_SIZE]
                        isFinish = false
                    } else {
                        sliceData = voiceData[i*AUX_SLICE_SIZE:]
                        isFinish = true
                    }

                    reportData := getReportData(&sessionId, &sliceData, isFinish)

                    if sendErr := conn.WriteMessage(websocket.BinaryMessage, reportData); sendErr != nil {
                        log.Printf("t-%v-r-%v, send msg err: %v\r\n", threadIndex, reqNum, sendErr)
                        return
                    }
                }
            }

        }
        if cmdType == 5 {
            sessionLenBytes := msg[15:16]
            sessionLen := sessionLenBytes[0] & 0xFF
            middleResultBytes := msg[18+sessionLen:]
            log.Printf("t-%v-r-%v, receive middle msg: %s\r\n", threadIndex, reqNum, middleResultBytes)
        }
        if cmdType == 9 {
            sessionLenBytes := msg[8:9]
            sessionLen := sessionLenBytes[0] & 0xFF
            otherDataBytes := msg[10+sessionLen:]
            log.Printf("t-%v-r-%v, receive other msg: %s\r\n", threadIndex, reqNum, otherDataBytes)
        }
        if cmdType == 8 {
            log.Printf("t-%v-r-%v, receive heartbeat msg: %x\r\n", threadIndex, reqNum, msg)
            time.Sleep(5 * time.Second)
            if sendErr := conn.WriteMessage(websocket.BinaryMessage, HEAT_BEAT_BYTES); sendErr != nil {
                log.Printf("t-%v-r-%v, send heartbeat msg err: %v\r\n", threadIndex, reqNum, sendErr)
                return
            }
        }

    }
    receiveChan <- 1
}

func getReportData(sessionId *[]byte, voiceData *[]byte, isFinish bool) []byte {
    var reportData []byte
    reportData = append(reportData, 0xa5, 0xa5)
    sessionIdLength := len(*sessionId)
    voiceDataLen := len(*voiceData)
    reportLen := 5 + 4 + sessionIdLength + voiceDataLen
    lenBytes := IntToBytes(reportLen)[:2]

    reportData = append(reportData, lenBytes...)
    reportData = append(reportData, 0x03)

    sessionLenBytes := IntToBytes(sessionIdLength)[:1]
    reportData = append(reportData, sessionLenBytes...)
    reportData = append(reportData, *sessionId...)
    if isFinish {
        reportData = append(reportData, 0x00)
    } else {
        reportData = append(reportData, 0x01)
    }
    voiceDataLenBytes := IntToBytes(voiceDataLen)[:2]
    reportData = append(reportData, voiceDataLenBytes...)
    reportData = append(reportData, *voiceData...)

    return reportData
}

func getConnection(u url.URL) (*websocket.Conn, error) {
    header := make(map[string][]string)
    header["X-Real-Ip"] = []string{"192.168.1.105"}
    dialer := websocket.DefaultDialer
    dialer.HandshakeTimeout = 5 * time.Second
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    conn, _, err := dialer.DialContext(ctx, u.String(), header)
    return conn, err
}

// 计算分片数目
func getSliceNumForAux(dataSize, sliceSize int) int {
    if dataSize%sliceSize == 0 {
        return dataSize / sliceSize
    } else {
        return dataSize/sliceSize + 1
    }
}

// 整数转bytes[] 小端
func IntToBytes(n int) []byte {
    x := int32(n)

    bytesBuffer := bytes.NewBuffer([]byte{})
    binary.Write(bytesBuffer, binary.LittleEndian, x)
    return bytesBuffer.Bytes()
}

// 计算字符串MD5值
func Md5EncodeForAux(str string) (strMd5 string) {
    strByte := []byte(str)
    strMd5Byte := md5.Sum(strByte)
    strMd5 = fmt.Sprintf("%x", strMd5Byte)
    return strMd5
}
