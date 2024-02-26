export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: 'login',
            path: '/user/login',
            component: './user/login',
          },
        ],
      },
      {
        component: './404',
      },
    ],
  },
  {
    path: '/welcome',
    name: 'welcome',
    icon: 'smile',
    component: './Welcome',
  },
  {
    path: '/admin',
    name: 'admin',
    icon: 'crown',
    access: 'canAdmin',
    component: './Admin',
    routes: [
      {
        path: '/admin/sub-page',
        name: 'sub-page',
        icon: 'smile',
        component: './Welcome',
      },
      {
        component: './404',
      },
    ],
  },
  {
    name: 'list.table-list',
    icon: 'table',
    path: '/list',
    component: './TableList',
  },
  {
    path: '/',
    redirect: '/welcome',
  },
  {
    path: '/system',
    name: '系统管理',
    icon: 'crown',
    routes: [
      {
        name: '用户列表',
        icon: 'table',
        path: '/system/user/list',
        component: './system/user',
      },
      {
        name: '角色列表',
        icon: 'table',
        path: '/system/role/list',
        component: './system/role',
      },
      {
        name: '菜单列表',
        icon: 'table',
        path: '/system/menu/list',
        component: './system/menu',
      },
    ],
  },
  {
    path: '/account',
    name: '账户相关',
    icon: 'crown',
    routes: [
      {
        name: '个人中心',
        icon: 'table',
        path: '/account/center',
        component: './account/settings',
      },
      {
        name: '个人设置',
        icon: 'table',
        path: '/account/settings',
        component: './account/settings',
      },
    ],
  },
  {
    component: './404',
  },
];
