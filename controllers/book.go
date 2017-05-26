package controllers

import (
	"strconv"

	"github.com/Renatdk/Bookshelf/forms"
	"github.com/Renatdk/Bookshelf/models"

)

//BookController ...
type BookController struct{}

var bookModel = new(models.BookModel)

//Create ...
func (ctrl BookController) Create(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	var bookForm forms.BookForm

	if c.BindJSON(&bookForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": bookForm})
		c.Abort()
		return
	}

	articleID, err := articleModel.Create(userID, bookForm)

	if articleID > 0 && err != nil {
		c.JSON(406, gin.H{"message": "Book could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Book created", "id": articleID})
}

//All ...
func (ctrl BookController) All(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	data, err := articleModel.All(userID)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the articles", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl BookController) One(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := articleModel.One(userID, id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Book not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl BookController) Update(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var bookForm forms.BookForm

		if c.BindJSON(&bookForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": bookForm})
			c.Abort()
			return
		}

		err := articleModel.Update(userID, id, bookForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Book could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Book updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl BookController) Delete(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := articleModel.Delete(userID, id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Book could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Book deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}