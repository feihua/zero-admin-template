create table hertzdb.sys_role_menu
(
    id          bigint auto_increment comment '主键'
        primary key,
    role_id     bigint                             not null comment '角色ID',
    menu_id     bigint                             not null comment '菜单ID',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '菜单角色关联表';

