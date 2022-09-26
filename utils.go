package grace

func Ptr[T any](v T) *T {
	return &v
}

func Repeat[T any](count int, v T) []T {
	data := make([]T, count)
	for i := 0; i < count; i++ {
		data[i] = v
	}
	return data
}

func MapKeys[T comparable, U any](m map[T]U) []T {
	data := make([]T, len(m))
	var i int
	for k := range m {
		data[i] = k
		i++
	}
	return data
}

func MapVals[T comparable, U any](m map[T]U) []U {
	data := make([]U, len(m))
	var i int
	for _, v := range m {
		data[i] = v
		i++
	}
	return data
}

func SliceToMap[T any, U comparable](s []T, fn func(v T) U) map[U]T {
	m := make(map[U]T, len(s))
	for _, v := range s {
		m[fn(v)] = v
	}
	return m
}
