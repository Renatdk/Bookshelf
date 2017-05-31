package models

import (
	"errors"
	"time"

	"github.com/Renatdk/Bookshelf/db"
	"github.com/Renatdk/Bookshelf/forms"
)

//Book ...
type Book struct {
	ID          int64    `db:"id, primarykey, autoincrement" json:"id"`
	LibraryID   int64    `db:"library_id" json:"-"`
	Title       string   `db:"title" json:"title"`
	Author      string   `db:"author", json:"author"`
	Description string   `db:"description" json:"description"`
	UpdatedAt   int64    `db:"updated_at" json:"updated_at"`
	CreatedAt   int64    `db:"created_at" json:"created_at"`
	Library     *JSONRaw `db:"library" json:"library"`
}

//BookModel ...
type BookModel struct{}

//Create ...
func (m BookModel) Create(UserID int64, form forms.BookForm) (bookID int64, err error) {
	getDb := db.GetDB()

	libraryModel := new(LibraryModel)

	checkLibrary, err := libraryModel.One(UserID, form.LibraryID)

	if err != nil && checkLibrary.ID > 0 {
		return 0, errors.New("Library doesn't exist")
	}

	_, err = getDb.Exec("INSERT INTO book(library_id, title, author, description, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", form.LibraryID, form.Title, form.Author, form.Description, time.Now().Unix(), time.Now().Unix())

	if err != nil {
		return 0, err
	}

	bookID, err = getDb.SelectInt("SELECT id FROM book WHERE library_id=$1 ORDER BY id DESC LIMIT 1", form.LibraryID)

	return bookID, err
}

//One ...
func (m BookModel) One(libraryID, id int64) (book Book, err error) {
	err = db.GetDB().SelectOne(&book, "SELECT b.id, b.title, b.author, b.description, b.updated_at, b.created_at, json_build_object('id', l.id, 'title', l.title, 'address', l.address) AS library FROM book b LEFT JOIN public.library l ON b.library_id = l.id WHERE b.library_id=$1 AND b.id=$2 GROUP BY b.id, b.title, b.author, b.description, b.updated_at, b.created_at, l.id, l.name, l.email LIMIT 1", libraryID, id)
	return book, err
}

//All ...
func (m BookModel) All(libraryID int64) (books []Book, err error) {
	_, err = db.GetDB().Select(&books, "SELECT b.id, b.title, b.author, b.description, b.updated_at, b.created_at, json_build_object('id', l.id, 'title', l.title, 'address', l.address) AS library FROM book b LEFT JOIN public.library l ON b.library_id = l.id WHERE b.library_id=$1 GROUP BY b.id, b.title, b.author, b.description, b.updated_at, b.created_at, l.id, l.name, l.email ORDER BY b.id DESC", libraryID)
	return books, err
}

//Update ...
func (m BookModel) Update(libraryID int64, id int64, form forms.BookForm) (err error) {
	_, err = m.One(libraryID, id)

	if err != nil {
		return errors.New("Book not found")
	}

	_, err = db.GetDB().Exec("UPDATE book SET title=$1, author=$2, description=$3, updated_at=$4 WHERE id=$5", form.Title, form.Author, form.Description, time.Now().Unix(), id)

	return err
}

//Delete ...
func (m BookModel) Delete(libraryID, id int64) (err error) {
	_, err = m.One(libraryID, id)

	if err != nil {
		return errors.New("Book not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM book WHERE id=$1", id)

	return err
}
