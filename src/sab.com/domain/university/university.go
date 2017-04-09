package university

type Language struct {
	Code string `json:"code" datastore:"code"`
	Name string `json:"name" datastore:"name"`
}

type Address struct {
	Line       string `json:"line" datastore:"line"`
	City       string `json:"city" datastore:"city"`
	State      string `json:"state" datastore:"state"`
	PostalCode string `json:"postalCode" datastore:"postalCode"`
}

type Tuition struct {
	Link   string `json:"link" datastore:"link"`
	Amount int    `json:"amount" datastore:"amount"`
}

type University struct {
	Id              int64      `json:"id" datastore:"-"`
	Name            string     `json:"name" datastore:"name"`
	Languages       []Language `json:"languages" datastore:"languages"`
	Website         string     `json:"website" datastore:"website"`
	ProgramListLink string     `json:"programListLink" datastore:"programListLink"`
	Address         Address    `json:"address" datastore:"address"`
	Tuition         Tuition    `json:"tuition" datastore:"tuition"`
}
