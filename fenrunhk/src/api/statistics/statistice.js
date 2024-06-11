import request from '@/utils/request'

// 彩票盈亏统计
export function lotteryProfitAndLoss(query) {
  return request({
    // url: '/api/v1/diceList',
    url: '/api/v1/lotteryProfitAndLoss',
    method: 'get',
    params: query
  })
}
// binary盈亏统计
export function binaryProfitAndLoss(query) {
  return request({
    url: '/api/v1/betProfitAndLoss',
    method: 'get',
    params: query
  })
}
// dice盈亏统计
export function diceProfitAndLoss(query) {
  return request({
    url: '/api/v1/diceProfitAndLoss',
    method: 'get',
    params: query
  })
}
// fomo盈亏统计
export function fomoProfitAndLoss(query) {
  return request({
    url: '/api/v1/fomoProfitAndLoss',
    method: 'get',
    params: query
  })
}
// 充币提币统计
export function chargeToMention(query) {
  return request({
    url: '/api/v1/chargeToMention',
    method: 'get',
    params: query
  })
}
