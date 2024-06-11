<template>
  <el-form id="bindfrom" ref="form" :model="user" :rules="rules" label-width="120px">
    <el-form-item label="扫码">
      <div id="qrcode" ref="qrCodeUrl" style="width: 100px;height: 100px;float: left;" />
    </el-form-item>
    <el-form-item label="私钥" prop="secret">
      <el-input v-model="user.googleSecret" type="text" />
      <button>复制</button>
    </el-form-item>
    <el-form-item label="谷歌验证码" prop="confirmPassword">
      <el-input v-model="user.code" placeholder="请输入谷歌验证码" type="text" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" size="mini" @click="submit">保存</el-button>
      <el-button type="danger" size="mini" @click="close">关闭</el-button>
    </el-form-item>
  </el-form>
</template>

<script>

import { bindVerification, getSecretCode } from '@/api/system/sysuser'
import QRCode from 'qrcodejs2'
export default {
  data() {
    return {
      user: {
        googleSecret: undefined,
        googleAccount: undefined,
        url: undefined,
        code: undefined,
        isBind: undefined
      },
      // 表单校验
      rules: {
        code: [
          { required: true, message: '验证码不能为空！', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getSecret()
  },

  methods: {
    getSecret() {
      getSecretCode().then(response => {
        // eslint-disable-next-line no-cond-assign
        this.user = response.data
        if (!this.user.isBind) {
          this.qrcode(this.user.url) // 调用二维码生成的方法
        }
      })
    },
    qrcode(url) {
      // eslint-disable-next-line no-unused-vars
      const qrCode = new QRCode('qrcode', {
        text: url, // 需要转换为二维码的内容
        width: 100,
        height: 100,
        colorDark: '#000000',
        colorLight: '#ffffff',
        correctLevel: QRCode.CorrectLevel.H
      })
    },
    submit() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          bindVerification(this.auth.googleSecret, this.auth.code).then(
            response => {
              if (response.code === 200) {
                this.msgSuccess(response.msg)
              } else {
                this.msgError(response.msg)
              }
            }
          )
        }
      })
    },
    close() {
      this.$store.dispatch('tagsView/delView', this.$route)
      this.$router.push({ path: 'profile/index' })
    }
  }
}
</script>
