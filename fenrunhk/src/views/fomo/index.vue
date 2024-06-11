<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="用户账号" prop="cid">
        <el-input
          v-model="queryParams.keyword"
          placeholder="请输入用户账号"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>

      <el-form-item label="投注时间">
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

    <!--添加修改删除操作开始-->
    <el-table v-loading="loading" :data="list" @selection-change="handleSelectionChange">
      <el-table-column label="用户账号" align="center" prop="userId" />
      <el-table-column label="投注数量" align="center" prop="count" />
      <el-table-column label="分红" align="center" prop="dividend" />
      <el-table-column label="分红是否翻倍" align="center" prop="is_double">
        <template #default="props">
          <div class="cell" v-if="props.row.is_double">是</div>
          <div class="cell" v-else>否</div>
        </template>

      </el-table-column>
      <el-table-column label="key的单价" align="center" prop="price" />
      <el-table-column label="购买总价" align="center" prop="total_price" />
      <el-table-column label="投注时间" align="center" prop="created_at" width="180" >
        <template #default="props">
          <div class="cell">{{new Date(props.row.created_at).format('yyyy-MM-dd hh:mm:ss')}}</div>
        </template>
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
import { fomoRecords, exportExcel } from '@/api/fomo/fomo'

export default {
  name: 'Fomo',
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
      list: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 类型数据字典
      typeOptions: [],
      // 日期范围
      dateRange: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
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
    // this.getDicts('sys_yes_no').then(response => {
    //   this.typeOptions = response.data
    // })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      fomoRecords(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.list = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    // 参数系统内置字典翻译
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions, row.configType)
    },
    // 日期处理
    dateFormat(row, column) {
      var date = row[column.property]
      if (date === undefined) {
        return ''
      }
      // eslint-disable-next-line no-undef
      return moment(date).format('YYYY-MM-DD HH:mm:ss')
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
