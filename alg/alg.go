package alg

type Alg interface {
	Get(key string) ([]byte, bool)
	Add(key string, value []byte)
	Delete(key string)
}
