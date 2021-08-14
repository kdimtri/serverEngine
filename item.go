package main

import (
	"image"
	"image/draw"
)

type Item struct {
	id   int
	rect image.Rectangle
}

// NewItem returns new Item
func NewItem(id int, width int, height int) *Item {
	return &Item{
		id:   id,
		rect: image.Rect(0, 0, width, height),
	}
}

// Draw impliments draw.Drawer interface
func (it *Item) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	it.rect = r
	draw.Draw(dst, it.rect, src, sp, draw.Src)
}
