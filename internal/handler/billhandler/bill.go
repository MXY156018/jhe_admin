package billhandler

import (
	"JHE_admin/internal/logic/bill"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/utils"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func BillHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BillReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := bill.NewBillLogic(r.Context(), ctx)
		resp, err := l.GetBillList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GameBillHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BillReq
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := bill.NewBillLogic(r.Context(), ctx)
		resp, err := l.GetGameBillList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DailyBillHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := bill.NewBillLogic(r.Context(), ctx)
		resp, err := l.GetDailyBill()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
