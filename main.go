package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/Renatdk/Bookshelf/controllers"
	"github.com/Renatdk/Bookshelf/db"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	store, _ := sessions.NewRedisStore(10, "tcp", "redis-17101.c3.eu-west-1-1.ec2.cloud.redislabs.com:17101", "", []byte("secret"))
	r.Use(sessions.Sessions("bookshelf-session", store))

	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/signin", user.Signin)
		v1.POST("/user/signup", user.Signup)
		v1.GET("/user/signout", user.Signout)

		/*** START Article ***/
		library := new(controllers.LibraryController)

		v1.POST("/library", library.Create)
		v1.GET("/articles", library.All)
		v1.GET("/library/:id", library.One)
		v1.PUT("/library/:id", library.Update)
		v1.DELETE("/library/:id", library.Delete)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/assets/img/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.02",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	r.Run(":" + port)
}
