import request from '@/utils/request'

// 查询资产列表
export function listLevel(query) {
  return request({
    url: '/api/v1/member/memberLevelsList',
    method: 'get',
    params: query
  })
}

// 查询会员详细
export function getMember(menberId) {
  return request({
    url: '/api/v1/memberList' + menberId,
    method: 'get'
  })
}

// 新增会员
export function addMenu(data) {
  return request({
    url: '/api/v1/menu',
    method: 'post',
    data: data
  })
}

// 修改会员
export function updateMenu(data) {
  return request({
    url: '/api/v1/menu',
    method: 'put',
    data: data
  })
}

// 删除会员资产
export function delAssets(menuId) {
  return request({
    url: '/api/v1/menu/' + menuId,
    method: 'delete'
  })
}

// 导出会员资产列表
export function exportExcel(query) {
  return request({
    url: '/api/v1/memberLevelsListExport',
    method: 'get',
    params: query
  })
}
