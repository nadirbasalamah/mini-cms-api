package contents

type contentUsecase struct {
	contentRepository Repository
}

func NewContentUseCase(repository Repository) UseCase {
	return &contentUsecase{
		contentRepository: repository,
	}
}

func (usecase *contentUsecase) GetAll() ([]Domain, error) {
	return usecase.contentRepository.GetAll()
}

func (usecase *contentUsecase) GetByID(id string) (Domain, error) {
	return usecase.contentRepository.GetByID(id)
}

func (usecase *contentUsecase) Create(contentReq *Domain) (Domain, error) {
	return usecase.contentRepository.Create(contentReq)
}

func (usecase *contentUsecase) Update(contentReq *Domain, id string) (Domain, error) {
	return usecase.contentRepository.Update(contentReq, id)
}

func (usecase *contentUsecase) Delete(id string) error {
	return usecase.contentRepository.Delete(id)
}
