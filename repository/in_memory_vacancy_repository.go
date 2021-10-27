package repository

import (
	"github.com/mayerkv/go-recruitmens/domain"
	"sync"
)

type InMemoryVacancyRepository struct {
	sync.Mutex
	items map[string]domain.Vacancy
}

func NewInMemoryVacancyRepository() domain.VacancyRepository {
	return &InMemoryVacancyRepository{
		items: map[string]domain.Vacancy{},
	}
}

func (r *InMemoryVacancyRepository) Save(vacancy *domain.Vacancy) error {
	r.Lock()
	defer r.Unlock()

	r.items[vacancy.Id] = *vacancy

	return nil
}

func (r *InMemoryVacancyRepository) FindById(id string) (*domain.Vacancy, error) {
	r.Lock()
	defer r.Unlock()

	vacancy, ok := r.items[id]
	if ok {
		return &vacancy, nil
	}

	return nil, nil
}

func (r *InMemoryVacancyRepository) FindAllByCustomer(customerId string, pageable domain.Pageable) (domain.VacancyPage, error) {
	r.Lock()
	defer r.Unlock()

	var vacanciesByCustomer []domain.Vacancy

	for _, item := range r.items {
		if item.CustomerId == customerId {
			vacanciesByCustomer = append(vacanciesByCustomer, item)
		}
	}

	// todo sort

	page := domain.VacancyPage{
		Items:      r.skip(vacanciesByCustomer, pageable),
		TotalCount: len(vacanciesByCustomer),
		Page:       pageable.Page,
		Size:       pageable.Size,
	}

	return page, nil
}

func (r *InMemoryVacancyRepository) FindAll(pageable domain.Pageable) (domain.VacancyPage, error) {
	r.Lock()
	defer r.Unlock()

	var vacancies []domain.Vacancy
	for _, item := range r.items {
		vacancies = append(vacancies, item)
	}

	page := domain.VacancyPage{
		Items:      r.skip(vacancies, pageable),
		TotalCount: len(vacancies),
		Page:       pageable.Page,
		Size:       pageable.Size,
	}

	return page, nil
}

func (r *InMemoryVacancyRepository) skip(vacancies []domain.Vacancy, pageable domain.Pageable) []domain.Vacancy {
	var items []domain.Vacancy
	skip := (pageable.Page - 1) * pageable.Size
	if len(vacancies) >= skip {
		items = vacancies[skip : skip+pageable.Size]
	}

	return items
}
