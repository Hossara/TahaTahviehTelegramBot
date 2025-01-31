package common

type Pagination[T interface{}] struct {
	Pages int
	Page  int
	Data  T
}
