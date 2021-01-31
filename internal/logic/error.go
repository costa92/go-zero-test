package logic

import  "image/internal/errors"

var (
	errorDuplicateUsername = errors.NewDefaultError("用户名已经注册")
	errorDuplicateMobile   = errors.NewDefaultError("手机号已经被占用")
	errorUnregisteredMobile   = errors.NewDefaultError("手机号未注册")
	errorPassword   = errors.NewDefaultError("错误的密码")
)