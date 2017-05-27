package models

import (
	"errors"
	"time"

	"github.com/Renatdk/Bookshelf/db"
	"github.com/Renatdk/Bookshelf/forms"
)

//Library ...
type Library struct {
	ID        int64    `db:"id, primarykey, autoincrement" json:"id"`
	UserID    int64    `db:"user_id" json:"-"`
	Title     string   `db:"title" json:"title"`
	Address   string   `db:"address" json:"address"`
	Phone     string   `db:"phone" json:"phone"`
	UpdatedAt int64    `db:"updated_at" json:"updated_at"`
	CreatedAt int64    `db:"created_at" json:"created_at"`
	User      *JSONRaw `db:"user" json:"user"`
}

//LibraryModel ...
type LibraryModel struct{}

//Create ...
func (m LibraryModel) Create(userID int64, form forms.LibraryForm) (libraryID int64, err error) {
	getDb := db.GetDB()

	userModel := new(UserModel)

	checkUser, err := userModel.One(userID)

	if err != nil && checkUser.ID > 0 {
		return 0, errors.New("User doesn't exist")
	}

	_, err = getDb.Exec("INSERT INTO library(user_id, title, address, phone, updated_at, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", userID, form.Title, form.Address, form.Phone, time.Now().Unix(), time.Now().Unix())

	if err != nil {
		return 0, err
	}

	libraryID, err = getDb.SelectInt("SELECT id FROM library WHERE user_id=$1 ORDER BY id DESC LIMIT 1", userID)

	return libraryID, err
}

//One ...
func (m LibraryModel) One(userID, id int64) (library Library, err error) {
	err = db.GetDB().SelectOne(&library, "SELECT l.id, l.title, l.address, l.phone, l.updated_at, l.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM library l LEFT JOIN public.user u ON l.user_id = u.id WHERE l.user_id=$1 AND l.id=$2 GROUP BY l.id, l.title, l.address, l.phone, l.updated_at, l.created_at, u.id, u.name, u.email LIMIT 1", userID, id)
	return library, err
}

//All ...
func (m LibraryModel) All(userID int64) (libraries []Library, err error) {
	_, err = db.GetDB().Select(&libraries, "SELECT l.id, l.title, l.address, l.phone, l.updated_at, l.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM library l LEFT JOIN public.user u ON l.user_id = u.id WHERE l.user_id=$1 GROUP BY l.id, l.title, l.address, l.phone, l.updated_at, l.created_at, u.id, u.name, u.email ORDER BY l.id DESC", userID)
	return libraries, err
}

//Update ...
func (m LibraryModel) Update(userID int64, id int64, form forms.LibraryForm) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("Library not found")
	}

	_, err = db.GetDB().Exec("UPDATE library SET title=$1, address=$2, phone=$3, updated_at=$4 WHERE id=$5", form.Title, form.Address, form.Phone, time.Now().Unix(), id)

	return err
}

//Delete ...
func (m LibraryModel) Delete(userID, id int64) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("Library not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM library WHERE id=$1", id)

	return err
}
