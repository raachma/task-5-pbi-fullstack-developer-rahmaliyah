// package middlewares

// import (
// 	"github.com/labstack/echo/v4"
// )

//	func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			// Implement your authentication logic here
//			// For example, check JWT or session
//			// If authenticated, proceed to the next handler
//			// If not, return an error response
//			return next(c)
//		}
//	}
// package middlewares

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // Middleware untuk autentikasi user
// func Auth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// token dari request header
// 		token := c.Request.Header.Get("Authorization")

// 		// validasi token
// 		if token == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak ditemukan"})
// 			c.Abort()
// 			return
// 		}

// 		// pecah token menjadi header dan payload
// 		header, payload, err := jwt.ParseRSA(token)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
// 			c.Abort()
// 			return
// 		}

//			// cek token masih valid atau tidak
//			claims := jwt.Claims{}
//			err = json.Unmarshal(payload, &claims)
//			if err != nil {
//				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
//				c.Abort()
//				return
//			}
//		}
//	}
package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware(c *gin.Context) {
	// Middleware untuk otentikasi
}
