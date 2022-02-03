package env

import (
	"coder.byzk.cn/golibs/common/logs"
	"os"
	"strconv"
	"strings"
)

type GetFn func() string
type GetIntFn func() int

// GetEnvOrDefault 获取环境变量或默认值
func GetEnvOrDefault(name, defaultValue string, callback ...func(envStr string) string) GetFn {
	envVal := os.Getenv(name)
	if envVal == "" {
		envVal = defaultValue
	}

	for _, c := range callback{
		envVal = c(envVal)
	}

	logs.Debugf("获取环境变量[%s], 对应的值为: %s", name, envVal)
	return func() string {
		return envVal
	}
}

// GetEnvReplaceNewlineChar 获取环境变量值内的\n进行真实换行替换
func GetEnvReplaceNewlineChar(name, defaultValue string, callback ...func(envStr string) string) GetFn {
	val := GetEnvOrDefault(name, defaultValue, callback...)()
	newVal := strings.ReplaceAll(val, "\\n", "\n")
	return func() string {
		return newVal
	}
}

// GetIntEnvOrDefault 获取int类型的环境变量
func GetIntEnvOrDefault(name string, defaultVal int, callback ...func(envInt int) int) GetIntFn {
	res := defaultVal
	val := os.Getenv(name)
	if val != "" {
		r, err := strconv.Atoi(val)
		if err == nil {
			res = r
		}
	}
	for _, c := range callback {
		res = c(res)
	}
	return func() int {
		return res
	}
}
