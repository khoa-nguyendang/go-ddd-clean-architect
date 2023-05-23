package controllers

import (
	config "app/core/configs"
	"app/core/interfaces/bizs"
	"app/pkg/logger"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type RestfulServer struct {
	logger logger.Logger
	cfg    ServerConfig
	jb     bizs.JobBiz
}
type ServerConfig struct {
	AppVersion               string
	Port                     string
	PprofPort                string
	Mode                     string
	JwtSecretKey             string
	JwtExpireInHour          int
	RefreshSecretKey         string
	RefreshTokenExpireInHour int
	CookieName               string
	ReadTimeout              time.Duration
	WriteTimeout             time.Duration
	SSL                      bool
	CtxDefaultTimeout        time.Duration
	CSRF                     bool
	Debug                    bool
	MaxConnectionIdle        time.Duration
	Timeout                  time.Duration
	MaxConnectionAge         time.Duration
	Time                     time.Duration
	CacheExpiryShort         time.Duration // default 10 mins
	CacheExpiryMedium        time.Duration // default 1 hours
	CacheExpiryLong          time.Duration // default 4 hours
	CacheExpiryDayLong       time.Duration // default 1 day
	HashKey                  string
	PassKey                  string
	IvKey                    string
}

func NewConfig(conf *config.Config) (ServerConfig, error) {
	if conf == nil {
		return ServerConfig{}, errors.New("invalid configs")
	}
	c := ServerConfig{
		AppVersion: conf.Server.AppVersion,
		Port:       conf.Server.Port,
		Debug:      conf.Server.Debug,
	}
	return c, nil
}

// NewService func initializes a service
func New(
	logger *logger.ApiLogger,
	cfg ServerConfig,
	jb bizs.JobBiz,
) AppServer {
	return &RestfulServer{
		logger: logger,
		cfg:    cfg,
		jb:     jb,
	}
}

// Run starts the server
func (s *RestfulServer) Run(port int) error {
	router := gin.Default()
	router.GET("/jobs/search", s.searchJob)
	router.GET("/jobs/search-by-db", s.searchJobDatabase)
	router.GET("/jobs/:id", s.getJob)
	router.POST("/jobs", s.addTestJob)
	router.POST("/jobs-test", s.addJob)
	router.PUT("/jobs/:id", s.updateJob)
	router.DELETE("/jobs/:id", s.deleteJob)
	router.PATCH("/jobs/:id", s.patchJob)
	router.HEAD("/jobs/:id", s.headJob)
	router.POST("/jobs/opensearch/create-index/:index", s.createIndex)
	router.POST("/jobs/opensearch/push-documents/:index", s.pushDocuments)
	router.GET("/ping", s.ping)
	return router.Run(fmt.Sprintf(":%v", port))
}
