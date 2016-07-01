package domains

import (
	//	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
)

type AutoRealEstate []struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `bson:"Name"`
	Type      string        `bson:"Type"`
	City      string        `bson:"City"`
	Country   string        `bson:"Country"`
	TotalArea int           `bson:"Total_area"`
	Actions   struct {
		Action []struct {
			Type   string `bson:"Type"`
			Price  int    `bson:"Price"`
			Images struct {
				Image []struct {
					ID          int    `bson:"Id"`
					Description string `bson:"Description"`
				} `bson:"Image"`
			} `bson:"Images"`
		} `bson:"Action"`
	} `bson:"Actions"`
}

type Image struct {
	Id          int
	Description string
}

type Action struct {
	Type   string
	Price  int
	Images []Image
}

type RealsEstateObjDump struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string
	Forecast interface{}
}

type RealsEstateObj struct {
	Name       string
	Type       string
	City       string
	Country    string
	Total_area int
	Actions    []Action
}

type ServerConfig struct {
	Dbmgo struct {
		Addrs     []string
		Database  string
		Username  string
		Password  string
		Mechanism string
	}
}
