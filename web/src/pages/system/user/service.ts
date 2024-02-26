import {request} from 'umi';
import {UpdatePasswordParams, UserListItem, UserListParams} from './data.d';

export async function queryUserList(params: UserListParams) {
  if (params.statusId) {
    params.statusId = Number(params.statusId)
  }
  return request('/api/system/user/queryUserList', {
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
  return request('/api/system/user/deleteUser', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function addUser(params: UserListItem) {
  return request('/api/system/user/addUser', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updateUser(params: UserListItem) {
  return request('/api/system/user/updateUser', {
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

export async function userRoleList(params: { userId: number }) {
  return request('/api/system/user/queryUserRoleList', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}


export async function updateUserRole(params: { userId: number, roleIds: number[] }) {
  return request('/api/system/user/updateUserRoleList', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}
