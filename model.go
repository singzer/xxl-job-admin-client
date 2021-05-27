package xxl

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/asmcos/requests"
	"github.com/goinggo/mapstructure"
)

type Admin struct {
	Url      string
	UserName string
	Password string
	Cookies  []*http.Cookie
}

type Rte struct {
	Code    int64       `json:"code"`
	Content interface{} `json:"content"`
	Msg     string      `json:"msg"`
}

type ChartInfo struct {
	TriggerCountFailTotal      int64    `json:"triggerCountFailTotal"`
	TriggerCountRunningTotal   int64    `json:"triggerCountRunningTotal"`
	TriggerCountSucTotal       int64    `json:"triggerCountSucTotal"`
	TriggerDayCountFailList    []int64  `json:"triggerDayCountFailList"`
	TriggerDayCountRunningList []int64  `json:"triggerDayCountRunningList"`
	TriggerDayCountSucList     []int64  `json:"triggerDayCountSucList"`
	TriggerDayList             []string `json:"triggerDayList"`
}

func New(Url string, UserName string, Password string) (a Admin, err error) {

	req := requests.Requests()

	data := requests.Datas{
		"userName": UserName,
		"password": Password,
	}

	resp, err := req.Post(Url+"/login", data)
	if err != nil {
		return a, err
	}

	a.Cookies = resp.Cookies()
	a.UserName = UserName
	a.Password = Password
	a.Url = Url

	var rte Rte
	err = json.Unmarshal(resp.Content(), &rte)
	if err != nil {
		return a, err
	}

	if rte.Code != 200 {
		return a, errors.New(rte.Msg)
	}

	return a, nil
}

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

	log.Println(resp.Text())

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
