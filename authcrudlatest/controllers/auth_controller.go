package controllers

import (
	"authcrudlatest/config"
	"authcrudlatest/models"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"message": "Fill the rgistrtaion form",
	})
}

func Register(c *gin.Context) {
	//var input models.User
	fmt.Println("-Post Register Is Running-")

	username := c.PostForm("Username")
	password := c.PostForm("Password")
	fmt.Println("-Username:", username)
	fmt.Println("-Password:", password)

	// Check if username already exists
	var existingUser models.User
	if err := config.DB.Where("username = ?", username).First(&existingUser).Error; err == nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Username already exists",
		})
		return
	}
	fmt.Println("No duplicate users found")

	// Create new user record
	input := models.User{
		Username: username,
		Password: password,
	}
	fmt.Println("Post data bind assigned with model")

	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	//fmt.Println("Password encrypted")

	/*
		user := models.User{
			Username: input.Username,
			Password: string(hashedPassword),
		}
	*/
	config.DB.Create(&input)

	//c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	c.HTML(http.StatusBadRequest, "register.html", gin.H{
		"success": "Registration successful for the user: ",
		"user":    username,
	})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Welcome to the server",
	})
}

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "Fill the form",
	})
}

func Login(c *gin.Context) {
	//var input models.User
	fmt.Println("-Post Login Is Running-")

	username := c.PostForm("Username")
	password := c.PostForm("Password")
	fmt.Println("-Username:", username)
	fmt.Println("-Password:", password)

	var user models.User
	config.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "User not found"})
		return
	}
	fmt.Println("User found")
	// Check password
	if user.Password != password {
		fmt.Println("user.password", user.Password)
		fmt.Println("password", password)
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": "Invalid password",
		})
		return
	}
	fmt.Println("Password matched")

	/*
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}
	*/
	// Create a session
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()
	fmt.Println("Session", user.ID)
	fmt.Println("Session saved")
	c.Redirect(http.StatusFound, "/dashboard")
	//c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"logout": "Logged successfully",
	})
	//c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
