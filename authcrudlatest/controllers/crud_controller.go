package controllers

import (
	"authcrudlatest/config"
	"authcrudlatest/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"message": "Welcome to dashboard",
	})
}

func CreateDataForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"message": "Fill the data",
	})
}
func CreateData(c *gin.Context) {

	// Parse form data
	name := c.PostForm("name")
	address := c.PostForm("address")
	state := c.PostForm("state")

	// Validate input
	if name == "" || address == "" || state == "" {
		c.HTML(http.StatusBadRequest, "create.html", gin.H{
			"error": "Fill all the data",
		})
		return
	}

	// Handle file upload
	file, err := c.FormFile("photo")
	if err != nil {
		c.HTML(http.StatusOK, "create.html", gin.H{
			"error": "Failed to upload photo",
		})
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload photo"})
		return
	}

	// Create directory if it doesn't exist
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm)
	}

	// Save the uploaded file to a specific path
	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.HTML(http.StatusInternalServerError, "create.html", gin.H{
			"error": "Failed to save photo",
		})
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
		return
	}

	data := models.Data{
		Name:    name,
		Address: address,
		State:   state,
	}

	// Convert file path to use forward slashes
	filePath = filepath.ToSlash(filePath)
	// Set file path to data
	data.PhotoPath = filePath

	config.DB.Create(&data)
	//c.JSON(http.StatusOK, data)
	c.Redirect(http.StatusFound, "/show")
}

func ReadData(c *gin.Context) {
	var data []models.Data
	config.DB.Find(&data)
	c.HTML(http.StatusOK, "read.html", gin.H{
		"data": data,
	})
	//c.JSON(http.StatusOK, data)
}

func UpdateDataById(c *gin.Context) {
	var data models.Data
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.HTML(http.StatusNotFound, "dashboard.html", gin.H{
			"error": "User not found",
		})
		return
	}
	fmt.Println("------fetched successfully-------")
	c.HTML(http.StatusOK, "update.html", gin.H{
		"data":    data,
		"message": "Please enter the details",
	})
	/*
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

		config.DB.Save(&data)
		c.JSON(http.StatusOK, data)
	*/
}

func UpdateDataByIdSave(c *gin.Context) {
	fmt.Println("----edit form submitted------")
	var data models.Data
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.HTML(http.StatusNotFound, "dashboard.html", gin.H{
			"error": "User not found",
		})
		return
	}
	fmt.Println("------Data exists-------")

	name := c.PostForm("name")
	address := c.PostForm("address")
	state := c.PostForm("state")

	if name == "" || address == "" || state == "" {
		c.HTML(http.StatusBadRequest, "dashboard.html", gin.H{
			"error": "Nothing to update",
		})
		return
	}

	data.Name = name
	data.Address = address
	data.State = state

	config.DB.Save(&data)
	c.Redirect(http.StatusFound, "/show")

	/*
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

		config.DB.Save(&data)
		c.JSON(http.StatusOK, data)
	*/
}

func DeleteData(c *gin.Context) {
	var data models.Data
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.HTML(http.StatusNotFound, "read.html", gin.H{
			"error": "Data not found",
		})
		//c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	config.DB.Delete(&data)
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"success": "Data deleted",
	})
	//c.JSON(http.StatusOK, gin.H{"message": "Data deleted"})
}
