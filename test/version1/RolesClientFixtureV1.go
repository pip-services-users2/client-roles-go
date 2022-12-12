package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-users2/client-roles-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type RolesClientFixtureV1 struct {
	Client version1.IRolesClientV1
}

var ROLES = []string{"Role 1", "Role 2", "Role 3"}

func NewRolesClientFixtureV1(client version1.IRolesClientV1) *RolesClientFixtureV1 {
	return &RolesClientFixtureV1{
		Client: client,
	}
}

func (c *RolesClientFixtureV1) clear() {
	page, _ := c.Client.GetRolesByFilter(context.Background(), "", data.NewEmptyFilterParams(), data.NewEmptyPagingParams())
	for _, roles := range page.Data {
		c.Client.RevokeRoles(context.Background(), "", roles.Id, roles.Roles)
	}
}

func (c *RolesClientFixtureV1) TestGetAndSetRoles(t *testing.T) {
	c.clear()
	defer c.clear()

	// Update party roles
	roles, err := c.Client.SetRoles(context.Background(), "", "1", ROLES)
	assert.Nil(t, err)

	assert.True(t, len(roles) == 3)

	// Read and check party roles
	roles, err = c.Client.GetRolesById(context.Background(), "", "1")
	assert.Nil(t, err)

	assert.True(t, len(roles) == 3)

	// Get roles by filter
	page, err1 := c.Client.GetRolesByFilter(context.Background(), "", data.NewFilterParamsFromTuples("roles", ROLES), data.NewEmptyPagingParams())
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 1)
}

func (c *RolesClientFixtureV1) TestGrantAndRevokeRoles(t *testing.T) {
	c.clear()
	defer c.clear()

	// Grant roles first time
	roles, err := c.Client.GrantRoles(context.Background(), "", "1", []string{"Role 1"})
	assert.Nil(t, err)

	assert.Len(t, roles, 1)
	assert.Contains(t, roles, "Role 1")

	// Grant roles second time
	roles, err = c.Client.GrantRoles(context.Background(), "", "1", []string{"Role 1", "Role 2", "Role 3"})
	assert.Nil(t, err)

	assert.Len(t, roles, 3)
	assert.Contains(t, roles, "Role 1")
	assert.Contains(t, roles, "Role 2")
	assert.Contains(t, roles, "Role 3")

	// Revoke roles first time
	roles, err = c.Client.RevokeRoles(context.Background(), "", "1", []string{"Role 1"})
	assert.Nil(t, err)

	assert.Len(t, roles, 2)
	assert.Contains(t, roles, "Role 2")
	assert.Contains(t, roles, "Role 3")

	// Get roles
	roles, err = c.Client.GetRolesById(context.Background(), "", "1")
	assert.Nil(t, err)

	assert.True(t, len(roles) == 2)
	assert.Contains(t, roles, "Role 2")
	assert.Contains(t, roles, "Role 3")

	// Revoke roles second time
	roles, err = c.Client.RevokeRoles(context.Background(), "", "1", []string{"Role 1", "Role 2"})
	assert.Nil(t, err)

	assert.Len(t, roles, 1)
	assert.Contains(t, roles, "Role 3")
}

func (c *RolesClientFixtureV1) TestAuthorize(t *testing.T) {
	c.clear()
	defer c.clear()

	// Grant roles
	roles, err := c.Client.GrantRoles(context.Background(), "", "1", []string{"Role 1", "Role 2"})
	assert.Nil(t, err)

	assert.Len(t, roles, 2)

	// Authorize positively
	auth, err1 := c.Client.Authorize(context.Background(), "", "1", []string{"Role 1"})
	assert.Nil(t, err1)

	assert.True(t, auth)

	// Authorize negatively
	auth, err1 = c.Client.Authorize(context.Background(), "", "1", []string{"Role 2", "Role 3"})
	assert.Nil(t, err1)

	assert.False(t, auth)
}
