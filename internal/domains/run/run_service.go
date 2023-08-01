package run

type RunService struct {
	Repo *RunRepository
}

func NewRunService(repo *RunRepository) *RunService {
	return &RunService{Repo: repo}
}

func (s *RunService) Create(payload RunPayload) (res Run, err error) {
	res = res.NewFromPayload(payload)
	err = s.Repo.Create(res)
	return
}

func (s *RunService) GetAll(offset, limit int, field, sort, location string) (res []Run, err error) {
	res, err = s.Repo.GetAll(limit, offset, sort, field, location)
	if err != nil {
		return
	}
	return
}

func (s *RunService) Update(id string, load RunPayload) (res Run, err error) {
	res, err = s.Repo.GetByID(id)
	if err != nil {
		return
	}
	res.Update(load)
	err = s.Repo.Update(res)
	return
}

func (s *RunService) Delete(id string) (err error) {
	_, err = s.Repo.GetByID(id)
	if err != nil {
		return
	}
	err = s.Repo.Delete(id)
	return
}
