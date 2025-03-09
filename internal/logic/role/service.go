package role

import (
	"context"
	roleDto "simple/internal/types/dto/role"
	"simple/internal/types/entity"
	"simple/pkg/resp"
)

/*
   @NAME    : service
   @author  : 清风
   @desc    :
   @time    : 2025/3/7 22:36
*/

type (
	IRoleService interface {
		// CreateRole 创建角色
		CreateRole(ctx context.Context, req *roleDto.CreateRoleReq) error
		// UpdateRole 更新角色
		UpdateRole(ctx context.Context, req *roleDto.UpdateRoleReq) error
		// DeleteRole 删除角色
		DeleteRole(ctx context.Context, req *roleDto.DeleteRoleReq) error
		// GetRole 获取角色
		GetRole(ctx context.Context, req *roleDto.GetRoleReq) (*entity.Role, error)
		// ListRole 角色列表
		ListRole(ctx context.Context, req *roleDto.ListRoleReq) (*resp.PageResp, error)
		// ListRoleItem 角色名列表用于创建管理员分配角色
		ListRoleItem(ctx context.Context) ([]*roleDto.ListRoleItemResp, error)
	}
)

const (
	SuperAdminCode = "super-admin"
)

var (
	localRole IRoleService
)

// Role 获取角色服务实例
func Role() IRoleService {
	if localRole == nil {
		localRole = newLogic()
	}
	return localRole
}
