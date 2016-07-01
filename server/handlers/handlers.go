package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/remotejob/real_estate_redux_go/dbhandler"
	"github.com/remotejob/real_estate_redux_go/domains"
	"gopkg.in/gcfg.v1"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"time"
)

var addrs []string
var database string
var username string
var password string
var mechanism string

func AllRecords(w http.ResponseWriter, r *http.Request) {
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

	results := dbhandler.GetAll(*dbsession)

	output, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Something went wrong!")
	} else {
		fmt.Fprintf(w, string(output))
	}

}
