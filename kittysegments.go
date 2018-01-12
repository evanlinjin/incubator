package incubator

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

const (
	ImageMaxSize = 1200
)

var PartsOrder = [...]PartName{
	PartTail,
	PartBody,
	PartEars,
	PartHead,
	PartEyes,
	PartBrows,
	PartNose,
	PartCap,
}

type KittySegments struct {
	imagesDir string
	config    *KittyConfig
	layers    []image.Image
}

func NewKittySegments(imagesDir string, config *KittyConfig) (*KittySegments, error) {
	ks := &KittySegments{
		imagesDir: imagesDir,
		config:    config,
	}
	for _, p := range PartsOrder {
		if e := ks.addLayer(p); e != nil {
			return nil, e
		}
	}
	return ks, nil
}

func (ks *KittySegments) Compile() (image.Image, error) {
	out := image.NewRGBA(image.Rect(0, 0, ImageMaxSize, ImageMaxSize))

	for _, img := range ks.layers {
		draw.Draw(out, out.Bounds(), img, img.Bounds().Min, draw.Over)
	}

	return out, nil
}

func (ks *KittySegments) CompileToFile(name string) error {
	img, e := ks.Compile()
	if e != nil {
		return e
	}
	f, e := os.Create(name)
	if e != nil {
		return e
	}
	return png.Encode(f, img)
}

func (ks *KittySegments) addLayer(part PartName) error {
	path, ok := ks.config.ImagePath(ks.imagesDir, part)
	if !ok {
		return nil
	}
	f, e := os.Open(path)
	if e != nil {
		return errors.New(fmt.Sprintf("failed to open file for kitty part: %v", e))
	}
	img, e := png.Decode(f)
	if e != nil {
		return errors.New(fmt.Sprintf("failed to decode image for kitty part: %v", e))
	}
	if size := img.Bounds().Size(); size.X != ImageMaxSize || size.Y != ImageMaxSize {
		return errors.New(fmt.Sprintf("kitty part image has invalid dimensions: %s", path))
	}
	ks.layers = append(ks.layers, img)
	return nil
}
