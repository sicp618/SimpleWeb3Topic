package controller

import (
	"SimpleBlog/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	COOKIE_NAME = "gin_cookie"
	REDIS_URL   = "localhost:6379"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: "",
		DB:       0,
	})
}

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Register(c *gin.Context) {
	var params User
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(params.Username) == 0 || len(params.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("params error")})
		return
	}
	user := model.User{Username: params.Username, Password: params.Password}
	if err := user.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func LoginGet(c *gin.Context) {
	log.Info("loginGet")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "*")
	c.JSON(http.StatusOK, gin.H{
		"login": "ok",
	})
}

func LoginOptions(c *gin.Context) {
	log.Info("loginOptions")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	//c.Header("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	c.Header("Access-Control-Expose-Headers", "*")
	c.JSON(http.StatusOK, gin.H{
		"login": "ok",
	})
}

func Login(c *gin.Context) {
	//cookie, _ := c.Cookie(COOKIE_NAME)
	//if cookie != "" {
	//	val, _ := rdb.Get(c, COOKIE_NAME).Result()
	//	if val == cookie {
	//		return
	//	}
	//}
	var params User
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(params.Username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username cant't be empty"})
		return
	}

	log.Info("params", params)
	user := model.User{Username: params.Username, Password: params.Password}
	if err := user.Find(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genCookie(params.Username, c)
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "*")
	c.JSON(http.StatusOK, gin.H{
		"login": "ok",
	})
}

func genCookie(username string, c *gin.Context) {
	cookie := uuid.NewString()
	c.SetCookie(COOKIE_NAME, cookie, 3600,
		"/", "localhost", false, true)
	log.Printf("gen cookie %s", cookie)
	if err := rdb.Set(c, username, cookie, time.Hour).Err(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": err.Error(),
		})
		return
	}
}

func Ping(c *gin.Context) {
	cookie, err := c.Cookie(COOKIE_NAME)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	val, err := rdb.Get(c, COOKIE_NAME).Result()
	if err != nil || val != cookie {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "response",
	})
}
