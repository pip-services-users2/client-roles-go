package build

import (
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
	clients1 "github.com/service-users/client-roles-go/version1"
)

type RolesClientFactory struct {
	cbuild.Factory
}

func NewRolesClientFactory() *RolesClientFactory {
	c := &RolesClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	// nullClientDescriptor := cref.NewDescriptor("service-roles", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("service-roles", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-roles", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-roles", "client", "grpc", "*", "1.0")
	memoryClientDescriptor := cref.NewDescriptor("service-roles", "client", "memory", "*", "1.0")

	// c.RegisterType(nullClientDescriptor, clients1.NewRolesNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewRolesDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewRolesHttpCommandableClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewRoleGrpcClientV1)
	c.RegisterType(memoryClientDescriptor, clients1.NewRolesMemoryClientV1)

	return c
}
