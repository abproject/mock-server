package rest

type restStorage struct {
	data map[string]*entityRest
}

type entityRest struct {
	Config EndpointRestDto
}
