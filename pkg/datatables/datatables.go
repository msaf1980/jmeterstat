package datatables

// Responce is a struct for JSON serialization
type Responce struct {
	Draw            int        `json:"draw"`
	RecordsTotal    int        `json:"recordsTotal"`
	RecordsFiltered int        `json:"recordsFiltered"`
	Data            [][]string `json:"data"`
}

// ResponceError is a struct for JSON serialization when error occured
type ResponceError struct {
	Draw  int    `json:"draw"`
	Error string `json:"error"`
}
