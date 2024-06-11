<template>
  <div class="ratio">
    <el-table :data="list">
      <el-table-column label="配置名" prop="name" align="center" show-overflow-tooltip  />
      <el-table-column label="配置值" prop="value" align="center" show-overflow-tooltip />
      <el-table-column label="状态" align="center">
        <template slot-scope="scope">
          <el-switch
            v-model="scope.row.status"
            active-value="0"
            inactive-value="1"
            active-text="启用" 
            inactive-text="禁用"
            active-color="#1890ff"
            inactive-color="#c0c0c0"
            @change="handleStatusChange(scope.row)"
          />
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:sysratio:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 添加或修改客户 -->
    <el-dialog title="修改配置" :visible.sync="open" width="500px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="配置名" prop="name">
          <el-input v-model="form.name" disabled placeholder="请输入配置名" />
        </el-form-item>
        <el-form-item label="配置值" prop="value">
          <el-input v-model="form.value" type="textarea" placeholder="请输入配置值" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { financeConfigList, financeConfigEdit } from '@/api/customer/customer'
export default {
  name: 'Ratio',
  components: {},
  props: {},
  data() {
    return {
      list: [],
      open: false,
      form: {},
      rules: {
        name: [
          { required: true, message: '配置名不能为空', trigger: 'blur' }
        ],
        value: [
          { required: true, message: '配置值不能为空', trigger: 'blur' },
          {
            pattern: /^\d+$|^\d*\.\d+$/g,
            message: '请输入数字和一个小数点',
            trigger: 'blur'
          }
        ],
      }
    }
  },
  computed: {},
  watch: {},
  created() {
    this.auto()
  },
  mounted() {},
  methods: {
    auto() {
      financeConfigList().then(res => {
        this.list = res.data
      })
    },
    handleUpdate(row) {
      this.auto()
      this.open = true
      this.form = row
      console.log(row.value.length)
    },
    submitForm() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          financeConfigEdit(this.form).then(res => {
            this.msgSuccess('修改成功')
            this.open = false
            this.auto()
          })
        }
      })
    },
    cancel() {
      this.open = false
    },
    // 状态修改
    handleStatusChange(row) {
      const text = row.status === '0' ? '启用' : '禁用'
      this.$confirm('确认要"' + text + '""' + row.name + '"配置吗?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return financeConfigEdit(row)
      }).then(() => {
        this.msgSuccess(text + '成功')
        this.auto()
      }).catch(function() {
        row.status = row.status === '0' ? '1' : '0'
      })
    },
  }
}
</script>

<style scoped>
.ratio {
    padding: 20px;
}
</style>
