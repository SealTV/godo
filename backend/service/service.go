package service

import (
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/SealTV/go-site/backend/data"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Config - server config params
type Config struct {
	SecretKey string `json:"secret"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
}

//Service - server object
type Service struct {
	db     data.DBConnector
	router *gin.Engine
	config Config
}

type response struct {
	Result interface{} `json:"result"`
	Error  error       `json:"error"`
}

// New - create new instance of servers
func New(db data.DBConnector, c Config) *Service {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	s := Service{db, router, c}

	initRouters(&s, router)
	return &s
}

func initRouters(s *Service, router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AddAllowHeaders(
		"ExposedHeader",
		"X-Requested-With",
		"Authorization")
	config.AddExposeHeaders(
		"Authorization",
	)
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte(s.config.SecretKey),
		Timeout:       time.Hour * 24,
		MaxRefresh:    time.Hour * 24,
		Authenticator: s.authenticator,
		Authorizator:  s.authorizator,
		Unauthorized:  s.unauthorized,
		PayloadFunc:   s.payloadFunc,
		TimeFunc:      time.Now,
	}

	router.POST("/auth/register", s.register)
	router.POST("/auth/login", authMiddleware.LoginHandler)

	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/user", s.verify)
		auth.GET("/refresh", authMiddleware.RefreshHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.POST("/logout", s.logout)
		auth.POST("/delete", s.delete)
	}
}

// Run - run web service
func (s *Service) Run() error {
	return s.router.Run(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
}
