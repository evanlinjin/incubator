package incubator

import (
	"sync"
	"errors"
)

type KittyPart int

func (kp KittyPart) Specs() *KittyPartSpecs {
	return kittyParts[kp]
}

const (
	Tail KittyPart = iota
	Body
	Ears
	Head
	Eyes
	Brows
	Nose
	Cap
	kittyPartsCount
)

type KittyPartSpecs struct {
	folderName string // Folder name of the part.
	fieldName  string // KittyConfig field folderName of the part.
	mutable    bool   // Can the part be changed once set.
	excludable bool   // Can the part be excluded.
}

func (kps *KittyPartSpecs) FolderName() string { return kps.folderName }
func (kps *KittyPartSpecs) FieldName() string  { return kps.fieldName }
func (kps *KittyPartSpecs) Mutable() bool      { return kps.mutable }
func (kps *KittyPartSpecs) Excludable() bool   { return kps.excludable }

var (
	kittyPartsByName sync.Map
	kittyParts       = [...]*KittyPartSpecs{
		Body:  {"body", "BodyID", false, false},
		Brows: {"brows", "BrowsID", false, false},
		Cap:   {"cap", "CapID", true, true},
		Ears:  {"ears", "EarsID", false, false},
		Eyes:  {"eyes", "EyesID", false, false},
		Head:  {"head", "HeadID", false, false},
		Nose:  {"nose", "NoseID", false, false},
		Tail:  {"tail", "TailID", false, false},
	}
)

func init() {
	for _, part := range kittyParts {
		kittyPartsByName.Store(part.folderName, part)
	}
}

var ReturnNoError = errors.New("returned action, no error")

type KittyPartAction func(part KittyPart) error

func RangeKittyParts(action KittyPartAction) error {
	for id := KittyPart(0); id < kittyPartsCount; id++ {
		if e := action(id); e != nil {
			switch e {
			case ReturnNoError:
				return nil
			default:
				return e
			}
		}
	}
	return nil
}