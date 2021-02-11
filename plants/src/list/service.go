package list

import (
	p "plants/src/plant"

	"github.com/pkg/errors"
)

type listService struct {
	r p.Repository
}

// NewListService initializes a create service with the necessary dependencies.
func NewListService(r p.Repository) p.ListService {
	return listService{r}
}

// ListAll lists all plants.
func (s listService) ListAll(uid string) ([]p.Plant, error) {
	result, err := s.r.FindAll(uid)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}

// ListOne lists a plant.
func (s listService) ListOne(uid string, id string) (p.Plant, error) {
	if id == "" {
		return p.Plant{}, p.ErrMissingData
	}

	result, err := s.r.FindOne(uid, id)
	if err != nil {
		return p.Plant{}, p.ErrNotFound
	}

	return result, nil
}
