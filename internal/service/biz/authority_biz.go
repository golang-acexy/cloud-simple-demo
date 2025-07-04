package biz

import (
	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/cloud-web/webcloud"
	"github.com/golang-acexy/starter-gin/ginstarter"
)

var UsrAuthorityFetch webcloud.AuthorityFetch[int64] = func(request *ginstarter.Request) webcloud.Authority[int64] {
	return &model.UsrUserAuthority[int64]{
		ID: 123456,
	}
}
