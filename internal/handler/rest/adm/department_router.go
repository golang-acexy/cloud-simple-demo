package adm

import (
	model "github.com/golang-acexy/cloud-simple-demo/internal/model"
	biz "github.com/golang-acexy/cloud-simple-demo/internal/service/biz"
	"github.com/golang-acexy/cloud-web/webcloud"
	"github.com/golang-acexy/starter-gin/ginstarter"
)

var departmentRouter = func() *DepartmentRouter {
	bizService := biz.NewDepartmentBizService()
	return &DepartmentRouter{
		BaseRouter: webcloud.NewBaseRouter[int64, model.DepartmentSDTO, model.DepartmentMDTO, model.DepartmentQDTO, model.DepartmentDTO](bizService),
		bizService: bizService,
	}
}()

func NewDepartmentRouter() *DepartmentRouter {
	return departmentRouter
}

type DepartmentRouter struct {
	*webcloud.BaseRouter[int64, model.DepartmentSDTO, model.DepartmentMDTO, model.DepartmentQDTO, model.DepartmentDTO]
	bizService *biz.DepartmentBizService
}

func (r *DepartmentRouter) Info() *ginstarter.RouterInfo {
	return &ginstarter.RouterInfo{
		GroupPath: "adm/department",
	}
}

func (r *DepartmentRouter) Handlers(router *ginstarter.RouterWrapper) {
	r.RegisterBaseHandlers(router, r)
	// 在此处注册自定义业务路由。
}
