import { request } from 'umi';
import { RoleListParams, RoleListItem } from './data.d';

export async function queryRole(params?: RoleListParams) {
  return request('/api/system/role/list', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function queryMenuByRoleId(params: { roleId?: number }) {
  return request('/api/system/menu/roleMenuList', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updateRoleMenu(params: { roleId: number ,menuIds:number[]}) {
  return request('/api/system/menu/roleMenuSave', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeRoleOne(params: { id: number }) {
  return request('/api/system/role/delete', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeRole(params: { ids: number[] }) {
  return request('/api/system/role/delete', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function addRole(params: RoleListItem) {
  return request('/api/system/role/save', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updateRole(params: RoleListItem) {
  return request('/api/system/role/update', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}
