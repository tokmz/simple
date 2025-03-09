package consts

import "errors"

var (
	ErrFail    = errors.New("fail") // 失败
	ErrUnknown = errors.New("未知错误") // 未知错误

	// 认证相关错误
	ErrUnauthorized     = errors.New("无权限访问") // 未授权
	ErrForbidden        = errors.New("禁止访问")  // 禁止访问
	ErrInvalidToken     = errors.New("无效的令牌") // 无效令牌
	ErrTokenExpired     = errors.New("令牌已过期") // 令牌过期
	ErrInvalidSignature = errors.New("无效的签名") // 签名无效

	// 请求相关错误
	ErrBadRequest       = errors.New("无效的请求")  // 无效请求
	ErrInvalidParam     = errors.New("无效的参数")  // 参数错误
	ErrMissingParam     = errors.New("缺少必要参数") // 缺少参数
	ErrResourceNotFound = errors.New("资源不存在")  // 资源不存在
	ErrMethodNotAllowed = errors.New("方法不允许")  // 方法不允许
	ErrTimeout          = errors.New("请求超时")   // 请求超时

	// 业务相关错误
	ErrUserNotFound    = errors.New("用户不存在") // 用户不存在
	ErrUserExists      = errors.New("用户已存在") // 用户已存在
	ErrInvalidPassword = errors.New("密码错误")  // 密码错误
	ErrAccountLocked   = errors.New("账号已锁定") // 账号锁定
	ErrOperationFailed = errors.New("操作失败")  // 操作失败

	// 系统相关错误
	ErrServer      = errors.New("系统错误") // 系统错误
	ErrServiceBusy = errors.New("服务繁忙") // 服务繁忙
	ErrConfig      = errors.New("配置错误") // 配置错误
	ErrNotFound    = errors.New("未找到")  // 未找到

	// 角色相关错误
	ErrRoleNotFound   = errors.New("角色不存在")      // 角色不存在
	ErrRoleNameExists = errors.New("角色名称已存在")    // 角色名称已存在
	ErrRoleCodeExists = errors.New("角色编码已存在")    // 角色编码已存在
	ErrRoleSuperAdmin = errors.New("超级管理员不允许删除") // 超级管理员不允许删除

	// 部门相关错误
	ErrDepartmentNotFound    = errors.New("部门不存在")      // 部门不存在
	ErrDepartmentNameExists  = errors.New("部门名称已存在")    // 部门名称已存在
	ErrDepartmentCodeExists  = errors.New("部门编码已存在")    // 部门编码已存在
	ErrDepartmentSuperAdmin  = errors.New("超级管理员不允许删除") // 超级管理员不允许删除
	ErrDepartmentHasChildren = errors.New("存在子部门")      // 存在子部门
	ErrDepartmentHasUsers    = errors.New("部门下存在用户")    // 部门下存在用户

	// 岗位相关错误
	ErrPositionNotFound   = errors.New("岗位不存在")   // 岗位不存在
	ErrPositionNameExists = errors.New("岗位名称已存在") // 岗位名称已存在
	ErrPositionCodeExists = errors.New("岗位编码已存在") // 岗位编码已存在
	ErrPositionHasUsers   = errors.New("岗位下存在用户") // 岗位下存在用户

	// 菜单相关错误
	ErrMenuNotFound       = errors.New("菜单不存在")    // 菜单不存在
	ErrMenuNameExists     = errors.New("菜单名称已存在")  // 菜单名称已存在
	ErrMenuParentNotFound = errors.New("父菜单不存在")   // 父菜单不存在
	ErrMenuHasChildren    = errors.New("菜单下存在子菜单") // 菜单下存在子菜单
	ErrMenuHasRoles       = errors.New("菜单被角色使用")  // 菜单被角色使用
)

// 错误码定义
var code = map[error]int{
	// 通用错误码 (0-99)
	ErrFail:    0,  // 失败
	ErrUnknown: -1, // 未知错误

	// 认证相关错误码 (1000-1999)
	ErrUnauthorized:     1001, // 未授权
	ErrForbidden:        1002, // 禁止访问
	ErrInvalidToken:     1003, // 无效令牌
	ErrTokenExpired:     1004, // 令牌过期
	ErrInvalidSignature: 1005, // 签名无效

	// 请求相关错误码 (2000-2999)
	ErrBadRequest:       2001, // 无效请求
	ErrInvalidParam:     2002, // 参数错误
	ErrMissingParam:     2003, // 缺少参数
	ErrResourceNotFound: 2004, // 资源不存在
	ErrMethodNotAllowed: 2005, // 方法不允许
	ErrTimeout:          2006, // 请求超时

	// 用户相关错误码 (3000-3100)
	ErrUserNotFound:    3001, // 用户不存在
	ErrUserExists:      3002, // 用户已存在
	ErrInvalidPassword: 3003, // 密码错误
	ErrAccountLocked:   3004, // 账号锁定
	ErrOperationFailed: 3005, // 操作失败

	// 系统相关错误码 (5000-5999)
	ErrServer:      5001, // 系统错误
	ErrServiceBusy: 5002, // 服务繁忙
	ErrConfig:      5003, // 配置错误
	ErrNotFound:    5004, // 未找到

	// 角色相关错误码 (3100-3200)
	ErrRoleNotFound:   3101, // 角色不存在
	ErrRoleNameExists: 3102, // 角色名称已存在
	ErrRoleCodeExists: 3103, // 角色编码已存在
	ErrRoleSuperAdmin: 3104, // 超级管理员不允许删除

	// 部门相关错误码 (3200-3300)
	ErrDepartmentNotFound:    3201, // 部门不存在
	ErrDepartmentNameExists:  3202, // 部门名称已存在
	ErrDepartmentCodeExists:  3203, // 部门编码已存在
	ErrDepartmentSuperAdmin:  3204, // 超级管理员不允许删除
	ErrDepartmentHasChildren: 3205, // 存在子部门
	ErrDepartmentHasUsers:    3206, // 部门下存在用户

	// 岗位相关错误码 (3300-3400)
	ErrPositionNotFound:   3301, // 岗位不存在
	ErrPositionNameExists: 3302, // 岗位名称已存在
	ErrPositionCodeExists: 3303, // 岗位编码已存在
	ErrPositionHasUsers:   3304, // 岗位下存在用户

	// 菜单相关错误码 (3400-3500)
	ErrMenuNotFound:       3401, // 菜单不存在
	ErrMenuNameExists:     3402, // 菜单名称已存在
	ErrMenuParentNotFound: 3403, // 父菜单不存在
	ErrMenuHasChildren:    3404, // 菜单下存在子菜单
	ErrMenuHasRoles:       3405, // 菜单被角色使用
}

// GC 获取错误码
func GC(err error) int {
	if err == nil {
		return 200
	}
	if c, ok := code[err]; ok {
		return c
	}
	return code[ErrUnknown]
}

// IsSuccess 判断是否成功
func IsSuccess(err error) bool {
	return err == nil
}

// GetMessage 获取错误信息
func GetMessage(err error) string {
	if err == nil {
		return "ok"
	}
	return err.Error()
}
