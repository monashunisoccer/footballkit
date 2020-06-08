package footballkit

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"

	_ "github.com/go-bindata/go-bindata"
)

// Given a string description of a football kit, render a rendered image of that kit
func RenderImage(stripDescription string) *image.Image {

	var bodyAsset, leftArmAsset, rightArmAsset, shortsAsset, socksAsset = decode_footballkit(stripDescription)

	/*
	This is not at all clever - we chose a fixed size (based on the wikipedia football kit images) and
	have hardcoded all our shapes to that.

	   100x135 is our image sizes

	   left arm 0,0.. 31x59
	   body     31,0 .. 38x59
	   right arm 69,0 .. 31x59
	   shorts    0,59 .. 100x36
	   socks    0,95 .. 100x40
	*/
	m := image.NewNRGBA(image.Rect(0, 0, 100, 135))

	imagePaste(bodyAsset.filename, m, 31, 0, bodyAsset.coloura, bodyAsset.colourb)
	imagePaste(leftArmAsset.filename, m, 0, 0, leftArmAsset.coloura, leftArmAsset.colourb)
	imagePaste(rightArmAsset.filename, m, 69, 0, rightArmAsset.coloura, rightArmAsset.colourb)
	imagePaste(shortsAsset.filename, m, 0, 59, shortsAsset.coloura, shortsAsset.colourb)
	imagePaste(socksAsset.filename, m, 0, 95, socksAsset.coloura, socksAsset.colourb)

	var img image.Image = m

	return &img
}

// Takes a name of an filename image file and replaces all the strong red/green pixels with the alternate
// specified colours, and pastes into a destination image object at 'destX' and 'destY'
func imagePaste(assetName string, destImage *image.NRGBA, destX int, destY int, colorA color.NRGBA, colorB color.NRGBA) {

	// find the raw filename bytes
	pngData, err := Asset("data/" + assetName)

	if err != nil {
		log.Println("filename error")
		return
	}

	// decode the raw filename png image
	pngImg, err := png.Decode(bytes.NewBuffer(pngData))

	if err != nil {
		log.Println("decode png error")
		return
	}

	// we are going to pixel by pixel place it into the destination
	// BUT are going to do some colour corrections on the way

	// this is not exactly high performance but these are very small images so meh..
	bounds := pngImg.Bounds()

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {

			pngColour := destImage.ColorModel().Convert(pngImg.At(x, y)).(color.NRGBA)

			if pngColour.A > 200 && pngColour.R == 255 && pngColour.G == 0 && pngColour.B == 0 {
				destImage.SetNRGBA(x + destX, y + destY, colorA)
			} else if pngColour.A > 200 && pngColour.R == 0 && pngColour.G == 255 && pngColour.B == 0  {
				destImage.SetNRGBA(x + destX, y + destY, colorB)
			} else {
				destImage.SetNRGBA(x + destX, y + destY, pngColour)
			}
		}
	}
}
