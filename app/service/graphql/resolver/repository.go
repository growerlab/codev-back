package resolver

import (
	"context"

	repoModel "github.com/growerlab/backend/app/model/repository"
	"github.com/growerlab/backend/app/service"
	"github.com/growerlab/backend/app/service/repository"
)

func (r *queryResolver) Repositories(ctx context.Context, namespaceID int) ([]*repoModel.Repository, error) {
	return repository.ListRepositories(ctx, int64(namespaceID))
}

func (r *mutationResolver) CreateRepository(ctx context.Context, input service.NewRepositoryPayload) (*service.Result, error) {
	ok, err := repository.CreateRepository(ctx, &input)
	return &service.Result{Ok: ok}, err
}
