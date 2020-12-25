package models

type Course struct {
	Name          string `json:"name"`
	QuantityPlace int    `json:"name"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"EndDate"`
	CreatedDate   string `json:"CreatedDate"`
	// Student       []Students `json:"students"`
}
