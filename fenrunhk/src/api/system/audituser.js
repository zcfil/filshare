import request from '@/utils/request'
export function GetPassingSysUserList(data) {
  return request({
    url: '/api/v1/passingSysUserList',
    method: 'get',
    params: data
  })
}
export function SubmitNewUser(data) {
  return request({
    url: '/api/v1/submitNewUser',
    method: 'post',
    data: data
  })
}
export function AllowNewUserPass(data) {
  return request({
    url: '/api/v1/allowNewUserPass',
    method: 'post',
    params: data
  })
}
export function userAuditRecord(data) {
  return request({
    url: '/api/v1/userAuditRecord',
    method: 'get',
    params: data
  })
}