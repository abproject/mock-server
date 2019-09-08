package file

type fileEntity struct {
	name string
	body []byte
}

type fileStorage struct {
	data map[string]*fileEntity
}
