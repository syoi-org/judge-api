package codeforces

type Service interface {
	GetProblem(contestID int, index string) (*Problem, error)
}

type service struct {
	Config *Config
}

type ServiceParams struct {
	Config *Config
}

func NewService(p ServiceParams) Service {
	return &service{
		Config: p.Config,
	}
}

func (s *service) GetProblem(contestID int, index string) (*Problem, error) {

	return nil, nil
}
