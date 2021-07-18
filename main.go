package main

import (
	"SKT_PhoneNumber/config"
	"SKT_PhoneNumber/lib/log"
	"SKT_PhoneNumber/lib/req"
	"net/http"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.Strong {
		return scrape.Attr(n, "class") == "cpTit"
	}
	return false
}

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

	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Err.Fatalln("HTML Parse 과정 중 오류")
	}

	phoneNumberList := scrape.FindAll(root, parseMainNodes)
	log.Info.Println(phoneNumberList)

	for _, phoneNumber := range phoneNumberList {
		log.Info.Println(scrape.Text(phoneNumber))
	}
}