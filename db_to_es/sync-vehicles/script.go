package main

import (
	"time"

	"github.com/gnufree/db_to_es/common"
	"github.com/gnufree/db_to_es/constant"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

const Token = "799eff86c0e741deb0e4eeeecf7c27b3"

// 执行前需设置 Token
func main() {
	// 初始化设置环境
	constant.Env = constant.Prod

	session, err := common.GetDBSession(constant.ServiceVehicle)
	if err != nil {
		log.Panic(err)
	}
	defer session.Close()

	collection := session.DB("vehicle").C("vehicle_bucket")

	var vehicles []common.Vehicle
	start := time.Date(2018, time.February, 1, 0, 0, 0, 0, time.UTC)
	// 条件过滤车源
	collection.Find(bson.M{"status.saleStatus": "on_sale", "createdAt": bson.M{"$gt": start}}).All(&vehicles)

	log.Infof("total vehicles is %d", len(vehicles))

	var vehicleIds []string
	for _, v := range vehicles {
		vehicleIds = append(vehicleIds, v.Id.Hex())
	}

	qns := common.QNS{Queries: 10, Seconds: 5}
	// 同步车源
	failedUrls := make(map[string]string)
	common.Sync(vehicleIds, Token, failedUrls, qns)

	log.Info("============================================")
	log.Infof("total failed requests is %d, retry: ", len(failedUrls))

	// 再次同步此次同步过程中失败的车源
	leftUrls := make(map[string]string)
	var cnt int
	for _, url := range failedUrls {
		common.Get(url, Token, leftUrls)
		cnt++
		if cnt%qns.Queries == 0 {
			time.Sleep(time.Second * time.Duration(qns.Seconds))
		}
	}
	log.Info("retried")
	log.Info("============================================")

	for _, url := range leftUrls {
		log.Info(url)
	}

	log.Info("Done")
}
