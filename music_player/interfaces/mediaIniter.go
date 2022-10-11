package interfaces

type MediaIniter interface {
	InitStreamer() error
	InitMediaInfo() error
}
