<template>
    <div class="transfer">
        <div style="margin-bottom:20px">
          日期：
          <el-date-picker
            v-model="value1"
            type="daterange"
            :clearable="false"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            @change="select">
          </el-date-picker>
        </div>
        <el-table :data="list">
          <el-table-column label="客户姓名" prop="name" align="center"  show-overflow-tooltip/>
          <el-table-column label="消息ID" prop="cid" align="center"  show-overflow-tooltip/>
          <el-table-column label="转出地址" prop="from" align="center"  show-overflow-tooltip/>
          <el-table-column label="收账地址" prop="to" align="center"  show-overflow-tooltip/>
          <el-table-column label="金额（FIL）" prop="amount" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.amount) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="状态" width="80" prop="status" align="center">
            <template slot-scope="scope">
                <span style="color: rgb(255, 80, 80);" v-show=" scope.row.status==0 ">确认中</span>
                <span style="color: #007AFF;" v-show=" scope.row.status==1 ">已到账</span>
            </template>
          </el-table-column>
          <el-table-column label="手续费" prop="service_charge" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.service_charge) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="转账时间" width="160" prop="create_time" align="center">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.create_time) }}</span>
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
import { getTransferList } from '@/api/customer/customer'
export default {
    name:'Transfer',
    props: {},
    data() {
        return {
            list:[],
            total: 0,
            queryParams: {
              pageIndex: 1,
              pageSize: 10,
              start:'',
              end:''
            },
            value1:'',
        };
    },
    computed: {},
    created() {
        const end = new Date();
        const start = new Date();
        start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
        this.queryParams.start=this.parseTime(start).slice(0,10)
        this.queryParams.end=this.parseTime(end).slice(0,10)
        this.value1=[start,end]
        this.auto()
    },
    mounted() {},
    watch: {},
    methods: {
        auto() {
            getTransferList(this.queryParams).then(res=> {
              this.list=res.list
              this.total = res.total
            })
        },
        select() {
          this.queryParams.start=this.parseTime(this.value1[0]).slice(0,10)
          this.queryParams.end=this.parseTime(this.value1[1]).slice(0,10)
          this.auto()
        },
    },
    components: {},
};
</script>

<style scoped>
.transfer {
    padding: 20px;
}
</style>
