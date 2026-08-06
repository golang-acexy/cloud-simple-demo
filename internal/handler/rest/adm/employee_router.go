package adm

import (
	model "github.com/golang-acexy/cloud-simple-demo/internal/model"
	biz "github.com/golang-acexy/cloud-simple-demo/internal/service/biz"
	"github.com/golang-acexy/cloud-web/webcloud"
	"github.com/golang-acexy/starter-gin/ginstarter"
)

var employeeRouter = func() *EmployeeRouter {
	bizService := biz.NewEmployeeBizService()
	return &EmployeeRouter{
		BaseRouter: webcloud.NewBaseRouter[int64, model.EmployeeSDTO, model.EmployeeMDTO, model.EmployeeQDTO, model.EmployeeDTO](bizService),
		bizService: bizService,
	}
}()

func NewEmployeeRouter() *EmployeeRouter {
	return employeeRouter
}

type EmployeeRouter struct {
	*webcloud.BaseRouter[int64, model.EmployeeSDTO, model.EmployeeMDTO, model.EmployeeQDTO, model.EmployeeDTO]
	bizService *biz.EmployeeBizService
}

func (r *EmployeeRouter) Info() *ginstarter.RouterInfo {
	return &ginstarter.RouterInfo{
		GroupPath: "adm/employee",
	}
}

func (r *EmployeeRouter) Handlers(router *ginstarter.RouterWrapper) {
	r.RegisterBaseHandlers(router, r)
	// 在此处注册自定义业务路由。
}
