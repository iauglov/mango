package mango

func Map[T, R any](collection []T, function func(T) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = function(item)
	}
	return result
}

func MapWithErr[T, R any](collection []T, function func(T) (R, error)) ([]R, error) {
	result := make([]R, len(collection))
	for i, item := range collection {
		tmp, err := function(item)
		if err != nil {
			return nil, err
		}
		result[i] = tmp
	}
	return result, nil
}
