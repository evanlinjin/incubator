package incubator

import (
	"fmt"
	"path/filepath"
)

type PartName string

const (
	PartBody  PartName = "body"
	PartBrows PartName = "brows"
	PartCap   PartName = "caps"
	PartEars  PartName = "ears"
	PartEyes  PartName = "eyes"
	PartHead  PartName = "head"
	PartNose  PartName = "nose"
	PartTail  PartName = "tail"
)

func (pn PartName) Secondary() string {
	var out string
	switch pn {
	case PartCap:
		out = "cap"
	case PartBrows:
		out = "brow"
	default:
		out = string(pn)
	}
	return out
}

type KittyConfig struct {
	KittyID string `json:"kitty_id"`
	Owner   string `json:"owner"`

	HasBrows bool `json:"has_brows"`
	HasCap   bool `json:"has_cap"`

	BodyID  string `json:"body_id"`
	BrowsID string `json:"brows_id"`
	CapID   string `json:"cap_id"`
	EarsID  string `json:"ears_id"`
	EyesID  string `json:"eyes_id"`
	HeadID  string `json:"head_id"`
	NoseID  string `json:"nose_id"`
	TailID  string `json:"tail_id"`
}

func (kc *KittyConfig) ImagePath(rootPath string, partName PartName) string {
	var partID string
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
	return filepath.Join(
		rootPath,
		fmt.Sprintf("Kitty_%s", kc.KittyID),
		string(partName),
		fmt.Sprintf("%s_%s.png", partName.Secondary(), partID),
	)
}
