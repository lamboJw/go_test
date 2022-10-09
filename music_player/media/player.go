package media

type player interface {
	init() error
	Play() error
	Pause() error
}
