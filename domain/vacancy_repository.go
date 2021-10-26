package domain

type VacancyPage struct {
	Items      []Vacancy
	TotalCount int
	Page       int
	Size       int
}

type VacancyRepository interface {
	Save(vacancy *Vacancy) error
	FindById(id string) (*Vacancy, error)
	FindAllByCustomer(customerId string, pageable Pageable) (VacancyPage, error)
	FindAll(pageable Pageable) (VacancyPage, error)
}
