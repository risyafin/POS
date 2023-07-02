package branch

type Usecase struct {
	Repo Repository
}

func (usecase Usecase) GetBrands(limit int, offset int, colum string, sort string, search string) ([]Branch, error) {
	brands, err := usecase.Repo.GetBrands(limit, offset, colum, sort, search)
	return brands, err
}
