package usrmgmt

import (
	"fmt"
)

var (
	ErrNotLogin    = fmt.Errorf("not login.")
	ErrFailedLogin = fmt.Errorf("failed login.")
)
