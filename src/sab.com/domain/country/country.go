package country

type Country struct {
	Code string `json:"code" datastore:"code"`
	Name string `json:"name" datastore:"name"`
}
