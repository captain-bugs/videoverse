package utils

func DotFind[In any](arr []In, predicate func(val In) bool) *In {
	for i := range arr {
		if predicate(arr[i]) {
			return &arr[i]
		}
	}
	return nil
}
