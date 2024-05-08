package common

func GetPtr[V any](v V) *V {
	return &v
}
