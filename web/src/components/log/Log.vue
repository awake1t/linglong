<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>设置</el-breadcrumb-item>
      <el-breadcrumb-item>运行日志</el-breadcrumb-item>
    </el-breadcrumb>

    <!--    <el-card>-->
<!--        <el-row :gutter="20">-->
<!--          <el-col :span="4">-->
<!--            <div class="grid-content bg-purple"></div>-->
<!--            <el-input-->
<!--                    placeholder="请输入指纹名称"-->
<!--                    class="input-with-select"-->
<!--                    v-model="queryInfo.name"-->
<!--                    clearable-->
<!--                    @clear="getIpList"-->
<!--                    @keyup.enter.native="getIpListSingle"-->
<!--            ></el-input>-->
<!--          </el-col>-->
<!--          <el-col :span="2">-->
<!--            <el-button type="primary" @click="getIpListSingle">搜索</el-button>-->
<!--          </el-col>-->

<!--          <el-col :span="3" >-->
<!--            <el-button type="primary" @click="addDiglogVisable = true">清空日志</el-button>-->
<!--          </el-col>-->
<!--        </el-row>-->

      <!-- 资产列表区 -->
      <el-table :data="iplist" border stripe>

        <el-table-column label="状态" prop="status">
          <template scope="scope">
            <div v-if="scope.row.status == 0">
              <el-tag type="info" close-transition>{{formatterStatus(scope.row.status)}}</el-tag>
            </div>
            <div v-else-if="scope.row.status == 1">
              <el-tag close-transition>{{formatterStatus(scope.row.status)}}</el-tag>
            </div>
          </template>
        </el-table-column>

        <el-table-column label="任务类型" prop="task_type"></el-table-column>
        <el-table-column label="成功数量" prop="all_num"></el-table-column>
        <el-table-column label="扫描耗时" prop="run_time"></el-table-column>
        <el-table-column label="完成时间" prop="created_time" align="center" width="180">
          <template scope="scope">{{scope.row.created_time| dateformat('YYYY-MM-DD HH:mm:ss')}}</template>
        </el-table-column>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.pagenum"
        :page-sizes="[10]"
        :page-size="queryInfo.pagesize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
      ></el-pagination>
    </el-card>
  </div>
</template>

<script>
import qs from "qs";
export default {
  data() {
    return {
      queryInfo: {
        ip: "",
        port: "",
        pagenum: 1,
        // 当前每页显示多少条数据
        pagesize: 10
      },

      iplist: [],
      total: 0,
      addDiglogVisable: false,
      addForm: {
        ip: ""
      }
    };
  },
  created() {
    this.getIpList();
  },
  methods: {
    formatterStatus(val) {
      if (val == '0') {
        return '运行中'
      } else if (val == 1) {
        return '已完成'
      } else if (val == '3') {
        return '任务失败'
      } else if (val == '4') {
        return '任务失败'
      }
    },
    async getIpList() {
      try {
        const { data: res } = await this.$http.get("/log", {
          params: this.queryInfo
        });
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

    // 监听pageSize改变事件
    handleSizeChange(newSize) {
      this.queryInfo.pagesize = newSize;
      console.log(newSize);
      this.getIpList();
    },
    // 监听页码值改变的事件
    handleCurrentChange(newPage) {
      this.queryInfo.pagenum = newPage;
      console.log(newPage);
      this.getIpList();
    }
  }
};
</script>

<style lang="less" scoped>
</style>
