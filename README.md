# gw2api

##[Guild Wars 2](https://www.guildwars2.com/en-gb/) - [API](http://wiki.guildwars2.com/wiki/API:Main) bindings for [Go](http://golang.org/).
[![GoDoc](https://godoc.org/github.com/yasvisu/gw2api?status.png)](https://godoc.org/github.com/yasvisu/gw2api)

Note:

ArenaNet has not finished development on [V2](http://wiki.guildwars2.com/wiki/API:2) yet.

## Requirements

* Developed with [Go](http://golang.org/) 1.3.3. 

## Installation
To get the gw2api package, simply hit in your console of choice:

    go get github.com/yasvisu/gw2api

Make sure you have [Go installed](http://golang.org/doc/install) and set up!

In your code, with your import path in mind:

    import "github.com/yasvisu/gw2api"

## Documentation
Documentation and examples hosted in [GoDoc](http://godoc.org/github.com/yasvisu/gw2api).


## Development state
###[V2](http://wiki.guildwars2.com/wiki/API:2)

* Polishing and debugging ongoing by default
* [x] Account
  * [x] Bank (Auth)
  * [x] Dyes (Auth)
  * [x] Materials (Auth)
  * [x] Skins (Auth)
  * [x] Characters (Auth)
* [x] Commerce
  * [x] Listings
  * [x] Exchange
  * [x] Prices
  * [x] Transactions (Auth)
* [x] PvP
  * [x] Stats (Auth)
  * [x] Games (Auth)
* [x] Items
  * [x] Recipes
    * [x] Search
  * [x] Items
  * [x] Skins
* [x] World vs World
  * [x] Matches
  * [x] Objectives
* [x] Game Mechanics
  * [x] Traits
  * [x] Specializations
* [ ] Map Information
  * [ ] Continents
  * [x] Maps
* [x] Misc
  * [x] Achievements
  * [x] Build
  * [x] Colors
  * [x] Currencies
  * [x] Quaggans
  * [x] Worlds


###[V1](http://wiki.guildwars2.com/wiki/API:1)

* Development stopped
  * I am stopping active development on the V1 of the API. If there is interest, I will continue along with my "not done" list. Otherwise, I will not be actively working on V1 unless requested.
* Done:
  * All endpoints wrapped
  * Basic tests for all endpoints
  * Documented code
  * [MIT license-d](https://github.com/yasvisu/gw2api/blob/master/LICENSE)
* Not done:
  * Make more thorough tests
   * Check for deep struct Unmarshalling
  * Write usage examples
  * Get feedback
