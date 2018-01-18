package http

import (
	"github.com/astaxie/beego"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"time"
	"strings"
	"net/http"
	"github.com/2young2simple/YiJob/model"
)


func ServerStart() {
	router := gin.Default()
	router.Use(Cors())
	router.Use(StaticFile())

	runmode := beego.AppConfig.DefaultString("runmode","debug")
	if runmode == "debug"{
		ginpprof.Wrapper(router)
	}

	router.POST("/api/job", DoJob)

	router.Run(beego.AppConfig.String("ip") + ":" + beego.AppConfig.String("port"))
}

func Cors() gin.HandlerFunc {
	return cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	})
}

func StaticFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Index(c.Request.URL.Path, "api/") >= 0 || strings.Index(c.Request.URL.Path, "debug/") >= 0{
			c.Next()
			return
		}
		http.ServeFile(c.Writer, c.Request, "static/web/"+c.Request.URL.Path)
	}
}

func WriteJson(ctx *gin.Context, code int, message string, data interface{}) {
	result := model.Result{Code: code, Message: message, Data: data}
	ctx.JSON(200, result)
}
