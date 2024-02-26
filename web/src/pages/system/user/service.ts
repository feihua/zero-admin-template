import { request } from 'umi';
import {UserListParams, UserListItem, UpdatePasswordParams} from './data.d';

export async function queryUserList(params?: UserListParams) {
  return request('/api/system/user/list', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function querySelectAllData(params?: UserListParams) {
  return request('/api/system/user/selectAllData', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}


export async function removeUser(params: { ids: number[] }) {
  return request('/api/system/user/delete', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function addUser(params: UserListItem) {
  return request('/api/system/user/save', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updateUser(params: UserListItem) {
  return request('/api/system/user/update', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updatePassword(params: UpdatePasswordParams) {
  return request('/api/system/user/updatePassword', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function userRoleList(params: { userId: number}) {
  return request('/api/system/role/userRoleList', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}


export async function updateUserRole(params: { userId: number ,roleIds: number[]}) {
  return request('/api/system/role/userRoleSave', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}
