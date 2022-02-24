/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-22 16:49:24
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-22 16:50:43
 */
package types

type Config struct {
	BuyVipCost      string `json:"buy_vip_cost,omitempty"`
	BuyVipSymbol    string `json:"buy_vip_symbol,omitempty"`
	DonateExpName   string `json:"donate_exp_name,omitempty"`
	DonateExpPerDay string `json:"donate_exp_per_day,omitempty"`
	RechargeFee     string `json:"recharge_fee,omitempty"`
	WithdrawFee     string `json:"withdraw_fee,omitempty"`
}
