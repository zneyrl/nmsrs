package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Eligibility struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (eligibility *Eligibility) Create() *Eligibility {
	if err := db.C("eligibilities").Insert(eligibility); err != nil {
		panic(err)
	}
	defer db.Close()
	return eligibility
}

func (eligibility Eligibility) Index(q string) []Eligibility {
	eligibilities := []Eligibility{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("eligibilities").Find(query).All(&eligibilities); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return eligibilities
}
