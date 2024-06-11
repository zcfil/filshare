<template>
    <div class="settlement">
        <el-table :data="customerlist">
          <el-table-column label="日期" prop="date" align="center">
            <template slot-scope="scope">
              <span style="color: #007AFF;cursor: pointer;" @click="$router.push({path: '/list/user?date=' + scope.row.date,})">{{ scope.row.date }}</span>
            </template>
          </el-table-column>
          <el-table-column label="金额" prop="amount" align="center" show-overflow-tooltip>
            <template slot-scope="scope">
              <span>{{ moneyFormat(scope.row.amount) }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button
                v-permisaction="['system:syssettlement:transfer']"
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
    </div>
</template>

<script>
import { getWeekList,transferWeek } from '@/api/customer/customer'
export default {
    name:'Settlement',
    props: {},
    data() {
        return {
            customerlist:[],
            queryParams: {
              pageIndex: 1,
              pageSize: 10,
            },
            total: 0,
            num:"",
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
          getWeekList(this.queryParams).then(res=> {
            this.num=null
            var a=Number(new Date());
            for (var i=0;i<res.list.length;i++) {
              if (a>=Date.parse(res.list[i].date.substring(0,10)+" 00:00:00")&&a<=Date.parse(res.list[i].date.substring(res.list[i].date.length-10,res.list[i].date.length)+" 23:59:59")) {
                  this.num=i
              }
            }
            if (this.num!=null) {
              res.list.splice(this.num,1)
              this.total = res.total-1
            } else {
              this.total = res.total
            }
            this.customerlist=res.list
            
          })
        },
        transfer(row) {
          var that=this
          this.$confirm('确认要转 "' + row.date  + '" 这周的帐吗?', '警告', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(function() {
            transferWeek({date:row.date}).then(res=> {
              that.msgSuccess('转账成功');
              that.auto()
            })
          }).catch(() => {
            // console.log("取消")
          })
        }
    },
    components: {},
};
</script>

<style scoped>
.settlement {
    padding: 20px;
}
</style>
