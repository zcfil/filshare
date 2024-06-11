<template>
	<view class="customer">
		<view class="tab" v-for="(item,index) in data" :key="index">
			<view>算力（TB）：{{item.storage}}</view>
			<view>备注：{{item.remark}}</view>
			<view>创建时间：{{parseTime(item.create_time)}}</view>
			<view>状态：
				<text style="color: rgb(31, 211, 0);" v-show="item.status==0">进行中</text>
				<text style="color: rgb(255, 80, 80);" v-show="item.status==1">已终止</text>
				<text style="color: rgb(31, 211, 0);" v-show="item.status==2">线性释放中</text>
				<text style="color: #007AFF;" v-show="item.status==3">已完成线性释放</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { investmentList } from '../../common/api.js';
	export default {
		data() {
			return {
				data:[],
				queryParams: {
				    pageIndex: 1,
				    pageSize: 10,
				},
				loadStatus:'loading',  //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				isLoadMore:false,  //是否加载中
			}
		},
		onShow() {
			this.data=[]
			this.auto()
		},
		onReachBottom(){  //上拉触底函数
		    if(!this.isLoadMore){  //此处判断，上锁，防止重复请求
		        this.isLoadMore=true
		        this.queryParams.pageIndex+=1
		        this.auto()
		    }
		},
		onPullDownRefresh() {
			this.data=[]
			this.queryParams.pageIndex=1
			this.auto()
		    uni.stopPullDownRefresh();
		},
		methods: {
			auto() {
				investmentList(this.queryParams).then(res=> {
					if(res.list){
					    this.data=this.data.concat(res.list)
					    if(res.list.length<this.queryParams.pageSize){  //判断接口返回数据量小于请求数据量，则表示此为最后一页
					        this.isLoadMore=true                                             
					        this.loadStatus='nomore'
					    }else{
					        this.isLoadMore=false
					    }
					}else{
					    this.isLoadMore=true
					    this.loadStatus='nomore'
					}
				})
			},
		}
	}
</script>

<style>
.tab {
	font-size: 14px;
	color: #444444;
	border-top: 1px solid #dddddd;
	padding: 11px 20px 15px;
}
.tab view {
	margin-top: 4px;
}	
</style>
