package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Saveindb(t Trip) {
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")
	c := session.DB("tests").C("Trips")
	err = c.Insert(&Trip{Id: t.Id,
		Status:                 t.Status,
		Startingfromlocationid: t.Startingfromlocationid,
		Bestroutelocationids:   t.Bestroutelocationids,
		Locationids:            t.Locationids,
		Totalubercosts:         t.Totalubercosts,
		Totaluberduration:      t.Totaluberduration,
		Totaldistance:          t.Totaldistance})
	if err != nil {
		log.Fatal(err)
	}
}
func Updateindb(t Trip) {
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")
	c := session.DB("tests").C("Trips")
	err = c.Remove(bson.M{"id": t.Id})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Insert(&Trip{Id: t.Id,
		Status:                    t.Status,
		Startingfromlocationid:    t.Startingfromlocationid,
		Nextdestinationlocationid: t.Nextdestinationlocationid,
		//Bestroutelocationids:      t.Bestroutelocationids,
		Locationids:       t.Locationids,
		Totalubercosts:    t.Totalubercosts,
		Totaluberduration: t.Totaluberduration,
		Totaldistance:     t.Totaldistance,
		Uberwaittime:      t.Uberwaittime})
	if err != nil {
		log.Fatal(err)
	}
}
