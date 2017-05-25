package forms

//LibraryForm ...
type LibraryForm struct {
	Title   string `form:"title" json:"title" binding:"required,max=100"`
	Address string `form:"address" json:"address" binding:"required,max=250"`
	Phone   string `form:"phone" json:"phone" binding:"required,max=25"`
}
