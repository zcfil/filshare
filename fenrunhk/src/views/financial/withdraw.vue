<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="搜索" prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="CID/提币地址"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>

      <el-form-item label="资产类型" prop="configType">
        <el-select v-model="queryParams.name" placeholder="资产类型" clearable size="small">
          <el-option
            v-for="dict in typeOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictLabel"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="提币时间">
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
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <!-- <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button> -->
      </el-form-item>

      <el-form-item>
        <el-col :span="1.5">
          <el-button
            v-permisaction="['system:sysconfig:export']"
            type="warning"
            icon="el-icon-download"
            size="mini"
            @click="handleExport"
          >导出</el-button>
        </el-col>

      </el-form-item>
    </el-form>
    <!--搜索结束-->

    <!-- 添加修改删除操作开始
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:edit']"
          type="success"
          icon="el-icon-edit"
          size="mini"
          :disabled="single"
          @click="handleUpdate"
        >修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:export']"
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
        >导出</el-button>
      </el-col>
    </el-row> -->
    <!--添加修改删除操作结束-->

    <el-table v-loading="loading" :data="drawList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="用户账号" align="center" prop="uid" />
      <el-table-column label="资产名称" align="center" prop="symbol" />
      <el-table-column label="数量" align="center" prop="value" />
      <el-table-column label="实际到账" align="center" prop="amount" />
      <el-table-column label="手续费" align="center" prop="free" />
      <el-table-column label="提币地址" align="center" prop="to_address" width="180" :show-overflow-tooltip="true" />
      <el-table-column label="TXID" align="center" prop="txid" width="180"  :show-overflow-tooltip="true" />
      <el-table-column label="提币时间" align="center" prop="create_time" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="确认时间" align="center" prop="confirm_time" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.confirm_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="审核人" align="center" prop="reviewer" class-name="small-padding fixed-width" />
      <el-table-column label="状态" align="center" prop="status" width="180" :formatter="statusFormat">
        <!-- <template v-if="1">
          <span> 提币成功 </span>
        </template> -->
      </el-table-column>
    </el-table>

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageIndex"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />

  </div>
</template>

<script>
import { lsitDraws, exportExcel } from '@/api/financial/withdraw'

export default {
  name: 'Draws',
  data() {
    return {
      // 遮罩层
      loading: true,
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
        status: '1,2,3,4',
        cid: undefined
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
      lsitDraws(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
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
    // 取消按钮
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
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.open = true
      this.title = '添加参数'
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.configId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 导出按钮操作 */
    handleExport() {
      const queryParams = this.queryParams
      this.$confirm('是否确认导出所有参数数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        exportExcel(queryParams).then(res => {
          location.href = process.env.VUE_APP_BASE_API + `/` + res.msg
        })
      })
    }
  }
}
</script>
