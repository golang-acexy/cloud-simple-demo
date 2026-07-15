package biz

import (
	"strconv"

	"github.com/golang-acexy/cloud-simple-demo/internal/model"
	"github.com/golang-acexy/cloud-web/webcloud"
	"github.com/golang-acexy/starter-gin/ginstarter"
)

var UsrAuthorityFetch webcloud.AuthorityFetch[int64] = func(request *ginstarter.Request) webcloud.Authority[int64] {
	userID, err := strconv.ParseInt(request.GetHeader("X-User-ID"), 10, 64)
	if err != nil || userID <= 0 {
		return nil
	}
	return &model.UsrUserAuthority{ID: userID}
}
