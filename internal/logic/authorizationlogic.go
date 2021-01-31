package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"image/internal/models"
	"image/internal/svc"
	"image/internal/types"
	"image/internal/utils"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type AuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthorizationLogic {
	return AuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthorizationLogic) Authorization(req types.UserOptReq) (*types.UserOptResp, error) {
	// todo: add your logic here and delete this line

	var user models.User
	if err := l.svcCtx.DbEngin.Where("mobile =? ", req.Mobile).First(&user).Error;err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, errorUnregisteredMobile
		}else{
			panic(err)
		}
	}
	// 验证密码
	if !utils.CheckPassword(req.Passwd,user.Passwd) {
		return nil, errorPassword
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, int64(user.ID))
	if err != nil {
		return nil, err
	}

	return &types.UserOptResp{
		Id:       user.ID,
		JwtToken: types.JwtToken{
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

func (l *AuthorizationLogic) getJwtToken(secretKey string,iat,seconds,userId int64) (string, error)  {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}