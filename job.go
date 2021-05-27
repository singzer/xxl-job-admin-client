package xxl

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/asmcos/requests"
	"github.com/goinggo/mapstructure"
)

func (a Admin) GetJobInfo(p GetJobInfoParam) (j JobInfoRte, err error) {

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

func (a Admin) GetJobsByGroup(GroupID uint) (j []JobInfo, err error) {

	req := requests.Requests()

	for _, cooike := range a.Cookies {
		req.SetCookie(cooike)
	}

	data := requests.Datas{
		"jobGroup": fmt.Sprint(GroupID),
	}

	resp, err := req.Post(a.Url+"/joblog/getJobsByGroup", data)
	if err != nil {
		return j, err
	}

	var rte Rte
	err = json.Unmarshal(resp.Content(), &rte)
	if err != nil {
		return j, err
	}

	if rte.Code != 200 {
		return j, errors.New(rte.Msg)
	}

	if err := mapstructure.Decode(rte.Content, &j); err != nil {
		return j, err
	}

	return j, nil

}

func (a Admin) AddJob(p AddJobParam) (id uint, err error) {

	req := requests.Requests()

	for _, cooike := range a.Cookies {
		req.SetCookie(cooike)
	}

	data := requests.Datas{
		"jobGroup":               fmt.Sprint(p.JobGroup),
		"jobDesc":                p.JobDesc,
		"executorRouteStrategy":  p.ExecutorRouteStrategy,
		"cronGen_display":        p.CronGenDisplay,
		"jobCron":                p.JobCron,
		"glueType":               p.GlueType,
		"executorHandler":        p.ExecutorHandler,
		"executorBlockStrategy":  p.ExecutorBlockStrategy,
		"childJobId":             idSilceToString(p.ChildJobId),
		"executorTimeout":        fmt.Sprint(p.ExecutorTimeout),
		"executorFailRetryCount": fmt.Sprint(p.ExecutorFailRetryCount),
		"author":                 p.Author,
		"alarmEmail":             p.AlarmEmail,
		"executorParam":          p.ExecutorParam,
		"glueRemark":             p.GlueRemark,
		"glueSource":             p.GlueSource,
	}

	resp, err := req.Post(a.Url+"/jobinfo/add", data)
	if err != nil {
		return id, err
	}

	var rte Rte
	err = json.Unmarshal(resp.Content(), &rte)
	if err != nil {
		return id, err
	}

	if rte.Code != 200 {
		return id, errors.New(rte.Msg)
	}

	theid, _ := strconv.Atoi(rte.Content.(string))

	id = uint(theid)

	return id, nil
}

func (a Admin) TriggerJob(jobID uint, executorParam string, addressList string) (err error) {
	req := requests.Requests()

	for _, cooike := range a.Cookies {
		req.SetCookie(cooike)
	}

	data := requests.Datas{
		"id":            fmt.Sprint(jobID),
		"executorParam": executorParam,
		"addressList":   addressList,
	}

	resp, err := req.Post(a.Url+"/jobinfo/trigger", data)
	if err != nil {
		return err
	}

	var rte Rte
	err = json.Unmarshal(resp.Content(), &rte)
	if err != nil {
		return err
	}

	if rte.Code != 200 {
		return errors.New(rte.Msg)
	}

	return nil
}

func (a Admin) RemoveJob(jobID uint) (err error) {
	req := requests.Requests()

	for _, cooike := range a.Cookies {
		req.SetCookie(cooike)
	}

	data := requests.Datas{
		"id": fmt.Sprint(jobID),
	}

	resp, err := req.Post(a.Url+"/jobinfo/remove", data)
	if err != nil {
		return err
	}

	var rte Rte
	err = json.Unmarshal(resp.Content(), &rte)
	if err != nil {
		return err
	}

	if rte.Code != 200 {
		return errors.New(rte.Msg)
	}

	return nil
}
