package env

import (
	"coder.byzk.cn/golibs/common/logs"
	"os"
	"strings"
)

type GetFn func() string

// GetEnvOrDefault 获取环境变量或默认值
func GetEnvOrDefault(name, defaultValue string) GetFn {
	envVal := os.Getenv(name)
	if envVal == "" {
		envVal = defaultValue
	}
	logs.Debugf("获取环境变量[%s], 对应的值为: %s", name, envVal)
	return func() string {
		return envVal
	}
}

// GetEnvReplaceNewlineChar 获取环境变量值内的\n进行真实换行替换
func GetEnvReplaceNewlineChar(name, defaultValue string) GetFn {
	val := GetEnvOrDefault(name, defaultValue)()
	newVal := strings.ReplaceAll(val, "\\n", "\n")
	return func() string {
		return newVal
	}
}
