package xxl

import (
	"encoding/json"
	"errors"
	"fmt"
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

type JobInfoParam struct {
	JobGroup        uint
	TriggerStatus   int
	JobDesc         string
	ExecutorHandler string
	Author          string
	Start           int
	Length          int
}

type JobInfoRte struct {
	Data []struct {
		AddTime                string `json:"addTime"`
		AlarmEmail             string `json:"alarmEmail"`
		Author                 string `json:"author"`
		ChildJobID             string `json:"childJobId"`
		ExecutorBlockStrategy  string `json:"executorBlockStrategy"`
		ExecutorFailRetryCount int64  `json:"executorFailRetryCount"`
		ExecutorHandler        string `json:"executorHandler"`
		ExecutorParam          string `json:"executorParam"`
		ExecutorRouteStrategy  string `json:"executorRouteStrategy"`
		ExecutorTimeout        int64  `json:"executorTimeout"`
		GlueRemark             string `json:"glueRemark"`
		GlueSource             string `json:"glueSource"`
		GlueType               string `json:"glueType"`
		GlueUpdatetime         string `json:"glueUpdatetime"`
		ID                     int64  `json:"id"`
		JobCron                string `json:"jobCron"`
		JobDesc                string `json:"jobDesc"`
		JobGroup               int64  `json:"jobGroup"`
		TriggerLastTime        int64  `json:"triggerLastTime"`
		TriggerNextTime        int64  `json:"triggerNextTime"`
		TriggerStatus          int64  `json:"triggerStatus"`
		UpdateTime             string `json:"updateTime"`
	} `json:"data"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	RecordsTotal    int64 `json:"recordsTotal"`
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

func (a Admin) GetJobInfo(p JobInfoParam) (j JobInfoRte, err error) {

	req := requests.Requests()

	for _, cooike := range a.Cookies {
		req.SetCookie(cooike)
	}

	/*
		jobGroup: 2
		triggerStatus: -1
		jobDesc:
		executorHandler:
		author:
		start: 0
		length: 10
	*/

	data := requests.Datas{
		"jobGroup":      fmt.Sprint(p.JobGroup),
		"triggerStatus": fmt.Sprint(p.TriggerStatus),
		"jobDesc":       p.JobDesc,
		"author":        p.Author,
		"start":         fmt.Sprint(p.Start),
		"length":        fmt.Sprint(p.Length),
	}

	resp, err := req.Post(a.Url+"/jobinfo/pageList", data)
	if err != nil {
		return j, err
	}

	err = json.Unmarshal(resp.Content(), &j)
	if err != nil {
		return j, err
	}

	return j, err
}
