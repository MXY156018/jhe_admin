package base

import (
	"JHE_admin/global"
	"JHE_admin/internal/middleware"
	"JHE_admin/internal/svc"
	"JHE_admin/internal/types"
	"JHE_admin/model"
	"JHE_admin/utils"
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/tal-tech/go-zero/core/logx"
	"go.uber.org/zap"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *LoginLogic) Login(req types.Login) (*types.Result, error) {
	if store.Verify(req.CaptchaId, req.Captcha, true) {
		var user types.SysUser
		req.Password = utils.MD5V([]byte(req.Password))
		u := &types.SysUser{Username: req.Username, Password: req.Password}
		err := global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
			return &types.Result{
				Code: 7,
				Msg:  "用户名不存在或密码错误",
			}, nil
		} else {
			return l.tokenNext(user), nil
		}
	} else {
		return &types.Result{
			Msg: "验证码错误",
		}, nil
	}
}

func (l *LoginLogic) JsonInBlacklist(r *http.Request) (*types.Result, error) {
	token := r.Header.Get("x-token")
	jwt := types.JwtBlacklist{
		Jwt:       token,
		Status:    1,
		CreatedAt: time.Now(),
	}
	if err := model.JwtServiceApp.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		return &types.Result{Code: 7, Msg: "jwt作废失败"}, nil
	}
	return &types.Result{Code: 0, Msg: "jwt作废成功"}, nil
}

func (l *LoginLogic) tokenNext(user types.SysUser) *types.Result {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := types.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.GVA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "获取token失败",
		}
	}
	if !global.GVA_CONFIG.System.UseMultipoint {

		return &types.Result{
			Code: 0,
			Data: types.LoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			},
			Msg: "登陆成功",
		}
	}
	if err, jwtStr := model.JwtServiceApp.GetRedisJWT(user.Username); err == redis.Nil {
		if err = model.JwtServiceApp.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Any("err", err))
			return &types.Result{
				Code: 7,
				Msg:  "设置登录状态失败",
			}
		}
		return &types.Result{
			Code: 0,
			Data: types.LoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			},
			Msg: "登陆成功",
		}
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Any("err", err))
		return &types.Result{
			Code: 7,
			Msg:  "设置登录状态失败",
		}
	} else {
		var blackJWT types.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := model.JwtServiceApp.JsonInBlacklist(blackJWT); err != nil {
			return &types.Result{
				Code: 7,
				Msg:  "jwt作废失败",
			}
		}
		if err := model.JwtServiceApp.SetRedisJWT(token, user.Username); err != nil {

			return &types.Result{
				Code: 7,
				Msg:  "设置登录状态失败",
			}
		}
		return &types.Result{
			Code: 0,
			Data: types.LoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			},
			Msg: "登陆成功",
		}
	}

}
