package db

import (
	"gopkg.in/mgo.v2"
)

const MONGOSERVER string = "192.168.99.100"
const MONGODB string = "rfpl"

var _database *mgo.Session

func getDatabase(name string) (collection *mgo.Collection) {

	if _database != nil {
		collection = _database.DB(MONGODB).C(name)
	} else {
		session, err := mgo.Dial(MONGOSERVER) // Коннектимся к серваку
		if err != nil {                       // Если ошибка, то ну его
			panic(err)
		}
		_database = session
		session.SetMode(mgo.Monotonic, true)
		collection = session.DB(MONGODB).C(name)
	}
	return collection
}
