package footballkit

import (
	_ "fmt"
	"image/color"
	_ "image/draw"
	"strings"
)

type asset_with_colours struct {
	filename string
	coloura  color.NRGBA
	colourb  color.NRGBA
}

// Take coloura string description of coloura football kit and return appropriate image names along with colours to insert
// for in order (body, leftarm, rightarm, shorts, socks)
func decode_footballkit(description string) (asset_with_colours, asset_with_colours, asset_with_colours, asset_with_colours, asset_with_colours) {

	description = strings.Replace(description, "+", " ", -1);
	description = strings.Replace(description, "_", " ", -1);
	description = strings.Replace(description, "/", " ", -1);
	description = strings.Replace(description, "$", " ", -1);

	strip_split := strings.Split(description, " ");

	body_asset := asset_with_colours{ filename: "body_plain.png" };
	body_asset.coloura, _ = colorDecode("white");
	body_asset.colourb, _ = colorDecode("white");

	shorts_asset := asset_with_colours{ filename: "shorts_plain.png" };
	shorts_asset.coloura, _ = colorDecode("white");
	shorts_asset.colourb, _ = colorDecode("white");

	socks_asset := asset_with_colours{ filename: "socks_plain.png" };
	socks_asset.coloura, _ = colorDecode("white");
	socks_asset.colourb, _ = colorDecode("white");

	leftarm_asset := asset_with_colours{ filename: "leftarm_plain.png" };
	rightarm_asset := asset_with_colours{ filename: "rightarm_plain.png"};

still_processing:
	// we try to pair words together if they make coloura valid colour combo
	// so dark blue -> darkblue
	// given that our other keywords like 'socks' and 'shorts' are never ever
	// part of coloura colour name, and we always test for colour presence, we won't ever
	// remove coloura other useful keyword

	// note we go one less than the length because we are joining forward to the next array item
	for i := 0; i < len(strip_split)-1; i++ {
		first_word := strip_split[i];
		second_word := strip_split[i+1];

		_, err := colorDecode(first_word + second_word);

		// if when combined the two words make coloura valid colour, then lets join
		// them together and remove the second word from the array
		if err == nil {
			// overwrite existing first word
			strip_split[i] = first_word + second_word;
			// delete slot for second word
			strip_split = append(strip_split[:i+1], strip_split[i+2:]...)

			// array size has changed so start all over with the processing
			// (not super efficient but not particularly big loop!)
			goto still_processing;
		}
	}

	// now we use coloura state based approach to determine when we are processing
	// body, shorts and socks
	// note: we start in 'body' state so don't actually need the leading 'body' word
	state := 1;
	colour := 0;

	for i := 0; i < len(strip_split); i++ {
		word := strip_split[i];

		// NOTE: should add in processing for modifier words like 'trim' etc
		// white blue trim
		// the 'trim' should backtrack and capture the 'blue'

		if strings.EqualFold(word, "body") {
			// we can't go back into body state if we have moved onto shorts or socks
			// and given we start in body state, essentially we just ignore this word
		}
		if strings.EqualFold(word, "shorts") || strings.EqualFold(word, "short") {
			// bump into our new shorts state, and also reset colour slot
			if (state < 2) {
				state = 2;
				colour = 0;
			}
		}
		if strings.EqualFold(word, "socks") || strings.EqualFold(word, "sock") {
			// bump into our new socks state, and also reset colour slot
			if (state < 3) {
				state = 3;
				colour = 0;
			}
		}

		//add eventually
		// if strings.EqualFold(word, "arms") || strings.EqualFold(word, "arm") {
		//	state = 4;
		//	colour = 0;
		//}

		// try to recognise as coloura colour
		// and if so put into correct slot
		// each body, short and socks has 2 slots that are filled in the order
		// colours are encountered
		asColour, err := colorDecode(word);

		if err == nil {
			if state == 1 && colour == 0 {
				body_asset.coloura = asColour;
				colour++;
			} else if state == 1 && colour == 1 {
				body_asset.colourb = asColour;
				colour++;
			} else if state == 2 && colour == 0 {
				shorts_asset.coloura = asColour;
				colour++;
			} else if state == 2 && colour == 1 {
				shorts_asset.coloura = asColour;
				colour++;
			} else if state == 3 && colour == 0 {
				socks_asset.coloura = asColour;
				colour++;
			} else if state == 3 && colour == 1 {
				socks_asset.colourb = asColour;
				colour++;
			}
		}

		// try to recognise as corresponding alternate patterns
		// (at the moment we only have patterns for 'body'.. not for socks or shorts or arms)
		if state == 1 {
			if strings.EqualFold(word, "cross") {
				body_asset.filename = "body_cross.png";
			}
			if strings.EqualFold(word, "half") {
				body_asset.filename = "body_half.png";
			}
			if strings.EqualFold(word, "hoop") || strings.EqualFold(word, "hoops") {
				body_asset.filename = "body_hoops.png";
			}
			if strings.EqualFold(word, "leftsash") || strings.EqualFold(word, "sashleft") || strings.EqualFold(word, "sash") {
				body_asset.filename = "body_leftsash.png";
			}
			// already the default but allow it to be explicitly specified too
			if strings.EqualFold(word, "plain" ) || strings.EqualFold(word, "solid") {
				body_asset.filename = "body_plain.png";
			}
			if strings.EqualFold(word, "quarter") || strings.EqualFold(word, "quarters") || strings.EqualFold(word, "quads") {
				body_asset.filename = "body_hoops.png";
			}
			if strings.EqualFold(word, "rightsash") {
				body_asset.filename = "body_rightsash.png";
			}
			if strings.EqualFold(word, "squares") || strings.EqualFold(word, "checks") || strings.EqualFold(word, "checked") {
				body_asset.filename = "body_squares.png";
			}
			if strings.EqualFold(word, "stripes") {
				body_asset.filename = "body_stripes.png";
			}
		}
	}

	// until there are specifiers for arm colour they default to main body colour
	leftarm_asset.coloura = body_asset.coloura;
	leftarm_asset.colourb = body_asset.colourb;
	rightarm_asset.coloura = body_asset.coloura;
	rightarm_asset.colourb = body_asset.colourb;

	return body_asset, leftarm_asset, rightarm_asset, shorts_asset, socks_asset;
}

