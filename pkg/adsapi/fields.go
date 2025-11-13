package adsapi

import "time"

// Field辅助函数

// String 创建字符串字段
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Int 创建整数字段
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Int64 创建int64字段
func Int64(key string, value int64) Field {
	return Field{Key: key, Value: value}
}

// Duration 创建时间段字段
func Duration(key string, value time.Duration) Field {
	return Field{Key: key, Value: value}
}

// Error 创建错误字段
func Error(err error) Field {
	return Field{Key: "error", Value: err}
}

// Bool 创建布尔字段
func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

// Float64 创建浮点数字段
func Float64(key string, value float64) Field {
	return Field{Key: key, Value: value}
}
