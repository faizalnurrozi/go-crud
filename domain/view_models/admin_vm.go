package view_models

import (
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type AdminVm struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//CreatedAt  string       `json:"created_at"`
	//UpdatedAt  string       `json:"updated_at"`
}

func (vm AdminVm) Build(model models.Admin) AdminVm {
	return AdminVm{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Credential.Email,
	}
}
