package internal

import (
	"ems/internal/controller"
	"github.com/kataras/iris/v12"
)

func Init() error {
	app := iris.New()
	controller.Init(app)
	return app.Listen(":8000")
}
