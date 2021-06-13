package env

import "os"

// GetEnvOrDefault 获取环境变量或默认值
func GetEnvOrDefault(envName, defaultValue string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}
