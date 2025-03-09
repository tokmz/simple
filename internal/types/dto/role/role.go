package role

/*
   @NAME    : role
   @author  : 清风
   @desc    :
   @time    : 2025/3/7 22:38
*/

// CreateRoleReq 创建角色请求
type CreateRoleReq struct {
	Name          string  `json:"name" binding:"required"`       // 角色名称
	Code          string  `json:"code" binding:"required"`       // 角色编码
	DefaultRouter *string `json:"default_router"`                // 默认路由
	Status        *int64  `json:"status"`                        // 状态 1:启用 2:禁用
	Remark        *string `json:"remark"`                        // 备注
	Sort          int64   `json:"sort" binding:"required,min=0"` // 排序
}

// UpdateRoleReq 更新角色请求
type UpdateRoleReq struct {
	ID            int64   `json:"id" binding:"required"`         // 角色ID
	Name          string  `json:"name" binding:"required"`       // 角色名称
	Code          string  `json:"code" binding:"required"`       // 角色编码
	DefaultRouter *string `json:"default_router"`                // 默认路由
	Status        *int64  `json:"status"`                        // 状态 1:启用 2:禁用
	Remark        *string `json:"remark"`                        // 备注
	Sort          int64   `json:"sort" binding:"required,min=0"` // 排序
}

// DeleteRoleReq 删除角色请求
type DeleteRoleReq struct {
	Ids []int64 `json:"ids" binding:"required,min=1"` // 角色ID列表
}

// GetRoleReq 获取角色请求
type GetRoleReq struct {
	ID int64 `json:"id" binding:"required"` // 角色ID
}

// ListRoleReq 角色列表请求
type ListRoleReq struct {
	Name   *string `json:"name"`                                   // 角色名称
	Code   *string `json:"code"`                                   // 角色编码
	Status *int64  `json:"status"`                                 // 状态
	Page   int     `json:"page" binding:"required,min=1"`          // 页码
	Size   int     `json:"size" binding:"required,min=10,max=100"` // 每页数量
}

// ListRoleItemResp 角色选项响应
type ListRoleItemResp struct {
	ID   int64  `json:"id"`   // 角色ID
	Name string `json:"name"` // 角色名称
}
