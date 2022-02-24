/*
 * @Descripttion: sys_configs 缓存
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-16 16:15:37
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-24 16:51:46
 */

package cache

import (
	"JHE_admin/table"
	"context"
	"fmt"
	"strconv"

	"go.uber.org/multierr"
	"gorm.io/gorm"
)

type CacheSysConfig struct {
	isLoad    bool
	parameter table.SysParameter
}

// 加载参数
func (l *CacheSysConfig) Load(db *gorm.DB) error {
	tx := db.WithContext(context.Background())
	params := []table.SysConfig{}
	err := tx.Select("param,value").Find(&params).Error
	if err != nil {
		return err
	}
	err = l.update(db, params, false)
	if err != nil {
		return err
	}
	l.isLoad = true
	return nil
}

func (l *CacheSysConfig) update(db *gorm.DB, params []table.SysConfig, isSave bool) error {
	var fvalue float64
	var ivalue int64
	var err error
	var finalErr error
	for i := 0; i < len(params); i++ {
		item := &params[i]
		err = nil
		switch item.Param {
		case "eth_sign_verify_url":
			l.parameter.EthSignVerifyURL = item.Value
		case "recharge_fee":
			fvalue, err = strconv.ParseFloat(item.Value, 32)
			l.parameter.RechargeFee = float32(fvalue)
		case "wallet_appid":
			ivalue, err = strconv.ParseInt(item.Value, 10, 32)
			l.parameter.WalletAppId = int(ivalue)
		case "wallet_appkey":
			l.parameter.WalletAppKey = item.Value
		case "wallet_url":
			l.parameter.WalletURL = item.Value
		case "wallet_withdraw_address":
			l.parameter.WalletWithdrawAddress = item.Value
		case "withdraw_fee":
			fvalue, err = strconv.ParseFloat(item.Value, 32)
			l.parameter.WithdrawFee = float32(fvalue)
		case "donate_exp_per_day":
			ivalue, err = strconv.ParseInt(item.Value, 10, 32)
			l.parameter.DonateExp = int(ivalue)
		case "donate_exp_name":
			l.parameter.DonateExpName = item.Value
		case "buy_vip_cost":
			fvalue, err = strconv.ParseFloat(item.Value, 32)
			l.parameter.BuyVipCost = float32(fvalue)
		case "buy_vip_symbol":
			l.parameter.BuyVipSymbol = item.Value
		case "game_cost":
			ivalue, err = strconv.ParseInt(item.Value, 10, 32)
			l.parameter.GameCost = int(ivalue)
		case "game_fee":
			fvalue, err = strconv.ParseFloat(item.Value, 32)
			l.parameter.GameFee = float32(fvalue)
		}

		if err != nil {
			if finalErr == nil {
				finalErr = err
			} else {
				finalErr = multierr.Append(finalErr, err)
			}
		}
	}
	if !isSave {
		return finalErr
	}
	db.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < len(params); i++ {
			item := &params[i]
			merr := tx.Exec(fmt.Sprintf("update sys_configs set value='%s' where param='%s'", item.Value, item.Param)).Error
			if merr != nil {
				if finalErr == nil {
					finalErr = merr
				} else {
					finalErr = multierr.Append(finalErr, merr)
				}
			}
		}
		return nil
	})
	return finalErr
}

// 保存参数
func (l *CacheSysConfig) Update(db *gorm.DB, params []table.SysConfig) error {
	return l.update(db, params, true)
}

func (l *CacheSysConfig) GetSysParameter() *table.SysParameter {
	if !l.isLoad {
		return nil
	}
	return &l.parameter
}
