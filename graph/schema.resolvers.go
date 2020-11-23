package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DavidNazareno/h_prueba/config"
	"github.com/DavidNazareno/h_prueba/graph/generated"
	"github.com/DavidNazareno/h_prueba/graph/model"
	"github.com/DavidNazareno/h_prueba/services"
	"github.com/DavidNazareno/h_prueba/utils"
)

func (r *mutationResolver) AddClient(ctx context.Context) (*model.Client, error) {
	var c services.Client
	c.Ticket = utils.Generate(4)
	result, err := c.Add()
	if err != nil {
		return nil, err
	}

	return &model.Client{ID: strconv.Itoa(int(result.ID)), Ticket: c.Ticket}, nil
}

func (r *mutationResolver) AddRequest(ctx context.Context, ticket model.NewRequest) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) GenerateRequest(ctx context.Context, ticket string) (string, error) {
	var c services.Client
	c.Ticket = ticket
	cli, validate, err := c.Check()

	if err != nil || validate == false {
		return "", err
	}

	randomTech, err := services.GetRandomTech()

	if err != nil {
		return "", err
	}
	req, err := services.NewRequest(*cli)

	if err != nil {
		return "", err
	}

	_, err = services.CreateOrder(req.Token, *randomTech)

	if err != nil {
		return "", err
	}

	return "localhost:" + config.PORT + "/status?token=" + req.Token + "&status=" + strconv.Itoa(req.Status) + "&tech=" + randomTech.Name, nil
}

func (r *queryResolver) GetRequest(ctx context.Context, token string) (*model.Request, error) {
	result, err := services.GetRequest(token)

	if err != nil {
		return nil, err
	}

	return &model.Request{Token: result.Token, Status: result.Status, Score: &result.Score}, nil
}

func (r *queryResolver) GetTechOrders(ctx context.Context, id string) (string, error) {
	result, err := services.GetTechOrder(id)

	if err != nil {
		return "", err
	}
	return strconv.FormatInt(*result, 10), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
