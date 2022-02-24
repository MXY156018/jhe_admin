/*
 * @Descripttion: users表定义
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 18:30:36
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-16 21:11:20
 */

package table

type User struct {
	// 用户UID
	UID int `gorm:"primaryKey;autoIncrement" json:"uid"`
	// 头像 编号
	Head int8 `json:"head"`
	// 上级
	Parent int `json:"parent"`
	//账户
	Account string `json:"account"`
	// 密码
	Password string `json:"-"`
	// 地址
	Address string `json:"address"`
	// 注册时间
	RegisterTime string `json:"registerTime"`
	// 上次登录IP
	LastLoginIp string `json:"lastLoginIp"`
	//是否是机器人
	IsBot int8 `json:"isBot"`
	//状态
	Status int8 `json:"status"`
	//账户类型
	Type int8 `json:"type"`
}
