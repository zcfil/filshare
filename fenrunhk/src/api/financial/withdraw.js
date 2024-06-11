import request from '@/utils/request'

// 查询提币记录
export function lsitDraws(query) {
  return request({
    url: '/api/v1/draws',
    method: 'get',
    params: query
  })
}
// 查询提币详情
export function getDraws(id) {
  return request({
    url: '/api/v1/getDraws/' + id,
    method: 'get'
  })
}
// 取消提币操作
export function drawCancel(params) {
  return request({
    url: '/api/v1/coinCheckCancel',
    method: 'post',
    params: params
  })
}
// 导出提币记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/withdrawExport',
    method: 'get',
    params: query
  })
}
// 提币审核
export function drawCheck(query) {
  return request({
    url: '/api/v1/coinCheck',
    method: 'Post',
    params: query
  })
}
