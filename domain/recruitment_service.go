package domain

import "errors"

var (
	ErrRecruitmentNotExists = errors.New("recruitment not exists")
)

type RecruitmentService struct {
	recruitmentRepository RecruitmentRepository
	vacancyRepository     VacancyRepository
}

func (s *RecruitmentService) getVacancy(vacancyId string) (*Vacancy, error) {
	vacancy, err := s.vacancyRepository.FindById(vacancyId)
	if err != nil {
		return nil, err
	}
	if vacancy == nil {
		return nil, ErrVacancyNotExists
	}

	return vacancy, nil
}

func (s *RecruitmentService) getRecruitment(recruitmentId string) (*Recruitment, error) {
	recruitment, err := s.recruitmentRepository.FindById(recruitmentId)
	if err != nil {
		return nil, err
	}
	if recruitment == nil {
		return nil, ErrRecruitmentNotExists
	}

	return recruitment, nil
}

func (s *RecruitmentService) ConsiderCandidate(vacancyId, candidateId, responsibleId string, settings []StageLineSettings) (*Recruitment, error) {
	vacancy, err := s.getVacancy(vacancyId)
	if err != nil {
		return nil, err
	}

	recruitment, err := vacancy.ConsiderCandidate(candidateId, responsibleId, settings)
	if err != nil {
		return nil, err
	}

	if err := s.recruitmentRepository.Save(recruitment); err != nil {
		return nil, err
	}

	return recruitment, nil
}

func (s *RecruitmentService) ConsiderCandidateAnotherVacancy(recruitmentId, vacancyId string, settings []StageLineSettings) (*Recruitment, error) {
	vacancy, err := s.getVacancy(vacancyId)
	if err != nil {
		return nil, err
	}

	from, err := s.getRecruitment(recruitmentId)
	if err != nil {
		return nil, err
	}

	to, err := from.TransferTo(vacancy, settings)
	if err != nil {
		return nil, err
	}

	if err := s.recruitmentRepository.SaveAll(from, to); err != nil {
		return nil, err
	}

	return to, nil
}

func (s *RecruitmentService) AcceptRecruitmentStage(recruitmentId, requestedStageId string) error {
	recruitment, err := s.getRecruitment(recruitmentId)
	if err != nil {
		return err
	}

	if err := recruitment.AcceptCurrentStage(requestedStageId); err != nil {
		return err
	}

	return s.recruitmentRepository.Save(recruitment)
}

func (s *RecruitmentService) DenyRecruitment(recruitmentId, reasonId, comment string) error {
	recruitment, err := s.getRecruitment(recruitmentId)
	if err != nil {
		return err
	}

	if err := recruitment.Deny(reasonId, comment); err != nil {
		return err
	}

	return s.recruitmentRepository.Save(recruitment)
}

func (s *RecruitmentService) GetRecruitment(recruitmentId string) (*Recruitment, error) {

}

func (s *RecruitmentService) ShowRecruitments(responsibleId string, pageable Pageable) (RecruitmentPage, error) {

}
