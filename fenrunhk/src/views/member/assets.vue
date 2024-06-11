<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="用户账号" prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="用户账号/充币地址"
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

    <el-table v-loading="loading" :data="assetsList" @selection-change="handleSelectionChange" @sort-change='handleOrderBy' >
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="用户账号" align="center" prop="user_id" />
      <el-table-column label="充币地址" align="center" prop="addr" :show-overflow-tooltip="true" />
      <el-table-column label="资产名称" align="center" prop="symbol" />
      <el-table-column label="数量" align="center" prop="quantity" :sortable="'custom'"/>
      <el-table-column label="冻结数量" align="center" prop="frozen" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:sysconfig:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleDetails(scope.row)"
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
    <el-dialog title="" :visible.sync="open" width="1140px"  center>
      <el-row>
       <el-col :span="6">
         <div class="grid-content bg-purple" style="height:140px;">
         <label>ID：</label>
         <span>{{id}}</span> <br><br>
         <label>资产：</label>
         <span >{{DetailsInfo.GOF}} GOF</span> <br><br>
         <label>资产：</label>
         <span >{{DetailsInfo.GVC}} GVC</span> <br><br>
         <label>资产：</label>
         <span >{{DetailsInfo.USDT}} USDT</span> <br><br>
         </div>
         </el-col>

        <el-col :span="18">
          <div class="grid-content bg-purple-light" style="height:50px;">
          <label>授权地址：</label>
         <span>{{addr}}</span> <br><br>
         <label>冻结资产：</label>
         <span >{{DetailsInfo.GOFfrozen}} GOF</span> <br><br>
         <label>冻结资产：</label>
         <span >{{DetailsInfo.GVCfrozen}} GVC</span> <br><br>
         <label>冻结资产：</label>
         <span >{{DetailsInfo.USDTfrozen}} USDT</span> <br><br>

          </div></el-col>
      </el-row>
      <!-- <div style="float：left;height:50px;width:50px;">
        
      </div>
      <div style="float：right;height:50px;width:50px;">
        
      </div> -->
      
    <el-tabs  v-model="activeName"  >
      <el-tab-pane label="充币记录"  name="coincharge">
        <coincharge :id="id"/>
      </el-tab-pane>
      <el-tab-pane label="提币记录"  name="withdraw">
        <withdraw :id="id"/>
      </el-tab-pane>
      <el-tab-pane label="财务记录"  name="moneylog">
        <moneylog :id="id"/>
      </el-tab-pane>
    </el-tabs>
    </el-dialog>
  </div>
</template>

<script>
import { listAssets, delAssets, exportExcel,AssetsById } from '@/api/member/assets'
import { lsitCoinCharg } from '@/api/financial/coincharge'
import { lsitDraws } from '@/api/financial/withdraw'
import { lsitMoneyLog } from '@/api/financial/moneylog'
import coincharge from './details/coincharge'
import withdraw from './details/withdraw'
import moneylog from './details/moneylog'

export default {
  components: { coincharge,withdraw,moneylog },
  name: 'Assets',
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      id:undefined,
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 参数表格数据
      assetsList: [],
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
        cid: undefined,
        keyword: '',
        sort:'',
        order:''
      },
      queryDetails:{
        cid:undefined
      },
      DetailsInfo:{
        GOF:'0',
        GVC:'0',
        USDT:'0',
        GOFfrozen:'0',
        GVCfrozen:'0',
        USDTfrozen:'0'
      },
      addr:'',
      // 删除参数
      deleParams: {
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
      },
       activeName: 'coincharge'
    }
  },
  created() {
    this.getList()
    this.getDicts('asset_type').then(response => {
      this.typeOptions = response.data
    })
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listAssets(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.assetsList = response.data.list
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
    handleDetails(row){
      this.id = row.user_id
      this.queryDetails.cid = row.user_id
      this.open = true

      AssetsById(this.queryDetails).then(response => {
        var res = response.data.list
        this.addr = response.data.addr
        for (var i = 0; i < response.data.list.length; i++) {
          const name = res[i].symbol
          switch (name.toUpperCase()) {
            case 'GOF':
              this.DetailsInfo.GOF = Number(res[i].quantity)
              this.DetailsInfo.GOFfrozen = Number(res[i].frozen)
              break
            case 'GVC':
              this.DetailsInfo.GVC = Number(res[i].quantity)
              this.DetailsInfo.GVCfrozen = Number(res[i].frozen)
              break
            case 'USDT':
              this.DetailsInfo.USDT = Number(res[i].quantity)
              this.DetailsInfo.USDTfrozen = Number(res[i].frozen)
              break
          }
        }
      }
      )
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const configIds = row.id || this.ids
      this.deleParams.ids = configIds
      var ids = ''
      for (var i = 0; i < configIds.length; i++) {
        ids += configIds[i]
        if (i < configIds.length - 1) {
          ids += ','
        } else {
          this.deleParams.ids = ids
        }
      }
      this.$confirm('是否确认删除参数编号为"' + configIds + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delAssets(this.params)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {})
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
    },
    handleOrderBy(column ) {
      var fieldName = column.prop;
      var sortingType = column.order;
      this.queryParams.sort = fieldName

      if (sortingType =='descending'){
        this.queryParams.order = 'DESC'
      }else if (sortingType =='ascending'){
        this.queryParams.order = 'ASC'
      }else{
        this.queryParams.order = ''
        this.queryParams.sort = ''
      }
      listAssets(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.assetsList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    }
  }
}
</script>
