package repository

type Authorization interface {

}

type MoneySource interface {

}

type Repository struct {
	Authorization
	MoneySource
}

func NewRepository() *Repository {
	return &Repository{}
}