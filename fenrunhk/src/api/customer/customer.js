import request from '@/utils/request'

// 获取客户列表
export function customerList(data) {
  return request({
    url: '/api/v1/customerList',
    method: 'get',
    params: data
  })
}
// 添加客户
export function customerAdd(data) {
  return request({
    url: '/api/v1/customerAdd',
    method: 'post',
    params: data
  })
}
// 编辑客户
export function customerEdit(data) {
  return request({
    url: '/api/v1/customerEdit',
    method: 'post',
    params: data
  })
}
export function customerDelete(data) {
  return request({
    url: '/api/v1/customerDelete',
    method: 'post',
    params: data
  })
}
// 获取客户投资列表
export function investmentList(data) {
  return request({
    url: '/api/v1/investmentList',
    method: 'get',
    params: data
  })
}
// 添加客户
export function investmentAdd(data) {
  return request({
    url: '/api/v1/investmentAdd',
    method: 'post',
    params: data
  })
}
// 编辑客户
export function investmentEdit(data) {
  return request({
    url: '/api/v1/investmentEdit',
    method: 'post',
    params: data
  })
}

export function investmentDelete(data) {
  return request({
    url: '/api/v1/investmentDelete',
    method: 'post',
    params: data
  })
}

export function investmentBreak(query) {
  return request({
    url: '/api/v1/investmentBreak',
    method: 'post',
    params: query
  })
}
// 获取配置
export function financeConfigList(data) {
  return request({
    url: '/api/v1/financeConfigList',
    method: 'get',
    params: data
  })
}
// 修改配置
export function financeConfigEdit(data) {
  return request({
    url: '/api/v1/financeConfigEdit',
    method: 'post',
    params: data
  })
}

// 获取周结算列表
export function getWeekList(data) {
  return request({
    url: '/api/v1/getWeekList',
    method: 'get',
    params: data
  })
}
// 获取周用户列表
export function getWeekCustomerList(data) {
  return request({
    url: '/api/v1/getWeekCustomerList',
    method: 'get',
    params: data
  })
}
// 获取周用户投资列表
export function getWeekCustomerInvestmentList(data) {
  return request({
    url: '/api/v1/getWeekCustomerInvestmentList',
    method: 'get',
    params: data
  })
}

// 转账记录
export function getTransferList(data) {
  return request({
    url: '/api/v1/getTransferList',
    method: 'get',
    params: data
  })
}
// 结算记录
export function getSettleList(data) {
  return request({
    url: '/api/v1/getSettleList',
    method: 'get',
    params: data
  })
}

// 周转账
export function transferWeek(data) {
  return request({
    url: '/api/v1/transferWeek',
    method: 'post',
    params: data
  })
}
// 周用户转账
export function transferWeekCustomer(data) {
  return request({
    url: '/api/v1/transferWeekCustomer',
    method: 'post',
    params: data
  })
}

// 首页
// export function officialInfo(data) {
//   return request({
//     url: '/api/v1/officialInfo',
//     method: 'get',
//     params: data
//   })
// }
export function userList(query) {
  return request({
    url: '/api/v1/userList',
    method: 'get',
    params: query
  })
}

// 上传文件
export function uploadfile(data) {
  return request({
    url: '/api/v1/uploadfile',
    method: 'post',
    params: data
  })
}
// 上传文件列表
export function getuploadfileList(data) {
  return request({
    url: '/api/v1/getuploadfileList',
    method: 'get',
    params: data
  })
}
// 删除文件
export function delfile(data) {
  return request({
    url: '/api/v1/delfile',
    method: 'get',
    params: data
  })
}

// 添加迁移用户数据
export function addmigrate(data) {
  return request({
    url: '/api/v1/addmigrate',
    method: 'post',
    params: data
  })
}
// 查询所有订单
export function getmigrate(data) {
  return request({
    url: '/api/v1/getmigrate',
    method: 'get',
    params: data
  })
}
// 删除订单
export function deletemigrate(data) {
  return request({
    url: '/api/v1/deletemigrate',
    method: 'post',
    params: data
  })
}
// 修改编辑订单
export function editmigrate(data) {
  return request({
    url: '/api/v1/editmigrate',
    method: 'post',
    params: data
  })
}
//终止订单
export function breakmigreate(data) {
  return request({
    url: '/api/v1/breakmigreate',
    method: 'post',
    params: data
  })
}
// 用户迁移数据列表
export function impressionslist(data) {
  return request({
    url: '/api/v1/impressionslist',
    method: 'get',
    params: data
  })
}