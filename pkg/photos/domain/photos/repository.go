package photos

type Repository interface {
	Save(*Photo) error
	Get() (*Photo, error)
}
