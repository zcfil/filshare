import request from '@/utils/request'

// 获取手动充币列表
export function manualCoinList(query) {
  return request({
    url: '/api/v1/manualCoinList',
    method: 'get',
    params: query
  })
}
// 手动充币
export function addManualCoin(query) {
  return request({
    url: '/api/v1/addManualCoin',
    method: 'post',
    params: query
  })
}
