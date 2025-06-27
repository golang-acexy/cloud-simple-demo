package model

import "github.com/golang-acexy/cloud-web/webcloud"

const (
	AppClient webcloud.Platform = "appClient"
)

type UsrUserAuthority[ID int64] struct {
	ID int64
}

func (u *UsrUserAuthority[ID]) GetIdentityID() int64 {
	return u.ID
}

func (u *UsrUserAuthority[ID]) GetPlatform() webcloud.Platform {
	return AppClient
}
