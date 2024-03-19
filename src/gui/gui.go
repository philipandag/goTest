package gui

import (
	"gotest/src/game"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameGui struct {
	Game game.Game
}

func (g *GameGui) Update() error {
	return nil
}
func (g *GameGui) Draw(screen *ebiten.Image) {
	x := g.Game.GameObject.X*50 + 50
	y := g.Game.GameObject.Y*50 + 50
	ebitenutil.DrawCircle(screen, x, y, 10, color.RGBA{30, 255, 250, 255})
	ebitenutil.DebugPrint(screen, g.Game.GameObject.ToString())
}
func (g *GameGui) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *GameGui) Run() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello")

	if err := ebiten.RunGame(g); err != nil {

	}
}
