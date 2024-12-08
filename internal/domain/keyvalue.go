package domain

type KeyValueRepository interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	List() (map[string]string, error)
}
