package loader

import (
	"github.com/x-hgg-x/goecsengine/loader"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/hajimehoshi/ebiten/audio"
)

// LoadSounds loads music and sfx
func LoadSounds(world w.World, sound bool) {
	world.Resources.AudioContext = loader.InitAudio(44100)
	audioPlayers := make(map[string]*audio.Player)
	audioPlayers["ufo1"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/ufo_highpitch.wav")
	audioPlayers["ufo2"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/ufo_lowpitch.wav")
	audioPlayers["invader1"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/fastinvader1.wav")
	audioPlayers["invader2"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/fastinvader2.wav")
	audioPlayers["invader3"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/fastinvader3.wav")
	audioPlayers["invader4"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/fastinvader4.wav")
	audioPlayers["shoot"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/shoot.wav")
	audioPlayers["killed"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/invaderkilled.wav")
	audioPlayers["explosion"] = loader.LoadAudio(world.Resources.AudioContext, "assets/audio/explosion.wav")
	world.Resources.AudioPlayers = &audioPlayers

	audioPlayers["music"].Play()
	if sound {
		EnableSound(world)
	} else {
		DisableSound(world)
	}
}

// EnableSound enables sound
func EnableSound(world w.World) {
	audioPlayers := *world.Resources.AudioPlayers
	audioPlayers["ufo1"].SetVolume(0.15)
	audioPlayers["ufo2"].SetVolume(0.15)
	audioPlayers["invader1"].SetVolume(0.15)
	audioPlayers["invader2"].SetVolume(0.15)
	audioPlayers["invader3"].SetVolume(0.15)
	audioPlayers["invader4"].SetVolume(0.15)
	audioPlayers["shoot"].SetVolume(0.15)
	audioPlayers["killed"].SetVolume(0.15)
	audioPlayers["explosion"].SetVolume(0.15)
}

// DisableSound disables sound
func DisableSound(world w.World) {
	audioPlayers := *world.Resources.AudioPlayers
	for key := range audioPlayers {
		audioPlayers[key].SetVolume(0)
	}
}
