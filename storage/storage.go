package storage

type IStorage interface {
	Save(string, string) error
	Load(string) (string, error)
}
