<template>
    <div class="settle">
        <el-table :data="list">
          <el-table-column label="投资者姓名" prop="customer_name" align="center"  show-overflow-tooltip/>
          <el-table-column label="总收益" prop="total_income" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.total_income) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="用户收益" prop="customer_income" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.customer_income) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="到用户余额" prop="to_customer_balance" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.to_customer_balance) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="到用户锁仓" prop="to_customer_lock" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.to_customer_lock) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="单T收益" prop="income" align="center"  show-overflow-tooltip/>
          <el-table-column label="状态" width="80" prop="is_transfer" align="center">
            <template slot-scope="scope">
                <span style="color: rgb(255, 80, 80);" v-show=" scope.row.is_transfer==0 ">未转账</span>
                <span style="color: #007AFF;" v-show=" scope.row.is_transfer==1 ">已转账</span>
            </template>
          </el-table-column>
          <el-table-column label="结算时间" width="160" prop="time" align="center">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.time) }}</span>
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
    </div>
</template>

<script>
import { getSettleList } from '@/api/customer/customer'
export default {
    name:'Settle',
    props: {},
    data() {
        return {
            list:[],
            total: 0,
            queryParams: {
              pageIndex: 1,
              pageSize: 10,
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
            getSettleList(this.queryParams).then(res=> {
              this.list=res.list
              this.total = res.total
            })
        },
    },
    components: {},
};
</script>

<style scoped>
.settle {
    padding: 20px;
}
</style>
