package usr

import (
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/cloud-simple-demo/internal/service/biz"
	"github.com/golang-acexy/cloud-web/webcloud"
	"github.com/golang-acexy/starter-gin/ginstarter"
)

var studentRouter = func() *StudentRouter[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO] {
	var bizService = biz.NewStudentBizService()

	return &StudentRouter[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO]{
		BaseRouter: webcloud.NewBaseRouterWithAuthority[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO](bizService, biz.UsrAuthorityFetch, "user_id"),
		bizService: bizService,
	}
}()

func NewUsrUserRouter() *StudentRouter[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO] {
	return studentRouter
}

type StudentRouter[ID webcloud.IDType, S, M, Q, D any] struct {
	*webcloud.BaseRouter[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO]
	bizService webcloud.BaseBizService[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO]
}

func (u *StudentRouter[ID, S, M, Q, T]) Info() *ginstarter.RouterInfo {
	return &ginstarter.RouterInfo{
		GroupPath: "usr/student",
	}
}

func (u *StudentRouter[ID, S, M, Q, T]) registerBaseHandler(router *ginstarter.RouterWrapper) {
	u.BaseRouter.RegisterBaseHandler(router, u.BaseRouter)
}

func (u *StudentRouter[ID, S, M, Q, T]) Handlers(router *ginstarter.RouterWrapper) {
	// 注册基础路由
	u.registerBaseHandler(router)

	// 自定义实现业务
	router.GET("test", u.test())
}

// 自定义实现业务

func (*StudentRouter[ID, S, M, Q, T]) test() ginstarter.HandlerWrapper {
	return func(request *ginstarter.Request) (ginstarter.Response, error) {
		return ginstarter.RespRestSuccess(), nil
	}
}
