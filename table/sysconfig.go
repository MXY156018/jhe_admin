/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 16:45:24
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 16:49:45
 */

package table

// sys_configs 表结构定义
type SysConfig struct {
	// 参数
	Param string `gorm:"primaryKey" json:"param"`
	// 参数值
	Value string `json:"value"`
	// 说明
	Remark string `json:"remark"`
}

func (l *SysConfig) TableName() string {
	return "sys_configs"
}

// 系统配置参数
type SysParameter struct {
	// 充值费率
	RechargeFee float32
	// 提币费率
	WithdrawFee float32
	// 钱包 AppId
	WalletAppId int
	// 钱包 AppKey
	WalletAppKey string
	// 钱包 提币地址
	WalletWithdrawAddress string
	// 钱包 API 地址
	WalletURL string
	// 钱包签名登录验证 URL
	EthSignVerifyURL string
	//每天赠送的体验币
	DonateExp int
	// 体验币是什么名称
	DonateExpName string
	//购买vip花费
	BuyVipCost float32
	//购买Vip花费币种
	BuyVipSymbol string
	//游戏门票
	GameCost int
	//游戏抽水
	GameFee float32
}
