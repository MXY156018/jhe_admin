package sysdictionaryhandler

import (
	"JHE_admin/internal/logic/sysdictionary"
	"JHE_admin/internal/types"
	"JHE_admin/utils"

	"github.com/tal-tech/go-zero/rest/httpx"

	"JHE_admin/internal/svc"
	"net/http"
)

func CreateSysDictionaryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionary
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionary.NewSysDictionaryLogic(r.Context(), ctx)
		resp, err := l.CreateSysDictionary(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func DeleteSysDictionaryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionary
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionary.NewSysDictionaryLogic(r.Context(), ctx)
		resp, err := l.DeleteSysDictionary(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func UpdateSysDictionaryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionary
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionary.NewSysDictionaryLogic(r.Context(), ctx)
		resp, err := l.UpdateSysDictionary(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func FindSysDictionaryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionary
		req.Type = r.URL.Query().Get("type")
		l := sysdictionary.NewSysDictionaryLogic(r.Context(), ctx)
		resp, err := l.FindSysDictionary(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
func GetSysDictionaryListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictionarySearch
		if err := utils.Bind(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := sysdictionary.NewSysDictionaryLogic(r.Context(), ctx)
		resp, err := l.GetSysDictionaryList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
