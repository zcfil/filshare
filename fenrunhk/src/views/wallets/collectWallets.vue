<template>
  <el-form ref="form" :model="balls" :rules="rules" label-width="80px">
    <el-form-item label="钱包地址" prop="address">
      <el-input v-model="balls.address" placeholder="请输入钱包地址" />
    </el-form-item>
    <el-form-item label="归集币种" prop="name">
      <el-select v-model="balls.name" placeholder="选择币种" clearable size="small">
        <el-option
          v-for="dict in typeOptions"
          :key="dict.dictValue"
          :label="dict.dictLabel"
          :value="dict.dictValue"
        />
      </el-select>
    </el-form-item>
    <el-form-item label="归集数量" prop="count">
      <el-input v-model="balls.count" placeholder="请输入归集数量" />
    </el-form-item>
    <el-form-item label="钱包密码" prop="password">
      <el-input v-model="balls.password" placeholder="请输入钱包密码" type="password" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" size="mini" @click="submit">归集</el-button>
      <el-button type="danger" size="mini" @click="resetQuery">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { collectWalletsAsset } from '@/api/wallets/wallets'

export default {
  data() {
    return {
      // 类型数据字典
      typeOptions: [],
      balls: {
        address: undefined,
        count: undefined,
        password: undefined,
        name: undefined
      },
      // 表单校验
      rules: {
        address: [
          { required: true, message: '钱包地址不能为空', trigger: 'blur' }
        ],
        count: [
          { required: true, message: '归集数量不能为空', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '钱包密码不能为空', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '请选择归集币种', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getDicts('asset_type').then(response => {
      this.typeOptions = response.data
    })
  },
  methods: {
    resetQuery() {
      this.balls = []
      this.resetForm('form')
    },
    submit() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          collectWalletsAsset(this.balls).then(
            response => {
              if (response.code === 200) {
                this.resetForm('form')
                this.msgSuccess(response.data)
              } else {
                this.msgError(response.msg)
              }
            }
          )
        }
      })
    }
  }
}
</script>
