package logic

import (
	"context"
	"gorm.io/gorm"
	"image/internal/models"
	"image/internal/svc"
	"image/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)



type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *RegisterLogic) Register(req types.UserOptReq) (*types.UserOptResp, error) {
	// todo: add your logic here and delete this line

	var usesModel models.User
	user := models.User{
		Mobile: req.Mobile,
		Passwd :req.Passwd,
	}

	if err := l.svcCtx.DbEngin.Where("mobile = ?", user.Mobile).First(&usesModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			l.svcCtx.DbEngin.Create(&user)
		}else{
			panic(err)
		}
	}else{
		user = usesModel
	}

	return &types.UserOptResp{
		Id: user.ID,
	},nil
}
