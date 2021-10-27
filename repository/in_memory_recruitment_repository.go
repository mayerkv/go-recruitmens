package repository

import (
	"github.com/mayerkv/go-recruitmens/domain"
	"sync"
)

type InMemoryRecruitmentRepository struct {
	sync.Mutex
	items map[string]domain.Recruitment
}

func NewInMemoryRecruitmentRepository() domain.RecruitmentRepository {
	return &InMemoryRecruitmentRepository{
		items: map[string]domain.Recruitment{},
	}
}

func (r *InMemoryRecruitmentRepository) Save(recruitment *domain.Recruitment) error {
	r.Lock()
	defer r.Unlock()

	r.items[recruitment.Id] = *recruitment

	return nil
}

func (r *InMemoryRecruitmentRepository) SaveAll(recruitment ...*domain.Recruitment) error {
	r.Lock()
	defer r.Unlock()

	for _, rec := range recruitment {
		r.items[rec.Id] = *rec
	}

	return nil
}

func (r *InMemoryRecruitmentRepository) FindById(id string) (*domain.Recruitment, error) {
	r.Lock()
	defer r.Unlock()

	if item, ok := r.items[id]; ok {
		return &item, nil
	}

	return nil, nil
}

// FindAllByResponsible todo sort
func (r *InMemoryRecruitmentRepository) FindAllByResponsible(responsibleId string, pageable domain.Pageable) (domain.RecruitmentPage, error) {
	r.Lock()
	defer r.Unlock()

	var recruitments []domain.Recruitment
	for _, item := range r.items {
		if item.ResponsibleId == responsibleId {
			recruitments = append(recruitments, item)
		}
	}

	page := domain.RecruitmentPage{
		Items:      r.skip(recruitments, pageable),
		TotalCount: len(recruitments),
		Page:       pageable.Page,
		Size:       pageable.Size,
	}

	return page, nil
}

// FindAll todo sort
func (r *InMemoryRecruitmentRepository) FindAll(pageable domain.Pageable) (domain.RecruitmentPage, error) {
	r.Lock()
	defer r.Unlock()

	var recruitments []domain.Recruitment
	for _, item := range r.items {
		recruitments = append(recruitments, item)
	}

	page := domain.RecruitmentPage{
		Items:      r.skip(recruitments, pageable),
		TotalCount: len(recruitments),
		Page:       pageable.Page,
		Size:       pageable.Size,
	}

	return page, nil
}

func (r *InMemoryRecruitmentRepository) skip(recruitments []domain.Recruitment, pageable domain.Pageable) []domain.Recruitment {
	var items []domain.Recruitment
	skip := (pageable.Page - 1) * pageable.Size
	if len(recruitments) >= skip {
		items = recruitments[skip : skip+pageable.Size]
	}

	return items
}
