package view_models

import "github.com/faizalnurrozi/go-crud/domain/models"

type CredentialVm struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	Password string `json:"password"`
}

func NewCredentialVm() CredentialVm {
	return CredentialVm{}
}

func (vm CredentialVm) Build(model *models.Credential) CredentialVm {
	return CredentialVm{
		ID:       model.ID,
		Email:    model.Email,
		IsActive: model.IsActive,
		Password: model.Password,
	}
}
