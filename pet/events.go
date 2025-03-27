package pet

import (
	"github.com/gofrs/uuid/v5"
)

type PetArrived struct {
	PetID   uuid.UUID `json:"petId"`
	Species string    `json:"species"`
	Name    string    `json:"name"`
	Price   int       `json:"price"`
}

type PetBorn struct {
	PetID    uuid.UUID `json:"petId"`
	Species  string    `json:"species"`
	Name     string    `json:"name"`
	Price    int       `json:"price"`
	MotherID uuid.UUID `json:"motherId"`
}

type PetSold struct {
	PetID      uuid.UUID `json:"petId"`
	CustomerID uuid.UUID `json:"customerId"`
	SoldPrice  int       `json:"soldPrice"`
}

type PetPriceChanged struct {
	PetID    uuid.UUID `json:"petId"`
	OldPrice int       `json:"oldPrice"`
	NewPrice int       `json:"newPrice"`
	Reason   string    `json:"reason,omitempty"`
}

type PetLost struct {
	PetID uuid.UUID `json:"petId"`
	Date  string    `json:"date"`
}

type PetFoundAgain struct {
	PetID     string `json:"petId"`
	Date      string `json:"date"`
	Condition string `json:"condition,omitempty"`
}
