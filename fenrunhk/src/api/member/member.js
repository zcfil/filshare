import request from '@/utils/request'
import qs from 'qs'
// 查询会员列表
export function listMember(query) {
  return request({
    url: '/api/v1/memberList',
    method: 'get',
    params: query
  })
}

// 查询会员详细
export function getMember(menberId) {
  return request({
    url: '/api/v1/memberList' + menberId,
    method: 'get'
  })
}

// 新增会员
export function addMenu(data) {
  return request({
    url: '/api/v1/menu',
    method: 'post',
    data: qs.stringify(data)
  })
}

// 修改会员
export function updateMenu(data) {
  return request({
    url: '/api/v1/member/memberEdit',
    method: 'post',
    params: data
  })
}

// 删除会员
export function delMenu(data) {
  return request({
    url: '/api/v1/member/memberDelete',
    method: 'post',
    params: data
  })
}

// 导出参数
export function exportUser(query) {
  return request({
    url: '/api/v1/memberExport',
    method: 'get',
    params: query
  })
}
// 邀请详情
export function userInvitePeople(query) {
  return request({
    url: '/api/v1/member/userInvitePeople',
    method: 'get',
    params: query
  })
}