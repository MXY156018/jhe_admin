package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"JHE_admin/utils"
	subType "JHE_admin/web/hall/types"
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tal-tech/go-zero/core/logx"
)

type FeedBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) FeedBackLogic {
	return FeedBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (f *FeedBackLogic) FeedBack(req types.FeedBack) (*types.Result, error) {
	req.CreateTime = time.Now()
	req.Status = "0"
	if err := global.GVA_DB.Model(&req).Create(&req).Error; err != nil {
		return &types.Result{
			Code: 7,
			Msg:  "反馈失败，请稍后再试",
		}, nil
	}
	return &types.Result{
		Code: 0,
		Msg:  "反馈成功",
	}, nil
}
func (f *FeedBackLogic) UserTree() (*types.Result, error) {

	treeList := model.GetUserTree(10001)

	return &types.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: treeList,
	}, nil
}
func (h *FeedBackLogic) UserLogin(req subType.HallUser) (*subType.LoginResp, error) {
	var count int64
	var user subType.User
	err := global.GVA_DB2.Table("user").Where("account = ? and password = ?", req.Account, utils.MD5V([]byte(req.PassWord))).Count(&count).Find(&user).Error
	if err != nil {
		return &subType.LoginResp{
			Code:    400,
			Message: "服务器内部错误",
		}, nil
	}
	if count == 0 {
		return &subType.LoginResp{
			Code:    400,
			Message: "账号或密码错误",
		}, nil
	}
	now := time.Now().Unix()
	accessExpire := h.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := h.getJwtToken(h.svcCtx.Config.Auth.AccessSecret, now, h.svcCtx.Config.Auth.AccessExpire, user.Uid)

	if err != nil {
		return nil, err
	}

	return &subType.LoginResp{
		Code:         200,
		Message:      "登陆成功",
		Id:           user.Uid,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
func (h *FeedBackLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
