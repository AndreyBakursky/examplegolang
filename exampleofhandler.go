package example_for_review

import "fmt"

func (srv *Service) OrdersCreateHandler(params apiOrders.OrdersCreateParams) middleware.Responder {
	processor := create.NewProcessor(
		srv.container.CreateExternalServiceClient(),
		srv.container.CreateOrderPersistRepository(),
	)

	response, err := processor.Process(
		params.HTTPRequest.Context(),
		create.ProcessorParams{
			OrderTitle: &apiModels.OrderTitle{
				Title: params.Body.Title,
			},
		},
	)

	if err != nil {
		logger.Error(params.HTTPRequest.Context(), err)

		return apiOrders.NewOrdersCreateInternalServerError().WithPayload(&apiModels.Error{
			Code:    500,
			Message: fmt.Sprintf("Internal server error: %s", err),
		})
	}

	if response == nil {
		return apiOrders.NewOrdersCreateBadRequest().WithPayload(&apiModels.Error{
			Code:    400,
			Message: "Bad request",
		})
	}

	return apiOrders.NewOrdersCreateOK().WithPayload(response)
}
