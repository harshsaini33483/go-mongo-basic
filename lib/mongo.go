package lib

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

var DefaultCollection *mgo.Collection

// DB connection
func connect(host string) (session *mgo.Session, err error) {
	session, err = mgo.Dial(host)
	return
}

// Init the database
func Init(host, db, collection string) {

	fmt.Println(host, db, collection)

	session, err := connect(host)
	if err != nil {
		log.Fatalf("Connection failed to DB", err)
	}

	DefaultCollection = session.DB(db).C(collection)

}
