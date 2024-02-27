import { request } from 'umi';
import { RoleListParams, RoleListItem } from './data.d';

export async function queryRole(params?: RoleListParams) {
  return request('/api/system/role/queryRoleList', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function queryMenuByRoleId(params: { roleId?: number }) {
  return request('/api/system/role/queryRoleMenuList', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updateRoleMenu(params: { roleId: number ,menuIds:number[]}) {
  return request('/api/system/role/updateRoleMenuList', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}


export async function removeRole(params: { ids: number[] }) {
  return request('/api/system/role/deleteRole', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function addRole(params: RoleListItem) {
  return request('/api/system/role/addRole', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updateRole(params: RoleListItem) {
  return request('/api/system/role/updateRole', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}
