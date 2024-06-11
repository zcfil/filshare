<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form
      ref="queryForm"
      :model="queryParams"
      :inline="true"
      label-width="68px"
    >
      <el-form-item label="搜索" prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="用户账号/提币地址"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="资产类型" prop="name">
        <el-select
          v-model="queryParams.name"
          placeholder="资产类型"
          clearable
          size="small"
        >
          <el-option
            v-for="dict in typeOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictLabel"
          />
        </el-select>
      </el-form-item>

      <!-- <el-form-item label="提币状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="提币状态" clearable size="small">
          <el-option
            v-for="dict in typeStatus"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictValue"
          />
        </el-select>
      </el-form-item>-->

      <el-form-item label="创建时间">
        <el-date-picker
          v-model="dateRange"
          size="small"
          style="width: 240px"
          value-format="yyyy-MM-dd"
          type="daterange"
          range-separator="-"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          icon="el-icon-search"
          size="mini"
          @click="handleQuery"
        >搜索</el-button>
        <el-button
          icon="el-icon-refresh"
          size="mini"
          @click="resetQuery"
        >重置</el-button>
      </el-form-item>
    </el-form>
    <!--搜索结束-->

    <!--添加修改删除操作开始-->
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="success"
          size="mini"
          :disabled="single"
          @click="handleCheck"
        >审核</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:edit']"
          type="success"
          icon="el-icon-remove"
          size="mini"
          :disabled="multiple"
          @click="handleCancel"
        >取消提币</el-button>
      </el-col>
    </el-row>
    <!--添加修改删除操作结束-->

    <el-table
      v-loading="loading"
      :data="drawList"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="用户账号" align="center" prop="uid" />
      <el-table-column label="资产名称" align="center" prop="symbol" />
      <el-table-column label="提币数量" align="center" prop="value" />
      <el-table-column label="实际到账" align="center" prop="amount" />
      <el-table-column label="手续费" align="center" prop="free" />
      <el-table-column
        label="提币状态"
        align="center"
        prop="status"
        :formatter="statusFormat"
      />
      <el-table-column
        label="提币地址"
        align="center"
        prop="to_address"
        :show-overflow-tooltip="true"
      />
      <el-table-column label="创建时间" align="center" prop="create_time" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="queryParams.pageIndex"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />

    <!-- 添加或修改参数配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="500px">
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="钱包密码" prop="pwd">
          <el-input
            v-model="form.pwd"
            show-password
            placeholder="请输入钱包密码"
          />
        </el-form-item>
        <el-form-item label="用户账号" prop="uid">
          <el-input v-model="form.uid" :disabled="true" />
        </el-form-item>
        <el-form-item label="币种名称" prop="symbol">
         <el-input v-model="form.symbol" :disabled="true" />
        </el-form-item>
        <el-form-item label="提币数量" prop="value">
          <el-input v-model="form.value" :disabled="true" />
        </el-form-item>
        <el-form-item label="实际到账" prop="amount">
          <el-input v-model="form.amount" :disabled="true" />
        </el-form-item>
        <el-form-item label="提币地址" prop="to_address">
          <el-input v-model="form.to_address" :disabled="true" />
        </el-form-item>
        <el-form-item label="备注" prop="memo">
          <el-input
            v-model="form.memo"
            type="textarea"
            placeholder="请输入内容"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm" :disabled="loading1">审核通过</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  lsitDraws,
  getDraws,
  drawCancel,
  drawCheck
} from '@/api/financial/withdraw'

export default {
  name: 'Draws',
  data() {
    return {
      // 遮罩层
      loading: true,
      loading1: false,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 参数表格数据
      drawList: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 类型数据字典
      typeOptions: [],
      // 提币状态数据字典
      typeStatus: [],
      // 日期范围
      dateRange: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        status: 0,
        cid: undefined
      },
      // 删除参数
      delParams: {
        ids: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        configName: [
          { required: true, message: '参数名称不能为空', trigger: 'blur' }
        ],
        configKey: [
          { required: true, message: '参数键名不能为空', trigger: 'blur' }
        ],
        configValue: [
          { required: true, message: '参数键值不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('asset_type').then(response => {
      this.typeOptions = response.data
    })
    this.getDicts('coin_status').then(response => {
      this.typeStatus = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      lsitDraws(this.addDateRange(this.queryParams, this.dateRange)).then(
        response => {
          this.drawList = response.data.list
          this.total = response.data.count
          this.loading = false
        }
      )
    },
    // 参数系统内置字典翻译
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions, row.status)
    },
    // 提币状态内置字典翻译
    statusFormat(row, column) {
      return this.selectDictLabel(this.typeStatus, row.status)
    },
    // 取消按钮`
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        configId: undefined,
        configName: undefined,
        configKey: undefined,
        configValue: undefined,
        configType: 'Y',
        remark: undefined
      }
      this.resetForm('form')
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    // 提交按钮
    submitForm: function() {
      this.loading1 = true
      drawCheck(this.form)
        .then(response => {
          this.loading1 = false
          if (response.code == 200) {
            this.msgSuccess('修改成功')
            this.open = false
            this.getList()
          } else {
            this.msgError(response.msg)
          }
        })
        .catch(res => {
          this.loading1 = false
          console.log(res)
        })
    },
    /** 审核提币操作 */
    handleCheck(row) {
      this.reset()
      const id = row.id || this.ids
      getDraws(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '审核提币'
      })
    },
    /** 取消按钮操作 */
    handleCancel(row) {
      const configIds = row.id || this.ids
      this.delParams.ids = configIds
      var ids = ''
      for (var i = 0; i < configIds.length; i++) {
        ids += configIds[i]
        if (i < configIds.length - 1) {
          ids += ','
        } else {
          this.delParams.ids = ids
        }
      }
      this.$confirm('是否确认取消"' + configIds + '"的提币申请?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        drawCancel(this.delParams)
          .then(() => {
            this.msgSuccess('取消成功')
            this.getList()
          })
          .catch(function() {})
      })
    }
  }
}
</script>
