package example_for_review

import (
	"context"
	"fmt"
)

type Processor struct {
	schoolGoExternalClient SchoolGoExternalClient
	orderPersistRepository OrderPersistRepository
}

func NewProcessor(
	schoolGoExternalClient SchoolGoExternalClient,
	orderPersistRepository OrderPersistRepository,
) *Processor {
	return &Processor{
		schoolGoExternalClient: schoolGoExternalClient,
		orderPersistRepository: orderPersistRepository,
	}
}

func (p *Processor) Process(ctx context.Context, params ProcessorParams) (*apiModels.OrderID, error) {
	order, err := p.orderPersistRepository.InsertOrder(
		ctx,
		params.OrderTitle.Title,
	)

	if err != nil {
		return nil, fmt.Errorf("can't insert the order: %w", err)
	}

	if order == nil {
		return nil, nil
	}

	_, err = p.schoolGoExternalClient.SendOrder(ctx, transformOrderEntityToGoSchoolExternalModel(order), order.ID)

	sentToExternal := orderRepository.SentToExternalSuccess
	if err != nil {
		sentToExternal = orderRepository.SentToExternalFailure
	}

	err = p.orderPersistRepository.UpdateSentToExternal(
		ctx,
		order.ID,
		sentToExternal,
	)
	if err != nil {
		return nil, fmt.Errorf("can't update order after adding: %w", err)
	}

	return transformOrderIdEntity(order.ID), nil
}
