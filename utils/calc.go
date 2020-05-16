package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func SwitchTimeStampToData(t int64) string {
	timeTemplate1 := "2006-01-02 15:04:05"
	time_str := time.Unix(t, 0).Format(timeTemplate1)
	fmt.Println(time_str)
	return time_str
}
