package view_models

type CurrentUserVm struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// TODO : call rpc svc-file-storage
func (vm CurrentUserVm) Build(id, name, email string) CurrentUserVm {
	return CurrentUserVm{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

//Factory interfaces to manage build
type FactoryCurrentUser interface {
	AdminCurrentUser(vm AdminVm, credentialID string) CurrentUserVm
}

type FactoryAction struct {
	CurrentUserVm
}

func NewFactoryCurrentUser(vm CurrentUserVm) FactoryCurrentUser {
	return &FactoryAction{CurrentUserVm: vm}
}

func (factory FactoryAction) AdminCurrentUser(vm AdminVm, credentialID string) CurrentUserVm {
	return factory.CurrentUserVm.Build(credentialID, vm.Name, vm.Email)
}
