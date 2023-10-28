// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.File("views/index.html")
// 	})

// 	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rachma1201/Task-5-pbi-btpns-Rahmaliyah-Kadir/database"
	"github.com/rachma1201/Task-5-pbi-btpns-Rahmaliyah-Kadir/router"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.File("views/index.html")
	})

	database.ConnectDB()
	router.InitRouter()
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
