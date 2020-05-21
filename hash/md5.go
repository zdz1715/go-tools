package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

func FileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(fmt.Sprintf("open error: %v", err))
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("copy error: %v", err))
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func StringMd5(str string) string {
	md5New := md5.New()
	md5New.Write([]byte(str))
	return hex.EncodeToString(md5New.Sum(nil))
}
