package main

import (
	"image"
	"image/color"
	"log"

	// "fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}
type Game struct {
	player  *Sprite
	sprites []*Sprite
}

func (g *Game) Update() error {

	// move the player based on keyboar input (left, right, up, down, wasd)
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.Y += 2
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)

	// Draw the player
	screen.DrawImage(
		g.player.Img.SubImage(image.Rect(0, 0, 64, 64)).(*ebiten.Image),
		opts,
	)
	for _, sprite := range g.sprites {
		opts.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Img.SubImage(image.Rect(0, 0, 70, 70)).(*ebiten.Image),
			opts,
		)
		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ready??? Prepare to fight!!!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/Characters(100x100)/Soldier/Soldier/Soldier.png")
	orcImg, _, err := ebitenutil.NewImageFromFile("assets/Characters(100x100)/Orc/Orc/Orc.png")

	// Screen and character dimensions
	// screenWidth := 1084
	// screenHeight := 720
	// characterWidth := 64
	// characterHeight := 64

	// Calculate the center position
	// initialX := float64((screenWidth - characterWidth) / 2)
	// initialY := float64((screenHeight - characterHeight) / 2)
	initialX := float64(0)
	initialY := float64(0)

	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(&Game{
		player: &Sprite{
			Img: playerImg,
			X:   initialX,
			Y:   initialY,
		},
		sprites: []*Sprite{
			{
				Img: orcImg,
				X:   100,
				Y:   100,
			},
			{
				Img: orcImg,
				X:   300,
				Y:   200,
			},
			{
				Img: orcImg,
				X:   190,
				Y:   150,
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
	ebiten.RunGame(game)
}
