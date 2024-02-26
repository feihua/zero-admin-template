import { request } from 'umi';
import { MenuListParams, MenuListItem } from './data.d';

export async function queryMenuList(params?: MenuListParams) {
  return request('/api/system/menu/list', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeMenuOne(params: { id: number }) {
  return request('/api/system/menu/delete', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function removeMenu(params: { ids: number[] }) {
  return request('/api/system/menu/delete', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function addMenu(params: MenuListItem) {
  return request('/api/system/menu/save', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}

export async function updateMenu(params: MenuListItem) {
  return request('/api/system/menu/update', {
    method: 'POST',
    data: {
      ...params,
    },
  });
}
