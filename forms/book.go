package forms

//BookForm ...
type BookForm struct {
	Title		string `form:"title" json:"title" binding:"required,max=100"`
	Autor		string `form:"autor" json:"autor" binding:"required,max=100"`
	Description string `form:"desctiption" json:"desctiption" binding:"required,max=1000"`
}
