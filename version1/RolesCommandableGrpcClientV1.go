package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type RolesCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewRolesCommandableGrpcClientV1() *RolesCommandableGrpcClientV1 {
	return NewRolesCommandableGrpcClientV1WithConfig(nil)
}

func NewRolesCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *RolesCommandableGrpcClientV1 {
	c := &RolesCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/roles"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *RolesCommandableGrpcClientV1) GetRolesByFilter(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*UserRolesV1], err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.get_roles_by_filter")
	defer timing.EndTiming(ctx, err)

	params := data.NewEmptyStringValueMap()
	c.AddFilterParams(params, filter)
	c.AddPagingParams(params, paging)

	res, err := c.CallCommand(ctx, "get_roles_by_filter", correlationId, data.NewAnyValueMapFromValue(params.Value()))

	if err != nil {
		return *data.NewEmptyDataPage[*UserRolesV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*UserRolesV1]](res, correlationId)
}

func (c *RolesCommandableGrpcClientV1) GetRolesById(ctx context.Context, correlationId string, userId string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.get_roles_by_id")
	defer timing.EndTiming(ctx, err)

	res, err := c.CallCommand(ctx, "get_roles_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"user_id", userId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesCommandableGrpcClientV1) SetRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.set_roles")
	defer timing.EndTiming(ctx, err)

	res, err := c.CallCommand(ctx, "set_roles", correlationId, data.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesCommandableGrpcClientV1) GrantRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.grant_roles")
	defer timing.EndTiming(ctx, err)

	res, err := c.CallCommand(ctx, "grant_roles", correlationId, data.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesCommandableGrpcClientV1) RevokeRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.revoke_roles")
	defer timing.EndTiming(ctx, err)

	res, err := c.CallCommand(ctx, "revoke_roles", correlationId, data.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[[]string](res, correlationId)
}

func (c *RolesCommandableGrpcClientV1) Authorize(ctx context.Context, correlationId string, userId string, roles []string) (result bool, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.authorize")
	defer timing.EndTiming(ctx, err)

	res, err := c.CallCommand(ctx, "authorize", correlationId, data.NewAnyValueMapFromTuples(
		"user_id", userId,
		"roles", roles,
	))

	if err != nil {
		return false, err
	}

	return clients.HandleHttpResponse[bool](res, correlationId)
}
