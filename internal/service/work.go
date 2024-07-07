package service

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
)

func (s *WorkService) StartWork(ctx context.Context, work *model.Work) (string, error) {
	return s.WorkRepository.AddWork(ctx, work)
}

func (s *WorkService) StopWork(ctx context.Context, id string) error {
	return s.WorkRepository.StopWork(ctx, id)
}

func (s *WorkService) GetWorkById(ctx context.Context, id string) (*model.Work, error) {
	workDB, err := s.WorkRepository.GetWorkById(ctx, id)
	if err != nil {
		return nil, err
	}
	return workDB.ToWork(), nil
}

func (s *WorkService) GetWorks(ctx context.Context, user string) ([]*model.Work, error) {
	works, err := s.WorkRepository.GetWorks(ctx, user)
	if err != nil {
		return nil, err
	}

	res := make([]*model.Work, 0, len(works))
	for _, work := range works {
		res = append(res, work.ToWork())
	}

	return res, nil
}
