package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type History struct {
	Timestamp int64
	Likes     int
	Reposts   int
	Comments  int
	Rating    float64
}

type News struct {
	_ID      bson.ObjectId `bson:"_id"`
	ID       int
	OwnerID  int
	Domein   string
	People   int
	Date     int
	Text     string
	Comments int
	Likes    int
	Reposts  int
	Rating   float64
	Trending float64
	History  []History `bson:"history"`
}

func (n *News) calculateRating() {
	dbn, _ := n.GetById()
	trending := 0.0
	if len(dbn.History) > 15 {
		l := dbn.History[len(dbn.History)-15:]
		trending = l[0].Rating - dbn.Rating
	}
	query := bson.M{"$set": bson.M{"trending": trending}}
	n.Change(query)
}

func (n *News) Exists() bool {
	db := getDatabase("news")
	err := db.Find(bson.M{"id": n.ID}).One(nil)
	return err == nil
}

func (n *News) GetById() (*News, error) {
	var ret News
	db := getDatabase("news")
	err := db.Find(bson.M{"id": n.ID}).One(&ret)
	return &ret, err
}

func (n *News) Insert() (err error) {
	db := getDatabase("news")
	err = db.Insert(n)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (n *News) Change(query bson.M) (err error) {
	db := getDatabase("news")
	change := mgo.Change{
		Update: query,
	}
	_, err = db.Find(bson.M{"id": n.ID}).Apply(change, nil)

	return err
}

func (n *News) UpdateOrInsert() (err error) {
	//db := getDatabase("news")
	indb := n.Exists()

	if indb != true {
		err = n.Insert()
	} else {
		query := bson.M{"$set": bson.M{"likes": n.Likes, "reposts": n.Reposts, "comments": n.Comments}}
		err = n.Change(query)
	}
	return err
}

func (n *News) UpdateHistory() (err error) {
	db := getDatabase("news")
	history := History{
		Likes:     n.Likes,
		Comments:  n.Comments,
		Reposts:   n.Reposts,
		Rating:    n.Rating,
		Timestamp: time.Now().Unix(),
	}

	change := mgo.Change{
		Update: bson.M{"$push": bson.M{"history": &history}},
	}
	n.calculateRating()
	_, err = db.Find(bson.M{"id": n.ID}).Apply(change, nil)
	return err
}
