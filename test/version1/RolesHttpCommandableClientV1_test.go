package test_version1

import (
	"context"
	"os"
	"testing"

	version1 "github.com/pip-services-users2/client-roles-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type rolesHttpCommandableClientV1Test struct {
	client  *version1.RolesHttpCommandableClientV1
	fixture *RolesClientFixtureV1
}

func newRolesHttpCommandableClientV1Test() *rolesHttpCommandableClientV1Test {
	return &rolesHttpCommandableClientV1Test{}
}

func (c *rolesHttpCommandableClientV1Test) setup(t *testing.T) *RolesClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewRolesHttpCommandableClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewRolesClientFixtureV1(c.client)

	return c.fixture
}

func (c *rolesHttpCommandableClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestHttpGetAndSetRoles(t *testing.T) {
	c := newRolesHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGetAndSetRoles(t)
}

func TestHttpGrantAndRevokeRoles(t *testing.T) {
	c := newRolesHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGrantAndRevokeRoles(t)
}

func TestHttpAuthorize(t *testing.T) {
	c := newRolesHttpCommandableClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestAuthorize(t)
}
