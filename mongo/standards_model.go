package mongo

import (
	"gopkg.in/mgo.v2"
	"ui-mockup-backend"
)

type standardModel struct {
	//ID bson.ObjectId `bson:"_id,omitempty"`
	Controls[] root.Controls
	StandardName string
	//StandardName string
}

func standardsIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"standardName"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func certificationIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"certificationName"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func userCertIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

// TODO: NO NEED OF THIS FUNCTION PROBABLY
func newStandardModel(std *root.Standard) (*standardModel) {
	standard_model := standardModel{StandardName: std.StandardName, Controls:std.Controls}
	return &standard_model
}