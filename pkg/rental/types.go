package rental

// Person defines a person
type Person struct {
	PW           string `json:"pw"`
	Name         string `json:"name"`
	Subscription string `json:"subscription"`
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

// Login contains the infos about the information
type Login struct {
	Correct bool `json:"login"`
}

// Reduction holds the information for the price reduction
type Reduction struct {
	Name   string `json:"name"`
	Clean  int    `json:"clean"`
	Fueled int    `json:"fueled"`
}
