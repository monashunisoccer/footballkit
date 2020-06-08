package footballkit

import (
	"testing"
)

func TestEasy(t *testing.T) {

	body, _, _, shorts, socks := decode_footballkit("red shorts black socks yellow")

	if body.filename != "body_plain.png" {
		t.Fatalf("Expected plain body but got %s", body.filename)
	}
	if body.coloura.R != 255 || body.coloura.G != 0 || body.coloura.B != 0 {
		t.Fatalf("Expected red but got %+v", body.coloura)
	}

	if shorts.filename != "shorts_plain.png" {
		t.Fatalf("Expected plain shorts but got %s", shorts.filename)
	}
	if shorts.coloura.R != 0 || shorts.coloura.G != 0 || shorts.coloura.B != 0 {
		t.Fatalf("Expected black but got %+v", shorts.coloura)
	}

	if socks.filename != "socks_plain.png" {
		t.Fatalf("Expected plain socks but got %s", socks.filename)
	}
	if socks.coloura.R != 255 || socks.coloura.G != 255 || socks.coloura.B != 0 {
		t.Fatalf("Expected yellow but got %+v", socks.coloura)
	}

}

func TestIgnoreExtraCharacters(t *testing.T) {

	body, _, _, shorts, socks := decode_footballkit("red+shorts_black nonsense socks $$$$ yellow")

	if body.filename != "body_plain.png" {
		t.Fatalf("Expected plain body but got %s", body.filename)
	}
	if body.coloura.R != 255 || body.coloura.G != 0 || body.coloura.B != 0 {
		t.Fatalf("Expected red but got %+v", body.coloura)
	}

	if shorts.filename != "shorts_plain.png" {
		t.Fatalf("Expected plain shorts but got %s", shorts.filename)
	}
	if shorts.coloura.R != 0 || shorts.coloura.G != 0 || shorts.coloura.B != 0 {
		t.Fatalf("Expected black but got %+v", shorts.coloura)
	}

	if socks.filename != "socks_plain.png" {
		t.Fatalf("Expected plain socks but got %s", socks.filename)
	}
	if socks.coloura.R != 255 || socks.coloura.G != 255 || socks.coloura.B != 0 {
		t.Fatalf("Expected yellow but got %+v", socks.coloura)
	}

}

func TestColourJoining(t *testing.T) {

	_, _, _, shorts, socks := decode_footballkit("body red shorts light blue socks dark green")

	if shorts.filename != "shorts_plain.png" {
		t.Fatalf("Expected plain shorts but got %s", shorts.filename)
	}
	if shorts.coloura.R != 173 || shorts.coloura.G != 216 || shorts.coloura.B != 230 {
		t.Fatalf("Expected light blue but got %+v", shorts.coloura)
	}

	if socks.filename != "socks_plain.png" {
		t.Fatalf("Expected plain socks but got %s", socks.filename)
	}
	if socks.coloura.R != 0 || socks.coloura.G != 100 || socks.coloura.B != 0 {
		t.Fatalf("Expected dark green but got %+v", socks.coloura)
	}

}
