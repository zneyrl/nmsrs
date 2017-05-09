package school

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type School struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]School, error) {
	schs := []School{}

	if err := db.Schools.Find(nil).Sort("+name").All(&schs); err != nil {
		return nil, err
	}
	return schs, nil
}

func (sch *School) Insert() (string, error) {
	sch.ObjectID = bson.NewObjectId()

	if err := db.Schools.Insert(sch); err != nil {
		return "", err
	}
	return sch.ObjectID.Hex(), nil
}

func FindByID(id string) (*School, error) {
	var sch School

	if !bson.IsObjectIdHex(id) {
		return &sch, models.ErrInvalidObjectID
	}

	if err := db.Schools.FindId(bson.ObjectIdHex(id)).One(&sch); err != nil {
		return &sch, err
	}
	return &sch, nil
}
