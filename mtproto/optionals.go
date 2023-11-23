package mtproto

func Ref[T any](val T) *T {
	return &val
}

func DerefOr[T any](val *T, defaultVal T) T {
	if val == nil {
		return defaultVal
	}
	return *val
}
