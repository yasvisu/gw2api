package gw2api

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

// Tile fetches a rendered JPEG tile for the game map on the requested
// continent/floor at zoom level and for the coordinates
func (gw2 *GW2Api) Tile(continent, floor, zoom, x, y int) (img image.Image, err error) {
	endpoint := fmt.Sprintf("https://tiles.guildwars2.com/%d/%d/%d/%d/%d.jpg", continent, floor, zoom, x, y)

	var body io.Reader
	if body, err = gw2.fetchRawEndpoint(endpoint); err != nil {
		return nil, fmt.Errorf("Failed to fetch tile at provided coordinates: %s", err)
	}
	return jpeg.Decode(body)
}

// Render fetches a pre-rendered PNG image for the requested ressource
func (gw2 *GW2Api) Render(signature string, fileID int) (img image.Image, err error) {
	endpoint := fmt.Sprintf("https://render.guildwars2.com/file/%s/%d.png", signature, fileID)

	var body io.Reader
	if body, err = gw2.fetchRawEndpoint(endpoint); err != nil {
		return nil, fmt.Errorf("Failed to fetch render target: %s", err)
	}
	return png.Decode(body)

}
