# Hello! – Simple event sourcing prototype with Go

A simple prototype of event sourcing architecture implemented in Go. The example models a pet store where we track various events related to pets - like their arrival at the store, being sold to customers, price changes, and more.

Inspired by [this article (german)](https://www.heise.de/blog/Event-Sourcing-Die-bessere-Art-zu-entwickeln-10258295.html) / [this video (german)](https://www.youtube.com/watch?v=ss9wnixCGRY) from Golo Roden.

## Overview

The project demonstrates event sourcing concepts by tracking the lifecycle of pets in a store:

- Storing events as the source of truth
- Building current state from event history  
- Maintaining read models for efficient querying
- Events are stored in a single event table with UUIDv7 IDs for sequential ordering
- We use Go generics to define event types and their payloads

## Example Domain Model

The example implementation models a simple pet store with the following events:


> In our little pet store, so cozy and bright,
> We track all the moments, from morning till night.
> But one special event you won't find in our code,
> Is when pets cross that final, rainbow road.
>
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
