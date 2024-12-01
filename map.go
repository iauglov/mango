package mango

func Keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}
