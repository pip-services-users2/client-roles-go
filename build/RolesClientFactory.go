package build

import (
	clients1 "github.com/pip-services-users2/client-roles-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type RolesClientFactory struct {
	cbuild.Factory
}

func NewRolesClientFactory() *RolesClientFactory {
	c := &RolesClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-roles", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("service-roles", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-roles", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-roles", "client", "grpc", "*", "1.0")
	commandableGrpcClientDescriptor := cref.NewDescriptor("service-roles", "client", "commandable-grpc", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-roles", "client", "mock", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewRolesNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewRolesDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewRolesCommandableHttpClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewRoleGrpcClientV1)
	c.RegisterType(commandableGrpcClientDescriptor, clients1.NewRolesCommandableGrpcClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewRolesMockClientV1)

	return c
}
