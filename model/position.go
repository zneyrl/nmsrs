package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/emurmotol/nmsrs/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func positionSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/positions.json")

	if err != nil {
		panic(err)
	}
	positions := []Position{}

	if err := json.Unmarshal(data, &positions); err != nil {
		panic(err)
	}
	log.Println("positionSeeder: todo")
}

type Position struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (position *Position) Create() *Position {
	if err := db.C("positions").Insert(position); err != nil {
		panic(err)
	}
	defer db.Close()
	return position
}

func (position Position) Index(q string) []Position {
	positions := []Position{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("positions").Find(query).All(&positions); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return positions
}
