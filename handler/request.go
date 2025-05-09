package handler

import "fmt"

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(name string) error {
	return fmt.Errorf("%s is required", name)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role")
	}

	if r.Company == "" {
		return errParamIsRequired("company")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote")
	}

	if r.Location == "" {
		return errParamIsRequired("location")
	}

	if r.Link == "" {
		return errParamIsRequired("link")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary")
	}

	return nil
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy

	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	// If none of the fields ere provided, return falsy
	return fmt.Errorf("at least one field must be provided")

}
