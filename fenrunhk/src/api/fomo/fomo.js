import request from '@/utils/request'

// 查询fomo投注记录
export function fomoRecords(query) {
  return request({
    // url: '/api/v1/diceList',
    url: '/api/v1/fomoRecords',
    method: 'get',
    params: query
  })
}
// 导出fomo投注记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/fomoRecordsExport',
    method: 'get',
    params: query
  })
}
