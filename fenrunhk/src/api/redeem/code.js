import request from '@/utils/request'

//获取兑换码列表
export function redeemCodeList(query) {
  return request({
    url: '/api/v1/redeemCodeList',
    method: 'get',
    params: query
  })
}
// 生成兑换码
export function addRedeemCode(query) {
  return request({
    url: '/api/v1/addRedeemCode',
    method: 'post',
    params: query
  })
}
export function delRedeemCode(query) {
    return request({
      url: '/api/v1/delRedeemCode',
      method: 'post',
      params: query
    })
  }
// 导出
export function exportExcel(query) {
    return request({
      url: '/api/v1/redeemCodeLisExport',
      method: 'get',
      params: query
    })
  }