import request from '@/utils/request'

// 走势图数据
export function GetStatGamesmoth(query) {
  return request({
    url: '/api/v1/statgamesmoth',
    method: 'get',
    params: query
  })
}
// 查询注册总人数and今日注册人数
export function GetUserCount(query) {
  return request({
    url: '/api/v1/homeusercount',
    method: 'get',
    params: query
  })
}
// 获取游戏下注总数
export function GetGamesCount(query) {
  return request({
    url: '/api/v1/gamescount',
    method: 'get',
    params: query
  })
}
// 会员资产总数
export function GetUserAssetCcount(query) {
  return request({
    url: '/api/v1/userassetscount',
    method: 'get',
    params: query
  })
}
// 区块信息
export function GetWalletInfo(query) {
  return request({
    url: '/api/v1/getWalletInfo',
    method: 'get',
    params: query
  })
}

// 彩票手续费总数
export function GetLotteryFreeCount(query) {
  return request({
    url: '/api/v1/LotteryFreeCount',
    method: 'get',
    params: query
  })

}
