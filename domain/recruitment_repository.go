package domain

type Predicate int

const (
	PredicateAsc Predicate = iota
	PredicateDesc
)

type Pageable struct {
	Page           int
	Size           int
	OrderBy        string
	OrderPredicate Predicate
}

type RecruitmentPage struct {
	Items      []Recruitment
	TotalCount int
	Page       int
	Size       int
}

type RecruitmentRepository interface {
	Save(recruitment *Recruitment) error
	SaveAll(recruitment ...*Recruitment) error
	FindById(id string) (*Recruitment, error)
	FindAllByResponsible(responsibleId string, pageable Pageable) (RecruitmentPage, error)
	FindAll(pageable Pageable) (RecruitmentPage, error)
}
