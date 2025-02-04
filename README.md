# Hello! – Simple event sourcing prototype with Go

A simple prototype of event sourcing architecture implemented in Go. The example models a pet store where we track various events related to pets - like their arrival at the store, being sold to customers, price changes, and more.

## Overview

The project demonstrates event sourcing concepts by tracking the lifecycle of pets in a store:

- Storing events as the source of truth
- Building current state from event history  
- Maintaining read models for efficient querying
- Using GORM and Postgres for persistence

## Example Domain

The pet store domain includes events like:

- `PetArrived` - A new pet arrives at the store
- `PetBorn` - A pet is born to a mother already in the store
- `PetSold` - A pet is sold to a customer
- `PetPriceChange` - The price of a pet is updated
- `PetLost` - A pet goes missing
- `PetFoundAgain` - A lost pet is found

> In our little pet store, so cozy and bright,
> We track all the moments, from morning till night.
> But one special event you won't find in our code,
> Is when pets cross that final, rainbow road.
> For this is a happy example, you see,
> Where pets live forever, as they should be.
> No PetDied events in our database here,
> Just joyful moments we hold so dear!
