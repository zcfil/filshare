<template>
	<view class="index">
		<view class="card">
			<view class="today">
				<view>总算力（TB）</view>
				<view>{{data.total_hashrate}}</view>
			</view>
			<view class="income">
				<view>
					<view>总收益</view>
					<view>{{data.total_income}}</view>
				</view>
				<view>
					<view>已释放</view>
					<view>{{data.release_fil}}</view>
				</view>
				<view>
					<view>锁仓</view>
					<view>{{data.lock_fil}}</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import { homepage } from '../../common/api.js';
	export default {
		data() {
			return {
				data:{}
			}
		},
		onShow() {
			this.auto()
		},
		methods: {
            auto() {
				homepage().then(res=> {
					this.data=res.data
				})
			}
		}
	}
</script>

<style>
.index {
	background: #f8f8f8;
	padding: 20px;
	box-sizing: border-box;
}
.card {
	height: 200px;
	background: #d4bf9d;
	border-radius: 10px;
}
.today {
	text-align: center;
	padding: 20px 0 32px;
}
.today view:nth-child(1) {
	font-size: 14px;
	color: #555555;
}
.today view:nth-child(2) {
	font-size: 36px;
	color: #333333;
}
.income {
	display: flex;
	text-align: center;
}
.income>view {
	flex: 1;
}
.income>view view:nth-child(1) {
	font-size: 14px;
	color: #555555;
}
.income>view view:nth-child(2) {
	font-size: 19px;
	color: #333333;
}
</style>
