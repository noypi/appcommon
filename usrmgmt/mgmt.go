package usrmgmt

import (
	"context"
	"net/http"

	"github.com/noypi/router"
)

type UserMgmt struct {
	SignUpUserField string
	SignUpPassField string

	LoginUserField string
	LoginPassField string

	OnNotLoginRedirectPath     string
	OnInvalidLoginRedirectPath string
}

func (this *UserMgmt) OnNotLogin(ctx context.Context) {
	router.AbortError(ctx, http.StatusUnauthorized, ErrNotLogin)
	router.Redirect(ctx, http.StatusTemporaryRedirect, this.OnNotLoginRedirectPath)
}

func (this *UserMgmt) OnInvalidLogin(ctx context.Context) {
	router.AbortError(ctx, http.StatusUnauthorized, ErrFailedLogin)
	router.Redirect(ctx, http.StatusTemporaryRedirect, this.OnInvalidLoginRedirectPath)
}

func (this *UserMgmt) SignUp(ctx context.Context) {
}

func (this *UserMgmt) VerifyLogin(ctx context.Context) {
	c := router.ToStore(ctx)

	var oUserLoginInfo *UserLoginInfo = &UserLoginInfo{}

	if oUserLoginInfo.ValidatePassword(pass) {
		c.Set(UserLoginInfoKey, oUserLoginInfo)
	} else {
		this.OnInvalidLogin(ctx)
		return
	}
}

func (this *UserMgmt) EnsureLogin(ctx context.Context) {
	c := router.ToStore(ctx)
	if _, has := c.Get(UserKey); !has {
		this.OnNotLogin(ctx)
		return
	}
}

func (this *UserMgmt) GetUserLoginInfo(ctx context.Context) {
	_ = router.Param(ctx, this.LoginUserField)
	pass := router.Param(ctx, this.LoginPassField)

}
