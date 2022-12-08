package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type RolesHttpCommandableClientV1 struct {
	*cclients.CommandableHttpClient
}

func NewRolesHttpCommandableClientV1() *RolesHttpCommandableClientV1 {
	c := &RolesHttpCommandableClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/roles"),
	}
	return c
}

func (c *RolesHttpCommandableClientV1) GetRolesByFilter(ctx context.Context, correlationId string, filter data.FilterParams,
	paging data.PagingParams) (result cdata.DataPage[*UserRolesV1], err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_roles_by_filter", correlationId, params)
	if err != nil {
		return *cdata.NewEmptyDataPage[*UserRolesV1](), err
	}

	return cclients.HandleHttpResponse[cdata.DataPage[*UserRolesV1]](res, correlationId)
}

func (c *RolesHttpCommandableClientV1) GetRolesById(ctx context.Context, correlationId string, userId string) (result []string, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
	)

	res, err := c.CallCommand(ctx, "get_roles_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesHttpCommandableClientV1) SetRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(ctx, "set_roles", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesHttpCommandableClientV1) GrantRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(ctx, "grant_roles", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesHttpCommandableClientV1) RevokeRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(ctx, "revoke_roles", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesHttpCommandableClientV1) Authorize(ctx context.Context, correlationId string, userId string, roles []string) (result bool, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	)

	res, err := c.CallCommand(ctx, "authorize", correlationId, params)
	if err != nil {
		return false, err
	}

	return cclients.HandleHttpResponse[bool](res, correlationId)
}
