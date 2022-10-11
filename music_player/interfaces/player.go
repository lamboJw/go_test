package interfaces

type player interface {
	Play() error
	Pause() error
}
