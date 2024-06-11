import request from '@/utils/request'

// 查询质押列表
export function listPledge(query) {
  return request({
    url: '/api/v1/listPledges',
    method: 'get',
    params: query
  })
}
// 查询质押分红列表
export function listPledgeLog(query) {
  return request({
    url: '/api/v1/listPledgeBonus',
    method: 'get',
    params: query
  })
}
