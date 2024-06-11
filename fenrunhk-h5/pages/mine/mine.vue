<template>
	<view class="mine">
		<view class="img">
			<view><image src="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80" mode=""></image></view>
			<button style="margin: 10px 0 20px;" type="warn" size="mini" @click="logut">退出登录</button>
		</view>
		<view class="menu" @click="pass">
			修改密码 
			<view><image src="../../static/arrow-right.png" mode=""></image></view>
		</view>
		<view class="menu" style="border-bottom: 1px solid #cccccc;margin-bottom: 20px;" @click="migrate">
			迁移数据列表
			<view><image src="../../static/arrow-right.png" mode=""></image></view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
		},
		methods: {
			logut() {
				this.$tip.success('退出成功')
				uni.removeStorageSync('token');
				uni.reLaunch({url:'/pages/login/login'});
			},
			pass() {
				uni.navigateTo({url:'/pages/pass/pass'})
			},
			migrate() {
				uni.navigateTo({url:'/pages/migrate/migrate'})
			},
		}
	}
</script>

<style>
.img {
	text-align: center;
	padding: 20px;
}
.img image{
	width: 100px;
	height: 100px;
	border-radius: 50px;
}
.menu {
	padding: 20px 20px;
	border-top: 1px solid #cccccc;
	position: relative;
	font-size: 15px;
}
.menu view {
	position: absolute;
	top: 20px;
	right: 20px;
}
.menu view image {
	width: 22px;
	height: 22px;
}
</style>
