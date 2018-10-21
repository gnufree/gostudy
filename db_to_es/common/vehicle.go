package common

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

const SyncVehicleUrl =  "https://p-api.kanche.com/v1/search/sync_index_vehicle/"

type Vehicle struct {
	Id bson.ObjectId `bson:"_id"`
	RelatedInfo VehicleRelatedInfo `bson:"relatedInfo"`
}

type VehicleRelatedInfo struct {
	TcSerialNumber string `bson:"tcSerialNumber"`
}

func Sync(vehicleIds []string, token string, failedUrls map[string]string, qns QNS) {
	var cnt int
	for _, vehicleId := range vehicleIds {
		url := SyncVehicleUrl + vehicleId + "?clusterName=kkche-pro&indexName=v&indexType=base&modelType=vehicle"
		Get(url, token, failedUrls)

		cnt ++
		if cnt % qns.Queries == 0 {
			time.Sleep(time.Second * time.Duration(qns.Seconds))
		}

	}
}

