/*
@Author :  
@Group : Rosefintech
@Description :
@Data: 2020/6/5 10:11
*/

package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	zcommon "rosefintech-rosl2/common"
)

func SendJSON(status int, success bool, data interface{}, c *gin.Context) {
	c.JSON(status, gin.H{
		"success": success,
		"data":    data,
	})
}

func route(engine *gin.Engine) {
	apiPrefix := "zk"
	app := engine.Group(apiPrefix)
	{
		app.POST("/verify", Verify)
	}
}

// new create a http server
func NewService(servicePort string) *http.Server {
	if true == zcommon.Conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	route(router)
	server := &http.Server{
		Addr:    servicePort,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return server
}
