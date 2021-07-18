package main

import (
	"SKT_PhoneNumber/config"
	"SKT_PhoneNumber/lib/log"
	"SKT_PhoneNumber/lib/req"
	"net/http"
)

func main() {
	client := &http.Client{
	}
	req := req.CreateRequest(config.Config.Skt.URL)

	resp, err := client.Do(req)
	if err != nil {
		log.Err.Fatalln("사용 브라우저의 T World 로그인 필요")
	}
	log.Info.Println(resp.Request.RequestURI)
	defer resp.Body.Close()
}