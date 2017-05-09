package otherskill

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type OtherSkill struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]OtherSkill, error) {
	oskills := []OtherSkill{}

	if err := db.OtherSkills.Find(nil).Sort("+name").All(&oskills); err != nil {
		return nil, err
	}
	return oskills, nil
}

func (oskill *OtherSkill) Insert() (string, error) {
	oskill.ObjectID = bson.NewObjectId()

	if err := db.OtherSkills.Insert(oskill); err != nil {
		return "", err
	}
	return oskill.ObjectID.Hex(), nil
}

func FindByID(id string) (*OtherSkill, error) {
	var oskill OtherSkill

	if !bson.IsObjectIdHex(id) {
		return &oskill, models.ErrInvalidObjectID
	}

	if err := db.OtherSkills.FindId(bson.ObjectIdHex(id)).One(&oskill); err != nil {
		return &oskill, err
	}
	return &oskill, nil
}
