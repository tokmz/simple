package role

import (
	"context"
	"errors"
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
   @time    : 2025/3/9 20:32
*/

type logic struct{}

func newLogic() *logic {
	return &logic{}
}

// CreateRole 创建角色
func (s *logic) CreateRole(ctx context.Context, req *roleDto.CreateRoleReq) error {
	// 使用事务进行所有操作，确保原子性
	return global.Query.Transaction(func(tx *query.Query) error {
		dao := tx.Role
		do := dao.WithContext(ctx)

		// 1. 检查角色名称是否已存在
		r, err := do.Where(dao.Name.Eq(req.Name)).Select(dao.ID).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("检查角色名称是否存在失败", zap.String("name", req.Name), zap.Error(err))
			return consts.ErrServer
		}

		if r.ID != 0 {
			return consts.ErrRoleNameExists
		}

		// 2. 检查角色编码是否已存在
		r, err = do.Where(dao.Code.Eq(req.Code)).Select(dao.ID).First()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("检查角色编码是否存在失败", zap.String("code", req.Code), zap.Error(err))
			return consts.ErrServer
		}
		if r.ID != 0 {
			return consts.ErrRoleCodeExists
		}

		// 3. 创建角色
		r = &entity.Role{
			Name:          req.Name,
			Code:          req.Code,
			DefaultRouter: req.DefaultRouter,
			Status:        req.Status,
			Remark:        req.Remark,
			Sort:          req.Sort,
		}

		// 使用事务中的DB进行创建
		if err := do.Create(r); err != nil {
			logger.Error("创建角色失败", zap.Any("role", r), zap.Error(err))
			return consts.ErrServer
		}

		return nil
	})
}

// UpdateRole 更新角色
func (s *logic) UpdateRole(ctx context.Context, req *roleDto.UpdateRoleReq) error {
	// 使用事务进行所有操作，确保原子性
	return global.Query.Transaction(func(tx *query.Query) error {
		dao := tx.Role
		do := dao.WithContext(ctx)

		// 1. 检查角色是否存在
		oldRole, err := do.Where(dao.ID.Eq(req.ID)).First()
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return consts.ErrRoleNotFound
			}
			logger.Error("查询角色失败", zap.Int64("id", req.ID), zap.Error(err))
			return consts.ErrServer
		}

		// 2. 检查角色名称是否与其他角色重复
		if oldRole.Name != req.Name {
			r, err := do.Where(dao.Name.Eq(req.Name)).Where(dao.ID.Neq(req.ID)).Select(dao.ID).First()
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				logger.Error("检查角色名称是否重复失败", zap.String("name", req.Name), zap.Error(err))
				return consts.ErrServer
			}
			if r.ID != 0 {
				return consts.ErrRoleNameExists
			}
		}

		// 3. 检查角色编码是否与其他角色重复
		if oldRole.Code != req.Code {
			r, err := do.Where(dao.Code.Eq(req.Code)).Where(dao.ID.Neq(req.ID)).Select(dao.ID).First()
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				logger.Error("检查角色编码是否重复失败", zap.String("code", req.Code), zap.Error(err))
				return consts.ErrServer
			}
			if r.ID != 0 {
				return consts.ErrRoleCodeExists
			}
		}

		// 4. 更新角色
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
		_, err = do.Where(dao.ID.Eq(req.ID)).Updates(r)
		if err != nil {
			logger.Error("更新角色失败", zap.Any("role", r), zap.Error(err))
			return consts.ErrServer
		}

		return nil
	})
}

// DeleteRole 批量删除角色
func (s *logic) DeleteRole(ctx context.Context, req *roleDto.DeleteRoleReq) error {
	if len(req.Ids) == 0 {
		return nil
	}

	// 使用事务进行所有操作，确保数据一致性
	return global.Query.Transaction(func(tx *query.Query) error {
		dao := tx.Role
		do := dao.WithContext(ctx)

		// 1. 检查角色是否存在并且不包含超级管理员
		roles, err := do.Where(dao.ID.In(req.Ids...)).Find()
		if err != nil {
			logger.Error("查询角色失败", zap.Any("Ids", req.Ids), zap.Error(err))
			return consts.ErrServer
		}

		// 检查是否所有角色都存在
		if len(roles) != len(req.Ids) {
			return consts.ErrRoleNotFound
		}

		// 检查是否包含超级管理员
		for _, role := range roles {
			if role.Code == SuperAdminCode {
				return consts.ErrRoleSuperAdmin
			}
		}

		// 2. 删除角色关联的用户信息
		_, err = tx.UserRole.WithContext(ctx).
			Where(query.UserRole.RoleID.In(req.Ids...)).Delete()
		if err != nil {
			logger.Error("删除角色关联用户失败", zap.Any("roleIds", req.Ids), zap.Error(err))
			return consts.ErrServer
		}

		// 3. 软删除角色
		_, err = do.Where(dao.ID.In(req.Ids...)).Delete()
		if err != nil {
			logger.Error("删除角色失败", zap.Any("roleIds", req.Ids), zap.Error(err))
			return consts.ErrServer
		}

		return nil
	})
}

// GetRole 获取角色
func (s *logic) GetRole(ctx context.Context, req *roleDto.GetRoleReq) (*entity.Role, error) {
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
func (s *logic) ListRole(ctx context.Context, req *roleDto.ListRoleReq) (*resp.PageResp, error) {
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
func (s *logic) ListRoleItem(ctx context.Context) ([]*roleDto.ListRoleItemResp, error) {
	var res []*roleDto.ListRoleItemResp
	if err := query.Role.WithContext(ctx).
		Where(query.Role.Status.Eq(1)). // 只查询启用的角色
		Order(query.Role.Sort).
		Scan(&res); err != nil {
		logger.Error("查询角色列表失败", zap.Error(err))
		return nil, consts.ErrServer
	}
	return res, nil
}
