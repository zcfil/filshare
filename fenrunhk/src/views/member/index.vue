<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="" prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="用户账号/邀请人/授权地址/邮箱"
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
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <!--<el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>-->
      </el-form-item>

      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="状态" clearable size="small">
          <el-option
            v-for="dict in typeOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictValue"
          />
        </el-select>
      </el-form-item>

    </el-form>
    <!--搜索结束-->

    <!--添加修改删除操作开始-->
    <el-row :gutter="10" class="mb8">
      <!-- <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增</el-button>
      </el-col> -->
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
      <!-- <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除</el-button>
      </el-col> -->
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:export']"
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
        >导出</el-button>
      </el-col>
    </el-row>
    <!--添加修改删除操作结束-->

    <el-table v-loading="loading" :data="memberList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="用户账号" align="center" prop="id" />
      <el-table-column label="邮箱" align="center" prop="email" :show-overflow-tooltip="true" />
      <el-table-column label="授权地址" align="center" prop="auth_addr" :show-overflow-tooltip="true" />
      <el-table-column label="授权时间" align="center" prop="auth_time" width="180">
        <template #default='props'>
          <div class='cell'>{{ new Date(props.row.auth_time).format('yyyy-MM-dd hh:mm:ss') }}</div>
        </template>
      </el-table-column>
      <el-table-column label="邀请码" align="center" prop="invite_code" width="180" />
      <el-table-column label="邀请人" align="center" prop="invite_people" width="180" />
      <el-table-column label="会员等级" align="center" prop="level" :formatter="typeLevelFormat" />
      <el-table-column label="状态" align="center" prop="status" :formatter="typeFormat" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:sysconfig:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >{{content}}</el-button>
          <el-button
            v-permisaction="['system:sysconfig:edit']"
            size="mini"
            type="text"
            icon="el-icon-tickets"
            @click="handleUpdate1(scope.row)"
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

    <!-- 添加或修改参数配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="500px">
      <el-form ref="form" :model="editParams" :rules="rules" label-width="80px">
        <el-form-item label="编号" prop="ids">
          <el-input v-model="editParams.id" placeholder="请输入参数名称" :disabled="true" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="editParams.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="修改密码" prop="addr">
          <el-input v-model="editParams.pwd" placeholder="请输入用户新密码（为空不修改）" />
        </el-form-item>
        <el-form-item label="授权地址" prop="addr">
          <el-input v-model="editParams.auth_addr" placeholder="请输入授权地址" />
        </el-form-item>
        <el-form-item label="会员等级" align="left" prop="level" :formatter="typeLevelFormat">
          <el-select v-model="editParams.level" placeholder="状态" size="medium">
            <el-option
              v-for="dict in typeLevel"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictValue"
            />
          </el-select>
        </el-form-item>
        <!-- <el-form-item label="充币地址" prop="money_address">
          <el-input v-model="editParams.money_address" placeholder="请输入充币地址" />
        </el-form-item> -->
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="editParams.status">
            <el-radio
              v-for="dict in typeOptions"
              :key="dict.dictValue"
              :label="dict.dictValue"
            >{{ dict.dictLabel }}</el-radio>
          </el-radio-group>
        </el-form-item>
        <!-- <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容" />
        </el-form-item> -->
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm" :disabled="epditUser">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
     <!-- 详情对话框-->
    <el-dialog :title="title1" :visible.sync="open1" width="1240px" center>
      <el-table :data="DetailsList">
        <el-table-column property="id" label="用户账号" width="150" />
        <el-table-column property="invite_code" label="邀请码" width="150" />
        <el-table-column property="level" label="会员等级" width="180" :formatter="typeLevelFormat" />
        <el-table-column property="phone" label="手机号" width="180" />
        <el-table-column property="email" label="邮箱" width="180" />
        <!-- <el-table-column property="total" label="投注总额">
          <template #default="props">
            <div class="cell">{{Number(props.row.total).toLocaleString()}}</div>
          </template>
        </el-table-column> -->
        <el-table-column property="created_time" label="注册时间" >
          <template #default="props">
          <div class="cell">{{new Date(props.row.created_time).format('yyyy-MM-dd hh:mm:ss')}}</div>
        </template>
        </el-table-column>

      </el-table>
      <pagination
        v-show="total>0"
        :total="detailsTotal"
        :page.sync="detailsParams.pageIndex"
        :limit.sync="detailsParams.pageSize"
        @pagination="handleUpdate1"
      />
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { listMember, delMenu, updateMenu, exportUser,userInvitePeople } from '@/api/member/member'

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
      epditUser: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 参数表格数据
      memberList: [],
      DetailsList:[],
      // 弹出层标题
      title: '',
      title1: '',
      content: '',
      // 是否显示弹出层
      open: false,
      open1: false,
      // 类型数据字典
      typeOptions: [],
      // 会员等级数据字典
      typeLevel: [],
      // 日期范围
      dateRange: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        status: '',
        keyword: ''
        // cid: undefined
      },
      detailsParams: {
        pageIndex: 1,
        pageSize: 10,
        cid: undefined
      },
      // 删除参数
      params: {
        ids: undefined
      },
      // 编辑参数
      editParams: {
        email: '',
        auth_addr: '',
        status: '0',
        level: undefined,
        pwd: undefined
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
      activeName: 'exceptionCoin'
    }
  },
  computed: {
    ...mapGetters([
      'roles'
    ])
  },
  created() {
    this.getList()
    this.getDicts('user_status').then(response => {
      this.typeOptions = response.data
    })
    this.getDicts('user_level').then(response => {
      this.typeLevel = response.data
    })
    if (this.roles.includes('admin')) {
      this.content  = '修改'
      this.title = '修改参数'
      this.epditUser = false
    }else{
      this.content  = '查看'
      this.title = '查看参数'
      this.epditUser = true
    }
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listMember(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.memberList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    // 参数系统内置字典翻译
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions, row.status)
    },
    typeLevelFormat(row, column) {
      return this.selectDictLabel(this.typeLevel, row.level)
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
      this.title = '添加参数'
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length

      this.editParams.id = String(selection.map(item => item.id))
      this.editParams.email = String(selection.map(item => item.email))
      this.editParams.auth_addr = String(selection.map(item => item.auth_addr))
      this.editParams.status = String(selection.map(item => item.status))
      this.editParams.level = String(selection.map(item => item.level))
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      const configId = row.id || this.ids
      this.editParams.id = configId
      if (this.editParams.id == row.id) {
        this.editParams.id = String(row.id)
        this.editParams.email = String(row.email)
        this.editParams.auth_addr = String(row.auth_addr)
        this.editParams.status = String(row.status)
        this.editParams.level = String(row.level)
      } else {
        this.editParams.id = String(this.editParams.id)
      }

      if (configId != null) {
        // this.reset()
        // this.form = response.data
        this.open = true
        //this.title = '修改参数'
      }
    },
    handleUpdate1(row){
      this.title1 = "用户"+row.id+"邀请列表"
      this.open1 = true
      this.detailsParams.id = row.id
      userInvitePeople(this.detailsParams).then(response => {
        if (response.code == 200) {
          this.DetailsList = response.data.list
          this.detailsTotal = response.data.count
        } else {
          this.msgError(response.msg)
        }
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      // this.$refs['form'].validate(valid => {
      // if (valid) {
      //   if (this.form.configId !== undefined) {
      updateMenu(this.editParams).then(response => {
        if (response.code == 200) {
          this.msgSuccess('修改成功')
          this.editParams.pwd = undefined
          this.open = false
          this.getList()
        } else {
          this.msgError(response.msg)
        }
      }).catch(res => {
        console.log(res)
      })
      // } else {
      //   addConfig(this.form).then(response => {
      //     if (response.code === 200) {
      //       this.msgSuccess('新增成功')
      //       this.open = false
      //       this.getList()
      //     } else {
      //       this.msgError(response.msg)
      //     }
      //   })
      // }
      // }
      // })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const configIds = row.id || this.ids
      this.params.ids = configIds
      var ids = ''
      for (var i = 0; i < configIds.length; i++) {
        ids += configIds[i]
        if (i < configIds.length - 1) {
          ids += ','
        } else {
          this.params.ids = ids
        }
      }
      this.$confirm('是否确认删除参数编号为"' + configIds + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delMenu(this.params)
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
        exportUser(queryParams).then(res => {
          location.href = process.env.VUE_APP_BASE_API + `/` + res.msg
        })
      })
    }
  }
}
</script>
