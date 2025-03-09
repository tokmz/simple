/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : simple

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 07/03/2025 00:05:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_department
-- ----------------------------
DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID|Primary key',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父部门ID|Parent department ID',
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门名称|Department name',
  `code` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门编码|Department code',
  `leader` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '部门负责人|Department leader',
  `phone` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '联系电话|Contact number',
  `email` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱|Email',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序|Sort',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:启用 2:禁用|Status 1:Enable 2:Disable',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注|Remark',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间|Created Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间|Updated Time',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间|Deleted Time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统部门表';

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID|Primary key',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父菜单ID|Parent menu ID',
  `title` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单标题|Menu title',
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称|Menu name',
  `path` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单路径|Menu path',
  `component` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '组件路径|Component path',
  `redirect` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '重定向|Redirect',
  `icon` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '图标|Icon',
  `type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '菜单类型 0:目录 1:菜单 2:按钮|Menu type 0:Directory 1:Menu 2:Button',
  `permission` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '权限标识|Permission',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序|Sort',
  `is_hidden` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否隐藏 1:隐藏 2:显示|Is hidden 1:Hide 2:Show',
  `is_cache` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否缓存 1:缓存 2:不缓存|Is cache 1:Yes 2:No',
  `is_affix` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否固定 1:固定 2:不固定|Is affix 1:Yes 2:No',
  `trans` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '多语言翻译|Translation',
  `level` int unsigned NOT NULL DEFAULT '1' COMMENT '菜单层级|Menu level',
  `hide_breadcrumb` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否隐藏面包屑 1:隐藏 2:显示|Hide breadcrumb 1:Yes 2:No',
  `hide_tab` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否隐藏标签页 1:隐藏 2:显示|Hide tab 1:Yes 2:No',
  `frame_src` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '内嵌iframe地址|Frame source',
  `carry_param` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否携带参数 1:是 2:否|Carry param 1:Yes 2:No',
  `hide_children_in_menu` tinyint(1) NOT NULL DEFAULT '2' COMMENT '是否在菜单中隐藏子节点 1:是 2:否|Hide children in menu 1:Yes 2:No',
  `dynamic_level` int unsigned DEFAULT '20' COMMENT '动态路由层级|Dynamic route level',
  `real_path` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '真实路径|Real path',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:启用 2:禁用|Status 1:Enable 2:Disable',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间|Created Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间|Updated Time',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间|Deleted Time',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统菜单表';

-- ----------------------------
-- Table structure for sys_position
-- ----------------------------
DROP TABLE IF EXISTS `sys_position`;
CREATE TABLE `sys_position` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID|Primary key',
  `department_id` bigint unsigned DEFAULT NULL COMMENT '部门ID|Department ID',
  `name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位名称|Position name',
  `code` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位编码|Position code',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序|Sort',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:启用 2:禁用|Status 1:Enable 2:Disable',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注|Remark',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间|Created Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间|Updated Time',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间|Deleted Time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`),
  KEY `idx_department_id` (`department_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_position_department` FOREIGN KEY (`department_id`) REFERENCES `sys_department` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统岗位表';

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID|Primary key',
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称|Role name',
  `code` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色编码|Role code',
  `default_router` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '/dashboard' COMMENT '默认的路由|Default router',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:启用 2:禁用|Status 1:Enable 2:Disable',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注|Remark',
  `sort` int unsigned NOT NULL DEFAULT '0' COMMENT '排序|Sort',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间|Created Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间|Updated Time',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间|Deleted Time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_code` (`code`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统角色表';

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID|Primary key',
  `uuid` char(36) COLLATE utf8mb4_general_ci NOT NULL COMMENT '唯一标识符|UUID',
  `username` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名|Username',
  `password` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码|Password',
  `salt` varchar(10) COLLATE utf8mb4_general_ci NOT NULL COMMENT '盐值|Salt',
  `name` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '姓名|Name',
  `nickname` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称|Nickname',
  `email` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '邮箱|Email',
  `mobile` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号|Mobile',
  `avatar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '头像|Avatar',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态 1:启用 2:禁用|Status 1:Enable 2:Disable',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注|Remark',
  `home_path` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '/dashboard' COMMENT '首页路径|Home Path',
  `department_id` bigint unsigned DEFAULT NULL COMMENT '部门ID|Department ID',
  `position_id` bigint unsigned DEFAULT NULL COMMENT '岗位ID|Position ID',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间|Last Login Time',
  `last_login_ip` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '最后登录IP|Last Login IP',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间|Created Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间|Updated Time',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间|Deleted Time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  UNIQUE KEY `idx_uuid` (`uuid`),
  KEY `idx_email` (`email`),
  KEY `idx_mobile` (`mobile`),
  KEY `idx_status` (`status`),
  KEY `idx_department_id` (`department_id`),
  KEY `idx_position_id` (`position_id`),
  KEY `idx_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_users_department` FOREIGN KEY (`department_id`) REFERENCES `sys_department` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_users_position` FOREIGN KEY (`position_id`) REFERENCES `sys_position` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统用户表';

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID|Primary key',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID|User ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID|Role ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间|Created Time',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间|Updated Time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_role` (`user_id`,`role_id`),
  KEY `idx_role_id` (`role_id`),
  CONSTRAINT `fk_user_roles_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `sys_user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户-角色关系表';

SET FOREIGN_KEY_CHECKS = 1;
