<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="充值账号" prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="用户账号组"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>

      <!--
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
      -->

      <el-form-item label="币种名称" prop="status">
        <el-select v-model="queryParams.name" placeholder="币种名称" clearable size="small">
          <el-option
            v-for="dict in typeOptions1"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictLabel"
          />
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <!--<el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>-->
      </el-form-item>

    </el-form>
    <!--搜索结束-->

    <!--添加修改删除操作开始-->
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >手动充币</el-button>
      </el-col>

    </el-row>
    <!--添加修改删除操作结束-->

    <el-table v-loading="loading" :data="memberList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="编号" width="80" align="center" prop="id" />
      <el-table-column label="充值账号组" align="center" prop="cids" />
      <el-table-column label="币种名称" align="center" prop="name" />
      <el-table-column label="充币数量" align="center" prop="value" />
      <el-table-column label="操作人" align="center" prop="operator" />
      <el-table-column label="充值类型" align="center" prop="chargetype" :formatter="typeFormat" :show-overflow-tooltip="true" />
      <el-table-column label="充值去向" align="center" prop="chargeto" :formatter="typeFormat3" />
      <el-table-column label="备注" align="center" prop="remark" :show-overflow-tooltip="true" />
      <el-table-column label="充值时间" align="center" prop="coin_time" width="200">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.coin_time) }}</span>
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

    <!-- 添加或修改参数配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" :rules="rules" width="520px">
      <el-form ref="form" :model="addParams" :rules="rules" label-width="80px">
        <el-form-item label="充值账号" prop="cids">
          <el-input v-model="addParams.cids" placeholder="请输入充值账号组如：1,2,3,4,5" type="textarea" :autosize="{ minRows: 4, maxRows: 6}" />
        </el-form-item>
        <el-form-item label="充币名称" prop="name">
          <el-select v-model="addParams.name" placeholder="充币名称" clearable size="small" style="width:100%">
            <el-option
              v-for="dict in typeOptions1"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictLabel"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="充值类型" prop="chargetype">
          <el-select v-model="addParams.chargetype" placeholder="充值类型" clearable size="small" style="width:100%">
            <el-option
              v-for="dict in typeOptions2"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictValue"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="充币数量" prop="value">
          <el-input v-model="addParams.value" placeholder="请输入充币数量" />
        </el-form-item>

        <el-form-item label="充币去向" prop="chargeto">
          <el-select v-model="addParams.chargeto" placeholder="充币去向" clearable size="small" style="width:100%">
            <el-option
              v-for="dict in typeOptions3"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictValue"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="addParams.remark" placeholder="" type="textarea" :autosize="{ minRows: 4, maxRows: 6}" />
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
import { manualCoinList, addManualCoin } from '@/api/financial/manualCoin'

export default {
  name: 'Member',
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
      memberList: [],
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      // 类型数据字典
      typeOptions: [],
      typeOptions1: [],
      typeOptions2: [],
      typeOptions3: [],
      // 会员等级数据字典
      typeLevel: [],
      // 日期范围
      dateRange: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        status: '1',
        keyword: ''
        // cid: undefined
      },
      // 删除参数
      delParams: {
        ids: undefined
      },
      // 增加参数
      addParams: {
        cids: undefined,
        name: undefined,
        chargetype: undefined,
        value: undefined,
        chargeto: undefined,
        remark: undefined,
        operator: undefined
      },
      // 编辑参数
      editParams: {
        email: '',
        address: '',
        status: '0',
        level: undefined
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        cids: [
          { required: true, message: '参数名称不能为空', trigger: 'blur' }
        ],
        name: [
          { required: true, message: '参数键名不能为空', trigger: 'blur' }
        ],
        chargetype: [
          { required: true, message: '参数键值不能为空', trigger: 'blur' }
        ],
        value: [
          { required: true, message: '参数名称不能为空', trigger: 'blur' }
        ],
        chargeto: [
          { required: true, message: '参数键名不能为空', trigger: 'blur' }
        ],
        remark: [
          { required: true, message: '参数键值不能为空', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts('asset_type').then(response => {
      this.typeOptions1 = response.data
    })
    this.getDicts('charge_type').then(response => {
      this.typeOptions2 = response.data
    })
    this.getDicts('charge_to').then(response => {
      this.typeOptions3 = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      manualCoinList(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.memberList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    // 参数系统内置字典翻译
    typeFormat1(row, column) {
      return this.selectDictLabel(this.typeOptions1, row.name)
    },
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions2, row.chargetype)
    },
    typeFormat3(row, column) {
      return this.selectDictLabel(this.typeOptions3, row.chargeto)
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
      // this.reset()
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
      this.title = '新增'
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          addManualCoin(this.addParams).then(response => {
            if (response.code == 200) {
              this.msgSuccess('添加成功')
              this.open = false
              this.getList()
            } else {
              this.msgError(response.msg)
            }
          }).catch(res => {
            console.log(res)
          })
        }
      })
    }

  }
}
</script>
