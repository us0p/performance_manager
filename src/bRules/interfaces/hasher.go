package interfaces

type Hasher interface {
	Hash(v string) (string, error)
}
