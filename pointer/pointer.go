package pointer

func To[T any](val T) *T {
	return &val
}

func ValueOf[T any](val *T) T {
	if val != nil {
		return *val
	}
	var def T
	return def
}

func ValueOrDefault[T any](val *T, defValue T) T {
	if val != nil {
		return *val
	}
	return defValue
}
