package footballkit

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func WriteOut(img *image.Image, name string) {
	buffer := new(bytes.Buffer)

	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	handle, err := os.Create(name)

	if err != nil {
		return
	}

	if _, err := handle.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func TestAndWriteExamples(t *testing.T) {

	WriteOut(RenderImage("body stripes skyblue white shorts navy socks navy"), "example-output/monashunisoccer.png")

	WriteOut(RenderImage("body red shorts white socks black"), "example-output/manutd.png")

	WriteOut(RenderImage("body claret shorts white socks light blue"), "example-output/astonvilla.png")

	WriteOut(RenderImage("hoops green white shorts white socks green"), "example-output/celtic.png")

	WriteOut(RenderImage("rightsash red white shorts white socks white"), "example-output/peru.png")

	WriteOut(RenderImage("checks red white shorts white socks blue"), "example-output/croatia.png")
}
