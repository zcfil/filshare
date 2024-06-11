<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form :inline="true" label-width="68px" :model="queryParams">
        <el-form-item label="状态" prop="status">
        <el-select
          v-model="queryParams.status"
          placeholder="状态"
          clearable
          size="small"
        >
          <el-option
            v-for="dict in typeOptions"
            :key="dict.dictValue"
            :label="dict.dictLabel"
            :value="dict.dictValue"
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
        <!--<el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>-->
      </el-form-item>

      <el-form-item label="" prop="status">
        <el-button type="primary" size="mini" @click="handleUpdate"
          >批量导入地址</el-button
        >
      </el-form-item>
    </el-form>
    <!--搜索结束-->

    <el-table
      v-loading="loading"
      :data="list"
      >
      <el-table-column label="编号" align="center" prop="id" />
      <el-table-column label="节点地址" align="center" prop="addr" />
      <el-table-column
        label="状态"
        align="center"
        prop="state"
        :formatter="typeFormat"
      />
      
    </el-table>

    <!-- 添加或修改参数配置对话框 -->
    <el-dialog
      title="请选择导入Excel表格文件"
      :visible.sync="open"
      width="380px"
    >
      <el-form ref="form" :model="form" :rules="rules">
        <el-form-item label="" prop="status">
          <el-upload
            class="upload-demo"
            ref="upload"
            :action="url"
            :auto-upload="false"
            accept=".xls,.xlsx"
            :on-success = "upSucc"
          >
            <el-button slot="trigger" size="small" type="primary"
              >选取文件</el-button
            >
            <el-button
              style="float:right;"
              size="small"
              type="success"
              @click="submitUpload"
              >开始导入</el-button
            >
            <div slot="tip" class="el-upload__tip" style="color:red">只能上传 xls/xlsx 文件</div>
          </el-upload>
        </el-form-item>
      </el-form>
    </el-dialog>
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
import { listUploadAddress } from "@/api/dice/dice";

export default {
  name: "Dice",
  data() {
    return {
      // 遮罩层
      loading: true,
      // 上传地址
      //url: "http://192.168.0.134:10124/api/v1/uploadAddress",
      url: '',
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
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        cid: undefined,
        status: '0'
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        configName: [
          { required: true, message: "参数名称不能为空", trigger: "blur" }
        ],
        configKey: [
          { required: true, message: "参数键名不能为空", trigger: "blur" }
        ],
        configValue: [
          { required: true, message: "参数键值不能为空", trigger: "blur" }
        ]
      }
    };
  },
  created() {
    this.getList();
    this.url = process.env.VUE_APP_BASE_API+'/api/v1/uploadAddress'
    //alert(this.url)
    this.getDicts("addr_state").then(response => {
      this.typeOptions = response.data;
    });
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = false;
      listUploadAddress(
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
      return this.selectDictLabel(this.typeOptions, row.state);
    },
    handleQuery() {
      this.queryParams.pageIndex = 1;
      this.getList();
    },
    handleUpdate(){
      this.open = true
    },
    submitUpload() {
        this.$refs.upload.submit()
    },
    upSucc(response, file, fileList){
      if (response.code==200){
        this.msgSuccess('导入成功')
        this.open = false
        this.getList()
      }else{
        this.msgError(response.msg)
      }
    }
  }
};
</script>
