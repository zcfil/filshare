<template>
<div>
  <el-form ref="form" :model="user" label-width="80px">
    <el-form-item label="txid" prop="txid">
      <el-input v-model="user.txid" placeholder="请输入交易txid" />
    </el-form-item>
    <!--<el-form-item label="区块高度" prop="blockhight">
      <el-input v-model="user.blockhight" placeholder="请输入区块高度blockhight" />
    </el-form-item>-->

    <el-form-item>
      <el-button type="primary" size="mini" @click="submit">保存</el-button>
      <el-button type="danger" size="mini" @click="resetQuery">重置</el-button>
    </el-form-item>
  </el-form>
  <!-- 详情对话框 -->
    <el-dialog title="充值成功" :visible.sync="open" width="800px" center>
      <div>
        <span>交易编号：{{coinDetail.hash}}</span><br><br>
        <span>币种名称：{{coinDetail.symbol}}</span><br><br>
        <span>交易数量：{{coinDetail.quantity}}</span><br><br>
        <span>交易地址：{{coinDetail.from}}</span><br><br>
        <span>充值地址：{{coinDetail.to}}</span><br><br>
        <span>充值时间：{{parseTime(coinDetail.confirm_time)}}</span><br>
      </div>
    </el-dialog>
 </div>
</template>
<script>
import { exceptionCoinCharge } from '@/api/wallets/wallets'

export default {
  data() {
    // const equalToPassword = (rule, value, callback) => {
    //   if (this.user.newpassword !== value) {
    //     callback(new Error('两次输入的密码不一致'))
    //   } else {
    //     callback()
    //   }
    // }
    return {
      test: '1test',
      user: {
        txid: undefined,
        blockhight: undefined
      },
      open: false,
      coinDetail:[]
      // 表单校验
      // rules: {
      //   oldpassword: [
      //     { required: true, message: '旧密码不能为空', trigger: 'blur' }
      //   ],
      //   newpassword: [
      //     { required: true, message: '新密码不能为空', trigger: 'blur' },
      //     { min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur' }
      //   ],
      //   confirmpassword: [
      //     { required: true, message: '确认密码不能为空', trigger: 'blur' },
      //     { required: true, validator: equalToPassword, trigger: 'blur' }
      //   ]
      // }
    }
  },
  methods: {
    submit() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          exceptionCoinCharge(this.user).then(
            response => {
              if (response.code === 200) {
                this.msgSuccess('充值成功')
                this.coinDetail = response.data
                this.open = true
                this.resetForm('form')  
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
