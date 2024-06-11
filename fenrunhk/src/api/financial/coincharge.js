import request from '@/utils/request'

// 查询充币记录
export function lsitCoinCharg(query) {
  return request({
    url: '/api/v1/coincharges',
    method: 'get',
    params: query
  })
}

// 导出充币记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/coinchargesExport',
    method: 'get',
    params: query
  })
}
