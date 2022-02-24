/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-14 19:15:03
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 00:42:24
 */
// 用户相关的接口

package logic

import (
	"JHE_admin/game/types"
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	mainTypes "JHE_admin/internal/types"
	"JHE_admin/utils"

	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

// 用户基本信息
type userInfo struct {
	UID    int  `gorm:"primarykey" json:"uid"`
	Head   int8 `json:"head"`
	Status int8 `json:"-"`
}

// 资产信息条目
type assetItem struct {
	// 币种
	Symbol string `json:"symbol"`
	// 余额
	Balance float32 `json:"balance"`
}

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLogic {
	return UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 根据 JWT token 获取用户信息
func (h *UserLogic) CheckJwt(req *types.CheckJwtReq) (*mainTypes.Result, error) {
	var result *mainTypes.Result = &mainTypes.Result{}
	if req.Token == "" {
		result.Code = 1
		result.Msg = "参数错误"
		return result, nil
	}

	token, err := utils.DecodeGameJwtToken(req.Token, h.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		result.Code = 2
		result.Msg = "token 参数错误"
		return result, nil
	}
	// 超时判断

	var users []userInfo
	db := global.GVA_DB.WithContext(context.Background())
	db = db.Table("users").Where("uid=?", token.UserID).Limit(1).Select("uid,head,status").Find(&users)
	if db.Error != nil {
		result.Code = 3
		result.Msg = db.Error.Error()
		h.Logger.Errorv(db.Error)
		return result, nil
	}
	if len(users) == 0 {
		result.Code = 4
		result.Msg = "用户不存在"
		return result, nil
	}
	info := &users[0]
	if info.Status == 0 {
		result.Code = 5
		result.Msg = "用户被冻结"
		return result, nil
	}
	result.Data = info

	return result, nil
}

// 根据 JWT token 获取用户信息
func (h *UserLogic) GetUserInfoByToken(req *types.GetUserInfoReq) (*mainTypes.Result, error) {
	var result *mainTypes.Result = &mainTypes.Result{}

	// 超时判断

	var users []userInfo
	db := global.GVA_DB.WithContext(context.Background())
	db = db.Table("users").Where("uid=?", req.UID).Limit(1).Select("uid,head,status").Find(&users)
	if db.Error != nil {
		result.Code = 3
		result.Msg = db.Error.Error()
		h.Logger.Errorv(db.Error)
		return result, nil
	}
	if len(users) == 0 {
		result.Code = 4
		result.Msg = "用户不存在"
		return result, nil
	}
	info := &users[0]
	if info.Status == 0 {
		result.Code = 5
		result.Msg = "用户被冻结"
		return result, nil
	}
	result.Data = info

	return result, nil
}

// 获取用户某个资产信息
func (h *UserLogic) GetUserAsset(req *types.GetUserAssetReq) (*mainTypes.Result, error) {
	var result *mainTypes.Result = &mainTypes.Result{}
	var assets []assetItem
	db := global.GVA_DB.WithContext(context.Background())
	db = db.Table("wallets").Where("uid=? and symbol=?", req.UID, req.Symbol).Limit(1).Select("symbol,balance").Find(&assets)
	if db.Error != nil {
		result.Code = 1
		result.Msg = db.Error.Error()
		h.Logger.Errorv(db.Error)
		return result, nil
	}
	if len(assets) == 0 {
		result.Code = 2
		result.Msg = "用户资产不存在"
		return result, nil
	}
	result.Data = &assets[0]
	return result, nil
}

// 获取用户所有资产
func (h *UserLogic) GetUserAssets(req *types.GetUserAssetsReq) (*mainTypes.Result, error) {
	var result *mainTypes.Result = &mainTypes.Result{}
	var assets []assetItem
	db := global.GVA_DB.WithContext(context.Background())
	db = db.Table("wallets").Where("uid=?", req.UID).Select("symbol,balance").Find(&assets)
	if db.Error != nil {
		result.Code = 1
		result.Msg = db.Error.Error()
		h.Logger.Errorv(db.Error)
		return result, nil
	}
	result.Data = &assets
	return result, nil
}
