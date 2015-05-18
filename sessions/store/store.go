package store

type Store interface {
	Get(key string) (string, error)
}
