<template>
  <div>
    <el-tabs v-model="activeName" type="border-card" style="margin-top: 30px;margin-left: 30px; height: 500px; width:1200px;">
      <!--<el-tab-pane label="修改钱包密码" name="resetPwd">
        <resetPwd />
      </el-tab-pane>-->
      <!--<el-tab-pane label="手动资金归集">
        <collectWallets />
      </el-tab-pane>-->
      <el-tab-pane label="异常充值处理"  name="exceptionCoin">
        <exceptionCoin />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import resetPwd from './resetPwd'
import collectWallets from './collectWallets'
import exceptionCoin from './exceptionCoin'
export default {
  components: { resetPwd, collectWallets, exceptionCoin },
  data() {
    return {
      activeName: 'exceptionCoin'
    }
  },
  methods: {

  }
}
</script>
