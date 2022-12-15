package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type RolesNullClientV1 struct {
}

func NewRolesNullClientV1() *RolesNullClientV1 {
	return &RolesNullClientV1{}
}

func (c *RolesNullClientV1) GetRolesByFilter(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*UserRolesV1], err error) {
	return *data.NewEmptyDataPage[*UserRolesV1](), nil
}

func (c *RolesNullClientV1) GetRolesById(ctx context.Context, correlationId string, userId string) (result []string, err error) {
	return make([]string, 0), nil
}

func (c *RolesNullClientV1) SetRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	return roles, nil
}

func (c *RolesNullClientV1) GrantRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	return roles, nil
}

func (c *RolesNullClientV1) RevokeRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	return make([]string, 0), nil
}

func (c *RolesNullClientV1) Authorize(ctx context.Context, correlationId string, userId string, roles []string) (result bool, err error) {
	return true, nil
}
