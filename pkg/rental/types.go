package rental

// Person defines a person
type Person struct {
	PW   string `json:"pw"`
	Name string `json:"name"`
}

// Rental specifies a car which can be rented
type Rental struct {
	Duration string `json:"duration"`
	Name     string `json:"name"`
	Car      Car    `json:"car"`
}

// Car defines a car
type Car struct {
	Rented bool   `json:"rented"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	ID     string `json:"id"`
}
