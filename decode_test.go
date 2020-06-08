package footballkit

import (
	"testing"
)

func TestEasy(t *testing.T) {

	body, _, _, shorts, socks := decodeFootballKit("red shorts black socks yellow")

	if body.filename != "body_plain.png" {
		t.Fatalf("Expected plain body but got %s", body.filename)
	}
	if body.colourOne.R != 255 || body.colourOne.G != 0 || body.colourOne.B != 0 {
		t.Fatalf("Expected red but got %+v", body.colourOne)
	}

	if shorts.filename != "shorts_plain.png" {
		t.Fatalf("Expected plain shorts but got %s", shorts.filename)
	}
	if shorts.colourOne.R != 0 || shorts.colourOne.G != 0 || shorts.colourOne.B != 0 {
		t.Fatalf("Expected black but got %+v", shorts.colourOne)
	}

	if socks.filename != "socks_plain.png" {
		t.Fatalf("Expected plain socks but got %s", socks.filename)
	}
	if socks.colourOne.R != 255 || socks.colourOne.G != 255 || socks.colourOne.B != 0 {
		t.Fatalf("Expected yellow but got %+v", socks.colourOne)
	}

}

func TestIgnoreExtraCharacters(t *testing.T) {

	body, _, _, shorts, socks := decodeFootballKit("red+shorts_black nonsense socks $$$$ yellow")

	if body.filename != "body_plain.png" {
		t.Fatalf("Expected plain body but got %s", body.filename)
	}
	if body.colourOne.R != 255 || body.colourOne.G != 0 || body.colourOne.B != 0 {
		t.Fatalf("Expected red but got %+v", body.colourOne)
	}

	if shorts.filename != "shorts_plain.png" {
		t.Fatalf("Expected plain shorts but got %s", shorts.filename)
	}
	if shorts.colourOne.R != 0 || shorts.colourOne.G != 0 || shorts.colourOne.B != 0 {
		t.Fatalf("Expected black but got %+v", shorts.colourOne)
	}

	if socks.filename != "socks_plain.png" {
		t.Fatalf("Expected plain socks but got %s", socks.filename)
	}
	if socks.colourOne.R != 255 || socks.colourOne.G != 255 || socks.colourOne.B != 0 {
		t.Fatalf("Expected yellow but got %+v", socks.colourOne)
	}

}

func TestColourJoining(t *testing.T) {

	_, _, _, shorts, socks := decodeFootballKit("body red shorts light blue socks dark green")

	if shorts.filename != "shorts_plain.png" {
		t.Fatalf("Expected plain shorts but got %s", shorts.filename)
	}
	if shorts.colourOne.R != 173 || shorts.colourOne.G != 216 || shorts.colourOne.B != 230 {
		t.Fatalf("Expected light blue but got %+v", shorts.colourOne)
	}

	if socks.filename != "socks_plain.png" {
		t.Fatalf("Expected plain socks but got %s", socks.filename)
	}
	if socks.colourOne.R != 0 || socks.colourOne.G != 100 || socks.colourOne.B != 0 {
		t.Fatalf("Expected dark green but got %+v", socks.colourOne)
	}

}
