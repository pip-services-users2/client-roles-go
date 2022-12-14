package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-roles-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type rolesGrpcClientV1Test struct {
	client  *version1.RoleGrpcClientV1
	fixture *RolesClientFixtureV1
}

func newRolesGrpcClientV1Test() *rolesGrpcClientV1Test {
	return &rolesGrpcClientV1Test{}
}

func (c *rolesGrpcClientV1Test) setup(t *testing.T) *RolesClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewRoleGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewRolesClientFixtureV1(c.client)

	return c.fixture
}

func (c *rolesGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGetAndSetRoles(t *testing.T) {
	c := newRolesGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGetAndSetRoles(t)
}

func TestGrantAndRevokeRoles(t *testing.T) {
	c := newRolesGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGrantAndRevokeRoles(t)
}

func TestAuthorize(t *testing.T) {
	c := newRolesGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestAuthorize(t)
}
