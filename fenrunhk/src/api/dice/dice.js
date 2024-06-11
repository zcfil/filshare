import request from '@/utils/request'

// 查询dice投注记录
export function diceRecords(query) {
  return request({
    // url: '/api/v1/diceList',
    url: '/api/v1/diceRecords',
    method: 'get',
    params: query
  })
}
// 导出dice投注记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/diceRecordsExport',
    method: 'get',
    params: query
  })
}
// 连胜
export function ContinuousWin(query) {
  return request({
    url: '/api/v1/continuouswin',
    method: 'get',
    params: query
  })
}
// 地址资产列表
export function listAddressBalls(query) {
  return request({
    url: '/api/v1/listAddressBalls',
    method: 'get',
    params: query
  })
}
// 更新地址资产数据
export function UpdateWalletsAssets(query) {
  return request({
    url: '/api/v1/updateWalletsAssets',
    method: 'post'
  })
}
// 更新地址资产数据
export function listUploadAddress(query) {
  return request({
    url: '/api/v1/listUploadAddress',
    method: 'get',
    params: query
  })
}
