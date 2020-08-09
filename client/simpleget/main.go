package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err) // エラーを表示してプログラムを終了させる
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
	log.Println("StatusCode:", resp.StatusCode)
	log.Println("Headers:", resp.Header)
	log.Println("Content-Length:", resp.Header.Get("Content-Length"))
	log.Println(string(body))
	log.Println(body) // []byte
}
