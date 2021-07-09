package main

import (
	"image"
	"image/color"
	_ "image/color"
	_ "image/jpeg"
	"log"
	"net/http"

	"github.com/oliamb/cutter"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	// "newtab/imtools"
)

const (
	screenWidth  = 412
	screenHeight = 732
	homeButton_w = 64
	homeButton_h = 64
	gutter = 16
)

const sampleText = `The quick brown fox jumps over the lazy dog.`

const (
	australiaPic1 = "https://images.unsplash.com/photo-1606302217493-6556ee49c419?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80"
	australiaPic2 = "https://images.unsplash.com/photo-1554629907-479bff71f153?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=600&q=80"
	australiaPic3 = "https://images.unsplash.com/photo-1540881625515-1393e7064d37?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=626&q=80"
)


var (
	bodyFont font.Face
	miniTitleFont font.Face
	titleFont font.Face
	gophersImage *ebiten.Image
	rectImage *ebiten.Image
)

type Game struct{}

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	bodyFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	miniTitleFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    18,
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
	screen.Fill(rgb(0,0,0))
	//Create Image Card
	cardPositions := &ebiten.DrawImageOptions{}
	cardPositions.GeoM.Translate(screenWidth/8, screenHeight/5)
	//Create Button
	buttonPosition := &ebiten.DrawImageOptions{}
	buttonPosition.GeoM.Translate(screenWidth/2 - (homeButton_w/2), screenHeight - (homeButton_h + gutter))
	//Draw to screen
	screen.DrawImage(gophersImage, cardPositions)
	screen.DrawImage(rectImage, buttonPosition)
	text.Draw(screen, "AUSTRALASIA", bodyFont, screenWidth/4, 40, rgb(255,255,255))
	text.Draw(screen, "Japan", bodyFont, 20, 80, rgb(255,255,255))
	text.Draw(screen, "Australia", bodyFont, 20, 80, rgb(255,255,255))
	text.Draw(screen, "New Zealand", bodyFont, 20, 80, rgb(255,255,255))

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	resp, err := http.Get(australiaPic3)
	defer resp.Body.Close()
	img, _, err := image.Decode(resp.Body)
	if err != nil {
	// handle error
	}
		
	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width: screenWidth/2 +  screenWidth/4,
		Height: screenHeight/2 + screenHeight/4,
		Mode: cutter.Centered,
	  })

	gophersImage = ebiten.NewImageFromImage(croppedImg)
	rectImage = ebiten.NewImageFromImage(generateRectangle(homeButton_w, homeButton_h, rgb(0, 87, 255)))
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Traverse App")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func generateRectangle(w, h int, pixelColor color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, pixelColor)
		}
	}
	return img
}

func rgb(r, g, b int) color.RGBA {
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}