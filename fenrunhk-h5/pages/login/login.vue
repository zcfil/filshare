<template>
	<view class="login">
		<view class="welcome">用户登录</view>
		<uni-forms class="content" ref="form" :modelValue="loginForm" :rules="rules">
			<uni-forms-item name="phone">
				<uni-easyinput type="text" v-model="loginForm.phone" maxlength="11" placeholder="请输入手机号" />
			</uni-forms-item>
			<uni-forms-item name="password">
				<uni-easyinput type="password" v-model="loginForm.password" maxlength="18" placeholder="请输入密码" />
			</uni-forms-item>
		</uni-forms>
		<button class="btn" @click="submit">登录</button>
	</view>
</template>

<script>
	import { login } from '../../common/api.js';
	export default {
		data() {
			return {
				loginForm:{
					phone:undefined,
					password:undefined,
				},
				rules: {
					phone: {
						rules: [
							{required: true,errorMessage: '请输入手机号',},
							{
							    pattern: /^1[3|4|5|7|8|9][0-9]\d{8}$/,
							    errorMessage: '请输入正确的手机号码',
							}
						]
					},
					password: {
						rules: [
							{required: true,errorMessage: '请输入密码',},
							{pattern: /^\d{6,18}$/,errorMessage: '密码长度为6到18位的数字',}
						]
					},
				}
			}
		},
		methods: {
			submit() {
				this.$refs.form.validate().then(valid => {
					login(this.loginForm).then(res=> {
						this.$tip.success('登录成功')
						uni.setStorageSync('token',res.data);
						uni.switchTab({url:'/pages/index/index'});
					})
				})
			},
		}
	}
</script>

<style>
.login {
	position: relative;
	width: 100vw;
	height: 94vh;
	overflow: hidden;
	background: #f8f8f8;
}
.welcome {
	margin: 65px 0 95px 27px;
	font-size: 25px;
	color: #333;
	text-shadow: 1px 0 1px rgb(0 0 0 / 30%);
}
.content {
	padding: 0 44px;
	font-size: 20px;
}
.btn {
	margin: 20px 44px;
	background: #beac8b;
	color: #ffffff;
	height: 40px;
	font-size: 16px;
}
</style>
