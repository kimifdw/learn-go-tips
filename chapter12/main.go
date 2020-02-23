package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

//  CopyFile 文件复制
func CopyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "open source")
	}

	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return errors.Wrap(err, "create dest")
	}

	nn, err := io.Copy(w, r)
	if err != nil {
		w.Close()
		os.Remove(dst)
		return errors.Wrap(err, "copy body")
	}

	if err := w.Close(); err != nil {
		os.Remove(dst)
		return errors.Wrapf(err, "close dest,nn=%v", nn)
	}
	return nil
}

// LoadSystem 系统加载
func LoadSystem() error {
	src, dst := "src.text", "dst.txt"
	if err := CopyFile(src, dst); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("load src=%v, dst=%v", src, dst))
	}
	return nil
}

func main() {
	if err := LoadSystem(); err != nil {
		fmt.Printf("err %+v\n", err)
	}
}
