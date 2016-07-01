package dbhandler

import (
//	"fmt"
	"github.com/remotejob/real_estate_redux_go/domains"
	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
	//	"github.com/fatih/structs"
	"log"
)

func GetAll(dbsession mgo.Session) domains.AutoRealEstate {

	dbsession.SetMode(mgo.Monotonic, true)

	c := dbsession.DB("real_estate").C("real_estate")
	var results domains.AutoRealEstate
	err := c.Find(nil).All(&results)
	if err != nil {

		log.Fatal(err)
	}

	//	fmt.Println(results)

//	for _, res := range results {
//
//		fmt.Println(res)
//
//	}

	//	var m Person
	//	err := c.Find(nil).One(&m)
	//	if err != nil {
	//
	//		log.Fatal(err)
	//	}
	//	for key, value := range m.Extra {
	//		fmt.Println(key, value)
	//	}

	//	image1 := domains.Image{Id: 0, Description: "image1.jpg"}
	//	image2 := domains.Image{Id: 1, Description: "image2.jpg"}
	//	images1 := []domains.Image{image1, image2}
	//
	//	action1 := domains.Action{Type: "rent", Price: 10000, Images: images1}
	//	actions1 := []domains.Action{action1}
	//
	//	realsEstateObj1 := domains.RealsEstateObj{Name: "Apartment", Type: "apartment", City: "Helsinki", Country: "Finland", Total_area:90,Actions: actions1}
	//
	//	err = c.Insert(realsEstateObj1)
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	fmt.Println(structs.Map(m.Extra) )

	//	var d bson.D
	//	err = c.Find(nil).One(&d)
	//	if err != nil {
	//
	//		log.Fatal(err)
	//	}
	//
	//	for _, elem := range d {
	//		fmt.Println(elem.Name, elem.Value)
	//	}

	return results
}
