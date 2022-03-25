//go:generate go run github.com/tc-hib/go-winres make --in "winres.json" --arch "amd64"

package main

import (
	"context"
	"time"

	"github.com/dulli/bbycrgo/pkg/common"
	"github.com/dulli/bbycrgo/pkg/hardware"
	"github.com/dulli/bbycrgo/pkg/lights"
	"github.com/dulli/bbycrgo/pkg/music"
	"github.com/dulli/bbycrgo/pkg/rest"
	"github.com/dulli/bbycrgo/pkg/shell"
	"github.com/dulli/bbycrgo/pkg/sounds"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: time.Stamp,
	})

	// Load configuration
	var cfg common.Config
	common.Configure(&cfg)

	// Prepare the lights command module and initialize the led groups
	lightPlayer, err := lights.NewRenderer("light-test", cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Renderer setup incomplete")
	} else {
		log.Info("Renderer setup complete")
	}
	log.Info("Gathering light effect files...")

	// Gather the light effect files
	err = lightPlayer.LoadEffects(cfg.Lights.Path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": cfg.Lights.Path,
			"err":  err,
		}).Fatal("Failed to load the light effect directory, is the path correct?")
	} else {
		log.WithFields(log.Fields{
			"num": lightPlayer.ListEffects(),
		}).Info("Loaded light effects")
	}

	driverLED, err := hardware.GetLEDDriver("ws281x")
	if err != nil {
		log.WithFields(log.Fields{
			"type":   "led",
			"driver": "ws281x",
			"err":    err,
		}).Error("Failed to load a driver")
	} else {
		driverLED.Setup(lightPlayer, cfg)
		defer driverLED.Close()
	}

	// Prepare the music command module and initialize the speaker
	musicPlayer, err := music.NewPlayer("music-rest", cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Player setup incomplete")
	} else {
		log.Info("Player setup complete")
	}
	log.Info("Gathering music files...")

	// Gather the music files
	err = musicPlayer.LoadPlaylists(cfg.Music.Path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": cfg.Music.Path,
			"err":  err,
		}).Fatal("Failed to load the music directory, is the path correct?")
	} else {
		log.WithFields(log.Fields{
			"num": musicPlayer.ListPlaylists(),
		}).Info("Loaded playlists")
	}
	musicPlayer.Play()

	// Prepare the sound command module and initialize the speaker
	soundPlayer, err := sounds.NewPlayer("sounds-rest", cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Player setup incomplete")
	} else {
		log.Info("Player setup complete")
	}
	log.Info("Gathering sound files...")

	// Gather the sound files
	err = soundPlayer.LoadSounds(cfg.Sounds.Path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": cfg.Sounds.Path,
			"err":  err,
		}).Fatal("Failed to load the sound directory, is the path correct?")
	} else {
		log.WithFields(log.Fields{
			"num": len(soundPlayer.ListSounds()),
		}).Info("Loaded sounds")
	}

	// Prepare shell command execution
	shellExec, err := shell.NewExecutor("shell-rest", cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Shell setup incomplete")
	} else {
		log.Info("Shell setup complete")
	}

	api := rest.Server{}
	srv := api.Start(musicPlayer, soundPlayer, lightPlayer, shellExec)

	go common.EventLoop()
	common.AwaitSignal()

	// Perform clean up
	srv.Shutdown(context.Background())
	log.Info("Closing")
}
