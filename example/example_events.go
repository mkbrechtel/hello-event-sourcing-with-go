package example

import (
	"hello-event-sourcing-with-go/event"
	"hello-event-sourcing-with-go/pet"

	"github.com/gofrs/uuid/v5"
)

func CreateSampleEvents() []any {
	// Generate UUIDs
	motherCatID, _ := uuid.NewV4()
	fluffyID, _ := uuid.NewV4()
	whiskersID, _ := uuid.NewV4()
	roverID, _ := uuid.NewV4()
	hammyID, _ := uuid.NewV4()
	bellaID, _ := uuid.NewV4()
	maxID, _ := uuid.NewV4()
	charlieID, _ := uuid.NewV4()
	axolotlID, _ := uuid.NewV4()
	octopusID, _ := uuid.NewV4()
	sandpuppyID, _ := uuid.NewV4()
	customer1ID, _ := uuid.NewV4()
	customer2ID, _ := uuid.NewV4()
	customer3ID, _ := uuid.NewV4()

	// Create array of events
	events := []any{
		// Mother cat arrives
		event.NewEvent(pet.PetArrived{
			PetID:   motherCatID,
			Species: "Cat",
			Name:    "Luna",
			Price:   150,
		}, "example", 1),

		// Fluffy arrives
		event.NewEvent(pet.PetArrived{
			PetID:   fluffyID,
			Species: "Cat",
			Name:    "Fluffy",
			Price:   100,
		}, "example", 1),

		// Whiskers is born
		event.NewEvent(pet.PetBorn{
			PetID:    whiskersID,
			Species:  "Cat",
			Name:     "Whiskers",
			Price:    80,
			MotherID: motherCatID,
		}, "example", 1),

		// Rover the dog arrives
		event.NewEvent(pet.PetArrived{
			PetID:   roverID,
			Species: "Dog",
			Name:    "Rover",
			Price:   200,
		}, "example", 1),

		// Hammy the hamster arrives
		event.NewEvent(pet.PetArrived{
			PetID:   hammyID,
			Species: "Hamster",
			Name:    "Hammy",
			Price:   25,
		}, "example", 1),

		// Bella the cat arrives
		event.NewEvent(pet.PetArrived{
			PetID:   bellaID,
			Species: "Cat",
			Name:    "Bella",
			Price:   130,
		}, "example", 1),

		// Max the dog arrives
		event.NewEvent(pet.PetArrived{
			PetID:   maxID,
			Species: "Dog",
			Name:    "Max",
			Price:   250,
		}, "example", 1),

		// Charlie the hamster arrives
		event.NewEvent(pet.PetArrived{
			PetID:   charlieID,
			Species: "Hamster",
			Name:    "Charlie",
			Price:   30,
		}, "example", 1),

		// Axolotl arrives
		event.NewEvent(pet.PetArrived{
			PetID:   axolotlID,
			Species: "Axolotl",
			Name:    "Axel",
			Price:   75,
		}, "example", 1),

		// Octopus arrives
		event.NewEvent(pet.PetArrived{
			PetID:   octopusID,
			Species: "Octopus",
			Name:    "Otto",
			Price:   200,
		}, "example", 1),

		// Sand puppy arrives
		event.NewEvent(pet.PetArrived{
			PetID:   sandpuppyID,
			Species: "Sand Puppy",
			Name:    "Sandy",
			Price:   150,
		}, "example", 1),

		// Price changes
		event.NewEvent(pet.PetPriceChanged{
			PetID:    fluffyID,
			OldPrice: 100,
			NewPrice: 120,
		}, "example", 1),

		event.NewEvent(pet.PetPriceChanged{
			PetID:    maxID,
			OldPrice: 250,
			NewPrice: 280,
		}, "example", 1),

		// Additional price changes
		event.NewEvent(pet.PetPriceChanged{
			PetID:    axolotlID,
			OldPrice: 75,
			NewPrice: 95,
		}, "example", 1),

		event.NewEvent(pet.PetPriceChanged{
			PetID:    octopusID,
			OldPrice: 200,
			NewPrice: 250,
		}, "example", 1),

		// Rover gets lost
		event.NewEvent(pet.PetLost{
			PetID: roverID,
			Date:  "2025-02-25",
		}, "example", 1),

		// Charlie gets lost
		event.NewEvent(pet.PetLost{
			PetID: charlieID,
			Date:  "2025-02-26",
		}, "example", 1),

		// Rover is found
		event.NewEvent(pet.PetFoundAgain{
			PetID:     roverID.String(),
			Date:      "2025-02-28",
			Condition: "Good",
		}, "example", 1),

		// Sales
		event.NewEvent(pet.PetSold{
			PetID:      fluffyID,
			CustomerID: customer1ID,
			SoldPrice:  120,
		}, "example", 1),

		event.NewEvent(pet.PetSold{
			PetID:      hammyID,
			CustomerID: customer2ID,
			SoldPrice:  25,
		}, "example", 1),

		event.NewEvent(pet.PetSold{
			PetID:      bellaID,
			CustomerID: customer3ID,
			SoldPrice:  130,
		}, "example", 1),

		event.NewEvent(pet.PetSold{
			PetID:      axolotlID,
			CustomerID: customer2ID,
			SoldPrice:  95,
		}, "example", 1),
	}

	return events
}
