import request from '@/utils/request'

// 查询兑换释放记录
export function listExchange(query) {
  return request({
    url: '/api/v1/exchanges',
    method: 'get',
    params: query
  })
}

// 导出兑换记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/exchangesExport',
    method: 'get',
    params: query
  })
}
