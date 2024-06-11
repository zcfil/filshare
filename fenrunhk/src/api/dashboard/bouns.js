import request from '@/utils/request'

export function bouns(data) {
  return request({
    url: '/api/v1/bonus',
    method: 'get',
    data
  })
}
