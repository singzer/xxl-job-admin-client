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

	jobInfo, err := xj.GetJobInfo(xxl.JobInfoParam{JobGroup: 2, Start: 0, Length: 20})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(jobInfo)

}
