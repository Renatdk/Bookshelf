package forms

//BookForm ...
type BookForm struct {
	Title       string `form:"title" json:"title" binding:"required,max=100"`
	Author      string `form:"author" json:"author" binding:"required,max=100"`
	Description string `form:"desctiption" json:"desctiption" binding:"required,max=1000"`
	LibraryID   int64  `db:"library_id, foreignkey" json:"library_id"`
}
