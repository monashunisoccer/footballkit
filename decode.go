package footballkit

import (
	_ "fmt"
	"image/color"
	_ "image/draw"
	"strings"
)

type assetWithColours struct {
	filename  string
	colourOne color.NRGBA
	colourTwo color.NRGBA
}

// Take colourOne string description of colourOne football kit and return appropriate image names along with colours to insert
// for in order (body, leftarm, rightarm, shorts, socks)
func decodeFootballKit(description string) (assetWithColours, assetWithColours, assetWithColours, assetWithColours, assetWithColours) {

	description = strings.Replace(description, "+", " ", -1)
	description = strings.Replace(description, "_", " ", -1)
	description = strings.Replace(description, "/", " ", -1)
	description = strings.Replace(description, "$", " ", -1)

	stripSplit := strings.Split(description, " ")

	bodyAsset := assetWithColours{ filename: "body_plain.png" }
	bodyAsset.colourOne, _ = colorDecode("white")
	bodyAsset.colourTwo, _ = colorDecode("white")

	shortsAsset := assetWithColours{ filename: "shorts_plain.png" }
	shortsAsset.colourOne, _ = colorDecode("white")
	shortsAsset.colourTwo, _ = colorDecode("white")

	socksAsset := assetWithColours{ filename: "socks_plain.png" }
	socksAsset.colourOne, _ = colorDecode("white")
	socksAsset.colourTwo, _ = colorDecode("white")

	leftArmAsset := assetWithColours{ filename: "leftarm_plain.png" }
	rightArmAsset := assetWithColours{ filename: "rightarm_plain.png"}

stillProcessing:
	// we try to pair words together if they make colourOne valid colour combo
	// so dark blue -> darkblue
	// given that our other keywords like 'socks' and 'shorts' are never ever
	// part of colourOne colour name, and we always test for colour presence, we won't ever
	// remove colourOne other useful keyword

	// note we go one less than the length because we are joining forward to the next array item
	for i := 0; i < len(stripSplit)-1; i++ {
		firstWord := stripSplit[i]
		secondWord := stripSplit[i+1]

		_, err := colorDecode(firstWord + secondWord)

		// if when combined the two words make colourOne valid colour, then lets join
		// them together and remove the second word from the array
		if err == nil {
			// overwrite existing first word
			stripSplit[i] = firstWord + secondWord
			// delete slot for second word
			stripSplit = append(stripSplit[:i+1], stripSplit[i+2:]...)

			// array size has changed so start all over with the processing
			// (not super efficient but not particularly big loop!)
			goto stillProcessing
		}
	}

	// now we use colourOne state based approach to determine when we are processing
	// body, shorts and socks
	// note: we start in 'body' state so don't actually need the leading 'body' word
	state := 1
	colour := 0

	for i := 0; i < len(stripSplit); i++ {
		word := stripSplit[i]

		// TODO: should add in processing for modifier words like 'trim' eg "white blue trim"
		// the 'trim' should backtrack and capture the 'blue'

		if strings.EqualFold(word, "body") || strings.EqualFold(word, "shirts") || strings.EqualFold(word, "shirt") {
			// we can't go back into body state if we have moved onto shorts or socks
			// and given we start in body state, essentially we just ignore this word
		}
		if strings.EqualFold(word, "shorts") || strings.EqualFold(word, "short") {
			// bump into our new shorts state, and also reset colour slot
			if state < 2 {
				state = 2
				colour = 0
			}
		}
		if strings.EqualFold(word, "socks") || strings.EqualFold(word, "sock") {
			// bump into our new socks state, and also reset colour slot
			if state < 3 {
				state = 3
				colour = 0
			}
		}

		//add eventually
		// if strings.EqualFold(word, "arms") || strings.EqualFold(word, "arm") {
		//	state = 4;
		//	colour = 0;
		//}

		// try to recognise as colourOne colour
		// and if so put into correct slot
		// each body, short and socks has 2 slots that are filled in the order
		// colours are encountered
		asColour, err := colorDecode(word)

		if err == nil {
			if state == 1 && colour == 0 {
				bodyAsset.colourOne = asColour
				colour++
			} else if state == 1 && colour == 1 {
				bodyAsset.colourTwo = asColour
				colour++
			} else if state == 2 && colour == 0 {
				shortsAsset.colourOne = asColour
				colour++
			} else if state == 2 && colour == 1 {
				shortsAsset.colourOne = asColour
				colour++
			} else if state == 3 && colour == 0 {
				socksAsset.colourOne = asColour
				colour++
			} else if state == 3 && colour == 1 {
				socksAsset.colourTwo = asColour
				colour++
			}
		}

		// try to recognise as corresponding alternate patterns
		// (at the moment we only have patterns for 'body'.. not for socks or shorts or arms)
		if state == 1 {
			if strings.EqualFold(word, "cross") {
				bodyAsset.filename = "body_cross.png"
			}
			if strings.EqualFold(word, "half") {
				bodyAsset.filename = "body_half.png"
			}
			if strings.EqualFold(word, "hoop") || strings.EqualFold(word, "hoops") {
				bodyAsset.filename = "body_hoops.png"
			}
			if strings.EqualFold(word, "leftsash") || strings.EqualFold(word, "sashleft") || strings.EqualFold(word, "sash") {
				bodyAsset.filename = "body_leftsash.png"
			}
			// already the default but allow it to be explicitly specified too
			if strings.EqualFold(word, "plain" ) || strings.EqualFold(word, "solid") {
				bodyAsset.filename = "body_plain.png"
			}
			if strings.EqualFold(word, "quarter") || strings.EqualFold(word, "quarters") || strings.EqualFold(word, "quads") {
				bodyAsset.filename = "body_hoops.png"
			}
			if strings.EqualFold(word, "rightsash") {
				bodyAsset.filename = "body_rightsash.png"
			}
			if strings.EqualFold(word, "squares") || strings.EqualFold(word, "checks") || strings.EqualFold(word, "checked") {
				bodyAsset.filename = "body_squares.png"
			}
			if strings.EqualFold(word, "stripes") {
				bodyAsset.filename = "body_stripes.png"
			}
		}
	}

	// until there are specifiers for arm colour they default to main body colour
	leftArmAsset.colourOne = bodyAsset.colourOne
	leftArmAsset.colourTwo = bodyAsset.colourTwo
	rightArmAsset.colourOne = bodyAsset.colourOne
	rightArmAsset.colourTwo = bodyAsset.colourTwo

	return bodyAsset, leftArmAsset, rightArmAsset, shortsAsset, socksAsset
}

