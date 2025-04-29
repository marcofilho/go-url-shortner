package database

type Repository interface {
	GetUrl(id string) (string, error)
	SaveUrl(id string, url string) error
}
