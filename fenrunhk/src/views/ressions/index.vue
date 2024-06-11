<template>
    <div class="ressions">
        <el-table :data="list">
          <el-table-column label="订单ID" prop="investment_id" align="center"  show-overflow-tooltip/>
          <el-table-column label="客户姓名" prop="customer_name" align="center"  show-overflow-tooltip/>
          <el-table-column label="线性释放" prop="customer_lock_release" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.customer_lock_release) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="状态" prop="name" align="center"  show-overflow-tooltip/>
          <el-table-column label="结算" prop="is_transfer" align="center"  show-overflow-tooltip>
             <template slot-scope="scope">
              <span v-if="scope.row.is_transfer=='已结算'" class="a">{{ scope.row.is_transfer }}</span>
              <span v-if="scope.row.is_transfer=='未结算'" class="b">{{ scope.row.is_transfer }}</span>
            </template>
          </el-table-column>
          <el-table-column label="释放时间" width="160" prop="time" align="center">
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
import { impressionslist } from '@/api/customer/customer'
export default {
    name:'Ressions',
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
            impressionslist(this.queryParams).then(res=> {
              this.list=res.list
              this.total = res.total
            })
        },
    },
    components: {},
};
</script>

<style scoped>
.ressions {
    padding: 20px;
}
.a {
   color: rgb(0, 255, 98);
}
.b {
  color: rgb(255, 80, 80);
}
</style>