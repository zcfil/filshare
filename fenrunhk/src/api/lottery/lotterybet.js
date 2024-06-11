import request from '@/utils/request'

// 查询彩票投注记录
export function listBet(query) {
  return request({
    url: '/api/v1/listLotteryBet',
    method: 'get',
    params: query
  })
}
// 导出投注记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/listLotteryBetExport',
    method: 'get',
    params: query
  })
}
