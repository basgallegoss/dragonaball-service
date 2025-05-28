package domain

type Repository interface {
	FindByName(name string) (*Character, error)
	Save(c Character) error
}
