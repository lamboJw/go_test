package media

type mediaInterface interface {
	player
	MediaInfoGetter
	mediaInfoSetter
	mediaFileGetSetter
	mediaIniter
}

type MediaPlayerInterface interface {
	player
	MediaInfoGetter
}
