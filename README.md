# Hello Event Sourcing with Go

An examplatory prototype implementation of Event Sourcing in Go.

Inspired by [this article (german)](https://www.heise.de/blog/Event-Sourcing-Die-bessere-Art-zu-entwickeln-10258295.html) / [this video (german)](https://www.youtube.com/watch?v=ss9wnixCGRY) from Golo Roden.

## Key Concepts

- Events are stored in a single event table with UUIDv7 IDs for sequential ordering
- We use Go generics to define event types and their payloads

## Example Model

The example implementation models a simple pet store with the following events:

- `PetArrived` - A new pet arrives at the store
- `PetBorn` - A pet is born to a mother already in the store
- `PetSold` - A pet is sold to a customer
- `PetPriceChanged` - The price of a pet is updated
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

## Entity Materialization

The events are materialized using a simple CQRS pattern in a read database. The materialized entities provide information about:

- Basic details (name, species)
- Current price
- Status (available, sold, lost)
- Ownership information
- Audit fields (created/updated timestamps)
