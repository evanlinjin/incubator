package incubator

import (
	"fmt"
	"path/filepath"
	"math/rand"
	"path"
)

type PartName string

const (
	PartBody  PartName = "body"
	PartBrows PartName = "brows"
	PartCap   PartName = "cap"
	PartEars  PartName = "ears"
	PartEyes  PartName = "eyes"
	PartHead  PartName = "head"
	PartNose  PartName = "nose"
	PartTail  PartName = "tail"
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

type KittyConfig struct {
	KittyID  uint64 `json:"kitty_id"`
	Version  uint64 `json:"version"`

	// -ve id represents that the kitty does not have that body part.
	BodyID   int64 `json:"body_id"`
	BrowsID  int64 `json:"brows_id"`
	CapID    int64 `json:"cap_id"`
	EarsID   int64 `json:"ears_id"`
	EyesID   int64 `json:"eyes_id"`
	HeadID   int64 `json:"head_id"`
	NoseID   int64 `json:"nose_id"`
	TailID   int64 `json:"tail_id"`
}

func (kc *KittyConfig) getPartIDPointer(name PartName) *int64 {
	var p *int64
	switch name {
	case PartBody:
		p = &kc.BodyID
	case PartBrows:
		p = &kc.BrowsID
	case PartCap:
		p = &kc.CapID
	case PartEars:
		p = &kc.EarsID
	case PartEyes:
		p = &kc.EyesID
	case PartHead:
		p = &kc.HeadID
	case PartNose:
		p = &kc.NoseID
	case PartTail:
		p = &kc.TailID
	}
	return p
}

func (kc *KittyConfig) ImagePath(rootPath string, partName PartName) (string, bool) {
	if partID := kc.getPartIDPointer(partName); partID == nil || *partID < 0  {
		return "", false
	} else {
		return filepath.Join(
			rootPath,
			fmt.Sprintf("Kitty_%d", kc.KittyID),
			string(partName),
			fmt.Sprintf("%d.png", *partID),
		), true
	}
}

func RandomKittyConfig(r *rand.Rand, rootPath string) (*KittyConfig, error) {
	kc := new(KittyConfig)

	kitties, e := filepath.Glob(path.Join(rootPath, "Kitty_*"))
	if e != nil {
		return nil, e
	}
	kc.KittyID = r.Uint64()%uint64(len(kitties))

	var kittyPath = filepath.Join(
		rootPath,
		fmt.Sprintf("Kitty_%d", kc.KittyID),
	)
	for _, part := range PartsOrder {
		list, _ := filepath.Glob(path.Join(kittyPath, string(part), "*.png"))
		if pID := kc.getPartIDPointer(part); pID != nil {
			if len(list) == 0 {
				*pID = -1
			} else {
				*pID = r.Int63()%int64(len(list))
			}
		}
	}

	return kc, nil
}