package role

import (
	"context"
	"simple/internal/global"
	roleDto "simple/internal/types/dto/role"
	"simple/internal/types/entity"
	"simple/internal/types/query"
	"simple/pkg/consts"
	"simple/pkg/logger"
	"simple/pkg/resp"

	"go.uber.org/zap"
	"gorm.io/gorm"
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

	roleService struct{}
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
		localRole = &roleService{}
	}
	return localRole
}

// CreateRole 创建角色
func (s *roleService) CreateRole(ctx context.Context, req *roleDto.CreateRoleReq) error {
	// 1. 检查角色名称是否已存在
	nameCount, err := query.Role.WithContext(ctx).Where(query.Role.Name.Eq(req.Name)).Count()
	if err != nil {
		logger.Error("检查角色名称是否存在失败", zap.String("name", req.Name), zap.Error(err))
		return consts.ErrServer
	}
	if nameCount > 0 {
		return consts.ErrRoleNameExists
	}

	// 2. 检查角色编码是否已存在
	codeCount, err := query.Role.WithContext(ctx).Where(query.Role.Code.Eq(req.Code)).Count()
	if err != nil {
		logger.Error("检查角色编码是否存在失败", zap.String("code", req.Code), zap.Error(err))
		return consts.ErrServer
	}
	if codeCount > 0 {
		return consts.ErrRoleCodeExists
	}

	// 3. 使用事务创建角色
	return global.DB.Transaction(func(tx *gorm.DB) error {
		r := &entity.Role{
			Name:          req.Name,
			Code:          req.Code,
			DefaultRouter: req.DefaultRouter,
			Status:        req.Status,
			Remark:        req.Remark,
			Sort:          req.Sort,
		}

		// 使用事务中的DB进行创建
		q := query.Use(tx).Role
		if err := q.WithContext(ctx).Create(r); err != nil {
			logger.Error("创建角色失败", zap.Any("role", r), zap.Error(err))
			return consts.ErrServer
		}

		return nil
	})
}

// UpdateRole 更新角色
func (s *roleService) UpdateRole(ctx context.Context, req *roleDto.UpdateRoleReq) error {
	// 1. 检查角色是否存在
	oldRole, err := query.Role.WithContext(ctx).Where(query.Role.ID.Eq(req.ID)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return consts.ErrRoleNotFound
		}
		logger.Error("查询角色失败", zap.Int64("id", req.ID), zap.Error(err))
		return consts.ErrServer
	}

	// 2. 检查角色名称是否与其他角色重复
	if oldRole.Name != req.Name {
		nameCount, err := query.Role.WithContext(ctx).
			Where(query.Role.Name.Eq(req.Name)).
			Where(query.Role.ID.Neq(req.ID)).
			Count()
		if err != nil {
			logger.Error("检查角色名称是否重复失败", zap.String("name", req.Name), zap.Error(err))
			return consts.ErrServer
		}
		if nameCount > 0 {
			return consts.ErrRoleNameExists
		}
	}

	// 3. 检查角色编码是否与其他角色重复
	if oldRole.Code != req.Code {
		codeCount, err := query.Role.WithContext(ctx).
			Where(query.Role.Code.Eq(req.Code)).
			Where(query.Role.ID.Neq(req.ID)).
			Count()
		if err != nil {
			logger.Error("检查角色编码是否重复失败", zap.String("code", req.Code), zap.Error(err))
			return consts.ErrServer
		}
		if codeCount > 0 {
			return consts.ErrRoleCodeExists
		}
	}

	// 4. 使用事务更新角色
	return global.DB.Transaction(func(tx *gorm.DB) error {
		r := &entity.Role{
			ID:            req.ID,
			Name:          req.Name,
			Code:          req.Code,
			DefaultRouter: req.DefaultRouter,
			Status:        req.Status,
			Remark:        req.Remark,
			Sort:          req.Sort,
		}

		// 使用事务中的DB进行更新
		q := query.Use(tx).Role
		_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Updates(r)
		if err != nil {
			logger.Error("更新角色失败", zap.Any("role", r), zap.Error(err))
			return consts.ErrServer
		}

		return nil
	})
}

// DeleteRole 批量删除角色
func (s *roleService) DeleteRole(ctx context.Context, req *roleDto.DeleteRoleReq) error {
	if len(req.Ids) == 0 {
		return consts.ErrInvalidParam
	}

	// 1. 检查待删除的角色是否都存在
	roleCount, err := query.Role.WithContext(ctx).Where(query.Role.ID.In(req.Ids...)).Count()
	if err != nil {
		logger.Error("检查角色是否存在失败", zap.Any("Ids", req.Ids), zap.Error(err))
		return consts.ErrServer
	}
	if roleCount != int64(len(req.Ids)) {
		return consts.ErrRoleNotFound
	}

	// 2. 检查是否包含超级管理员角色
	superAdminExists, err := query.Role.WithContext(ctx).
		Where(query.Role.ID.In(req.Ids...)).
		Where(query.Role.Code.Eq(SuperAdminCode)).
		Count()
	if err != nil {
		logger.Error("检查是否包含超级管理员失败", zap.Any("Ids", req.Ids), zap.Error(err))
		return consts.ErrServer
	}
	if superAdminExists > 0 {
		return consts.ErrRoleSuperAdmin
	}

	// 3. 使用事务批量删除角色
	return global.DB.Transaction(func(tx *gorm.DB) error {
		q := query.Use(tx).Role
		// 删除角色
		_, err := q.WithContext(ctx).Where(q.ID.In(req.Ids...)).Delete()
		if err != nil {
			logger.Error("批量删除角色失败", zap.Any("Ids", req.Ids), zap.Error(err))
			return consts.ErrServer
		}

		// 同时也可以删除关联数据，例如角色-用户关联
		uq := query.Use(tx).UserRole
		_, err = uq.WithContext(ctx).Where(uq.RoleID.In(req.Ids...)).Delete()
		if err != nil {
			logger.Error("批量删除角色用户关联失败", zap.Any("roleIds", req.Ids), zap.Error(err))
			return consts.ErrServer
		}

		return nil
	})
}

// GetRole 获取角色
func (s *roleService) GetRole(ctx context.Context, req *roleDto.GetRoleReq) (*entity.Role, error) {
	role, err := query.Role.WithContext(ctx).Where(query.Role.ID.Eq(req.ID)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, consts.ErrRoleNotFound
		}
		logger.Error("查询角色失败", zap.Int64("id", req.ID), zap.Error(err))
		return nil, consts.ErrServer
	}
	return role, nil
}

// ListRole 角色列表
func (s *roleService) ListRole(ctx context.Context, req *roleDto.ListRoleReq) (*resp.PageResp, error) {
	q := query.Role.WithContext(ctx)

	// 条件查询
	if req.Name != nil {
		q = q.Where(query.Role.Name.Like("%" + *req.Name + "%"))
	}
	if req.Code != nil {
		q = q.Where(query.Role.Code.Like("%" + *req.Code + "%"))
	}
	if req.Status != nil {
		q = q.Where(query.Role.Status.Eq(*req.Status))
	}

	// 分页查询
	result, count, err := q.Order(query.Role.Sort, query.Role.ID.Desc()).
		FindByPage((req.Page-1)*req.Size, req.Size)
	if err != nil {
		logger.Error("查询角色列表失败", zap.Any("req", req), zap.Error(err))
		return nil, consts.ErrServer
	}

	return &resp.PageResp{
		Total: count,
		List:  result,
	}, nil
}

// ListRoleItem 角色名列表
func (s *roleService) ListRoleItem(ctx context.Context) ([]*roleDto.ListRoleItemResp, error) {
	result, err := query.Role.WithContext(ctx).
		Where(query.Role.Status.Eq(1)). // 只查询启用的角色
		Order(query.Role.Sort).
		Find()
	if err != nil {
		logger.Error("查询角色列表失败", zap.Error(err))
		return nil, consts.ErrServer
	}

	list := make([]*roleDto.ListRoleItemResp, 0, len(result))
	for _, r := range result {
		list = append(list, &roleDto.ListRoleItemResp{
			ID:   r.ID,
			Name: r.Name,
		})
	}
	return list, nil
}
