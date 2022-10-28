package env

import "os"

const (
	Product = "product" // 生产环境
	Gray    = "gray"    // 灰度环境
	Test    = "test"    // 测试环境
	Develop = "develop" // 开发环境
)

const (
	APP_ENV_KEY = "GOWINS_RUN_ENV"
)

var (
	APP_ENV string
)

func GetRunEnv() string {
	return APP_ENV
}

func Setup() {
	if APP_ENV = os.Getenv(APP_ENV_KEY); len(APP_ENV) == 0 {
		APP_ENV = Develop
	} else {
		switch APP_ENV {
		case Product, Gray, Test, Develop:
		default:
			panic("unknow env " + APP_ENV)
		}
	}
}

func IsDev() bool {
	return APP_ENV == Develop
}

func IsTest() bool {
	return APP_ENV == Test
}

func IsProduct() bool {
	return APP_ENV == Product
}
