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
		BaseRouter: webcloud.NewBaseRouterWithAuthority[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO](bizService, biz.UsrAuthorityFetch, "UserID"),
		bizService: bizService,
	}
}()

func NewStudentRouter() *StudentRouter[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO] {
	return studentRouter
}

type StudentRouter[ID webcloud.IDType, S, M, Q, D any] struct {
	*webcloud.BaseRouter[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO]
	bizService webcloud.BaseBizService[int64, model.StudentSDTO, model.StudentMDTO, model.StudentQDTO, model.StudentDTO]
}

func (u *StudentRouter[ID, S, M, Q, D]) Info() *ginstarter.RouterInfo {
	return &ginstarter.RouterInfo{
		GroupPath: "usr/student",
	}
}

func (u *StudentRouter[ID, S, M, Q, D]) Handlers(router *ginstarter.RouterWrapper) {
	// 自定义路由业务
}

// 自定义路由业务
