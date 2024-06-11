<template>
    <div class="user">
        <el-table :data="list">
          <el-table-column label="用户名" prop="name" align="center"  show-overflow-tooltip />
          <el-table-column label="手机号" prop="phone" align="center"  show-overflow-tooltip />
          <!-- <el-table-column label="公司收益" prop="company_income" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.company_income) }}</span>
            </template>
          </el-table-column> -->
          <el-table-column label="直接释放收益" prop="to_balance" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.to_balance) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="线性释放收益" prop="lock_release" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.lock_release) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="总收益" prop="amount" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(Number(scope.row.amount)) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['system:sysuser:lists']"
                size="mini"
                type="text"
                icon="el-icon-document"
                @click="investment(scope.row)"
              >收益列表</el-button>
              <el-button
                v-permisaction="['system:sysuser:transfer']"
                size="mini"
                type="text"
                icon="el-icon-document"
                @click="transfer(scope.row)"
              >转账</el-button>
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
        <!-- 投资列表 -->
        <el-dialog title="收益详情" :visible.sync="open" width="1000px">
          <el-table :data="lists">
            <el-table-column label="用户名" prop="name" align="center"  show-overflow-tooltip/>
            <el-table-column label="订单号" width="180" prop="investment_id" align="center" />
            <el-table-column label="手机号" prop="phone" align="center"  show-overflow-tooltip/>
            <el-table-column label="金额" prop="amount" align="center" show-overflow-tooltip>
              <template slot-scope="scope">
                <span>{{ moneyFormat(scope.row.amount) }}</span>
              </template>
            </el-table-column>
            <el-table-column label="奖励类型" width="180" prop="types_of" align="center" />
            <el-table-column label="创建时间" width="180" align="center" prop="create_time">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.create_time) }}</span>
            </template>
          </el-table-column>
          </el-table>
          <pagination
            v-show="totals>0"
            :total="totals"
            :page.sync="params.pageIndex"
            :limit.sync="params.pageSize"
            @pagination="investment"
          />
        </el-dialog>
    </div>
</template>

<script>
import { getWeekCustomerList,getWeekCustomerInvestmentList,transferWeekCustomer } from '@/api/customer/customer'
export default {
    name:'User',
    props: {},
    data() {
        return {
            list:[],
            total: 0,
            queryParams: {
              pageIndex: 1,
              pageSize: 10,
              date:'',
            },
            open:false,
            lists:[],
            totals: 0,
            params:{
              pageIndex: 1,
              pageSize: 10,
              date:'',
              customer_id:""
            }
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
            this.queryParams.date=this.$route.query.date
            this.params.date=this.$route.query.date
            getWeekCustomerList(this.queryParams).then(res=> {
              this.list=res.list
              this.total = res.total
            })
        },
        investment(row) {
            this.open=true
            if (row.customer_id==undefined) {
              this.params.customer_id=localStorage.getItem("customer_id")
            }else {
              localStorage.setItem("customer_id",row.customer_id)
              this.params.customer_id=row.customer_id
            }
            getWeekCustomerInvestmentList(this.params).then(res=> {
              this.lists=res.list
              this.totals = res.total
            })
        },
        transfer(row) {
          var param={
            customer_id:row.customer_id,
            date:this.params.date
          }
          this.$confirm('确认要转 "' + row.name  + '" 用户的帐吗?', '警告', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(function() {
            return transferWeekCustomer(param)
          }).then(() => {
            this.msgSuccess('转账成功');
            this.auto()
          })
        }
    },
    components: {},
};
</script>

<style scoped>
.user {
    padding: 20px;
}
</style>
