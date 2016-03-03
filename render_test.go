package gw2api

import (
	"image"
	"testing"
)

func TestTile(t *testing.T) {
	var err error
	api := NewGW2Api()
	var tile image.Image

	if _, err = api.Tile(1024, 1024, 1024, 1024, 1024); err == nil {
		t.Error("Failed to fetch error for not existant tile")
	}

	if tile, err = api.Tile(1, 1, 5, 15, 14); err != nil {
		t.Error("Failed to fetch Lions Arch tile: ", err)
	}
	if tile.Bounds().Max.X != 256 && tile.Bounds().Max.Y != 256 {
		t.Error("Fetched tile has incorrect size ", tile.Bounds().Max)
	}
}

func TestRender(t *testing.T) {
	var err error
	api := NewGW2Api()
	var render image.Image

	if _, err = api.Render("FOOBAR", 123); err == nil {
		t.Error("Failed to fetch error for non existant render")
	}

	if render, err = api.Render("943538394A94A491C8632FBEF6203C2013443555", 102478); err != nil {
		t.Error("Failed to fetch icon render for dungeon: ", err)
	}
	if render.Bounds().Max.X != 32 && render.Bounds().Max.Y != 32 {
		t.Error("Fetched icon has incorrect size ", render.Bounds().Max)
	}
}
