package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User structure for working with users
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Email    string        `bson:"email" json:"email"`
	Created  time.Time     `bson:"created" json:"created"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}

// FindUserByID searches for an existing user with the passed ID
func FindUserByID(id string, db *mgo.Database) (*User, error) {
	user := User{}
	err := db.C("users").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByUsername searches for an existing user with the passed Username
func FindUserByUsername(username string, db *mgo.Database) (*User, error) {
	user := User{}
	err := db.C("users").Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert creates a new User and saves it in the database
func (u User) Insert(db *mgo.Database) error {
	//verify the uname isn't already being used
	user, _ := FindUserByUsername(u.Username, db)
	if user != nil {
		return errors.New("user: The user already exists")
	}

	u.Created = time.Now()

	passHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("stack|user: %+v", err)
	}
	u.Password = string(passHash)

	err = db.C("users").Insert(u)
	if err != nil {
		return err
	}

	return nil
}

// Update is a method on User that updates the copy in the db
func (u User) Update(db *mgo.Database) error {
	err := db.C("users").UpdateId(u.ID, u)
	return err
}

// Delete is a method on User that will delete the User
func (u User) Delete(db *mgo.Database) error {
	err := db.C("users").RemoveId(u.ID)
	return err
}

// CheckPassword will check a passed password string with the stored hash
func (u User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
