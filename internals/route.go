package internals

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	
	app.Router.GET("/", app.IndexPageHandler())
	app.Router.POST("/", app.createTodoHandler())
}
