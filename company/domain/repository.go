package company

type Repository interface {
	GetAll() (companies []Company, err error)
	Find(id string) (company *Company,err error)
	Store(company *Company) (companyStored *Company,err error)
	Delete(companyID string) (err error)
}
