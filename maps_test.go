package gw2api

import "testing"

func TestMaps(t *testing.T) {
	api := NewGW2Api()

	if maps, err := api.Maps(); err != nil {
		t.Error("Failed to fetch map ids: ", err)
	} else if len(maps) < 1 {
		t.Error("Fetched an unlikely number of map ids")
	}
	var continent int
	if conts, err := api.Continents(); err != nil {
		t.Error("Failed to fetch continents: ", err)
	} else {
		continent = conts[0]
	}
	if _, err := api.ContinentIds("en", continent); err != nil {
		t.Error("Failed to fetch continent details: ", err)
	}
	var floor int
	if floors, err := api.ContinentFloors(continent); err != nil {
		t.Error("Failed to fetch floor: ", err)
	} else {
		floor = floors[0]
	}
	var region int
	if regions, err := api.ContinentFloorRegions(continent, floor); err != nil {
		t.Error("Failed to fetch region: ", err)
	} else {
		region = regions[0]
	}
	var mapID int
	if maps, err := api.ContinentFloorRegionMaps(continent, floor, region); err != nil {
		t.Error("Failed to fetch maps: ", err)
	} else {
		mapID = maps[0]
	}
	if _, err := api.MapIds("en", mapID); err != nil {
		t.Error("Failed to fetch map details: ", err)
	}

	if sectors, err := api.ContinentFloorRegionMapSectors(continent, floor, region, mapID); err != nil || len(sectors) < 1 {
		t.Error("Failed to fetch sectors: ", err)
	}

	if pois, err := api.ContinentFloorRegionMapPois(continent, floor, region, mapID); err != nil || len(pois) < 1 {
		t.Error("Failed to fetch POIs: ", err)
	}

	if tasks, err := api.ContinentFloorRegionMapTasks(continent, floor, region, mapID); err != nil || len(tasks) < 1 {
		t.Error("Failed to fetch tasks: ", err)
	}
}
