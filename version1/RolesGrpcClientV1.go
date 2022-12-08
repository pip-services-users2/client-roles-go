package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
	"github.com/service-users/client-roles-go/protos"
)

type RoleGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewRoleGrpcClientV1() *RoleGrpcClientV1 {
	return &RoleGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("roles_v1.Roles"),
	}
}

func (c *RoleGrpcClientV1) GetRolesByFilter(ctx context.Context, correlationId string, filter data.FilterParams,
	paging data.PagingParams) (result data.DataPage[*UserRolesV1], err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.get_roles_by_filter")
	defer timing.EndTiming(ctx, err)

	req := &protos.RolesPageRequest{
		CorrelationId: correlationId,
	}

	req.Filter = filter.Value()

	req.Paging = &protos.PagingParams{
		Skip:  paging.GetSkip(0),
		Take:  (int32)(paging.GetTake(100)),
		Total: paging.Total,
	}

	reply := new(protos.RolesPageReply)
	err = c.CallWithContext(ctx, "get_roles_by_filter", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*UserRolesV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*UserRolesV1](), err
	}

	result = toUserRolesPage(reply.Page)

	return result, nil
}

func (c *RoleGrpcClientV1) GetRolesById(ctx context.Context, correlationId string, userId string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.get_roles_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.RoleIdRequest{
		CorrelationId: correlationId,
		UserId:        userId,
	}

	reply := new(protos.RolesReply)
	err = c.CallWithContext(ctx, "get_roles_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) SetRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.set_roles")
	defer timing.EndTiming(ctx, err)

	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.RolesReply)
	err = c.CallWithContext(ctx, "set_roles", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) GrantRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.grant_roles")
	defer timing.EndTiming(ctx, err)

	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.RolesReply)
	err = c.CallWithContext(ctx, "grant_roles", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) RevokeRoles(ctx context.Context, correlationId string, userId string, roles []string) (result []string, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.revoke_roles")
	defer timing.EndTiming(ctx, err)

	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.RolesReply)
	err = c.CallWithContext(ctx, "revoke_roles", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = reply.Roles

	return result, nil
}

func (c *RoleGrpcClientV1) Authorize(ctx context.Context, correlationId string, userId string, roles []string) (result bool, err error) {
	timing := c.Instrument(ctx, correlationId, "roles_v1.authorize")
	defer timing.EndTiming(ctx, err)

	req := &protos.RolesRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		Roles:         roles,
	}

	reply := new(protos.AuthorizeReply)
	err = c.CallWithContext(ctx, "authorize", correlationId, req, reply)
	if err != nil {
		return false, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return false, err
	}

	result = reply.Authorized

	return result, nil
}
