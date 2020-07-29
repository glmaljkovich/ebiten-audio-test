package main

import (
	"log"

	"github.com/glmaljkovich/ebiten-audio-test/sounds"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"

	// "github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/vorbis"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	sampleRate   = 22050
	screenWidth  = 240
	screenHeight = 320
)

var sound *audio.Player

type Game struct {
}

func logErrorAndExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	sound = loadAudio(sounds.A)
}

func loadAudio(src []byte) *audio.Player {
	audioContext, _ := audio.NewContext(sampleRate)
	file := audio.BytesReadSeekCloser(src)

	// sound, err := mp3.Decode(audioContext, file)
	sound, err := vorbis.Decode(audioContext, file)
	logErrorAndExit(err)
	player, err := audio.NewPlayer(audioContext, sound)
	logErrorAndExit(err)
	return player
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		sound.Rewind()
		sound.Play()
	}

	if len(ebiten.TouchIDs()) > 0 {
		sound.Rewind()
		sound.Play()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := "Click/touch to play sound"
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("sound-test")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
