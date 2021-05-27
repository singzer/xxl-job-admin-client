package xxl

import (
	"encoding/json"
	"errors"

	"github.com/asmcos/requests"
	"github.com/goinggo/mapstructure"
)

func (a Admin) GetChartInfo(startDate string, endDate string) (c ChartInfo, err error) {

	req := requests.Requests()

	for _, cooike := range a.Cookies {
		req.SetCookie(cooike)
	}

	data := requests.Datas{
		"startDate": startDate,
		"endDate":   endDate,
	}

	resp, err := req.Post(a.Url+"/chartInfo", data)
	if err != nil {
		return c, err
	}

	var rte Rte
	err = json.Unmarshal(resp.Content(), &rte)
	if err != nil {
		return c, err
	}

	if rte.Code != 200 {
		return c, errors.New(rte.Msg)
	}

	if err := mapstructure.Decode(rte.Content, &c); err != nil {
		return c, err
	}

	return c, err
}
