package rest

type restStorage struct {
	data   map[string]*entityRest
	global *entityRest
}

type entityRest struct {
	Config  EndpointRestDto
	created int64
}
