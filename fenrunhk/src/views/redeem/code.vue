<template>
  <div class="app-container">
    <!--搜索开始-->
    <el-form
      ref="queryForm"
      :model="queryParams"
      :inline="true"
      label-width="68px"
    >
      <el-form-item label="" prop="keyword">
        <el-input
          v-model="queryParams.keyword"
          placeholder="兑换码/面额"
          clearable
          size="small"
          style="width: 240px"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>

      <!--
      <el-form-item label="创建时间">
        <el-date-picker
          v-model="dateRange"
          size="small"
          style="width: 240px"
          value-format="yyyy-MM-dd"
          type="daterange"
          range-separator="-"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        />
      </el-form-item>
      -->
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
    </el-form>
    <!--搜索结束-->

    <!--添加修改删除操作开始-->
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
          >新增</el-button
        >
      </el-col>
      <!-- <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:edit']"
          type="success"
          icon="el-icon-edit"
          size="mini"
          :disabled="single"
          @click="handleUpdate"
        >修改</el-button>
      </el-col> -->
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
          >删除</el-button
        >
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['system:sysconfig:export']"
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
          >导出</el-button
        >
      </el-col>
    </el-row>
    <!--添加修改删除操作结束-->

    <el-table
      v-loading="loading"
      :data="memberList"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="编号" width="80" align="center" prop="id" />
      <el-table-column label="兑换码" align="center" prop="code" />
      <el-table-column
        label="面额"
        align="center"
        prop="usdt"
        :show-overflow-tooltip="true"
      />
      <el-table-column
        label="兑换用户"
        align="center"
        prop="user_id"
        :show-overflow-tooltip="true"
      >
        <template #default='scope'>
          {{scope.row.user_id!=='0'?scope.row.user_id:'未兑换'}}
        </template>
      </el-table-column>
      <el-table-column
        label="兑换时间"
        align="center"
        prop="coin_time"
        width="200"
      >
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.coin_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="生成时间"
        align="center"
        prop="create_time"
        width="200"
      >
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.create_time) }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="状态"
        align="center"
        prop="status"
        :formatter="typeFormat"
      />
      <el-table-column
        label="操作"
        align="center"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button
            v-permisaction="['system:sysconfig:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleDelete(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="queryParams.pageIndex"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />

    <!-- 添加或修改参数配置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="460px">
      <el-form ref="form" :model="addParams" :rules="rules" label-width="80px">
        <el-form-item label="生成面额" prop="value">
          <el-input v-model="addParams.value" placeholder="请输入面额" />
        </el-form-item>
        <el-form-item label="生成数量" prop="count">
          <el-input v-model="addParams.count" placeholder="请输入数量" />
        </el-form-item>

        <!-- <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容" />
        </el-form-item> -->
      </el-form>
      <div slot="footer" class="dialog-footer">
        <span style="float:left;color:red;">当前兑换码总额：{{ Number(sum).toLocaleString() }}</span>
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  redeemCodeList,
  addRedeemCode,
  delRedeemCode,
  exportExcel
} from "@/api/redeem/code";

export default {
  name: "Member",
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
      memberList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // 类型数据字典
      typeOptions: [],
      // 会员等级数据字典
      typeLevel: [],
      // 日期范围
      dateRange: [],
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        status: "1",
        keyword: ""
        // cid: undefined
      },
      sum: 0,
      // 删除参数
      delParams: {
        ids: undefined
      },
      // 增加参数
      addParams: {
        count: undefined,
        value: undefined
      },
      // 编辑参数
      editParams: {
        email: "",
        address: "",
        status: "0",
        level: undefined
      },
      // 表单参数
      form: [],
      // 表单校验
      rules: {
        count: [{ required: true, message: "数量不能为空", trigger: "blur" }],
        value: [
          {
            required: true,
            validator: (rule, value, callback) => {
              if (value === ""||typeof(value) == "undefined") {
                callback(new Error("面额不能为空"));
              } else if (value > 50000) {
                callback(new Error("面额最大为50,000"));
              }else{
                callback();
              }
            },
            trigger: "blur"
          }
        ]
      }
    };
  },
  created() {
    this.getList();
    this.getDicts("code_status").then(response => {
      this.typeOptions = response.data;
    });
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true;
      redeemCodeList(this.addDateRange(this.queryParams, this.dateRange)).then(
        response => {
          this.memberList = response.data.list;
          this.total = response.data.count;
          this.sum = response.data.total;
          this.loading = false;
        }
      );
    },
    // 参数系统内置字典翻译
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions, row.status);
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
      // this.reset()
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
      //this.reset();
      this.open = true;
      this.title = "生成兑换码";
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id);
      this.single = selection.length !== 1;
      this.multiple = !selection.length;
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (
            this.addParams.value * this.addParams.count + this.sum >
            50000
          ) {
            this.$alert(
              "兑换码总USDT面额已超过50,000，请重新输入!",
              "警告",
              {}
            );
            return;
          }
          this.addParams.max = 50000
          addRedeemCode(this.addParams)
            .then(response => {
              if (response.code == 200) {
                this.msgSuccess("添加成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            })
            .catch(res => {
              console.log(res);
            });
        }
      });
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const configIds = row.id || this.ids;
      this.delParams.ids = configIds;
      var ids = "";
      if (configIds.constructor == Array) {
        for (var i = 0; i < configIds.length; i++) {
          ids += configIds[i];
          if (i < configIds.length - 1) {
            ids += ",";
          } else {
            this.delParams.ids = ids;
          }
        }
      }

      this.$confirm(
        '是否确认删除参数编号为"' + configIds + '"的数据项?',
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }
      )
        .then(() => {
          delRedeemCode(this.delParams);
        })
        .then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        })
        .catch(function() {});
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
