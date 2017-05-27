package controllers

import (
	"strconv"

	gin "gopkg.in/gin-gonic/gin.v1"

	"github.com/Renatdk/Bookshelf/forms"
	"github.com/Renatdk/Bookshelf/models"
)

//LibraryController ...
type LibraryController struct{}

var libraryModel = new(models.LibraryModel)

//Create ...
func (ctrl LibraryController) Create(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	var libraryForm forms.LibraryForm

	if c.BindJSON(&libraryForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": libraryForm})
		c.Abort()
		return
	}

	libraryID, err := libraryModel.Create(userID, libraryForm)

	if libraryID > 0 && err != nil {
		c.JSON(406, gin.H{"message": "Library could not be created", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Library created", "id": libraryID})
}

//All ...
func (ctrl LibraryController) All(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	data, err := libraryModel.All(userID)

	if err != nil {
		c.JSON(406, gin.H{"Message": "Could not get the libraries", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl LibraryController) One(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := libraryModel.One(userID, id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Library not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Update ...
func (ctrl LibraryController) Update(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		var libraryForm forms.LibraryForm

		if c.BindJSON(&libraryForm) != nil {
			c.JSON(406, gin.H{"message": "Invalid parameters", "form": libraryForm})
			c.Abort()
			return
		}

		err := libraryModel.Update(userID, id, libraryForm)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Library could not be updated", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Library updated"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

//Delete ...
func (ctrl LibraryController) Delete(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		err := libraryModel.Delete(userID, id)
		if err != nil {
			c.JSON(406, gin.H{"Message": "Library could not be deleted", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"message": "Library deleted"})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
