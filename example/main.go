package main

import (
	"log"

	"github.com/singzer/xxl-job-admin-client"
)

func main() {
	xj, err := xxl.New("http://127.0.0.1:8080/xxl-job-admin/", "admin", "123456")
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(xj.Cookies)

	chartInfo, err := xj.GetChartInfo("2021-05-20 00:00:00", "2021-05-27 23:59:59")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(chartInfo)

	job := xxl.AddJobParam{
		JobGroup:              2,
		JobDesc:               "TEST",
		ExecutorRouteStrategy: "FIRST",
		CronGenDisplay:        "* * * * * ?",
		JobCron:               "* * * * * ?",
		GlueType:              "BEAN",
		ExecutorHandler:       "test.task",
		ExecutorBlockStrategy: "SERIAL_EXECUTION",
		//ChildJobId:             []unit{1,2},
		ExecutorTimeout:        0,
		ExecutorFailRetryCount: 0,
		Author:                 "icepie",
		// AlarmEmail:             "",
		// ExecutorParam:          "",
		// GlueRemark:             "",
		// GlueSource:             "",
	}

	jobID, err := xj.AddJob(job)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("add job:", jobID)

	jobInfo, err := xj.GetJobInfo(xxl.GetJobInfoParam{JobGroup: 2, Start: 0, Length: 20})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(jobInfo)

	err = xj.RemoveJob(jobID)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("remove job:", jobID)

	jobInfoList, err := xj.GetJobsByGroup(2)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(jobInfoList)

	err = xj.TriggerJob(2, "test2", "")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("tigger job...")

}
