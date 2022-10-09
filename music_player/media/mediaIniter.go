package media

type mediaIniter interface {
	initStreamer() error
	initMediaInfo() error
}
