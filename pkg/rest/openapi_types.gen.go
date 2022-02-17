// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package rest

// AudioLevel defines model for AudioLevel.
type AudioLevel struct {
	Level int `json:"level"`
}

// SoundActions defines model for SoundActions.
type SoundActions struct {
	Loop   string `json:"loop"`
	Play   string `json:"play"`
	Unloop string `json:"unloop"`
}

// SoundDetails defines model for SoundDetails.
type SoundDetails struct {
	BufferCount int          `json:"buffer-count"`
	Links       SoundActions `json:"links"`
	Name        string       `json:"name"`
}

// LightEffect defines model for LightEffect.
type LightEffect string

// Playlist defines model for Playlist.
type Playlist string

// Sound defines model for Sound.
type Sound string

// EntityList defines model for EntityList.
type EntityList struct {
	Entity []string `json:"entity"`
}

// PlaylistChance defines model for PlaylistChance.
type PlaylistChance struct {
	Chance int `json:"chance"`
}

// PlaylistPosition defines model for PlaylistPosition.
type PlaylistPosition struct {
	Position int `json:"position"`
}

// SongInfo defines model for SongInfo.
type SongInfo struct {
	Artist   string  `json:"artist"`
	Image    *string `json:"image,omitempty"`
	Playlist string  `json:"playlist"`
	Title    string  `json:"title"`
}

// SoundList defines model for SoundList.
type SoundList struct {
	Sounds *[]SoundDetails `json:"sounds,omitempty"`
}

// PostSystemIntensityJSONBody defines parameters for PostSystemIntensity.
type PostSystemIntensityJSONBody AudioLevel

// PostSystemVolumeJSONBody defines parameters for PostSystemVolume.
type PostSystemVolumeJSONBody AudioLevel

// PostSystemIntensityJSONRequestBody defines body for PostSystemIntensity for application/json ContentType.
type PostSystemIntensityJSONRequestBody PostSystemIntensityJSONBody

// PostSystemVolumeJSONRequestBody defines body for PostSystemVolume for application/json ContentType.
type PostSystemVolumeJSONRequestBody PostSystemVolumeJSONBody
