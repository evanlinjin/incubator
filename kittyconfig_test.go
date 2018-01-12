package incubator

import (
	"fmt"
	"testing"
)

func TestKittyConfig_ImagePath(t *testing.T) {
	kc := &KittyConfig{
		KittyID:  "1",
		HasBrows: true,
		HasCap:   true,
		BodyID:   "02",
		BrowsID:  "00",
		CapID:    "00",
		EarsID:   "01",
		EyesID:   "04",
		HeadID:   "02",
		NoseID:   "00",
		TailID:   "01",
	}
	cases := []struct {
		PN  PartName
		Exp string
	}{
		{PartBody, "Kitty_1/body/body_02.png"},
		{PartBrows, "Kitty_1/brows/brow_00.png"},
		{PartCap, "Kitty_1/caps/cap_00.png"},
		{PartEars, "Kitty_1/ears/ears_01.png"},
		{PartEyes, "Kitty_1/eyes/eyes_04.png"},
		{PartHead, "Kitty_1/head/head_02.png"},
		{PartNose, "Kitty_1/nose/nose_00.png"},
		{PartTail, "Kitty_1/tail/tail_01.png"},
	}
	for i, c := range cases {
		got := kc.ImagePath("", c.PN)
		out := fmt.Sprintf("[%d:%s] expected(%s) got(%s)", i, c.PN, c.Exp, got)
		if got != c.Exp {
			t.Error(out)
		} else {
			t.Log(out)
		}
	}
}
