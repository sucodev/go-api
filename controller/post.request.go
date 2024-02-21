package controller

import (
	"fmt"
)

type CreatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Public      *bool  `json:"public"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required ", name, typ)
}

func (r *CreatePostRequest) Validate() error {

	if r.Title == "" && r.Description == "" && r.Company == "" && r.Public == nil {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Public == nil {
		return errParamIsRequired("public", "bool")
	}
	return nil
}
