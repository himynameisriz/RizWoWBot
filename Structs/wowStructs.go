package wowStructs

// Class : these are the wow classes
type Class struct {
	Id   int    `json:"id"`
	Mask int    `json:"mask"`
	Side string `json:"side"`
	Name string `json:"name"`
}
