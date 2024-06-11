import request from '@/utils/request'

// 查询彩票列表
export function listLottery(query) {
  return request({
    url: '/api/v1/listLottery',
    method: 'get',
    params: query
  })
}

// 导出开奖记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/listLotteryExport',
    method: 'get',
    params: query
  })
}

// 获取彩票详情
export function lotteryDetails(query) {
  return request({
    url: '/api/v1/lotterydetails',
    method: 'get',
    params: query
  })
}
