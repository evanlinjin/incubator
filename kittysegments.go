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
	ImageMin = 0
	ImageMax = 1200
)

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
	return ks, RangeKittyParts(func(part KittyPart) error {
		return ks.addLayer(part)
	})
}

func (ks *KittySegments) Compile() (image.Image, error) {
	out := image.NewRGBA(image.Rect(ImageMin, ImageMin, ImageMax, ImageMax))

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

func (ks *KittySegments) addLayer(part KittyPart) error {
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
	if size := img.Bounds().Size(); size.X != ImageMax || size.Y != ImageMax {
		return errors.New(fmt.Sprintf("kitty part image has invalid dimensions: %s", path))
	}
	ks.layers = append(ks.layers, img)
	return nil
}
