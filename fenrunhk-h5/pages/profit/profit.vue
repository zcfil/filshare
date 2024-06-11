<template>
	<view class="profit">
		<view class="btm">
			<view :class="{active:num=='2'}" @click="carry">收账</view>
			<view :class="{active:num=='1'}" @click="charge">转账</view>
		</view>
		<view v-if="num=='1'">
			<view class="tab" v-for="(item,index) in data" :key="index">
				<view>金额：<text style="color: #007AFF;">{{item.amount}}</text></view>
				<view>到账时间：{{parseTime(item.create_time)}}</view>
			</view>
		</view>
		<view v-if="num=='2'">
			<view class="tab" v-for="(item,index) in data" :key="index">
				<view>客户收益：<text style="color: #007AFF;">{{item.customer_income}}</text></view>
				<view>到客户余额：<text style="color: #007AFF;">{{item.to_customer_balance}}</text></view>
				<view>到客户锁仓：<text style="color: #007AFF;">{{item.to_customer_lock}}</text></view>
				<!-- <view>客户锁仓释放：<text style="color: #007AFF;">{{item.customer_lock_release}}</text></view> -->
				<view>结算时间：{{parseTime(item.time)}}</view>
				<view class="is_transfer">状态：
					<text style="color: rgb(255, 80, 80);" v-show="item.is_transfer==0">未结算</text>
					<text style="color: #007AFF;" v-show="item.is_transfer==1">已结算</text>
				</view>
			</view>
		</view>
		<view v-show="isLoadMore">
		    <uni-load-more :status="loadStatus" ></uni-load-more>
		</view>
	</view>
</template>

<script>
	import { customerTransferList,settlementList } from '../../common/api.js';
	export default {
		data() {
			return {
				num:'2',
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
			this.auto();
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
				if (this.num=='1') {
					customerTransferList(this.queryParams).then(res=> {
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
				}else if(this.num=='2') {
					settlementList(this.queryParams).then(res=> {
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
				}
			},
			charge() {
				this.num=1
				this.data=[]
				this.queryParams.pageIndex=1
				this.auto()
			},
			carry() {
				this.num=2
				this.data=[]
				this.queryParams.pageIndex=1
				this.auto()
			},
		}
	}
</script>

<style>
.btm {
	height: 44px;
	line-height: 44px;
	display: flex;
	font-size: 15px;
	padding: 0 20px;
}
.btm view {
	width: 70px;
	text-align: center;
}
.tab {
    font-size: 14px;
	color: #444444;
	border-top: 1px solid #dddddd;
	padding: 11px 20px 15px;
	position: relative;
}
.tab view {
	margin-top: 4px;
}
.active {
	font-size: 18px;
	color: #007AFF;
}
</style>
