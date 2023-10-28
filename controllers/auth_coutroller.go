// controllers/auth_controller.go
// package controllers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/rachma1201/task-5-pbi-btpns-Rahmaliyah-Kadir/helpers"
// 	"github.com/rachma1201/task-5-pbi-btpns-Rahmaliyah-Kadir/models"
// )

// func Login(c *gin.Context) {
// 	var credentials models.Credentials
// 	if err := c.ShouldBindJSON(&credentials); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := helpers.AuthenticateUser(credentials.Email, credentials.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	token, err := helpers.GenerateToken(user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"token": token})
// }

package controllers

import (
	// "github.com/rachma1201/task-5-pbi-btpns-Rahmaliyah-Kadir/database"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	// Implementasi endpoint POST /photos
}

func GetPhotos(c *gin.Context) {
	// Implementasi endpoint GET /photos
}

func UpdatePhoto(c *gin.Context) {
	// Implementasi endpoint PUT /photos/:photoId
}

func DeletePhoto(c *gin.Context) {
	// Implementasi endpoint DELETE /photos/:photoId
}
