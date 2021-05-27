package xxl

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/asmcos/requests"
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

type AddJobParam struct {
	JobGroup               uint
	JobDesc                string
	ExecutorRouteStrategy  string
	CronGenDisplay         string
	JobCron                string
	GlueType               string
	ExecutorHandler        string
	ExecutorBlockStrategy  string
	ChildJobId             []uint
	ExecutorTimeout        int
	ExecutorFailRetryCount int
	Author                 string
	AlarmEmail             string
	ExecutorParam          string
	GlueRemark             string
	GlueSource             string
}

type GetJobInfoParam struct {
	JobGroup        uint
	TriggerStatus   int
	JobDesc         string
	ExecutorHandler string
	Author          string
	Start           int
	Length          int
}

type JobInfo struct {
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
}

type JobInfoRte struct {
	Data            []JobInfo `json:"data"`
	RecordsFiltered int64     `json:"recordsFiltered"`
	RecordsTotal    int64     `json:"recordsTotal"`
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
