import request from '@/utils/request'

// 查询需要审核的提币记录
export function lsitDraws(query) {
  return request({
    url: '/api/v1/draws',
    method: 'get',
    params: query
  })
}
