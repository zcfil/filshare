<template>
    <div class="audituser">
        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              v-permisaction="['system:sysaudituser:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增用户</el-button>
          </el-col>
        </el-row>
        <el-table :data="list">
          <el-table-column label="用户昵称" prop="nick_name" align="center"/>
          <el-table-column label="用户名" prop="username" align="center" />
          <el-table-column label="手机号" prop="phone" align="center" />
          <el-table-column label="推荐人" prop="referrer_name" align="center"/>
          <el-table-column label="创建时间" width="180" align="center" prop="create_time">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.create_time) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="130" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['system:sysaudituser:adopt']"
                size="mini"
                type="text"
                @click="adopt(scope.row)"
              >通过</el-button>
              <el-button
                v-permisaction="['system:sysaudituser:refuse']"
                size="mini"
                type="text"
                @click="refuse(scope.row)"
              >拒绝</el-button>
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
        <el-dialog :title="title" :visible.sync="open" width="700px">
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-row>
              <el-col :span="12">
                <el-form-item label="姓名" prop="nickName">
                  <el-input v-model="form.nickName" placeholder="请输入用户昵称" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="归属部门" prop="deptId">
                  <treeselect
                    v-model="form.deptId"
                    :options="deptOptions"
                    :normalizer="normalizer"
                    placeholder="请选择归属部门"
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="手机号码" prop="phone">
                  <el-input v-model="form.phone" placeholder="请输入手机号码" maxlength="11" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="邮箱" prop="email">
                  <el-input v-model="form.email" placeholder="请输入邮箱" maxlength="50" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="用户名" prop="username">
                  <el-input v-model="form.username" placeholder="请输入用户名称" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item v-if="form.userId == undefined" label="用户密码" prop="password">
                  <el-input v-model="form.password" placeholder="请输入用户密码" type="password" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="用户性别" prop="sex">
                  <el-select v-model="form.sex" style="width:100%"  placeholder="请选择">
                    <el-option
                      v-for="dict in sexOptions"
                      :key="dict.dictValue"
                      :label="dict.dictLabel"
                      :value="dict.dictValue"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="岗位" prop="postId">
                  <el-select v-model="form.postId" style="width:100%" placeholder="请选择" @change="$forceUpdate()">
                    <el-option
                      v-for="item in postOptions"
                      :key="item.postId"
                      :label="item.postName"
                      :value="item.postId"
                      :disabled="item.status == 1"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="角色" prop="roleId">
                  <el-select v-model="form.roleId" style="width:100%" placeholder="请选择" @change="$forceUpdate()">
                    <el-option
                      v-for="item in roleOptions"
                      :key="item.roleId"
                      :label="item.roleName"
                      :value="item.roleId"
                      :disabled="item.status == 1"
                    />
                  </el-select>
                </el-form-item>
              </el-col>
               <el-col :span="12">
                <el-form-item label="银行户名" prop="bank_user_name">
                  <el-input v-model="form.bank_user_name" placeholder="请输入银行户名"/>
                </el-form-item>
              </el-col>
               <el-col :span="12">
                <el-form-item label="银行卡号" prop="bank_num">
                  <el-input v-model="form.bank_num" placeholder="请输入银行卡号" maxlength="19" />
                </el-form-item>
              </el-col>
               <el-col :span="12">
                <el-form-item label="开户行" prop="open_bank">
                  <el-input v-model="form.open_bank" placeholder="请输入开户行" />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="状态">
                  <el-radio-group v-model="form.status">
                    <el-radio
                      v-for="dict in statusOptions"
                      :key="dict.dictValue"
                      :label="dict.dictValue"
                    >{{ dict.dictLabel }}</el-radio>
                  </el-radio-group>
                </el-form-item>
              </el-col>
              <el-col :span="24">
                <el-form-item label="备注">
                  <el-input v-model="form.remark" type="textarea" placeholder="请输入内容" />
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
    </div>
</template>

<script>
import { GetPassingSysUserList,SubmitNewUser,AllowNewUserPass } from '@/api/system/audituser'
import {  getUserInit } from '@/api/system/sysuser'
import { treeselect } from '@/api/system/dept'
import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
export default {
    name:'Audituser',
    props: {},
    data() {
        return {
          total:0,
            list:[],
            queryParams: {
              pageIndex: 1,
              pageSize: 10,
            },
            title:'',
            open:false,
            form:{},
            // 部门树选项
            deptOptions: undefined,
            // 状态数据字典
            statusOptions: [],
            // 性别状态字典
            sexOptions: [],
            // 岗位选项
            postOptions: [],
            // 角色选项
            roleOptions: [],
            // 表单校验
            rules: {
              username: [
                { required: true, message: '用户名称不能为空', trigger: 'blur' }
              ],
              nickName: [
                { required: true, message: '用户昵称不能为空', trigger: 'blur' }
              ],
              deptId: [
                { required: true, message: '归属部门不能为空', trigger: 'blur' }
              ],
              sex: [
                { required: true, message: '用户性别不能为空', trigger: 'blur' }
              ],
              postId: [
                { required: true, message: '岗位不能为空', trigger: 'blur' }
              ],
              roleId: [
                { required: true, message: '角色不能为空', trigger: 'blur' }
              ],
              password: [
                { required: true, message: '用户密码不能为空', trigger: 'blur' }
              ],
              bank_user_name: [
                { required: true, message: '银行户名不能为空', trigger: 'blur' }
              ],
              bank_num: [
                { required: true, message: '银行卡号不能为空', trigger: 'blur' },
                {
                  pattern: /^([1-9]{1})(\d{15}|\d{18})$/,
                  message: '请输入正确的银行卡号',
                  trigger: 'blur'
                }
              ],
              open_bank: [
                { required: true, message: '开户行不能为空', trigger: 'blur' }
              ],
              email: [
                // { required: true, message: '邮箱地址不能为空', trigger: 'blur' },
                {
                  type: 'email',
                  message: "'请输入正确的邮箱地址",
                  trigger: ['blur', 'change']
                }
              ],
              phone: [
                { required: true, message: '手机号码不能为空', trigger: 'blur' },
                {
                  pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/,
                  message: '请输入正确的手机号码',
                  trigger: 'blur'
                }
              ]
            }
        };
    },
    computed: {},
    created() {
        this.auto()
        this.getDicts('sys_normal_disable').then(response => {
          this.statusOptions = response.data
        })
        this.getDicts('sys_user_sex').then(response => {
          this.sexOptions = response.data
        })
    },
    mounted() {},
    watch: {},
    methods: {
        auto() {
            GetPassingSysUserList(this.queryParams).then(res=> {
                this.list=res.data.list
              this.total=Number(res.data.totalCount)
            })
        },
        // 表单重置
        reset() {
          this.form = {
            userId: undefined,
            deptId: undefined,
            username: undefined,
            nickName: undefined,
            password: undefined,
            phone: undefined,
            email: undefined,
            sex: undefined,
            status: '0',
            remark: undefined,
            postIds: undefined,
            roleIds: undefined,
            bank_user_name:undefined,
            bank_num:undefined,
            open_bank:undefined
          }
          this.resetForm('form')
        },
        // 取消按钮
        cancel() {
          this.open = false
          this.reset()
        },
        /** 查询部门下拉树结构 */
        getTreeselect() {
          treeselect().then(response => {
            this.deptOptions = response.data
          })
        },
        /** 转换菜单数据结构 */
        normalizer(node) {
          if (node.children && !node.children.length) {
            delete node.children
          }
          return {
            id: node.deptId,
            label: node.deptname,
            children: node.children
          }
        },
        handleAdd() {
          this.getTreeselect()
          getUserInit().then(response => {
            this.postOptions = response.data.posts
            this.roleOptions = response.data.roles
            this.open = true
            this.title = '新增用户'
          })
        },
        submitForm() {
            this.$refs['form'].validate(valid => {
              if (valid) {
                SubmitNewUser(this.form).then(response => {
                  if (response.code === 200) {
                    this.msgSuccess('新增成功')
                    this.open = false
                    this.auto()
                  } else {
                    this.msgError(response.msg)
                  }
                })
              }
            })
        },
        adopt(row) {
            var params={
                user_id:row.user_id,
                is_allow:1
            }
            AllowNewUserPass(params).then(res=> {
                this.msgSuccess('操作成功')
                this.auto()
            })
        },
        refuse(row) {
            var params={
                user_id:row.user_id,
                is_allow:0
            }
            AllowNewUserPass(params).then(res=> {
                this.msgSuccess('操作成功')
                this.auto()
            })
        }
    },
    components: { Treeselect },
};
</script>

<style scoped>
.audituser {
    padding: 20px;
}
</style>
