package incubator

import (
	"fmt"
	"math/rand"
	"path"
	"path/filepath"
	"reflect"
	"encoding/json"
	"github.com/skycoin/skycoin/src/cipher"
)

type KittyConfig struct {
	KittyID uint64 `json:"kitty_id"`
	Version uint64 `json:"version"`

	// -ve id represents that the kitty does not have that body part.
	BodyID  int64 `json:"body_id" part:"body"`
	BrowsID int64 `json:"brows_id" part:"brows"`
	CapID   int64 `json:"cap_id" part:"cap"`
	EarsID  int64 `json:"ears_id" part:"ears"`
	EyesID  int64 `json:"eyes_id" part:"eyes"`
	HeadID  int64 `json:"head_id" part:"head"`
	NoseID  int64 `json:"nose_id" part:"nose"`
	TailID  int64 `json:"tail_id" part:"tail"`
}

func (kc *KittyConfig) getPartIDPointer(partSpecs *KittyPartSpecs) *int64 {
	return reflect.
		ValueOf(kc).
		Elem().
		FieldByName(partSpecs.FieldName()).
		Addr().
		Interface().(*int64)
}

func (kc *KittyConfig) ImagePath(rootPath string, part KittyPart) (string, bool) {
	var partSpecs = part.Specs()

	if partID := kc.getPartIDPointer(partSpecs); partID == nil || *partID < 0 {
		return "", false
	} else {
		return filepath.Join(
			rootPath,
			fmt.Sprintf("Kitty_%d", kc.KittyID),
			partSpecs.FolderName(),
			fmt.Sprintf("%d.png", *partID),
		), true
	}
}

func (kc *KittyConfig) Hash() cipher.SHA256 {
	data, _ := json.Marshal(kc)
	return cipher.SumSHA256(data)
}

func (kc *KittyConfig) Print(pretty bool) string {
	var data []byte
	if pretty {
		data, _ = json.MarshalIndent(kc, "", "  ")
	} else {
		data, _ = json.Marshal(kc)
	}
	return string(data)
}

func RandomKittyConfig(r *rand.Rand, rootPath string) (*KittyConfig, error) {
	kc := new(KittyConfig)

	kitties, e := filepath.Glob(path.Join(rootPath, "Kitty_*"))
	if e != nil {
		return nil, e
	}
	kc.KittyID = r.Uint64() % uint64(len(kitties))

	var kittyPath = filepath.Join(
		rootPath,
		fmt.Sprintf("Kitty_%d", kc.KittyID),
	)

	RangeKittyParts(func(part KittyPart) error {
		partSpecs := part.Specs()
		list, _ := filepath.Glob(path.Join(kittyPath, partSpecs.FolderName(), "*.png"))

		if pID := kc.getPartIDPointer(partSpecs); pID != nil {
			if len(list) == 0 || partSpecs.Excludable() && r.Int()%2 == 0 {
				*pID = -1
			} else {
				*pID = r.Int63() % int64(len(list))
			}
		}
		return nil
	})

	return kc, nil
}

func AllKittyConfigs(rootPath string) ([]*KittyConfig, error) {
	return nil, nil
}
