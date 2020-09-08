<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/task' }">任务列表</el-breadcrumb-item>
      <el-breadcrumb-item>任务结果</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <!-- 搜索 添加区域 -->
      <el-row :gutter="20">
        <el-col :span="14">
          <template v-for="item in taskstatus">
            <el-tag type="success" close-transition>{{formatterTaskStatus(item.status)}}</el-tag>
            <el-tag>总量: {{item.all_num}}</el-tag>
            <el-tag>成功: {{item.succes_num}}</el-tag>
            <!-- <el-tag>用户名字典: {{item.userdict}}</el-tag>
            <el-tag>密码字典: {{item.passdict}}</el-tag> -->
            <el-tag>运行耗时: {{item.run_time}}</el-tag>
            <el-tag>运行时间: {{item.created_time}}</el-tag>
          </template>
          
        </el-col>

        <el-col :span="6">
          <div id="app">
            <el-select
              v-model="selectForm.created_time"
              placeholder="选择历史记录"
              @change="getTaskTime(selectForm.created_time)"
              style="width:100%"
            >
              <el-option
                v-for="item in iplist"
                :key="item.created_time"
                :label="item.created_time"
                :value="item.created_time"
              >{{ item.created_time | dateformat('YYYY-MM-DD HH:mm:ss')}}</el-option>
            </el-select>
          </div>
          <!-- <div><a :href="downLoad()" > <el-button type="primary">下载报告</el-button></a></div> -->
        </el-col>
        <el-col :span="4">
          <div><a :href="downLoad()" > <el-button type="primary">下载报告</el-button></a></div>
        </el-col>
      </el-row>

      <!-- 资产列表区 -->
      <el-table :data="tasks" border stripe>
        <el-table-column label="ip" prop="ip"></el-table-column>
        <el-table-column label="port" prop="port"></el-table-column>
                <!-- <el-table-column label="漏洞类型" prop="vulntype" width="130px" align="center">
          <template scope="scope">
            <el-tag type="success" close-transition>{{formatterVulnType(scope.row.vulntype)}}</el-tag>
          </template>
        </el-table-column> -->

        <el-table-column label="用户名" prop="user"></el-table-column>
        <el-table-column label="密码" prop="pass"></el-table-column>

        <el-table-column label="发现时间" prop="created_time" align="center" width="180">
          <template scope="scope">{{scope.row.created_time| dateformat('YYYY-MM-DD HH:mm:ss')}}</template>
        </el-table-column>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.pagenum"
        :page-sizes="[200]"
        :page-size="queryInfo.pagesize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="taskstotal"
      ></el-pagination>

      <!-- 卡片结尾 -->
    </el-card>
  </div>
</template>

<script>
import qs from "qs";
import axios from "axios";

export default {
  data() {
    return {
      // 参数列表
      queryInfo: {
        taskid: "",
        ip: "",
        port: "",
        pagenum: 1,
        // 当前每页显示多少条数据
        pagesize: 10
      },
      iplist: [],
      total: 0,
      tasks: [],
      taskstotal: 0,
      taskstatus: [],
      taskstatustotal: 0,
      addDiglogVisable: false,
      addForm: {
        ip: ""
      },

      selectForm: {
        created_time: ""
      },
      editForm: {},
      editDiglogVisable: false,

      formRules: {}
    };
  },
  created() {
       this.getTaskLog();
    this.getTaskTime(0);
    this.getTaskStatus(0);
  },
  methods: {
      formatterVulnType(val) {
      if (val == 1) {
        return "弱口令";
      } else if (val == 2) {
        return "未授权";
      } else if (val == 3) {
        return "空口令";
      } else  {
        return "未知";
      }
    },

    formatterTaskStatus(val) {
      if (val == 1) {
        return "状态: 爆破中";
      } else if (val == 2) {
        return "状态: 已完成";
      }
    },

    async getTaskLog() {
      
      try {
        let id = this.$route.params.id;
        const { data: res } = await this.$http.get("/masstasklog/" + id);
        if (res.code !== 200) {
          return this.$message.error("获取列表失败");
        }
        this.iplist = res.data.lists;
        this.total = res.data.total;
      } catch (error) {
        return this.$message.error("cookie失效，请点击右上角退出重新登陆");
        error.message; // "Oops!"
      }
    },

    //获取任务时间记录，一个任务会有很多时间
    async getTaskTime(time) {
      let id = this.$route.params.id;
      const { data: resp } = await this.$http.get(
        "/masstasktime?taskid=" + id + "&tasktime=" + time
      );
      if (resp.code !== 200) {
        return this.$message.error(resp.msg);
        return this.$message.error("获取列表失败");
      }
      this.tasks = resp.data.lists;
      this.taskstotal = resp.data.total;
      this.getTaskStatus(time);
    },

    downLoad() {
       return axios.defaults.baseURL+'downtasklog/'+this.$route.params.id
    },


    async getTaskStatus(time) {
      let id = this.$route.params.id;
      const { data: resp } = await this.$http.get(
        "/masstaskstatus?taskid=" + id + "&tasktime=" + time
      );
      if (resp.code !== 200) {
        return this.$message.error(resp.msg);
        return this.$message.error("获取列表失败");
      }
      this.taskstatus = resp.data.lists;
      console.log(task);
    },

    // 监听pageSize改变事件
    handleSizeChange(newSize) {
      this.queryInfo.pagesize = newSize;
      console.log(newSize);
      this.getTaskList();
    },
    // 监听页码值改变的事件
    handleCurrentChange(newPage) {
      this.queryInfo.pagenum = newPage;
      console.log(newPage);
      this.getTaskList();
    },
    //监听对话框关闭事件
    diglogClose(operation) {
      if (operation === "add") {
        this.$refs.addFormRef.resetFields();
      } else if (operation === "edit") {
        this.$refs.editFormRef.resetFields();
      }
    }
  }
};
</script>

<style lang="less" scoped>
</style>
