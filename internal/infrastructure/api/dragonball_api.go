package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/basgallegoss/dragonball-service/internal/domain"
)

type DragonballAPI struct {
	baseURL string
}

func NewDragonballAPI(baseURL string) *DragonballAPI {
	return &DragonballAPI{baseURL: baseURL}
}
func (d *DragonballAPI) FetchByName(name string) (domain.Character, error) {
	u, err := url.Parse(d.baseURL + "/characters")
	if err != nil {
		return domain.Character{}, fmt.Errorf("%w: %s", ErrInvalidURL, err)
	}
	q := u.Query()
	q.Set("name", name)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return domain.Character{}, fmt.Errorf("%w: %s", ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Character{}, fmt.Errorf("%w: status %d", ErrUnexpectedStatus, resp.StatusCode)
	}

	var results []characterResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return domain.Character{}, fmt.Errorf("%w: %s", ErrDecodeResponse, err)
	}
	if len(results) == 0 {
		return domain.Character{}, fmt.Errorf("%w: %q", ErrCharacterNotFound, name)
	}

	cr := results[0]
	return domain.Character{
		ID:          cr.ID.String(),
		Affiliation: cr.Affiliation,
		DeletedAt:   cr.DeletedAt,
		Description: cr.Description,
		Gender:      cr.Gender,
		Image:       cr.Image,
		Ki:          cr.Ki,
		MaxKi:       cr.MaxKi,
		Name:        cr.Name,
	}, nil
}
