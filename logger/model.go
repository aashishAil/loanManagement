package logger

import "go.uber.org/zap"

type Field = zap.Field

func String(key string, val string) Field {
	return zap.String(key, val)
}

func Int(key string, val int) Field {
	return zap.Int(key, val)
}

func Int64(key string, val int64) Field {
	return zap.Int64(key, val)
}

func Float64(key string, val float64) Field {
	return zap.Float64(key, val)
}

func Bool(key string, val bool) Field {
	return zap.Bool(key, val)
}

func Error(error error) Field {
	return zap.Error(error)
}

func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}
