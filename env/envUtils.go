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

// GetIntEnvOrDefault 获取int类型的环境变量
func GetIntEnvOrDefault(name string, defaultVal int) GetIntFn {
	res := defaultVal
	val := os.Getenv(name)
	if val != "" {
		r, err := strconv.Atoi(val)
		if err == nil {
			res = r
		}
	}
	return func() int {
		return res
	}
}
