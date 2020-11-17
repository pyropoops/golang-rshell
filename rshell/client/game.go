package main

import (
	"bytes"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	// Settings
	screenWidth  = 960
	screenHeight = 540
)

var (
	idleSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	// Preload images
	img, _, err := image.Decode(bytes.NewReader(peeblePng))
	if err != nil {
		panic(err)
	}
	idleSprite, _ = ebiten.NewImageFromImage(img, 0)

	img, _, err = image.Decode(bytes.NewReader(backgroundPng))
	if err != nil {
		panic(err)
	}
	backgroundImage, _ = ebiten.NewImageFromImage(img, 0)
}

const (
	unit    = 16
	groundY = 445
)

type char struct {
	x  int
	y  int
	vx int
	vy int
}

func (c *char) tryJump() {
	// Now the character can jump anytime, even when the character is not on the ground.
	// If you want to restrict the character to jump only when it is on the ground, you can add an 'if' clause:
	//
	//     if peeble.y == groundY * unit {
	//         ...
	if c.y >= groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *char) Update() {
	c.x += c.vx
	c.y += c.vy
	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 4
	} else if c.vx < 0 {
		c.vx += 4
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

func (c *char) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2, 2)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(idleSprite, op)
}

type game struct {
	peeble *char
}

func (g *game) Update(_ *ebiten.Image) error {
	if g.peeble == nil {
		g.peeble = &char{x: 50 * unit, y: groundY * unit}
	}

	// Controls
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.peeble.vx = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.peeble.vx = 4 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.peeble.tryJump()
	}
	g.peeble.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// Draws Background Image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(backgroundImage, op)

	// Draws the peeble
	g.peeble.draw(screen)

	// Show the message
	msg := fmt.Sprintf("TPS: %0.2f\nPress the space key to jump.", ebiten.CurrentTPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func startGame() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Peeble Platformer Beta v0.1")
	if err := ebiten.RunGame(&game{}); err != nil {
		panic(err)
	}
}
