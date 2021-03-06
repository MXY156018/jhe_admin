module JHE_admin

go 1.16

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.1.10+incompatible
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/casbin/casbin/v2 v2.37.4
	github.com/casbin/gorm-adapter/v3 v3.4.4
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.5.1
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/mojocn/base64Captcha v1.3.5
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/qiniu/api.v7/v7 v7.8.2
	github.com/robfig/cron v1.2.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.21.9+incompatible
	github.com/spf13/viper v1.9.0
	github.com/tal-tech/go-zero v1.2.4
	github.com/tencentyun/cos-go-sdk-v5 v0.7.31
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/xuri/excelize/v2 v2.4.1
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.17.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
)

replace github.com/casbin/casbin/v2 v2.37.4 => github.com/casbin/casbin/v2 v2.11.0

replace github.com/casbin/gorm-adapter/v3 v3.4.4 => github.com/casbin/gorm-adapter/v3 v3.0.2
