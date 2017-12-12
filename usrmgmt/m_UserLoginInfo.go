package usrmgmt

import (
	"fmt"

	"github.com/noypi/util"
)

type UserLoginInfo struct {
	ID           string
	PasswordHash string
	Salt         []byte
}

func (usr *UserLoginInfo) ValidatePassword(pass string) (bValid bool) {
	return getPassHash(pass, usr.Salt) == usr.PasswordHash
}

func (usr *UserLoginInfo) UpdatePassword(pass string) {
	usr.Salt = util.GenSalt(10)
	usr.PasswordHash = getPassHash(pass, usr.Salt)
}

func getPassHash(pass string, salt []byte) (hash string) {
	bb, _ := util.SaltedHash256([]byte(pass), salt)
	return fmt.Sprintf("%x", bb)
}
