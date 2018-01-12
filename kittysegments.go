package incubator

import (
	"image"
	"os"
	"errors"
	"fmt"
	"image/png"
	"image/draw"
)

const (
	ImageMaxSize = 1200
)

type KittySegments struct {
	imagesDir string
	config *KittyConfig
	layers []image.Image
}

func NewKittySegments(imagesDir string, config *KittyConfig) (*KittySegments, error) {
	ks := &KittySegments{
		imagesDir: imagesDir,
		config:   config,
	}
	if e := ks.getImage(PartTail); e != nil {
		return nil, e
	}
	if e := ks.getImage(PartBody); e != nil {
		return nil, e
	}
	if e := ks.getImage(PartEars); e != nil {
		return nil, e
	}
	if e := ks.getImage(PartHead); e != nil {
		return nil, e
	}
	if e := ks.getImage(PartEyes); e != nil {
		return nil, e
	}
	if e := ks.getImage(PartBrows); e != nil && !config.HasBrows {
		return nil, e
	}
	if e := ks.getImage(PartNose); e != nil {
		return nil, e
	}
	if e := ks.getImage(PartCap); e != nil && !config.HasCap {
		return nil, e
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

func (ks *KittySegments) getImage(part PartName) error {
	path := ks.config.ImagePath(ks.imagesDir, part)
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