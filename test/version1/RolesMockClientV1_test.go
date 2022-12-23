package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-roles-go/version1"
)

type rolesMockClientV1Test struct {
	client  *version1.RolesMockClientV1
	fixture *RolesClientFixtureV1
}

func newRolesMockClientV1Test() *rolesMockClientV1Test {
	return &rolesMockClientV1Test{}
}

func (c *rolesMockClientV1Test) setup(t *testing.T) *RolesClientFixtureV1 {

	c.client = version1.NewRolesMockClientV1()
	c.fixture = NewRolesClientFixtureV1(c.client)
	return c.fixture
}

func (c *rolesMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMockGetAndSetRoles(t *testing.T) {
	c := newRolesMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGetAndSetRoles(t)
}

func TestMockGrantAndRevokeRoles(t *testing.T) {
	c := newRolesMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestGrantAndRevokeRoles(t)
}

func TestMockAuthorize(t *testing.T) {
	c := newRolesMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestAuthorize(t)
}
