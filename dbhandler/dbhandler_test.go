package dbhandler

import (
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"	
	"testing"
	"fmt"
	"time"
	"github.com/remotejob/real_estate_redux_go/domains"
	"log"	
)

var addrs []string
var database string
var username string
var password string
var mechanism string

func init() {

	var cfg domains.ServerConfig
	if err := gcfg.ReadFileInto(&cfg, "/home/juno/neonworkspace/real_estate_redux_go/config.gcfg"); err != nil {
		log.Fatalln(err.Error())

	} else {

		addrs = cfg.Dbmgo.Addrs
		database = cfg.Dbmgo.Database
		username = cfg.Dbmgo.Username
		password = cfg.Dbmgo.Password
		mechanism = cfg.Dbmgo.Mechanism

	}

}

func TestGetAll(t *testing.T) {

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Timeout:   60 * time.Second,
		Database:  database,
		Username:  username,
		Password:  password,
		Mechanism: mechanism,
	}

	//	dbsession, err := mgo.Dial("127.0.0.1")
	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()
		
	
	results := GetAll(*dbsession)
//	GetAll(*dbsession)	
	
	for _,res := range results {
		
		fmt.Println(res)		
	}
	
//	fmt.Println(result)

}


