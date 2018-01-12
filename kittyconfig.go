package incubator

import (
	"fmt"
	"path/filepath"
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

func (kc *KittyConfig) ImagePath(rootPath string, partName PartName) (string, bool) {
	var partID int64
	switch partName {
	case PartBody:
		partID = kc.BodyID
	case PartBrows:
		partID = kc.BrowsID
	case PartCap:
		partID = kc.CapID
	case PartEars:
		partID = kc.EarsID
	case PartEyes:
		partID = kc.EyesID
	case PartHead:
		partID = kc.HeadID
	case PartNose:
		partID = kc.NoseID
	case PartTail:
		partID = kc.TailID
	}
	if partID < 0 {
		return "", false
	} else {
		return filepath.Join(
			rootPath,
			fmt.Sprintf("Kitty_%d", kc.KittyID),
			string(partName),
			fmt.Sprintf("%d.png", partID),
		), true
	}
}