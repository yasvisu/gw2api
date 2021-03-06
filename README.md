# gw2api
[Guild Wars 2](https://www.guildwars2.com/en-gb/) - [API](http://wiki.guildwars2.com/wiki/API:Main) bindings for [Go](http://golang.org/).
[![GoDoc](https://godoc.org/github.com/yasvisu/gw2api?status.png)](https://godoc.org/github.com/yasvisu/gw2api)
[![Travis](https://travis-ci.org/lhw/gw2api.svg)](https://travis-ci.org/lhw/gw2api)
[![Coverage Status](https://coveralls.io/repos/lhw/gw2api/badge.svg?branch=master&service=github)](https://coveralls.io/github/lhw/gw2api?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/yasvisu/gw2api)](https://goreportcard.com/report/github.com/yasvisu/gw2api)


Note:

ArenaNet has not finished development on [V2](http://wiki.guildwars2.com/wiki/API:2) yet.

## Requirements

* Developed with [Go](http://golang.org/) 1.6
* CI tests done with Go 1.5 and 1.6

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
* [x] New tests for the updated API
* [x] Examples for the updated API
* [x] Endpoints
    * [x] Account
        * [x] Bank (Auth)
        * [x] Dyes (Auth)
        * [x] Materials (Auth)
        * [x] Skins (Auth)
        * [x] Characters (Auth)
        * [x] Shared Inventory (Auth)
    * [x] Achievements
      * [x] Achievements
      * [x] Achievements Daily
      * [x] Achievement Groups
      * [x] Achievement Categories
    * [x] Commerce
      * [x] Listings
      * [x] Exchange
      * [x] Prices
      * [x] Transactions (Auth)
    * [x] Guilds
      * [x] Guild Upgrades
      * [x] Guild Permissions
      * [x] Guild Members (Auth Guild Leader)
      * [x] Guild Ranks (Auth Guild Leader)
      * [x] Guild Stash (Auth Guild Leader)
      * [x] Guild Treasury (Auth Guild Leader)
      * [x] Log (Auth Guild Leader)
      * [x] Emblems
      *  [x] Teams
    * [x] PvP
      * [x] Stats (Auth)
      * [x] Games (Auth)
      * [x] Standings (Auth)
      * [x] Season
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
    * [x] Map Information
      * [x] Continents
      * [x] Maps
    * [x] Misc
      * [x] Build
      * [x] Colors
      * [x] Currencies
      * [x] Quaggans
      * [x] Worlds
      * [x] Minis
