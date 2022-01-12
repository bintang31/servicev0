package config

import (
	"math"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//CronJob : Struct Load Config CronJob
type CronJob struct {
	CronTime string
}

//Database : Struct Load Config
type Database struct {
	Driver            string
	Host              string
	User              string
	Password          string
	DBName            string
	QueueName         string
	WorkerLimit       int
	Port              string
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
}

// LoadCronJob load module loket configuration
func LoadCronJob(name string) CronJob {
	db := viper.Sub(name)
	conf := CronJob{
		CronTime: db.GetString("daily"),
	}
	return conf
}

// Param 分页参数
type Param struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	OrderBy []string
	Select  string
	Join    string
	Where   string
	Group   string
	ShowSQL bool
}

// Paginator 分页返回
type Paginator struct {
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

//Loket : Struct Load Config Loket
type Loket struct {
	EndPoint  string
	User      string
	Password  string
	UserLoket string
}

//EndPointPayment : Struct Load Config EndPointPayment
type EndPointPayment struct {
	Secret           string
	ClientID         string
	StoreExtID       string
	MerchantExtID    string
	Currency         string
	TerminalID       string
	QrValidityPeriod int32
	Host             string
	QrDirectory      string
}

// LoadAppConfig load database configuration
func LoadAppConfig(name string) Database {
	db := viper.Sub("database." + name)
	conf := Database{
		Driver:            db.GetString("driver"),
		Host:              db.GetString("host"),
		User:              db.GetString("user"),
		Password:          db.GetString("password"),
		DBName:            db.GetString("db_name"),
		QueueName:         db.GetString("queue_name"),
		WorkerLimit:       db.GetInt("worker_limit"),
		Port:              db.GetString("port"),
		ReconnectRetry:    db.GetInt("reconnect_retry"),
		ReconnectInterval: db.GetInt64("reconnect_interval"),
		DebugMode:         db.GetBool("debug"),
	}
	return conf
}

// LoadModuleLoketConfig load module loket configuration
func LoadModuleLoketConfig(name string) Loket {
	db := viper.Sub(name)
	conf := Loket{
		EndPoint:  db.GetString("endpoint"),
		User:      db.GetString("user"),
		Password:  db.GetString("password"),
		UserLoket: db.GetString("userloket"),
	}
	return conf
}

// LoadPaymentConfig load module loket configuration
func LoadPaymentConfig(name string) EndPointPayment {
	db := viper.Sub(name)
	conf := EndPointPayment{
		Secret:           db.GetString("secret"),
		ClientID:         db.GetString("client_id"),
		StoreExtID:       db.GetString("store_ext_id"),
		MerchantExtID:    db.GetString("merchant_ext_id"),
		Currency:         db.GetString("currency"),
		TerminalID:       db.GetString("terminal_id"),
		QrValidityPeriod: db.GetInt32("qr_validity_period"),
		QrDirectory:      db.GetString("qr_directory"),
		Host:             db.GetString("host"),
	}
	return conf
}

// Paging 分页
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	//db.Select(p.Select).Joins(p.Join).Where(p.Where).Group(p.Group).Limit(p.Limit).Offset(offset).Find(result)
	//<-done

	db.Raw(p.Select).Limit(p.Limit).Offset(offset).Find(result)
	<-done

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = p.Page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
