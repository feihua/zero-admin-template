create table hertzdb.sys_menu
(
    id          bigint auto_increment comment '主键'
        primary key,
    menu_name   varchar(50)                            not null comment '菜单名称',
    menu_type   tinyint      default 1                 not null comment '菜单类型(1：目录   2：菜单   3：按钮)',
    status_id   tinyint      default 1                 not null comment '状态(1:正常，0:禁用)',
    sort        int          default 1                 not null comment '排序',
    parent_id   bigint                                 not null comment '父ID',
    menu_url    varchar(255) default ''                null comment '路由路径',
    api_url     varchar(255) default ''                null comment '接口URL',
    menu_icon   varchar(255)                           null comment '菜单图标',
    remark      varchar(255)                           null comment '备注',
    create_time datetime     default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime                               null on update CURRENT_TIMESTAMP comment '修改时间',
    constraint menu_name
        unique (menu_name)
)
    comment '菜单信息';


INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (1, '首页', 1, 1, 0, 0, '/home', '', 'SmileOutlined', '首页', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (2, '权限管理', 1, 1, 1, 0, '', '', 'SettingOutlined', '权限管理', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (3, '用户管理', 2, 1, 1, 2, '/system/user/list', '/api/system/user/queryUserList', 'setting', '用户管理', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (4, '角色管理', 2, 1, 2, 2, '/system/role/list', '/api/system/role/queryRoleList', '', '角色管理', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (5, '菜单管理', 2, 1, 3, 2, '/system/menu/list', '/api/system/menu/queryMenuList', 'setting', '菜单管理', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (6, '更新用户状态接口', 3, 1, 1, 3, '', '/api/system/user/update_user_status', '', '更新用户状态接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (7, '保存用户弹窗', 3, 1, 1, 3, '', '/api/system/user/userSaveView', '', '保存用户弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (8, '保存用户接口', 3, 1, 1, 3, '', '/api/system/user/addUser', '', '保存用户接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (9, '删除用户接口', 3, 1, 1, 3, '', '/api/system/user/deleteUser', '', '删除用户接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (10, '更新用户弹窗', 3, 1, 1, 3, '', '/api/system/user/userUpdateView', '', '更新用户弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (11, '更新用户接口', 3, 1, 1, 3, '', '/api/system/user/updateUser', '', '更新用户接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (12, '更新用户密码弹窗', 3, 1, 1, 3, '', '/api/system/user/update_user_passwordView', '', '更新用户密码弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (13, '更新用户密码', 3, 1, 1, 3, '', '/api/system/user/update_user_password', '', '更新用户密码接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (14, '设置角色弹窗', 3, 1, 1, 3, '', '/api/system/user/updateUserRoleView', '', '设置角色弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (15, '保存用户角色', 3, 1, 1, 3, '', '/api/system/user/updateUserRoleList', '', '保存用户角色接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (16, '用户关联的角色', 3, 1, 1, 3, '', '/api/system/user/queryUserRoleList', '', '获取用户关联的角色', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (17, '查询用户菜单接口', 3, 1, 1, 3, '', '/api/system/user/queryUserMenu', '', '查询用户菜单接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (18, '更新角色状态接口', 3, 1, 1, 4, '', '/api/system/role/updateRoleStatus', '', '更新角色状态接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (19, '保存角色弹窗', 3, 1, 1, 4, '', '/api/system/role/roleSaveView', '', '保存角色弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (20, '保存角色接口', 3, 1, 1, 4, '', '/api/system/role/addRole', '', '保存角色接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (21, '删除角色接口', 3, 1, 1, 4, '', '/api/system/role/deleteRole', '', '删除角色接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (22, '修改角色弹窗', 3, 1, 1, 4, '', '/api/system/role/roleUpdateView', '', '修改角色弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (23, '更新角色接口', 3, 1, 1, 4, '', '/api/system/role/updateRole', '', '更新角色接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (24, '设置权限弹窗', 3, 0, 1, 4, '', '/api/system/role/queryRoleMenuView', '', '设置权限弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (25, '菜单角色关联', 3, 1, 1, 4, '', '/api/system/role/queryRoleMenuList', '', '菜单角色关联', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (26, '保存角色菜单关联', 3, 1, 1, 4, '', '/api/system/role/updateRoleMenuList', '', '角色菜单关联接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (27, '更新菜单状态接口', 3, 1, 1, 5, '', '/api/system/role/updateMenuStatus', '', '更新菜单状态接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (28, '保存菜单弹窗', 3, 1, 1, 5, '', '/api/system/menu/menuSaveView', '', '保存菜单弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (29, '保存菜单接口', 3, 1, 1, 5, '', '/api/system/menu/addMenu', '', '保存菜单接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (30, '删除菜单接口', 3, 1, 1, 5, '', '/api/system/menu/deleteMenu', '', '删除菜单接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (31, '修改菜单弹窗', 3, 1, 1, 5, '', '/api/system/menu/menuUpdateView', '', '修改菜单弹窗', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (32, '更新菜单接口', 3, 1, 1, 5, '', '/api/system/menu/updateMenu', '', '更新菜单接口', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (63, '日志管理', 1, 1, 2, 0, '/log1', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (64, '登录日志', 2, 1, 1, 63, '/log', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (65, '常用图表', 1, 1, 3, 0, '/line1', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (66, '饼图', 2, 1, 1, 65, '/bar', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (67, '线图', 2, 1, 1, 65, '/line', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (68, '柱状图', 2, 1, 1, 65, '/pie', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (69, '个人中心', 1, 0, 1, 0, '/us', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (72, '个人信息', 2, 1, 1, 69, '/center', '', 'Setting', '', current_timestamp, null);
INSERT INTO sys_menu (id, menu_name, menu_type, status_id, sort, parent_id, menu_url, api_url, menu_icon, remark, create_time, update_time) VALUES (73, '个人设置', 2, 1, 1, 69, '/setting', '', 'Setting', '', current_timestamp, null);
