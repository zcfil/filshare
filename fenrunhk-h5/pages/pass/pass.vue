<template>
	<view class="pass">
		<uni-forms ref="form" :modelValue="form" :rules="rules">
			<uni-forms-item label="旧密码" required name="oldPassword">
				<uni-easyinput type="password" v-model="form.oldPassword" maxlength="18" placeholder="请输入旧密码" />
			</uni-forms-item>
			<uni-forms-item label="新密码" required name="newPassword">
				<uni-easyinput type="password" v-model="form.newPassword" maxlength="18" placeholder="请输入新密码" />
			</uni-forms-item>
			<uni-forms-item label="确认密码" required name="password">
				<uni-easyinput type="password" v-model="form.password" maxlength="18" placeholder="请再次输入新密码" />
			</uni-forms-item>
		</uni-forms>
		<button type="primary" @click="submit">修改</button>
	</view>
</template>

<script>
	import { changePassword } from '../../common/api.js';
	export default {
		data() {
			return {
				form:{
					oldPassword:undefined,
					newPassword:undefined,
					password:undefined
				},
				rules: {
					oldPassword: {
						rules: [
							{required: true,errorMessage: '请输入旧密码',},
							{pattern: /^\d{6,18}$/,errorMessage: '密码长度为6到18位的数字',}
						]
					},
					newPassword: {
						rules: [
							{required: true,errorMessage: '请输入新密码',},
							{pattern: /^\d{6,18}$/,errorMessage: '密码长度为6到18位的数字',}
						]
					},
					password: {
						rules: [
							{required: true,errorMessage: '请输入新密码',},
							{pattern: /^\d{6,18}$/,errorMessage: '密码长度为6到18位的数字',}
						]
					},
				}
			}
		},
		methods: {
			submit() {
				if (this.form.newPassword!=this.form.password) {
					this.$tip.error('两次密码不一致')
					return
				}
				changePassword(this.form).then(res=> {
					this.$tip.success('密码修改成功')
					this.form={
						oldPassword:undefined,
						newPassword:undefined,
						password:undefined
					}
				})
			}
		}
	}
</script>

<style>
.pass {
	padding: 20px;
}
</style>
