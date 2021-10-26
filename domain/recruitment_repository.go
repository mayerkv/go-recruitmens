package domain

type Predicate int

const (
	PredicateAsc Predicate = iota
	PredicateDesc
)

type Pageable struct {
	Page int
	Size int

}

type RecruitmentPage struct {
	Items      []Recruitment
	TotalCount int
	Page       int
	Size       int
}

type RecruitmentRepository interface {
	Save(recruitment *Recruitment) error
	FindById(id string) (*Recruitment, error)
	FindAllByResponsible(responsibleId string, page int, size int, orderBy string, orderPredicate int) (RecruitmentPage, error)
	FindAll(page int, size int, orderBy string, orderPredicate int)
}
