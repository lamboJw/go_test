package types

import (
	"strings"
)

type MusicType int

const (
	Unknown MusicType = iota
	Mp3
	Wav
)

func (t MusicType) String() string {
	switch t {
	case Mp3:
		return "Mp3"
	case Wav:
		return "Wav"
	default:
		return "Unknown"
	}
}

func ExtToMusicType(Ext string) MusicType {
	Ext = strings.TrimLeft(Ext, ".")
	Ext = strings.ToLower(Ext)
	switch Ext {
	case "mp3":
		return Mp3
	case "wav":
		return Wav
	default:
		return Unknown
	}
}
