<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="玩法名称" prop="name">
        <el-input
          v-model="queryParams.keyword"
          placeholder="请输入参数名称"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>

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
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <!-- <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button> -->
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
      <!-- <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:export']"
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
        >导出</el-button>
      </el-col> -->
    </el-row>
    <!--添加修改删除操作结束-->

    <el-table v-loading="loading" :data="LotterySetsList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="编号" width="80" align="center" prop="id" />
      <el-table-column label="玩法名称" align="center" prop="name" />
      <el-table-column label="赔率" align="center" prop="odds" />
      <el-table-column label="平台手续费" align="center" prop="lotteryFree" />
      <el-table-column label="规则提示" align="center" prop="tulesTip" :show-overflow-tooltip="true" />
      <el-table-column label="设置时间" align="center" prop="createTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="修改时间" align="center" prop="updateTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.updateTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="180">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:sysmenu:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改</el-button>
          <el-button
            v-permisaction="['system:sysmenu:remove']"
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除</el-button>
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
    <el-dialog :title="title" :visible.sync="open" width="500px">
      <el-form ref="form" :model="editParams" :rules="rules" label-width="80px">
        <el-form-item label="编号" prop="ids">
          <el-input v-model="editParams.id" placeholder="编号自动生成" :disabled="true" />
        </el-form-item>
        <el-form-item label="玩法名称" prop="name">
          <el-input v-model="editParams.name" placeholder="请输入玩法名称" />
        </el-form-item>
        <el-form-item label="返奖倍数" prop="odds">
          <el-input v-model="editParams.odds" placeholder="请输入赔率" />
        </el-form-item>
        <el-form-item label="等级数字" prop="level">
          <el-input v-model="editParams.level" placeholder="请输入等级" @change="handleNumberChange" />
        </el-form-item>
        <el-form-item label="手续费" prop="lotteryfree">
          <el-input v-model="editParams.lotteryfree" placeholder="请输入手续费" />
        </el-form-item>
        <el-form-item label="规则提示" prop="tulestip">
          <el-input v-model="editParams.tulestip" type="textarea" rows="4" placeholder="请输入规则提示" />
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
import { listLotterySets, lotterySet, delLotterySet } from '@/api/lottery/lotteryset'

export default {
  name: 'Release',
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
      LotterySetsList: [],
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
      // 编辑参数
      editParams: {
        id: undefined,
        level: undefined,
        lotteryfree: undefined,
        name: undefined,
        odds: undefined,
        tulestip: undefined
      },
      // 删除参数
      delParams: {
        ids: ''
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
      listLotterySets(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.LotterySetsList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    // 参数系统内置字典翻译
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions, row.configType)
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.editParams = {
        id: undefined,
        level: undefined,
        lotteryfree: undefined,
        name: undefined,
        odds: undefined,
        tulestip: undefined
      }
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
    /** 修改按钮操作 */
    handleUpdate(row) {
      const configId = row.id || this.ids
      this.editParams.id = configId
      if (this.editParams.id === row.id) {
        this.editParams.level = String(row.id || '')
        this.editParams.lotteryfree = String(row.lotteryFree || '')
        this.editParams.name = String(row.name || '')
        this.editParams.odds = String(row.odds || '')
        this.editParams.tulestip = String(row.tulesTip || '')
      } else {
        this.editParams.id = String(this.editParams.id || '')
      }
      if (configId != null) {
        // this.form = response.data
        this.open = true
        this.title = '修改参数'
      }
    },
    /** 删除按钮操作 */
    handleDelete(row) {
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
      this.$confirm('是否确认删除参数编号为"' + configIds + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delLotterySet(this.delParams)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {})
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.editParams.id !== undefined) {
            lotterySet(this.editParams).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            }).catch(res => {
              console.log(res)
            })
          } else {
            lotterySet(this.editParams).then(response => {
              if (response.code === 200) {
                this.msgSuccess('新增成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },

    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length

      this.editParams.level = String(selection.map(item => item.level))
      this.editParams.lotteryfree = String(selection.map(item => item.lotteryFree))
      this.editParams.name = String(selection.map(item => item.name))
      this.editParams.odds = String(selection.map(item => item.odds))
      this.editParams.tulestip = String(selection.map(item => item.tulesTip))
    },
    // 文本框改变数据
    handleNumberChange(selection) {
      // alert("sdfj")
      // val=selection.replace(/[\u4e00-\u9fa5]/ig,'')
      // this.$nextTick(()=>{this.editParams.level=val});
    }
  }
}
</script>
