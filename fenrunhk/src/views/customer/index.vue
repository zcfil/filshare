<template>
  <div class="customer">
    <el-form ref="queryForm" :model="queryParams" :inline="true">
      <el-form-item prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="客户姓名/手机号"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:syscustomer:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增</el-button>
      </el-col>
    </el-row>
    <el-table :data="customerlist">
      <el-table-column label="姓名" width="100" prop="name" align="center" show-overflow-tooltip />
      <el-table-column label="手机号" width="130" prop="phone" align="center" />
      <el-table-column label="身份证号" width="190" prop="identity" align="center"/>
      <el-table-column label="钱包地址" prop="wallet" align="center" show-overflow-tooltip />
      <el-table-column label="锁仓余额" width="150" prop="locked_balance" align="center">
        <template slot-scope="scope">
          <span>{{ moneyFormat(scope.row.locked_balance) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="钱包余额" width="150" prop="wallet_balance" align="center">
        <template slot-scope="scope">
          <span>{{ moneyFormat(scope.row.wallet_balance) }}</span>
        </template>
      </el-table-column>
      <!-- <el-table-column label="锁仓余额" width="150" prop="locked_balance" align="center" />
      <el-table-column label="钱包余额" width="150" prop="wallet_balance" align="center" /> -->
      <el-table-column label="创建时间" width="170" align="center" prop="create_time">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="客户文件" width="110" align="center" prop="create_time">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:syscustomer:upload']"
            type="text"
            icon="el-icon-tickets"
            size="mini"
            @click="upload(scope.row)"
          >文件管理</el-button>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="150" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:syscustomer:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改</el-button>
          <el-button
            v-permisaction="['system:syscustomer:remove']"
            type="text"
            icon="el-icon-delete"
            size="mini"
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
      @pagination="auto"
    />

    <!-- 添加或修改客户 -->
    <el-dialog :title="title" :visible.sync="open" width="500px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item v-show="show" label="姓名" prop="name">
          <el-input v-model="form.name" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input v-model="form.phone" placeholder="请输入手机号" maxlength="11" />
        </el-form-item>
        <el-form-item v-show="show" label="密码" prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入密码" />
        </el-form-item>
        <el-form-item v-show="show" label="身份证" prop="identity">
          <el-input v-model="form.identity" placeholder="请输入身份证" />
        </el-form-item>
        <el-form-item label="钱包地址" prop="wallet">
          <el-input v-model="form.wallet" type="textarea" placeholder="请输入钱包地址" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>

    <el-dialog title="文件管理" :visible.sync="show1" width="1100px">
      <el-upload
        action="/api/v1/uploadfile"
        :http-request="imageChange"
        :show-file-list="false"
        multiple
        :limit="20"
      >
        <el-button size="small" type="primary">点击上传</el-button>
      </el-upload>
      <el-table :data="list">
        <el-table-column label="文件名称" width="160" prop="filename" align="center" show-overflow-tooltip />
        <el-table-column label="文件路径" prop="url" align="center" show-overflow-tooltip>
          <template slot-scope="scope">
            <a :href="scope.row.url" style="color:rgb(0, 110, 255);">{{ scope.row.url }}</a>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="160" prop="datetime" align="center" show-overflow-tooltip />
        <el-table-column label="操作" align="center" width="100" class-name="small-padding fixed-width">
          <template slot-scope="scope">
            <el-button
              v-permisaction="['system:syscustomer:delete']"
              type="text"
              icon="el-icon-delete"
              size="mini"
              @click="uploadDelete(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

  </div>
</template>

<script>
import { customerList, customerAdd, customerEdit, customerDelete, delfile, getuploadfileList } from '@/api/customer/customer'
import { getUserProfile } from '@/api/system/sysuser'
import axios from 'axios'
import { getToken } from '@/utils/auth'
export default {
  name: 'Customer',
  components: {},
  props: {},
  data() {
    return {
      open: false,
      title: '',
      form: {},
      customerlist: [],
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        keyword: undefined
      },
      total: 0,
      // 表单校验
      rules: {
        name: [
          { required: true, message: '姓名不能为空', trigger: 'blur' },
          {
            pattern: /^(?!.*?[\u3000-\u303F\u4DC0-\u4DFF\u2800-\u28FF\u3200-\u32FF\u3300-\u33FF\u2700-\u27BF\u2600-\u26FF\uFE10-\uFE1F\uFE30-\uFE4F])[\u4e00-\u9fbb\u2E80-\uFE4Fa-zA-Z.`·]+$/,
            message: '请输入正确的姓名',
            trigger: 'blur'
          }
        ],
        phone: [
          { required: true, message: '手机号码不能为空', trigger: 'blur' },
          {
            pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
            message: '请输入正确的手机号码',
            trigger: 'blur'
          }
        ],
        identity: [
          { required: true, message: '身份证不能为空', trigger: 'blur' },
          {
            pattern: /^[0-9a-zA-Z]{6,}$/,
            message: '请输入正确的身份证号',
            trigger: 'blur'
          }
        ],
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          {
            pattern: /^(\w){6,18}$/,
            message: '请输入正确的密码',
            trigger: 'blur'
          }
        ],
        wallet: [
          { required: true, message: '钱包地址不能为空', trigger: 'blur' }
        ]
      },
      show: false,
      show1: false,
      list: [],
      name: '',
      username: ''
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
      customerList(this.queryParams).then(res => {
        this.customerlist = res.list
        this.total = res.total
      })
    },
    reset() {
      this.form = {
        name: undefined,
        phone: undefined,
        password: undefined,
        wallet: undefined,
        identity: undefined
      }
      this.resetForm('form')
    },
    handleAdd() {
      this.reset()
      this.open = true
      this.show = true
      this.title = '添加客户'
    },
    handleUpdate(row) {
      this.auto()
      this.form = row
      this.open = true
      this.show = false
      this.title = '修改客户'
    },
    handleDelete(row) {
      var that = this
      var params = {
        id: row.id
      }
      this.$confirm('是否确认删除姓名为 ' + row.name + ' 的客户！', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        customerDelete(params).then(res => {
          that.msgSuccess('删除成功')
          that.auto()
        })
      }).catch(() => {
        // console.log("取消")
      })
    },
    submitForm() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            customerEdit(this.form).then(res => {
              this.msgSuccess('修改成功')
              this.open = false
              this.auto()
            })
          } else {
            customerAdd(this.form).then(res => {
              this.msgSuccess('新增成功')
              this.open = false
              this.auto()
            })
          }
        }
      })
    },
    cancel() {
      this.open = false
      this.reset()
    },
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.auto()
    },
    resetQuery() {
      this.resetForm('queryForm')
      this.handleQuery()
    },
    upload(row) {
      this.show1 = true
      if (row == undefined) {
        this.name = localStorage.getItem('name')
      } else {
        localStorage.setItem('name', row.name)
        this.name = row.name
      }
      getuploadfileList({ username: this.name }).then(res => {
        this.list = res.data
      })
      getUserProfile().then(res => {
        this.username = res.data.username
      })
    },
    // 提交图片
    imageChange(param, type) {
      // var params={
      //   file:param.file,
      //   userid:this.name,
      //   adminid:this.username
      // }
      // console.log(params);
      const formData = new FormData()
      formData.append('file', param.file)
      formData.append('userid', this.name)
      formData.append('adminid', this.username)
      // uploadfile(formData).then(res => {
      //     if (res.IsSuccess) {
      //         // this.imgList.push(res.Data.Data)
      //         console.log(res);
      //     }
      // });
      axios.post('http://47.112.99.108:8005/api/v1/uploadfile', formData, { headers: {
        Authorization: 'Bearer ' + getToken()
      }}).then(res => {
        this.msgSuccess('添加成功')
        this.upload()
      })
    },
    uploadDelete(row) {
      var that = this
      var params = {
        filename: row.filename
      }
      this.$confirm('是否确认删除文件名为 ' + row.filename + ' 的文件！', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        delfile(params).then(res => {
          that.msgSuccess('删除成功')
          that.upload()
        })
      }).catch(() => {
        // console.log("取消")
      })
    }
  }
}
</script>

<style scoped>
.customer {
    padding: 20px;
}
</style>
