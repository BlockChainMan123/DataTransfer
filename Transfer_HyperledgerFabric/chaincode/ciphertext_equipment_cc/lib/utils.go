package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//ruturn a randomed six-digit captcha
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

func GetMD5String(b []byte) (result string) {
	res := md5.Sum(b)
	result = hex.EncodeToString(res[:])
	return
}

func BJtime(timeString string) (string) {
	const longForm = "2006-01-02 15:04:05"
	t, _ := time.Parse(longForm, timeString)
	realtime := t.Add(8 * time.Hour)
	return realtime.Format("2006-01-02 15:04:05")
}

func BJtimeShort(timeString string) (string) {
	const longForm = "2006-01-02"
	t, _ := time.Parse(longForm, timeString)
	realtime := t.Add(8 * time.Hour)
	return realtime.Format("2006-01-02")
}
