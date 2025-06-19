package adm

import (
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/cloud-simple-demo/internal/service/biz"
	"github.com/golang-acexy/cloud-web/webcloud"
	"github.com/golang-acexy/starter-gin/ginstarter"
)

var teacherRouter = func() *TeacherRouter[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO] {
	var bizService = biz.NewTeacherBizService()

	return &TeacherRouter[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO]{
		BaseRouter: webcloud.NewBaseRouter[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO](bizService),
		bizService: bizService,
	}
}()

func NewTeacherRouter() *TeacherRouter[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO] {
	return teacherRouter
}

type TeacherRouter[ID webcloud.IDType, S, M, Q, D any] struct {
	*webcloud.BaseRouter[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO]
	bizService webcloud.BaseBizService[int64, model.TeacherSDTO, model.TeacherMDTO, model.TeacherQDTO, model.TeacherDTO]
}

func (u *TeacherRouter[ID, S, M, Q, D]) Info() *ginstarter.RouterInfo {
	return &ginstarter.RouterInfo{
		GroupPath: "adm/teacher",
	}
}
func (u *TeacherRouter[ID, S, M, Q, D]) registerBaseHandler(router *ginstarter.RouterWrapper) {
	u.BaseRouter.RegisterBaseHandler(router, u.BaseRouter)
}

func (u *TeacherRouter[ID, S, M, Q, D]) Handlers(router *ginstarter.RouterWrapper) {
	// 注册基础路由
	u.registerBaseHandler(router)
	// 自定义路由业务
}

// 自定义路由业务
