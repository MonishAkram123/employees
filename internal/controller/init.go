package controller

import (
	"ems/internal/repository"
	"ems/internal/service"
	"github.com/kataras/iris/v12"
)

func Init(app *iris.Application) {
	employeeController := NewEmployeeController(service.NewEmployeeServiceImpl(repository.NewEmployeeInMemory()))
	employeeRoute := app.Party("/employees")

	employeeRoute.Get("", employeeController.List)
	employeeRoute.Post("", employeeController.Create)
	employeeRoute.Put("/{id:uint64}", employeeController.Update)
	employeeRoute.Delete("/{id:uint64}", employeeController.Delete)
	employeeRoute.Get("/{id:uint64}", employeeController.Get)

}
