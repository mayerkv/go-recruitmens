package domain

import "errors"

var (
	ErrVacancyNotExists = errors.New("vacancy not exists")
)

type VacancyService struct {
	vacancyRepository VacancyRepository
}

func (s *VacancyService) PostVacancy(positionId, customerId string) (*Vacancy, error) {
	vacancy := CreateVacancy(positionId, customerId)

	if err := s.vacancyRepository.Save(vacancy); err != nil {
		return nil, err
	}

	return vacancy, nil
}

func (s *VacancyService) ShowVacancies(customerId string, pageable Pageable) (VacancyPage, error) {
	return s.vacancyRepository.FindAllByCustomer(customerId, pageable)
}

func (s *VacancyService) SearchVacancies(pageable Pageable) (VacancyPage, error) {
	return s.vacancyRepository.FindAll(pageable)
}

func (s *VacancyService) ChangeVacancyPosition(vacancyId, positionId string) error {
	vacancy, err := s.getVacancy(vacancyId)
	if err != nil {
		return err
	}

	if err := vacancy.ChangePosition(positionId); err != nil {
		return err
	}

	return s.vacancyRepository.Save(vacancy)
}

func (s *VacancyService) ApproveVacancy(vacancyId string) error {
	vacancy, err := s.getVacancy(vacancyId)
	if err != nil {
		return err
	}

	if err := vacancy.Approve(); err != nil {
		return err
	}

	return s.vacancyRepository.Save(vacancy)
}

func (s *VacancyService) CloseVacancy(vacancyId string) error {
	vacancy, err := s.getVacancy(vacancyId)
	if err != nil {
		return err
	}

	if err := vacancy.Close(); err != nil {
		return err
	}

	return s.vacancyRepository.Save(vacancy)
}

func (s *VacancyService) RejectVacancy(vacancyId string) error {
	vacancy, err := s.getVacancy(vacancyId)
	if err != nil {
		return err
	}

	if err := vacancy.Reject(); err != nil {
		return err
	}

	return s.vacancyRepository.Save(vacancy)
}

func (s *VacancyService) TakeInWorkVacancy(vacancyId string) error {
	vacancy, err := s.getVacancy(vacancyId)
	if err != nil {
		return err
	}

	if err := vacancy.TakeInWork(); err != nil {
		return err
	}

	return s.vacancyRepository.Save(vacancy)
}

func (s *VacancyService) GetVacancy(vacancyId string) (*Vacancy, error) {
	return s.getVacancy(vacancyId)
}

func (s *VacancyService) getVacancy(vacancyId string) (*Vacancy, error) {
	vacancy, err := s.vacancyRepository.FindById(vacancyId)
	if err != nil {
		return nil, err
	}

	if vacancy == nil {
		return nil, ErrVacancyNotExists
	}

	return vacancy, err
}
