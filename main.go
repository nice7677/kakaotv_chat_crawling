package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net"
)

func main() {

	client := resty.New()
	resp, err := client.R().
		//SetHeader("Content-Type", "application/json").
		SetHeaders(map[string]string{
			"Content-Type":   "application/x-www-form-urlencoded",
			"Origin":         "https://live-tv.kakao.com",
			"Referer":        "https://live-tv.kakao.com/kakaotv/live/chat/user/6219964",
			"Sec-Fetch-Mode": "cors",
			"User-Agent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36",
		}).
		SetFormData(map[string]string{
			"groupid": "3217974",
		}).
		Post("https://play.kakao.com/chat/service/api/room")
	if err != nil {
		log.Println(err)
	}

	byt := []byte(resp.Body())
	var jsonMap map[string]interface{}

	if err := json.Unmarshal(byt, &jsonMap); err != nil {
		panic(err)
	}
	enter := jsonMap["enter"]
	s := fmt.Sprintf("ENTER %s\n", enter)

	conn, err := net.Dial("tcp", "203.133.165.105:9003")
	if nil != err {
		log.Println(err)
	}

	conn.Write([]byte(s))

	data := make([]byte, 4096)

	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n]))
	}
}
