package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	db *mgo.Database
}

func NewMgo(url, db string) *Mongo {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("mongodb dial error : " + err.Error())
	}
	return &Mongo{
		db: session.DB(db),
	}
}

func (m *Mongo) Insert(c string, docs ...interface{}) error {
	return m.db.C(c).Insert(docs)
}

func (m *Mongo) AddToSet(c string, selector, update interface{}) error {
	change := bson.M{"$addToSet": update}
	return m.db.C(c).Update(selector, change)
}

func (m *Mongo) RemoveFromSet(c string, selector, update interface{}) error {
	change := bson.M{"$pull": update}
	return m.db.C(c).Update(selector, change)
}

func (m *Mongo) Find(c string, query, result interface{}) error {
	return m.db.C(c).Find(query).One(result)
}

func (m *Mongo) FindAll(c string, query, result interface{}) error {
	return m.db.C(c).Find(query).All(result)
}

func (m *Mongo) Update(c string, selector, update interface{}) error {
	return m.db.C(c).Update(selector, update)
}

func (m *Mongo) UpdateAll(c string, seletor, update interface{}) error {
	_, err := m.db.C(c).UpdateAll(seletor, update)
	return err
}

func (m *Mongo) Remove(c string, selector interface{}) error {
	return m.db.C(c).Remove(selector)
}
