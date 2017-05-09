package position

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Position struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Position, error) {
	poss := []Position{}

	if err := db.Positions.Find(nil).Sort("+name").All(&poss); err != nil {
		return nil, err
	}
	return poss, nil
}

func (pos *Position) Insert() (string, error) {
	pos.ObjectID = bson.NewObjectId()

	if err := db.Positions.Insert(pos); err != nil {
		return "", err
	}
	return pos.ObjectID.Hex(), nil
}

func FindByID(id string) (*Position, error) {
	var pos Position

	if !bson.IsObjectIdHex(id) {
		return &pos, models.ErrInvalidObjectID
	}

	if err := db.Positions.FindId(bson.ObjectIdHex(id)).One(&pos); err != nil {
		return &pos, err
	}
	return &pos, nil
}

func Search(query interface{}) ([]Position, error) {
	poss := []Position{}

	if err := db.Positions.Find(query).Sort("+name").All(&poss); err != nil {
		return nil, err
	}
	return poss, nil
}
