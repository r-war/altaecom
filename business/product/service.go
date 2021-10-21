package product

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) findProductByID(id int) (*Product, error) {
	return s.repository.findProductByID(id)
}
