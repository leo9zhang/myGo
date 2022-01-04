package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//func main() {
//	//global.Viper = core.Viper()   // 初始化Viper
//	//global.Logger = core.Zap()    // 初始化zap日志库
//	//global.Db = initialize.Gorm() // gorm连接数据库
//	//
//	//global.Redis = initialize.Redis()
//	//
//	//if global.Db != nil {
//	//	initialize.RegisterTables(global.Db) // 初始化表
//	//	// 程序结束前关闭数据库链接
//	//	db, _ := global.Db.DB()
//	//	defer db.Close()
//	//}
//	//core.RunServer()
//
//
//}
//

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{
	"version": 1,
	"componentName": "qcbuy",
	"eventId": 143371,
	"timestamp": 1551854215344,
	"user": "auto",
	"interface": {
		"interfaceName": "qcloud.Deal.getDealsByCreatTime",
		"para": {
			"cond": {
				"ownerUin": "2407912486",
				"order": 1,
				"page": 0,
				"rows": 1,
				"payMode": 1,
				"statusList": [3, 4],
				"invalidActivityIdList": [1],
				"payEndTimeStart": "2018-05-04 00:00:00",
				"taskStartTimeStart": "2018-05-10 17:52:38",
				"productCodeList": ["p_cds_audit", "p_cvm"],
				"actionList": ["t"],
				"invalidStatusList": [5],
				"creatTimeStart": "2017-03-04 00:00:00"
			}
		}
	},
	"seqId": "647ea242-f654-965d-a62a-eabe0289d954",
	"spanId": "https://buy.qcloud.com;61911"
}`)
	req, err := http.NewRequest("POST", "http://9.139.137.40/interfaces/interface.php", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "trade.itd.com")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
