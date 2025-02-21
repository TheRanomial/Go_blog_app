package internals

import (
	"context"
	"go-fullstack/views"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

const appTimeout=time.Second*10

func render(ctx *gin.Context,status int,template templ.Component)  error{
	ctx.Status(status)
	return template.Render(ctx.Request.Context(),ctx.Writer)
}

func (app *Config) IndexPageHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        _, cancel := context.WithTimeout(context.Background(), appTimeout)
        defer cancel()

        todos, err := app.getAllTodosService()
        if err != nil {
            ctx.JSON(http.StatusBadRequest, err.Error())
            return
        }

        var viewsTodos []*views.Todo
        for _, todo := range todos {
            viewsTodo := &views.Todo{
                Id:          todo.Id,
                Description: todo.Description,
            }
            viewsTodos = append(viewsTodos, viewsTodo)
        }

        render(ctx, http.StatusOK, views.Index(viewsTodos))
    }
}

func (app *Config) createTodoHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        _, cancel := context.WithTimeout(context.Background(), appTimeout)
        description := ctx.PostForm("description")
        defer cancel()

        newTodo := TodoRequest{
            Description: description,
        }

        data, err := app.CreateTodoService(&newTodo)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, err.Error())
            return
        }

        ctx.JSON(http.StatusCreated, data)
    }
}