package media

type mediaInfoGetter interface {
	Name() string
	Size() int64
	Id() string
	Title() string
	Artist() string
	Year() string
	Comment() string
	Genre() uint8
	Sort() int64
}
