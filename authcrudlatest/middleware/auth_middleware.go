package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Unauthorized"})

			c.Abort()
			return
		}

		c.Next()
	}
}

// Middleware to check if user is already authenticated

func AlreadyAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID != nil {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		c.Next()
	}
}
