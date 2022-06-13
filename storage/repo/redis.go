package repo

type RedisRepoStorage interface {
	Get(key string) (interface{}, error)
	Set(key, value string) error
	SetWithTTL(key, value string, sec int64) error
}