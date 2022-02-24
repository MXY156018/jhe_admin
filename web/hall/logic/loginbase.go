/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 01:26:18
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 09:56:16
 */
package logic

import (
	"JHE_admin/global"
	"JHE_admin/internal/svc"
	"JHE_admin/table"
	"JHE_admin/utils"
	subType "JHE_admin/web/hall/types"

	"context"
	"fmt"
	"time"
)

// 登录成功
func onLoginSuccess(ctx *svc.ServiceContext, resp *subType.LoginResp, user *table.User) {
	now := time.Now().Unix()
	accessExpire := ctx.Config.Auth.AccessExpire
	jwtToken, err := utils.GetGameJwtToken(ctx.Config.Auth.AccessSecret, now, ctx.Config.Auth.AccessExpire, int64(user.UID))
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		return
	}

	resp.Data = user
	resp.AccessExpire = now + accessExpire
	resp.RefreshAfter = now + accessExpire/2
	resp.AccessToken = jwtToken
}

// 登录之后的处理逻辑
func afterLogin(uid int, ip string) {
	// 赠送 体验币
	donateExpCoin(uid)
	// 更新登录信息
	updateLoginInfo(uid, ip)
}

func donateExpCoin(uid int) {
	param := global.GVA_CacheSysConfig.GetSysParameter()
	if param == nil {
		return
	}
	if param.DonateExp <= 0 {
		return
	}
	if param.DonateExpName == "" {
		return
	}

	today, _ := utils.GetToday(time.Now(), true)
	users := []table.UserDailyDonate{}
	db := global.GVA_DB.WithContext(context.Background())
	err := db.Where("uid=?", uid).Limit(1).Find(&users).Error
	if err != nil {
		return
	}
	isDoate := false
	if len(users) == 0 {
		isDoate = true
	} else {
		user := &users[0]
		todayT := user.Date.Format("2006-01-02")
		if todayT != today {
			isDoate = true
		}
	}
	if !isDoate {
		return
	}
	db.Exec(fmt.Sprintf(
		"insert into user_daily_donates(uid, date,amount) values(%d,'%s',%d) on duplicate key update date='%s',amount=%d",
		uid, today, param.DonateExp, today, param.DonateExp,
	))
	db.Exec(fmt.Sprintf(
		"insert into wallets(uid,symbol,balance) values(%d,'%s',%d) on duplicate key update  balance=%d",
		uid, param.DonateExpName, param.DonateExp, param.DonateExp,
	))
}

// 更新登录信息
func updateLoginInfo(uid int, ip string) {
	db := global.GVA_DB.WithContext(context.Background())
	db.Exec(fmt.Sprintf("update users set last_login_time=CURRENT_TIMESTAMP,last_login_ip='%s' where uid=%d", ip, uid))
}
