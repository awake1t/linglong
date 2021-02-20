<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>任务列表</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <!-- 搜索 添加区域 -->
      <el-row :gutter="20">
        <el-col :span="7">
          <div class="grid-content bg-purple"></div>
          <el-input
            placeholder="请输入title"
            class="input-with-select"
            v-model="queryInfo.title"
            clearable
            @clear="getIpList"
            @keyup.enter.native="getIpList"
          ></el-input>
        </el-col>

        <el-col :span="8">
          <el-button type="primary" @click="getIpList">搜索</el-button>
        </el-col>
      </el-row>

      <!-- 资产列表区 -->
      <el-table :data="iplist" border stripe>
        <el-table-column label="title" min-width="90" prop="title" sortable="custom">
          <template slot-scope="scope">
            <a
              :href="scope.row.url"
              target="_blank"
              prop
              style="color:#606266;"
            >{{ scope.row.title }}</a>
          </template>
        </el-table-column>
        <el-table-column label="url" prop="url" width="280px" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="loginurl" prop="loginurl"></el-table-column>
        <el-table-column label="ip" prop="ip"></el-table-column>
        <el-table-column label="port" prop="port"></el-table-column>

        <el-table-column
          label="创建时间"
          prop="created_time"
        >{{ created_at | dateformat('YYYY-MM-DD HH:mm:ss')}}</el-table-column>
        <el-table-column
          label="更新时间"
          prop="updated_time"
        >{{ updated_time | dateformat('YYYY-MM-DD HH:mm:ss')}}</el-table-column>
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

      <!-- 添加列表对话框 -->
      <el-dialog
        title="新增任务"
        :visible.sync="addDiglogVisable"
        width="50%"
        @close="diglogClose('add')"
      >
        <!-- 内容主体区域  :rules="formRules"  -->
        <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
          <el-form-item label="任务名称" prop="taskname">
            <el-col :span="14">
              <el-input v-model="addForm.taskname" placeholder="请输入任务名称"></el-input>
            </el-col>
          </el-form-item>
          <el-form-item label="描述" prop="description">
            <el-col :span="14">
              <el-input v-model="addForm.description" placeholder="任务描述,非必填"></el-input>
            </el-col>
          </el-form-item>

          <el-form-item label="任务类型" prop="brute">
            <el-select v-model="addForm.brute">
              <el-option label="爆破SSH" value="ssh"></el-option>
              <el-option label="爆破MYSQL" value="mysql"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="任务周期" prop="cronspec">
            <el-tabs v-model="addForm.cronspec" type="card" @tab-click="handleClick">
              <el-tab-pane label="立即执行" name="now">执行一次,立即执行</el-tab-pane>
              <el-tab-pane label="每天一次" name="day">
                每天
                <el-input-number
                  v-model="addForm.hour"
                  @change="handleChange"
                  :min="0"
                  :max="23"
                  size="mini"
                  label="描述文字"
                ></el-input-number>点执行
              </el-tab-pane>
              <el-tab-pane label="每周一次" name="week">
                每周
                <el-input-number
                  v-model="addForm.day"
                  @change="handleChange"
                  :min="1"
                  :max="7"
                  size="mini"
                  label="描述文字"
                ></el-input-number>点,每天
                <el-input-number
                  v-model="addForm.hour"
                  @change="handleChange"
                  :min="0"
                  :max="23"
                  size="mini"
                  label="描述文字"
                ></el-input-number>点执行
              </el-tab-pane>
              <el-tab-pane label="自定义" name="cmd">
                <el-input
                  type="textarea"
                  v-model="addForm.cronspecmd"
                  placeholder="请输入自定义的crontab表达式,参考:https://crontab.guru/"
                ></el-input>
              </el-tab-pane>
            </el-tabs>
          </el-form-item>

          <el-form-item label="任务资产" prop="source">
            <el-tabs v-model="addForm.source" type="card" @tab-click="handleClick">
              <el-tab-pane
                label="自动匹配"
                name="1"
              >推荐默认使用，跟爆破类型匹配数据库对应资产，比如爆破类型是mysql 会匹配数据中 port=3306+server=mysql的资产</el-tab-pane>

              <el-tab-pane label="手动输入列表" name="2">
                <el-input
                  type="textarea"
                  v-model="addForm.sourcecontent"
                  placeholder="10.10.10.10:22
                  10.10.10.11:1433
                  10.10.10.10:23|SSH
                  10.10.10.10:3307|MYSQL"
                ></el-input>
              </el-tab-pane>
            </el-tabs>
          </el-form-item>

          <el-form-item label="线程" prop="thread">
            <el-select v-model="addForm.thread">
              <el-option label="100" value="100"></el-option>
              <el-option label="200" value="200"></el-option>
              <el-option label="500" value="500"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <!-- 底部区域 -->
        <span slot="footer" class="dialog-footer">
          <el-button @click="addDiglogVisable = false">取 消</el-button>
          <el-button type="primary" @click="addList">确 定</el-button>
        </span>
      </el-dialog>

      <!-- 编辑列表对话框 -->
      <el-dialog
        title="修改列表"
        :visible.sync="editDiglogVisable"
        width="50%"
        @close="diglogClose('edit')"
      >
        <!-- 内容主体区域 -->
        <el-form :model="editForm" ref="editFormRef" label-width="70px">
          <el-form-item label="ip">
            <el-input v-model="editForm.ip"></el-input>
          </el-form-item>
        </el-form>
        <!-- 底部区域 -->
        <span slot="footer" class="dialog-footer">
          <el-button @click="editDiglogVisable = false">取 消</el-button>
          <el-button type="primary" @click="editListInfo">确 定</el-button>
        </span>
      </el-dialog>
      <!-- 卡片结尾 -->
    </el-card>
  </div>
</template>

<script>
import qs from "qs";
export default {
  data() {
    return {
      num: 1,
      // 参数列表
      queryInfo: {
        title: "",
        pagenum: 1,
        // 当前每页显示多少条数据
        pagesize: 10
      },
      iplist: [],
      total: 0,
      addDiglogVisable: false,
      addForm: {
        taskname: "",
        description: "",
        cronspec: "now",
        day: 1,
        hour: 1,
        cronspecmd: "",
        thread: "100",
        brute: "",
        source: "1",
        sourcecontent: ""
      },
      editForm: {},
      editDiglogVisable: false,

      // 前端验证
      addFormRules: {
        taskname: [
          { required: true, message: "请输入任务名称", trigger: "blur" },
          {
            min: 2,
            max: 20,
            message: "任务名的长度在2～20个字",
            trigger: "blur"
          }
        ],
        brute: [
          { required: true, message: "请选择任务类型", trigger: "blur" },
          {
            min: 2,
            max: 18,
            message: "任务类型长度在2～18个字",
            trigger: "blur"
          }
        ],
        codetype: [
          { required: true, message: "请选择代码类型", trigger: "blur" },
          {
            min: 2,
            max: 18,
            message: "代码类型长度在2～18个字",
            trigger: "blur"
          }
        ],
        filename: [
          { required: true, message: "文件上传未完成", trigger: "blur" }
        ]
      }
    };
  },
  created() {
    this.getIpList();
  },
  methods: {
    formatterTaskType(val) {
      if (val == "ssh") {
        return "ssh爆破";
      } else if (val == "mysql") {
        return "mysql爆破";
      } else if (val == "3") {
        return "已完成";
      } else if (val == "4") {
        return "任务失败";
      }
    },
    formatterTaskCycle(val) {
      if (val == "now") {
        return "只运行一次";
      } else {
        return val;
      }
    },
    async getIpList() {
      try {
        const { data: res } = await this.$http.get("/webloginlist", {
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
    async getIpListSingle() {
      const { data: res } = await this.$http.get("/webloginlistsearch", {
        params: this.queryInfo
      });
      if (res.code !== 200) {
        return this.$message.error("获取列表失败");
      }
      this.iplist = res.data.lists;
      this.total = res.data.total;
      //   console.log(res)
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
    },

    //监听对话框关闭事件
    diglogClose(operation) {
      if (operation === "add") {
        this.$refs.addFormRef.resetFields();
      } else if (operation === "edit") {
        this.$refs.editFormRef.resetFields();
      }
    },

    // 添加事件
    addList() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return;
        const { data: resp } = await this.$http.post(
          "/addcron",
          qs.stringify(this.addForm),
          { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
        );

        if (resp.code == 401) {
          this.$message.success("cookie失效");
          console.log("resp");
          console.log(resp);
          return this.$message.error(resp.msg);
        }
        this.$refs.addFormRef.resetFields();
        this.addDiglogVisable = false;
        this.$message.success("添加成功");
        this.getIpList();
      });
    },

    editDiglogInit(userInfo) {
      this.editForm = JSON.parse(JSON.stringify(userInfo));
      this.editDiglogVisable = true;
    },

    editListInfo() {
      this.$refs.editFormRef.validate(async valid => {
        if (!valid) return;
        const {
          data: resp
        } = await this.$http.put(
          `/jobiplist/${this.editForm.id}`,
          qs.stringify(this.editForm),
          { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
        );
        if (resp.code !== 200) {
          return this.$message.error(respeta.msg);
        }
        this.$refs.editFormRef.resetFields();
        this.editDiglogVisable = false;
        this.$message.success("修改成功");
        this.getIpList();
      });
    },

    async deleteById(id) {
      const confirmResult = await this.$confirm(
        "此操作将永久删除, 是否继续?",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }
      ).catch(err => err);
      if (confirmResult !== "confirm") {
        return this.$message.info("已取消删除");
      }
      const { data: resp } = await this.$http.delete("/delcron/" + id);
      if (resp.code !== 200) {
        return this.$message.error(resp.msg);
      }
      this.$message.success("删除成功");
      this.getIpList();
    },
    gototasklog(id) {
      this.$router.push({ path: `/tasklog/${id}` });
    }
  }
};
</script>

<style lang="less" scoped>
a{text-decoration:none}
</style>
