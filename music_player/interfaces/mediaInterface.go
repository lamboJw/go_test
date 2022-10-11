package interfaces

type MediaInterface interface {
	player
	InfoGetter
	InfoSetter
	FileGetSetter
	MediaIniter
}
