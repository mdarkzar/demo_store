package restapi

import (
	"net/http"
	"store/tools/logger"
	"store/uimport"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RestAPI struct {
	log *logrus.Logger
	gin *gin.Engine
	uimport.UsecaseImports
}

func NewRestAPI(ui uimport.UsecaseImports,
	log *logrus.Logger,
) *RestAPI {
	api := &RestAPI{
		gin:            gin.Default(),
		UsecaseImports: ui,
		log:            logger.NewUsecaseLogger(log, "restapi"),
	}

	api.gin.Use(gzip.Gzip(gzip.DefaultCompression))

	apiGroup := api.gin.Group("/api/v1")

	u := apiGroup.Group("/user")
	u.POST("/auth", api.Auth)
	u.GET("/profile", api.AuthRequired(), api.Profile)
	u.POST("/logout", api.AuthRequired(), api.Logout)

	p := apiGroup.Group("/product")
	p.POST("/create", api.AuthRequired(), api.CreateProduct)
	p.POST("/remove", api.AuthRequired(), api.RemoveProduct)
	p.GET("/find/:id", api.AuthRequired(), api.FindProduct)
	p.GET("/load", api.AuthRequired(), api.LoadAllProduct)
	p.GET("/storage_list", api.AuthRequired(), api.LoadStorageList)

	n := apiGroup.Group("/notification")
	n.GET("/new", api.AuthRequired(), api.LoadMessages)

	return api
}

func (e *RestAPI) RunServer() {
	e.gin.Run(e.Config.ApiURL())
}

func (e *RestAPI) errorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}

func (e *RestAPI) returnResult(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// setTokenInCookie установка token в header ответа
func (e *RestAPI) setTokenInCookie(c *gin.Context, key, token string) {
	c.SetCookie(key, token, 60*60*24*365, "/", "", false, true)
}
