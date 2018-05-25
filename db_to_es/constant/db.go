package constant

import "errors"

type DBConfig struct {
	DBHosts    []string
	Source     string
	DBName     string
	DBUsername string
	DBPassword string
	Tag        string
}

const (
	ServiceVehicle = "service_vehicle"
	// ServiceMerchant = "service_merchant"
	// ServiceBilling  = "service_billing"
	// ServiceFinance  = "service_finance"
)

const (
	Prod = "PROD"
)

// 车源服务生产环境数据库配置
var DBServiceVehicleProdConfig = DBConfig{
	//DBHosts: []string{"dds-m5e04fc8a96627c41.mongodb.rds.aliyuncs.com:3717", "dds-m5e04fc8a96627c42.mongodb.rds.aliyuncs.com:3717"},
	DBHosts:    []string{"10.3.9.32:27017"},
	Source:     "admin",
	DBName:     "vehicle",
	DBUsername: "vehicle_rw",
	DBPassword: "Rfvbhu#!7620",
	Tag:        "PROD",
}

// 门店服务生产环境数据库配置
var DBServiceMerchantProdConfig = DBConfig{
	DBHosts:    []string{"dds-m5e509e61d19fa541.mongodb.rds.aliyuncs.com:3717", "dds-m5e509e61d19fa542.mongodb.rds.aliyuncs.com:3717"},
	DBName:     "service-merchant",
	DBUsername: "service-merchant",
	DBPassword: "101010Pp",
	Tag:        "PROD",
}

// 签约服务生产环境数据库配置
var DBServiceBillingProdConfig = DBConfig{
	DBHosts:    []string{"rm-m5ej5d2034yz1x9qb.mysql.rds.aliyuncs.com:3306"},
	DBName:     "kanche-billing",
	DBUsername: "billing",
	DBPassword: "WorkFlow$(20161201!",
	Tag:        "PROD",
}

// 金融服务生产环境数据库配置
var DBServiceFinanceProdConfig = DBConfig{
	DBHosts:    []string{"rm-m5e3my8m426s2m447.mysql.rds.aliyuncs.com:3306"},
	DBName:     "finance_sales",
	DBUsername: "finance",
	DBPassword: "WorkFlow$(20160804!",
	Tag:        "PROD",
}

var DBConfigMap = map[string]DBConfig{
	ServiceVehicle + Prod: DBServiceVehicleProdConfig,
	// ServiceMerchant + Prod: DBServiceMerchantProdConfig,
	// ServiceBilling + Prod:  DBServiceBillingProdConfig,
	// ServiceFinance + Prod:  DBServiceFinanceProdConfig,
}

func GetDBConfig(serviceName string) (DBConfig, error) {
	if c, ok := DBConfigMap[serviceName+Env]; ok {
		return c, nil
	}
	return DBConfig{}, errors.New("获取数据库配置失败")
}
