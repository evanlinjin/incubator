package incubator

import (
	"fmt"
	"testing"
)

func TestKittyConfig_ImagePath(t *testing.T) {
	kc := &KittyConfig{
		KittyID:  1,
		HasBrows: true,
		HasCap:   true,
		BodyID:   2,
		BrowsID:  0,
		CapID:    0,
		EarsID:   1,
		EyesID:   4,
		HeadID:   2,
		NoseID:   0,
		TailID:   1,
	}
	cases := []struct {
		PN  PartName
		Exp string
	}{
		{PartBody, "Kitty_1/body/2.png"},
		{PartBrows, "Kitty_1/brows/0.png"},
		{PartCap, "Kitty_1/cap/0.png"},
		{PartEars, "Kitty_1/ears/1.png"},
		{PartEyes, "Kitty_1/eyes/4.png"},
		{PartHead, "Kitty_1/head/2.png"},
		{PartNose, "Kitty_1/nose/0.png"},
		{PartTail, "Kitty_1/tail/1.png"},
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
