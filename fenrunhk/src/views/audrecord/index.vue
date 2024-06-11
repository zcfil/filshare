<template>
    <div class="audrecord">
        <el-table :data="list">
          <el-table-column label="姓名" prop="nick_name" align="center" />
          <el-table-column label="用户名" prop="username" align="center" />
          <el-table-column label="手机号" prop="phone" align="center" />
          <el-table-column label="推荐人" prop="referrer_name" align="center" />
          <el-table-column label="状态" prop="is_pass" align="center">
            <template slot-scope="scope">
                <span class="a" v-show=" scope.row.is_pass==1 ">已通过</span>
                <span class="b" v-show=" scope.row.is_pass==0 ">已拒绝</span>
            </template>
          </el-table-column>
          <el-table-column label="操作时间" width="160" align="center" prop="pass_time">
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.pass_time) }}</span>
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
import { userAuditRecord } from '@/api/system/audituser'
export default {
    name:'Audrecord',
    props: {},
    data() {
        return {
            queryParams: {
              pageIndex: 1,
              pageSize: 10,
            },
            total: 0,
            list:[],
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
           userAuditRecord(this.queryParams).then(res=> {
              this.list=res.data.list
              this.total=Number(res.data.totalCount)
           }) 
        }
    },
    components: {},
};
</script>

<style scoped>
.audrecord {
    padding: 20px;
}
.a {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(170, 255, 170);
  color: rgb(0, 255, 0);
  background-color: rgb(235, 255, 235);
}
.b {
  display: inline-block;
  padding: 5px 14px;
  border: 1px solid rgb(255, 180, 180);
  color: rgb(255, 80, 80);
  background-color: rgb(255, 221, 221);
}
</style>
