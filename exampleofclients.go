package example_for_review

func (c *Container) CreateExternalServiceClient() *school_go_external.Client {
	transportConfig := client.DefaultTransportConfig()
	transportConfig.Host = c.config.ExternalService.Host

	loggedTransport := swagger_client_transport.New(
		c.metrics.client.schoolGoExternal,
		transportConfig.Host,
		transportConfig.BasePath,
		transportConfig.Schemes,
	)

	return school_go_external.NewSchoolGoExternalClient(
		client.New(loggedTransport, nil),
		c.config.ExternalService.Timeout,
		c.metrics.client.schoolGoExternal,
	)
}
