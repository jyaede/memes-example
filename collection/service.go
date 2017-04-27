package collection

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Manager ...
type Manager interface {
	All() ([]Item, error)
	Create(i *Item) (*Item, error)
	Get(id bson.ObjectId) (*Item, error)
	Delete(id bson.ObjectId) error
}

type manager struct {
	sess *mgo.Session
}

//NewManager ...
func NewManager(sess *mgo.Session) Manager {
	return manager{sess}
}

func (m manager) col() *mgo.Collection {
	return m.sess.DB("memes").C("my-meme-collection")
}

func (m manager) All() ([]Item, error) {
	var items []Item
	if err := m.col().Find(nil).All(&items); err != nil {
		return nil, err
	}
	return items, nil
}

func (m manager) Create(i *Item) (*Item, error) {
	i.ID = bson.NewObjectId()
	if err := m.col().Insert(i); err != nil {
		return nil, err
	}
	return i, nil
}

func (m manager) Get(id bson.ObjectId) (*Item, error) {
	var item *Item
	if err := m.col().FindId(id).All(&item); err != nil {
		return nil, err
	}
	return item, nil
}

func (m manager) Delete(id bson.ObjectId) error {
	return m.col().RemoveId(id)
}
