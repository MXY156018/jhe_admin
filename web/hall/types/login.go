/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-18 01:15:00
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-18 01:17:04
 */
package types

type AccountLoginReq struct {
	Account  string `json:"account"`
	PassWord string `json:"passWord"`
}
