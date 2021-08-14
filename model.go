package main

import (
	"fmt"
	"image"
	"image/draw"
	"net/http"

	"github.com/kdimtri/serverEngine/engine"
)

type Model struct {
	Name    string `json:"name"`
	Storage []Item `json:"storage"`
}

// ServeHTTP method impliments net/http.Handler interface
// to serve collage over network requests
func (m *Model) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := m.doSome(); err != nil {
		engine.Respond(w, r, http.StatusInternalServerError, err)
	} else {
		engine.Respond(w, r, http.StatusOK, m.respond())
	}
	return
}

func (m *Model) doSome() error {
	return nil
}
func NewModel() *Model {
	m := &Model{}
	m.Name = "ModelName"
	m.Storage = make([]Item, 0, 1)
	return m
}

// Draw impliments image/draw.Drawer interface
// Concurently draws items in storage
func (m *Model) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	b := src.Bounds()
	ch := make(chan int)
	for i := 0; i < m.len(); i++ {
		go func(item Item, done chan int, dstc draw.Image, rcc image.Rectangle, srcc image.Image, spc image.Point) {
			item.Draw(dstc, rcc, srcc, spc)
			done <- item.id
		}(m.Storage[i], ch, dst, m.Storage[i].rect, src, sp)
	}
	i := 0
	for range ch {
		i++
		if i >= len(m.Storage) {
			close(ch)
			return
		}
	}
	return
}

type modelRespond struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func (m *Model) respond() ModelRespond {
	return ModelRespond{"ok", m.Name}
}
func (m *Model) newError(msg string) error {
	return fmt.Errorf("Model error: %v", msg)
}
