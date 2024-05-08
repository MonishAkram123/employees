package dto

type Employee struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
	Position string  `json:"position,omitempty"`
}

type EmployeeCreateUpdateRequest struct {
	Name     string  `json:"name"`
	Salary   float64 `json:"salary"`
	Position string  `json:"position"`
}
