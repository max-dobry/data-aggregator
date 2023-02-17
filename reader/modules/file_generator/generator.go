package file_generator

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"reader/modules/domain"
	"reader/modules/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateNewFile(length int) error {
	f, err := os.Create("data.json")
	utils.Check(err)
	defer f.Close()

	byteFile := prepareFileContent(length)

	n, err := f.Write(byteFile)
	utils.Check(err)

	fmt.Println("Bytes written: ", n)
	return nil
}

func prepareFileContent(count int) []byte {
	defer utils.TimeTrack(time.Now(), "prepareFileContent")

	bt := []byte("[")
	for i := 1; i <= count; i++ {
		m, err := json.Marshal(newJson())
		utils.Check(err)
		bt = append(bt, m...)
		if i == count {
			continue
		}
		bt = append(bt, ","...)
	}
	bt = append(bt, "]"...)

	return bt
}

func newJson() (d domain.FooData) {
	d.ID = uuid.NewString()
	d.Time = randTime()
	d.Signal = randSignal()
	d.Data = randString(14)
	return
}

func randTime() (t int) {
	minTime := int(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
	maxTime := int(time.Date(2020, 12, 0, 0, 0, 0, 0, time.UTC).Unix())

	rand.Seed(time.Now().UnixNano())
	t = minTime + (rand.Intn(maxTime - minTime))
	return
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func randSignal() string {
	s := "signal_"
	rs := rand.Intn(10)
	s += strconv.Itoa(rs)
	return s
}
