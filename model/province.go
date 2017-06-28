package model

import (
	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type Province struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgcCode" bson:"psgcCode"`
	RegCode  string        `json:"regCode" bson:"regCode"`
}

func (province *Province) Create() *Province {
	if err := db.C("provinces").Insert(province); err != nil {
		panic(err)
	}
	defer db.Close()
	return province
}
