package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IRolesClientV1 interface {
	GetRolesByFilter(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*UserRolesV1], err error)

	GetRolesById(ctx context.Context, correlationId string, userId string) (result []string, err error)

	SetRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error)

	GrantRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error)

	RevokeRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error)

	Authorize(ctx context.Context, correlationId string, userId string, roles []string) (result bool, err error)
}
