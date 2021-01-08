package services

import (
	"errors"
	"ginauth101/database"
	"ginauth101/models"

	"github.com/goonode/mogo"
	"labix.org/v2/mgo/bson"
)

//Userservice is to handle user relation db query
type Userservice struct{}

//Create is to register new user
func (userservice Userservice) Create(user *(models.User)) error {
	conn := database.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(models.User{}).(*(models.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)
	if err == nil {
		return errors.New("Already Exist")
	}
	userModel := mogo.NewDoc(user).(*(models.User))
	err = mogo.Save(userModel)
	if vErr, ok := err.(*mogo.ValidationError); ok {
		return vErr
	}
	return err
}

// Delete a user from DB
func (userservice Userservice) Delete(email string) error {
	user, _ := userservice.FindByEmail(email)
	conn := database.GetConnection()
	defer conn.Session.Close()
	err := user.Remove()
	return err
}

//Find user
func (userservice Userservice) Find(user *(models.User)) (*models.User, error) {
	conn := database.GetConnection()
	defer conn.Session.Close()

	doc := mogo.NewDoc(models.User{}).(*(models.User))
	err := doc.FindOne(bson.M{"email": user.Email}, doc)

	if err != nil {
		return nil, err
	}
	return doc, nil
}

//FindByEmail is to find user from email
func (userservice Userservice) FindByEmail(email string) (*models.User, error) {
	conn := database.GetConnection()
	defer conn.Session.Close()

	user := new(models.User)
	user.Email = email
	return userservice.Find(user)
}
