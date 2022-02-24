/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 01:13:56
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 18:12:55
 */

package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/table"
	"JHE_admin/utils"
	subType "JHE_admin/web/hall/types"
	"context"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) AccountLogin(req *subType.AccountLoginReq, ip string) *subType.LoginResp {
	resp := &subType.LoginResp{
		Code:    200,
		Message: "登录成功",
	}
	var count int64
	var user table.User
	err := global.GVA_DB.Table("users").Where("account = ? and password = ?", req.Account, utils.MD5V([]byte(req.PassWord))).Count(&count).Find(&user).Error
	if err != nil {
		resp.Code = 400
		resp.Message = "服务器内部错误"
		return resp
	}
	if count == 0 {
		resp.Code = 400
		resp.Message = "账号或密码错误"
		return resp
	}
	onLoginSuccess(l.svcCtx, resp, &user)
	go afterLogin(int(user.UID), ip)
	return resp
}
