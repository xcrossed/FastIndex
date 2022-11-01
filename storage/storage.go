package storage

// supply storage feature

type Storage interface {
	Add(key []byte) (bool, error)
	Delete(key []byte) (bool, error)
	Get(Key []byte) ([]byte, error)
	Set(key []byte) ([]byte, error)
}

type DefaultStorage struct {
}
