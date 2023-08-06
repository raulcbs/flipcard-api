package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raulcbs/flipcard-api/internal/database"
)

var db = make(map[string]string)

type apiConfig struct {
	DB *database.Queries
}

func getEnv() map[string]string {
	envMap := make(map[string]string)

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment")
	}
	envMap["port"] = portString

	dbUrlString := os.Getenv("DB_URL")
	if dbUrlString == "" {
		log.Fatal("DB_URL is not found in the enviroment")
	}
	envMap["db"] = dbUrlString

	return envMap
}

func setupRouter() *gin.Engine {

	envs := getEnv()

	connect, err := sql.Open("postgres", envs["port"])
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	apiCfg := apiConfig{
		DB: database.New(connect),
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	router.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Post user
	router.POST("/user", func(c *gin.Context) {
		data, err := apiCfg.DB.CreateUser(c, database.CreateUserParams{ID: uuid.UUID{}, Name: "Manolo", Email: sql.NullString{String: "manolo@gmail.com"}, Password: "1234", CreatedAt: time.Now(), UpdatedAt: time.Now()})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something was wrong to create a user"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "user created", "user": data})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
