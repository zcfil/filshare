<template>
    <div class="bequeath">
      <el-form ref="queryForm" :model="queryParams" :inline="true">
        <el-form-item prop="keyword">
          <el-input
            v-model="queryParams.keyword"
            placeholder="客户姓名/订单ID"
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
              v-permisaction="['system:sysbequeath:add']"
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
            >新增</el-button>
          </el-col>
      </el-row>
      <el-table :data="bequeathlist">
          <el-table-column label="订单ID" width="190" prop="order_mid" align="center" />
          <el-table-column label="姓名" prop="user_name" align="center"  show-overflow-tooltip />
          <el-table-column label="锁仓币" prop="balance" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.balance) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="订单周期" prop="totalDay" align="center"  show-overflow-tooltip/>
          <el-table-column label="备注" prop="remark" align="center" show-overflow-tooltip  />
          <el-table-column label="开始时间" width="170" align="center" prop="start_time">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.start_time) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="结束时间" width="170" align="center" prop="end_time">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.end_time) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="状态" width="160" prop="status" align="center">
            <template slot-scope="scope">
                <span class="a" v-show=" scope.row.status==0 ">进行中</span>
                <span class="b" v-show=" scope.row.status==1 ">已终止</span>
                <span class="a" v-show=" scope.row.status==2 ">线性释放中</span>
                <span class="c" v-show=" scope.row.status==3 ">已完成线性释放</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="200" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['system:sysbequeath:edit']"
                :disabled="scope.row.disabled==0"
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >修改</el-button>
              <el-button
                v-permisaction="['system:sysbequeath:stop']"
                size="mini"
                type="text"
                icon="el-icon-circle-close"
                @click="handleStop(scope.row)"
              >终止</el-button>
              <el-button
                v-permisaction="['system:sysbequeath:remove']"
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
            <el-form-item v-if="show1" label="姓名" prop="user_name">
              <el-input v-model="form.user_name" placeholder="请输入客户姓名查找用户绑定订单" @input="chang"/>
            </el-form-item>
            <el-form-item label="锁仓币" prop="balance">
              <el-input v-model="form.balance" placeholder="请输入锁仓币"/>
            </el-form-item>
            <el-form-item label="订单周期" prop="totalDay">
              <el-input v-model="form.totalDay" placeholder="请输入订单周期（默认为180天）"/>
            </el-form-item>
            <el-form-item v-if="!show1" label="开始时间" prop="time">
              <el-date-picker
               :picker-options="pickerOptions"
                v-model="form.time"
                type="datetime"
                placeholder="选择日期时间">
              </el-date-picker>
            </el-form-item>
            <el-form-item label="备注" prop="remark">
              <el-input v-model="form.remark" type="textarea" placeholder="请输入备注" />
            </el-form-item>
          </el-form>
          <div class="small" v-show="show">
            <div v-for="(item,index) in data" :key="index" @click="choice(item)">{{item.name}}  ({{item.phone}})</div>
          </div>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>

    </div>
</template>

<script>
import { addmigrate,getmigrate,deletemigrate,editmigrate,breakmigreate,userList } from '@/api/customer/customer'
export default {
    name:'Bequeath',
    props: {},
    data() {
        return {
            pickerOptions:{                                 //禁用当前日期之前的日期
                disabledDate(time) {
                    return time.getTime() < Date.now() - 8.64e7;
                },
            },
            open:false,
            show:false,
            show1:false,
            data:[],
            title:'',
            form:{},
            bequeathlist:[],
            queryParams: {
              pageIndex: 1,
              pageSize: 10,
              keyword: undefined,
            },
            total: 0,
            // 表单校验
            rules: {
              user_name: [
                { required: true, message: '姓名不能为空', trigger: 'blur' },
                {
                  pattern: /^(?!.*?[\u3000-\u303F\u4DC0-\u4DFF\u2800-\u28FF\u3200-\u32FF\u3300-\u33FF\u2700-\u27BF\u2600-\u26FF\uFE10-\uFE1F\uFE30-\uFE4F])[\u4e00-\u9fbb\u2E80-\uFE4Fa-zA-Z.`·]+$/,
                  message: '请输入正确的姓名',
                  trigger: 'blur'
                }
              ],
              totalDay: [
                { required: true, message: '订单周期不能为空', trigger: 'blur' },
                {
                  pattern: /^[0-9]*[1-9][0-9]*$/,
                  message: '请输入正整数',
                  trigger: 'blur'
                }
              ],
              balance: [
                { required: true, message: '锁仓币不能为空', trigger: 'blur' },
                {
                  pattern: /^\d+$|^\d*\.\d+$/g,
                  message: '请输入数字和一个小数点',
                  trigger: 'blur'
                }
              ],
              time: [
                { required: true, message: '日期不能为空', trigger: 'blur' }
              ],
            },
        };
    },
    computed: {},
    created() {
        this.auto()
    },
    mounted() {},
    watch: {},
    methods: {
        auto() {
          getmigrate(this.queryParams).then(res=> {
            this.bequeathlist=res.data.list
            for (var i=0;i<this.bequeathlist.length;i++) {
              if (Date.parse(new Date())<=Date.parse(this.parseTime(this.bequeathlist[i].start_time.substring(0,10)+" 23:30:00"))) {
                this.bequeathlist[i].disabled="1"
              }else {
                this.bequeathlist[i].disabled="0"
              }
            }
            this.total = res.data.total
          })
        },
        chang() {
            if(this.form.user_name.length!='') {
                userList({keyword:this.form.user_name}).then(res=> {
                    this.data=res.list
                    if (this.data.length>0) {
                      this.show=true
                    }else {
                      this.show=false
                    }
                })
            }else{
              this.data=[]
              this.show=false
            }
        },
        choice(item) {
          this.form.user_name=item.name
          this.form.customer_id=item.id
          this.form.phone=item.phone
          localStorage.setItem("user_name",item.name)
          this.show=false
        },
        reset() {
            this.form={
              user_name:undefined,
              customer_id:undefined,
              phone:undefined,
              balance:undefined,
              totalDay: "180",
              remark:undefined
            }
            this.resetForm('form')
        },
        handleAdd() {
            this.reset()
            this.open=true
            this.show1=true
            this.title = '添加订单'
        },
        handleUpdate(row) {
          if (row.status=="1") {
            this.msgError('该订单以终止！');
            return
          }
            this.auto()
            this.form=row
            localStorage.setItem("user_name",row.user_name)
            this.open = true
            this.show1=false
            this.title = '修改订单'
        },
        handleDelete(row) {
          if (row.status=="0") {
            this.msgError('该订单正在进行中，终止后才能删除！');
            return
          }
            var that=this
            var params={
              order_mid:row.order_mid
            }
            this.$confirm('是否确认删除订单号为 '+row.order_mid+ ' 的订单', '警告', {
               confirmButtonText: '确定',
               cancelButtonText: '取消',
               type: 'warning'
            }).then(function() {
              deletemigrate(params).then(res => {
                that.msgSuccess('删除成功')
                that.auto()
              })
            }).catch(() => {
              //console.log("取消")
            })
        },
        handleStop(row) {
          if (row.status=="1") {
            this.msgError('该订单以终止！');
            return
          }
            var that=this
            var params={
              order_mid:row.order_mid
            }
            this.$confirm('是否确认终止订单号为 '+row.order_mid+ ' 的订单', '警告', {
               confirmButtonText: '确定',
               cancelButtonText: '取消',
               type: 'warning'
            }).then(function() {
              breakmigreate(params).then(res => {
                that.msgSuccess('终止成功')
                that.auto()
              })
            }).catch(() => {
              //console.log("取消")
            })
        },
        submitForm() {
          this.$refs['form'].validate(valid => {
            if (valid) {
              if (this.form.totalDay<180) {
                this.msgError('订单周期不能小于180天！');
                return
              }
              if(this.form.order_mid!== undefined) {
                this.form.time=this.parseTime(this.form.time)
                editmigrate(this.form).then(res=> {
                  this.msgSuccess('修改成功');
                  this.open = false
                  this.auto()
                })
              }else {
                if (this.form.customer_id==undefined||localStorage.getItem("user_name")!=this.form.user_name) {
                  this.msgError('没有选中绑定客户或者该用户不存在');
                  return
                }
                addmigrate(this.form).then(res=> {
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
    },
    components: {},
};
</script>

<style scoped>
.bequeath {
    padding: 20px;
}
.small {
  position: absolute;
  top: 130px;
  left: 100px;
  width: 380px;
  height: 280px;
  border: 1px solid rgb(0, 119, 255);
  background-color: #fff;
  border-radius: 4px;
  overflow: hidden;
  overflow-y: auto;
}
.small div {
  padding: 8px 20px;
}
.small div:hover {
  background-color: #eeeeee;
}
.b {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(255, 180, 180);
  color: rgb(255, 80, 80);
  background-color: rgb(255, 221, 221);
}
.a {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(149, 195, 255);
  color: rgb(0, 110, 255);
  background-color: rgb(231, 241, 255);
}
.c {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(170, 255, 170);
  color: rgb(0, 255, 98);
  background-color: rgb(235, 255, 235);
}
</style>