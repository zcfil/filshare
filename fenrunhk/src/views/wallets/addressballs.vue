<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form
      ref="queryForm"
      :model="queryParams"
      :inline="true"
      label-width="68px"
    >
      <el-form-item label="币种名称" prop="status">
        <el-select
          v-model="queryParams.tokens"
          placeholder="币种名称"
          clearable
          size="small"
        >
          <el-option
            v-for="dict in typeOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictLabel"
          />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          icon="el-icon-search"
          size="mini"
          @click="handleQuery"
          >搜索</el-button
        >
      </el-form-item>

      <el-form-item >
        <el-button
          style ="margin-left:30px"
          v-permisaction="['system:sysconfig:edit']"
          type="success"
          icon="el-icon-edit"
          size="mini"
          @click="UpdateData"
          >更新数据</el-button
        >
      </el-form-item>
      <el-form-item label="上次更新时间:" prop='updatetime' label-width = '300'>
        <el-input v-model="updatetime" :disabled="true">
        </el-input>
      </el-form-item>
    </el-form>
    <!--搜索结束-->

    <el-table
      v-loading="loading"
      :data="list"
      @selection-change="handleSelectionChange"
      >>
      <el-table-column label="编号" align="center" prop="id" />
      <el-table-column label="节点地址" align="center" prop="addres" />
      <el-table-column label="币种类型" align="center" prop="symbol" />
      <el-table-column label="数量" align="center" prop="quantity" />
      <!-- <el-table-column label="raw" align="center" prop="raw" />
       <el-table-column label="assetref" align="center" prop="assetref" /> -->
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="queryParams.pageIndex"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />
  </div>
</template>

<script>
import { listAddressBalls, exportExcel,UpdateWalletsAssets } from "@/api/dice/dice";

export default {
  name: "Dice",
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 参数表格数据
      list: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // 类型数据字典
      typeOptions: [],
      // 日期范围
      dateRange: [],
      // 查询参数
      updatetime: '',
      //判断是否跟updatetime是同一小时，避免短时间内重复更新
      flagtime: "2020-05-20 00:00:00",
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        cid: undefined,
        tokens: undefined
      },
      // 表单参数
      form: {}
    };
  },
  created() {
    this.getList();
    this.getDicts("asset_type").then(response => {
      this.typeOptions = response.data;
    });
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = false;
      listAddressBalls(
        this.addDateRange(this.queryParams, this.dateRange)
      ).then(response => {
        // var data = response.data
        // var arr = []
        // for (const key in data) {
        //   if (key !== 'total') {
        //     data[key].map(item => {
        //       item.key = key
        //     })
        //     arr.push(...data[key])
        //   }
        // }
        if  (response.data.list.length>0){
            this.updatetime = response.data.list[0].updatetime;
        }
        this.list = response.data.list;
        this.total = response.data.count;
        this.loading = false;
      });
    },
    // 参数系统内置字典翻译
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions, row.configType);
    },
    // 日期处理
    dateFormat(row, column) {
      var date = row[column.property];
      if (date === undefined) {
        return "";
      }
      // eslint-disable-next-line no-undef
      return moment(date).format("YYYY-MM-DD HH:mm:ss");
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },
    //更新数据库时间
    UpdateData(){
      if (this.updatetime.substring(13)===this.flagtime.substring(13)){
        //alert('更新周期过短，请稍后再试！');
        this.$alert('更新周期过短，请稍后再试！');
        return 
      }
      UpdateWalletsAssets().then(res => {
         this.flagtime = this.updatetime
      });
    },
    // 表单重置
    reset() {
      this.form = {
        configId: undefined,
        configName: undefined,
        configKey: undefined,
        configValue: undefined,
        configType: "Y",
        remark: undefined
      };
      this.resetForm("form");
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1;
      this.getList();
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = [];
      this.resetForm("queryForm");
      this.handleQuery();
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.open = true;
      this.title = "添加参数";
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.configId);
      this.single = selection.length !== 1;
      this.multiple = !selection.length;
    },
    /** 导出按钮操作 */
    handleExport() {
      const queryParams = this.queryParams;
      this.$confirm("是否确认导出所有参数数据项?", "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(function() {
        exportExcel(queryParams).then(res => {
          location.href = process.env.VUE_APP_BASE_API + `/` + res.msg;
        });
      });
    }
  }
};
</script>
