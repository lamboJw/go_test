package media

type player interface {
	Play() error
	Pause() error
}
