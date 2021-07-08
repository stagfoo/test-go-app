package main

import (
	"image"
	"image/color"
	_ "image/color"
	_ "image/jpeg"
	"log"
	"net/http"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 411
	screenHeight = 731
)

const sampleText = `The quick brown fox jumps over the lazy dog.`

var (
	mplusNormalFont font.Face
	gophersImage *ebiten.Image
)

type Game struct{}

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xFF, 0xFF, 0xFF, 0xff})
	text.Draw(screen, sampleText, mplusNormalFont, 20, 80, color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(gophersImage, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	resp, err := http.Get("https://images.unsplash.com/photo-1444090695923-48e08781a76a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=800&q=80")
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
	// handle error
	}

	gophersImage = ebiten.NewImageFromImage(img)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Traverse App")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
