package product

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetProductsByCategoryID(CategoryID int) ([]Product, error) {
	product, err := s.repository.GetProductsByCategoryID(CategoryID)
	if err != nil {
		return []Product{}, err
	}

	return product, err
}
