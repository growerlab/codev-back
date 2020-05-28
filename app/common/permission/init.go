package permission

import (
	context2 "github.com/growerlab/backend/app/common/context"
	"github.com/growerlab/backend/app/common/userdomain"
	"github.com/growerlab/backend/app/model/db"
)

var permHub *Hub

func InitPermission() error {
	permHub = NewPermissionHub(db.DB, db.PermissionDB)

	if err := initRules(); err != nil {
		return err
	}
	if err := initUserDomains(); err != nil {
		return err
	}
	if err := initContexts(); err != nil {
		return err
	}
	return nil
}

func initUserDomains() error {
	userDomains := []UserDomainDelegate{
		&userdomain.SuperAdmin{},
		&userdomain.Person{},
		&userdomain.RepositoryOwner{},
		&userdomain.Visitor{},
	}
	return permHub.RegisterUserDomains(userDomains)
}

func initContexts() error {
	contexts := make([]ContextDelegate, 0)
	contexts = append(contexts, &context2.Repository{})
	return permHub.RegisterContexts(contexts)
}

func initRules() error {
	rules := []*Rule{
		{
			Code:                  ViewRepository,
			ConstraintUserDomains: []int{userdomain.TypePerson},
			BuiltInUserDomains:    []int{userdomain.TypeRepositoryOwner},
		},
		{
			Code:                  CloneRepository,
			ConstraintUserDomains: []int{userdomain.TypePerson},
			BuiltInUserDomains:    []int{userdomain.TypeRepositoryOwner},
		},
		{
			Code:                  PushRepository,
			ConstraintUserDomains: []int{userdomain.TypePerson},
			BuiltInUserDomains:    []int{userdomain.TypeRepositoryOwner},
		},
	}
	return permHub.RegisterRules(rules)
}
