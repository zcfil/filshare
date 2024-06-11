import request from '@/utils/request'

// 查询兑换释放记录
export function listRelease(query) {
  return request({
    url: '/api/v1/betreleases',
    method: 'get',
    params: query
  })
}
// 导出兑换释放记录
export function exportExcel(query) {
  return request({
    url: '/api/v1/betreleasesExport',
    method: 'get',
    params: query
  })
}
