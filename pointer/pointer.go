package pointer

func To[T any](val T) *T {
	return &val
}

func Of[T any](val *T) T {
	if val != nil {
		return *val
	}
	var def T
	return def
}
