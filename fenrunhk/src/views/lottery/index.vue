<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="搜索" prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="期数/彩票HASH"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="开奖时间">
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

    <!--添加修改删除操作结束-->

    <el-table v-loading="loading" :data="lotteryList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="期数" align="center" prop="lottery_period" />
      <el-table-column label="开奖区块高度" align="center" prop="block_height" />
      <el-table-column label="彩票HASH" align="center" prop="block_hash" :show-overflow-tooltip="true" />
      <el-table-column label="开奖号码" align="center" prop="lottery_number_array" />
      <el-table-column label="号码位置" align="center" prop="lottery_number_index" />
      <el-table-column label="总投注金额" align="center" prop="bet_amount" />
      <el-table-column label="平台盈利" align="center" prop="profit" />
      <el-table-column label="开奖时间" align="center" prop="lottery_date" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.lottery_date) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:sysconfig:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row.lottery_period)"
          >详情</el-button>
          <!-- <el-button
            v-permisaction="['system:sysconfig:remove']"
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除</el-button> -->
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

    <!-- 详情对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="1140px" center>
      <el-table :data="lotteryDetailsList">
        <el-table-column property="cid" label="UID" width="150" />
        <el-table-column property="bet_quantity" label="投注数量" width="150" />
        <el-table-column property="bet_amount" label="投注金额" width="150" />
        <el-table-column property="bet_content" label="投注详情" width="150" />
        <el-table-column property="winning_amount" label="中奖数量" width="150" />
        <el-table-column property="free" label="平抽佣" width="150" />
        <el-table-column property="bet_date" label="投注时间" width="180" />

      </el-table>
      <pagination
        v-show="total>0"
        :total="detailsTotal"
        :page.sync="detailsParams.pageIndex"
        :limit.sync="detailsParams.pageSize"
        @pagination="handleUpdate"
      />
    </el-dialog>
  </div>
</template>

<script>
import { listLottery, exportExcel, lotteryDetails } from '@/api/lottery/lottery'


export default {
  name: 'CoinCharge',
  data() {
    return {
      open: false,
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
      lotteryList: [],
      // 参数表格数据
      lotteryDetailsList: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      dialogTableVisible: false,
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
      detailsParams: {
        pageIndex: 1,
        pageSize: 10,
        cid: undefined
      },
      detailsId: undefined,
      detailsTotal: 0,
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
      listLottery(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.lotteryList = response.data.list
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
    // 详情
    handleUpdate(row) {
      // this.reset()
      if (this.detailsId === undefined) {
        this.detailsId = row
      }

      const configId = this.detailsId
      this.detailsParams.period = configId
      lotteryDetails(this.detailsParams).then(response => {
        if (response.code == 200) {
          // this.msgSuccess('修改成功')
          // this.open = false
          // this.getDetails
          this.lotteryDetailsList = response.data.list
          // for (var i=0;i<lotteryDetailsList;i++){
          //   this.lotteryDetailsList
          // }
          this.detailsTotal = response.data.count
        } else {
          this.msgError(response.msg)
        }
      })
      if (configId != null) {
        // this.reset()
        // this.form = response.data
        this.open = true
        this.title = '第' + configId + '期'
      }
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
