package infra

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"github.com/gin-contrib/cors"
	"log"
	"net/http"
	"time"
)

type ServerHttp interface {
	StartServer(port string) error
	GetGinRouterGroup(relativePath string) *gin.RouterGroup
}

type serverHttp struct {
	ginEngine *gin.Engine
}

func NewServerHttp() ServerHttp{
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.MaxMultipartMemory = 8 << 20

	return &serverHttp{ginEngine: engine }
}

func (ref *serverHttp) StartServer(port string) error {
	var g errgroup.Group
	g.Go(func() error {
		address := fmt.Sprintf(":%s", port)
		err := http.ListenAndServe(address, ref.ginEngine)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (ref *serverHttp) GetGinRouterGroup(relativePath string) *gin.RouterGroup {
	ref.ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	return ref.ginEngine.Group(relativePath)
}