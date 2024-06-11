<template>
  <el-form ref="form" :model="user" :rules="rules" label-width="80px">
    <el-form-item label="旧密码" prop="oldpassword">
      <el-input v-model="user.oldpassword" placeholder="请输入旧密码" type="password" />
    </el-form-item>
    <el-form-item label="新密码" prop="newpassword">
      <el-input v-model="user.newpassword" placeholder="请输入新密码" type="password" />
    </el-form-item>
    <el-form-item label="确认密码" prop="confirmpassword">
      <el-input v-model="user.confirmpassword" placeholder="请确认密码" type="password" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" size="mini" @click="submit">保存</el-button>
      <el-button type="danger" size="mini" @click="resetQuery">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { updateUserPwd } from '@/api/wallets/wallets'

export default {
  data() {
    const equalToPassword = (rule, value, callback) => {
      if (this.user.newpassword !== value) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    }
    return {
      test: '1test',
      user: {
        oldpassword: undefined,
        newpassword: undefined,
        confirmpassword: undefined
      },
      // 表单校验
      rules: {
        oldpassword: [
          { required: true, message: '旧密码不能为空', trigger: 'blur' }
        ],
        newpassword: [
          { required: true, message: '新密码不能为空', trigger: 'blur' },
          { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
        ],
        confirmpassword: [
          { required: true, message: '确认密码不能为空', trigger: 'blur' },
          { required: true, validator: equalToPassword, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submit() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          updateUserPwd(this.user).then(
            response => {
              if (response.code === 200) {
                this.resetForm('form')
                this.msgSuccess('修改成功')
              } else {
                this.msgError(response.msg)
              }
            }
          )
        }
      })
    },
    resetQuery() {
      this.user = []
      this.resetForm('form')
    }
  }
}
</script>
