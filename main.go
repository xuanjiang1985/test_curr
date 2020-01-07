package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	host       = "http://caifu-local.colourlife.com/api/activity/test/insert"
	currNumber = 9
	httpChan   = make(chan string, currNumber)
	ctx        = context.Background()
	queryData  = map[string]string{
		"a":  "1",
		"ab": "122",
		"c":  "1234",
	}
)

func main() {
	httpGet(ctx, httpChan)

	//httpPost(ctx, httpChan)
	count := 0
	for {
		if data, ok := <-httpChan; ok {
			count++
			fmt.Println(data)
		}

		if count == currNumber {
			break
		}
	}
}

func httpGet(ctx context.Context, done chan string) {

	for i := 0; i < currNumber; i++ {
		go httpGetChild(ctx, i, done)
	}
}

func httpGetChild(ctx context.Context, i int, done chan string) {

	start := "start" + strconv.Itoa(i) + "   " + time.Now().Format("01-02 15:04:05.000")

	// init get query
	params := url.Values{}
	for k, v := range queryData {
		params.Set(k, v)
	}

	httpUrl, _ := url.Parse(host)
	httpUrl.RawQuery = params.Encode()
	httpPath := httpUrl.String()

	res, err := http.Get(httpPath)
	if err != nil {
		done <- start + err.Error()
		return
	}
	defer res.Body.Close()

	start += "   " + time.Now().Format("01-02 15:04:05.000")
	done <- start

}

func httpPost(ctx context.Context, done chan string) {

	for i := 0; i < currNumber; i++ {
		go httpPostChild(ctx, i, done)
	}
}

func httpPostChild(ctx context.Context, i int, done chan string) {

	start := "start" + strconv.Itoa(i) + "   " + time.Now().Format("01-02 15:04:05.000")

	bytesData, _ := json.Marshal(queryData)
	res, err := http.Post(host, "application/json", bytes.NewReader(bytesData))
	if err != nil {
		done <- start + err.Error()
		return
	}
	defer res.Body.Close()

	start += "   " + time.Now().Format("01-02 15:04:05.000")
	done <- start
}
