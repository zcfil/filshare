import request from '@/utils/request'

// 查询资金记录
export function lsitMoneyLog(query) {
  return request({
    url: '/api/v1/moneyLogs',
    method: 'get',
    params: query
  })
}
// 导出财务记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/usermoneylogExport',
    method: 'get',
    params: query
  })
}
