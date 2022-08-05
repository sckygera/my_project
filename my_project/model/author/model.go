package author

type Author struct {
	AuthorId int    `json:"id"`
	Name     string `json:"name"`
	Born     string `json:"born"`
	Death    string `json:"death"`
}
