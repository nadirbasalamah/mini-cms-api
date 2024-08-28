package categories

type CategoryUseCase struct {
	categoryRepository Repository
}

func NewCategoryUseCase(repository Repository) UseCase {
	return &CategoryUseCase{
		categoryRepository: repository,
	}
}

func (usecase *CategoryUseCase) GetAll() ([]Domain, error) {
	return usecase.categoryRepository.GetAll()
}

func (usecase *CategoryUseCase) GetByID(id string) (Domain, error) {
	return usecase.categoryRepository.GetByID(id)
}

func (usecase *CategoryUseCase) Create(contentReq *Domain) (Domain, error) {
	return usecase.categoryRepository.Create(contentReq)
}

func (usecase *CategoryUseCase) Update(contentReq *Domain, id string) (Domain, error) {
	return usecase.categoryRepository.Update(contentReq, id)
}

func (usecase *CategoryUseCase) Delete(id string) error {
	return usecase.categoryRepository.Delete(id)
}
