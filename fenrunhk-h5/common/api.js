import request from './request.js'
import qs from 'qs'
// 登录
export function login(data) {
  data = qs.stringify(data)
  return request({
    url: '/customer/login?'+data,
    method: 'post',
  })
}
// 修改密码
export function changePassword(data) {
  data = qs.stringify(data)
  return request({
    url: '/customer/changePassword?'+data,
    method: 'post',
  })
}
// 投资列表
export function investmentList(data) {
  return request({
    url: '/customer/investmentList',
    method: 'get',
	data
  })
}
// 转账列表
export function customerTransferList(data) {
  return request({
    url: '/customer/customerTransferList',
    method: 'get',
	data
  })
}
// 结算列表
export function settlementList(data) {
  return request({
    url: '/customer/settlementList',
    method: 'get',
	data
  })
}


export function homepage(data) {
  return request({
    url: '/customer/homepage',
    method: 'get',
	data
  })
}

export function impressionsList(data) {
  return request({
    url: '/customer/impressionsList',
    method: 'get',
	data
  })
}