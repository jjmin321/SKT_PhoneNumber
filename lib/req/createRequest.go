package req

import (
	"SKT_PhoneNumber/config"
	"SKT_PhoneNumber/lib/log"
	"net/http"

	"github.com/skratchdot/open-golang/open"
)

func CreateRequest(url string) *http.Request{
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Err.Fatalln("Request 생성 중 오류가 발생함.")
	}
	query := req.URL.Query()
	query.Add("viewId", config.Config.Skt.ViewId)
	query.Add("serviceId", config.Config.Skt.ServiceId)
	query.Add("chg_cd", config.Config.Skt.Chg_cd)
	query.Add("co_cd", config.Config.Skt.Co_cd)
	query.Add("line_num_from", config.Config.Skt.Line_num_from)
	query.Add("line_num_to", config.Config.Skt.Line_num_to)
	req.URL.RawQuery = query.Encode()
	open.Start(req.URL.String()) 
	return req
}