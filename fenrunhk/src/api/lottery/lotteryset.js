import request from '@/utils/request'

// 查询彩票倍率设置
export function listLotterySets(query) {
  return request({
    url: '/api/v1/listLotterySet',
    method: 'get',
    params: query
  })
}

// 添加修改彩票倍率设置
export function lotterySet(query) {
  return request({
    url: '/api/v1/lotteryset',
    method: 'post',
    params: query
  })
}

// 删除彩票倍率设置
export function delLotterySet(query) {
  return request({
    url: '/api/v1/delLotterySet',
    method: 'post',
    params: query
  })
}
