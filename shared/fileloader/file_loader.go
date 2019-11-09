package sharedfileloader

type FileLoader interface {
	Load(path string, filename string) []byte
}
