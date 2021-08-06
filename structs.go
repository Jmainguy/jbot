package main

// QuoteRequest : A quote request
type QuoteRequest struct {
	Services  string `json:"services"`
	Address   string `json:"address"`
	Unit      string `json:"unit"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zipcode   string `json:"zipcode"`
	Prefix    string `json:"prefix"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Notes     string `json:"notes"`
}
