package model

import (
	"JHE_admin/global"
	"JHE_admin/internal/types"
	"JHE_admin/table"
	"fmt"

	"time"

	"github.com/araddon/dateparse"
)

func GetBill(req types.BillReq) (total int64, list interface{}, sum float64, err error) {
	if req.Type == 1 {
		total, list, sum, err := GetRecharge(req)
		return total, list, sum, err
	} else if req.Type == 2 {
		total, list, sum, err := GetWithdrawls(req)
		return total, list, sum, err
	} else if req.Type == 3 {
		total, list, sum, err := GetVipProfit(req)
		return total, list, sum, err
	}
	return total, list, sum, err
}

func GetRecharge(req types.BillReq) (total int64, list interface{}, sum float64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list2 []types.UserRecharge
	var list1 []types.UserRecharge
	db1 := global.GVA_DB.Model(&list1)
	db2 := global.GVA_DB.Model(&list2)

	if req.Uid != 0 {
		db1 = db1.Where("uid = ?", req.Uid)
		db2 = db2.Where("uid = ?", req.Uid)
	}
	if req.StartTime != "" && req.EndTime != "" {
		// start, err := dateparse.ParseAny(req.StartTime)
		end, err := dateparse.ParseAny(req.EndTime)
		if err != nil {
			return 0, list, 0, err
		}

		ad, _ := time.ParseDuration("24h")
		end = end.Add(ad)

		db1 = db1.Where("date >= ? AND date < ?", req.StartTime, end)
		db2 = db2.Where("date >= ? AND date < ?", req.StartTime, end)
	}

	if err := db2.Count(&total).Error; err != nil {
		return total, list, 0, err
	}
	var sum1 []float64

	err = db1.Limit(limit).Offset(offset).Order("date desc").Find(&list1).Error
	var sum2 = 0.00
	db2.Pluck("amount", &sum1).Scan(&list2)
	for i := 0; i < len(sum1); i++ {
		sum2 = sum2 + sum1[i]
	}
	sum = sum2

	return total, list1, sum, err
}
func GetWithdrawls(req types.BillReq) (total int64, list interface{}, sum float64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list2 []types.UserWithdrawl
	var list1 []types.UserWithdrawl
	db1 := global.GVA_DB.Model(&list1).Where("status = 2")
	db2 := global.GVA_DB.Model(&list2).Where("status = 2")

	if req.Uid != 0 {
		db1 = db1.Where("uid = ?", req.Uid)
		db2 = db2.Where("uid = ?", req.Uid)
	}
	if req.StartTime != "" && req.EndTime != "" {
		end, err := dateparse.ParseAny(req.EndTime)
		if err != nil {
			return 0, list, 0, err
		}
		ad, _ := time.ParseDuration("24h")
		end = end.Add(ad)
		db1 = db1.Where("finish_date >= ? AND finish_date < ?", req.StartTime, end)
		db2 = db2.Where("finish_date >= ? AND finish_date < ?", req.StartTime, end)
	}

	if err := db2.Count(&total).Error; err != nil {
		return total, list1, 0, err
	}
	var sum1 []float64

	err = db1.Limit(limit).Offset(offset).Order("finish_date desc").Find(&list1).Error
	var sum2 = 0.00
	db2.Pluck("amount", &sum1).Scan(&list2)
	for i := 0; i < len(sum1); i++ {
		sum2 = sum2 + sum1[i]
	}
	sum = sum2
	return total, list1, sum, err
}

func GetVipProfit(req types.BillReq) (total int64, list interface{}, sum float64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list2 []types.CustomerOperator
	var list1 []types.CustomerOperator
	db1 := global.GVA_DB.Model(&list1).Where("type = 4")
	db2 := global.GVA_DB.Model(&list2).Where("type = 4")

	if req.Uid != 0 {
		db1 = db1.Where("uid = ?", req.Uid)
		db2 = db2.Where("uid = ?", req.Uid)
	}
	if req.StartTime != "" && req.EndTime != "" {
		end, err := dateparse.ParseAny(req.EndTime)
		if err != nil {
			return 0, list, 0, err
		}
		ad, _ := time.ParseDuration("24h")
		end = end.Add(ad)
		db1 = db1.Where("create_time >= ? AND create_time < ?", req.StartTime, end)
		db2 = db2.Where("create_time >= ? AND create_time < ?", req.StartTime, end)
	}
	if err := db2.Count(&total).Error; err != nil {
		return total, list, 0, err
	}
	var sum1 []float64

	err = db1.Limit(limit).Offset(offset).Order("create_time desc").Find(&list1).Error
	var sum2 = 0.00
	db2.Pluck("num", &sum1).Scan(&list2)
	for i := 0; i < len(sum1); i++ {
		sum2 = sum2 + sum1[i]
	}
	sum = sum2
	return total, list1, sum, err
}

func GetGameBill(req types.BillReq) (total int64, list []table.GameCommission, sum float64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	var list2 []table.GameCommission
	db1 := global.GVA_DB.Model(&list).Select("uid,symbol,commission")
	db2 := global.GVA_DB.Model(&list2)

	if req.Uid != 0 {
		db1 = db1.Where("uid = ?", req.Uid)
		db2 = db2.Where("uid = ?", req.Uid)
	}

	if err := db2.Count(&total).Error; err != nil {
		return total, list, 0, err
	}
	var sum1 []float64

	err = db1.Limit(limit).Offset(offset).Find(&list).Error
	var sum2 = 0.00
	db2.Pluck("commission", &sum1).Scan(&list2)
	for i := 0; i < len(sum1); i++ {
		sum2 = sum2 + sum1[i]
	}
	sum = sum2
	return total, list, sum, err
}
func GetDailyBill() (recharge float64, reward float64, platform float64, err error) {
	var start_date = time.Now().Format("2006-01-02")
	start, _ := dateparse.ParseAny(start_date)
	ad, _ := time.ParseDuration("24h")
	end := start.Add(ad)

	// var count int64
	var recharge1 []types.UserRecharge
	sartStr := start.Format("2006-01-02 00:00:00")
	endStr := end.Format("2006-01-02 23:59:59")
	err = global.GVA_DB.Raw(fmt.Sprintf(
		"SELECT * FROM `user_recharges` WHERE `date` >= '%s' AND `date` < '%s'",
		sartStr, endStr,
	)).Scan(&recharge1).Error
	// err = db.Raw("").Model(&recharge1).Where("`date` >= ? AND `date` < ?", start, end).Count(&count).Scan(&recharge1).Error
	// fmt.Printf("count = %d\n%+v\n", count, recharge1)
	if err != nil {
		return recharge, reward, platform, err
	}
	var withdraw []types.UserWithdrawl
	err = global.GVA_DB.Model(&types.UserWithdrawl{}).Where("status = 2").Where("finish_date BETWEEN ? AND ?", start, end).Find(&withdraw).Error
	if err != nil {
		return recharge, reward, platform, err
	}
	rechargeFee := 0.00
	for k, _ := range recharge1 {
		recharge += recharge1[k].Amount
		rechargeFee += recharge1[k].Fee
	}
	rewardFee := 0.00
	for k, _ := range withdraw {
		reward += withdraw[k].Amount
		rewardFee += withdraw[k].Fee
	}
	var gameCount int64
	var gameSum float32

	err = global.GVA_DB.Table("game_record_g1").Where("win > 0").Where("create_time >= ? AND create_time < ?", start, end).Count(&gameCount).Error
	if err != nil {
		return recharge, reward, platform, err
	}
	param := global.GVA_CacheSysConfig.GetSysParameter()

	gameSum = float32(gameCount) * param.GameFee
	platform = float64(gameSum) + rechargeFee + rewardFee
	return recharge, reward, platform, err
}

func GetSumPlatformProfit() (sum float64, err error) {
	var (
		rechargeList []float64
		rewardList   []float64
		recharge     float64
		reward       float64
		platform     float64
	)
	err = global.GVA_DB.Model(&types.UserRecharge{}).Pluck("fee", &rechargeList).Error
	if err != nil {
		return sum, err
	}
	err = global.GVA_DB.Model(&types.UserWithdrawl{}).Where("status = 2").Pluck("fee", &rewardList).Error
	if err != nil {
		return sum, err
	}
	for _, v := range rechargeList {
		recharge = recharge + v
	}
	for _, v := range rewardList {
		reward = reward + v
	}
	var platformList []float64
	err = global.GVA_DB.Model(&types.GameRecordG1{}).Pluck("rebate", &platformList).Error
	if err != nil {
		return sum, err
	}
	for _, v := range platformList {
		platform = platform + v
	}
	platform = platform + recharge + reward
	return platform, err
}
