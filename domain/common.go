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
