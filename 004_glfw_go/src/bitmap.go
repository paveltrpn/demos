package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

type bitmapType struct {
	width  int32
	height int32
	bpp    int32
	data   []uint8
}

func (bmp *bitmapType) loadFromJpeg(fname string) {
	reader, err := os.Open(fname)
	if err != nil {
		fmt.Println("loadFromJpeg(): error loading file!")
		log.Fatal(err)
	}
	defer reader.Close()

	img, err := jpeg.Decode(reader)
	if err != nil {
		fmt.Println("loadFromJpeg(): error decoding jpeg!")
		log.Fatal(err)
	}

	rgba := image.NewRGBA(img.Bounds())

	// draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	bmp.data = rgba.Pix
	bmp.width = int32(rgba.Rect.Size().X)
	bmp.height = int32(rgba.Rect.Size().Y)
	bmp.bpp = 32
}

func (bmp *bitmapType) loadFromPng(fname string) {
	reader, err := os.Open(fname)
	if err != nil {
		fmt.Println("loadFromPng(): error loading file!")
		log.Fatal(err)
	}
	defer reader.Close()

	img, err := png.Decode(reader)
	if err != nil {
		fmt.Println("loadFromPng(): error decoding png!")
		log.Fatal(err)
	}

	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	bmp.data = rgba.Pix
	bmp.width = int32(rgba.Rect.Size().X)
	bmp.height = int32(rgba.Rect.Size().Y)
	bmp.bpp = 32
}
