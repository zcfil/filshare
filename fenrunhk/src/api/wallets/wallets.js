import request from '@/utils/request'
// 用户密码重置
export function updateUserPwd(user) {
  return request({
    url: '/api/v1/updateWalletsPass',
    method: 'post',
    params: user
  })
}

export function collectWalletsAsset(user) {
  return request({
    url: '/api/v1/collectWalletsAsset',
    method: 'post',
    params: user
  })
}
// 处理异常充值
export function exceptionCoinCharge(user) {
  return request({
    url: '/api/v1/exception/coin_charge',
    method: 'get',
    params: user
  })
}
// 获取资金归集记录
export function getCollectList(user) {
  return request({
    url: '/api/v1/getCollectList',
    method: 'get',
    params: user
  })
}
// 获取资金归集记录地址
export function getCollectAddressList(user) {
  return request({
    url: '/api/v1/getCollectAddressList',
    method: 'get',
    params: user
  })
}
