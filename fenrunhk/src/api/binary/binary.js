import request from '@/utils/request'

// 查询二元期权投注记录
export function listBinary(query) {
  return request({
    url: '/api/v1/listBinary',
    method: 'get',
    params: query
  })
}
// 导出
export function exportExcel(query) {
  return request({
    url: '/api/v1/listBinaryExport',
    method: 'get',
    params: query
  })
}
