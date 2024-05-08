package controller

import (
	"ems/internal/dto"
	"ems/internal/service"
	"github.com/kataras/iris/v12"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (ctrl *EmployeeController) Create(ctx iris.Context) {
	request := &dto.EmployeeCreateUpdateRequest{}
	if err := ctx.ReadJSON(request); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	employee, err := ctrl.service.Create(request)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	if err = ctx.JSON(employee); err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
}

func (ctrl *EmployeeController) Get(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	employee, err := ctrl.service.Get(id)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	if err = ctx.JSON(employee); err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
}

func (ctrl *EmployeeController) Update(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	request := &dto.EmployeeCreateUpdateRequest{}
	if err := ctx.ReadJSON(request); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	employee, err := ctrl.service.Update(id, request)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	if err = ctx.JSON(employee); err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

}

func (ctrl *EmployeeController) Delete(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	err = ctrl.service.Delete(id)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
}

func (ctrl *EmployeeController) List(ctx iris.Context) {
	pageNumber := ctx.URLParamIntDefault("pageNumber", 1)
	size := ctx.URLParamIntDefault("pageSize", 10)

	employees, err := ctrl.service.List(pageNumber, size)
	if err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}

	if err = ctx.JSON(employees); err != nil {
		ctx.StopWithError(iris.StatusInternalServerError, err)
		return
	}
}
