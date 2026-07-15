package model

import "github.com/golang-acexy/cloud-web/webcloud"

const (
	AppClient webcloud.Platform = "appClient"
)

type UsrUserAuthority struct {
	ID int64
}

func (u *UsrUserAuthority) GetIdentityID() int64 {
	return u.ID
}

func (u *UsrUserAuthority) GetPlatform() webcloud.Platform {
	return AppClient
}
