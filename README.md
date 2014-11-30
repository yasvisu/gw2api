# gw2api

##[Guild Wars 2](https://www.guildwars2.com/en-gb/) - [API](http://wiki.guildwars2.com/wiki/API:Main) bindings for [Go](http://golang.org/).

The latest version of this package will be put here. 
Until [v2](http://wiki.guildwars2.com/wiki/API:2) is complete, 
the latest version will be a merged package of v1 and v2, 
with v2 functions replacing v1. For pure v1 and v2, check /v1 and /v2.

## Requirements

* Developed with Go 1.3.3. 

## Development state

###[V1](http://wiki.guildwars2.com/wiki/API:1)

* Development ongoing
* Done:
  * All endpoints wrapped - gw2api.go
  * Basic tests for all endpoints - gw2api_test.go
  * Documented code
  * [MIT license-d](https://github.com/yasvisu/gw2api/blob/master/LICENSE)
* To do:
  * Make more thorough tests
    * Check for deep struct Unmarshalling
  * Write usage examples
  * Get feedback

###[V2](http://wiki.guildwars2.com/wiki/API:2)

* Development ongoing
* Done:
* To do:
 * Wrap all endpoints
 * Write tests
 * Write usage samples
 * Get feedback
 * Merge with v1 for a common version
