<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>指纹管理</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <!-- 搜索 添加区域 -->
      <el-row :gutter="20">
        <el-col :span="4">
          <div class="grid-content bg-purple"></div>
          <el-input
            placeholder="请输入指纹名称"
            class="input-with-select"
            v-model="queryInfo.name"
            clearable
            @clear="getIpList"
            @keyup.enter.native="getIpListSingle"
          ></el-input>
        </el-col>

        <el-col :span="4">
          <div class="grid-content bg-purple"></div>
          <el-input
                  placeholder="请输入指纹详细内容"
                  class="input-with-select"
                  v-model="queryInfo.finger"
                  clearable
                  @clear="getIpList"
                  @keyup.enter.native="getIpListSingle"
          ></el-input>
        </el-col>

        <el-col :span="2">
          <el-button type="primary" @click="getIpListSingle">搜索</el-button>
        </el-col>


        <el-col :span="3" >
            <el-button type="primary" @click="addDiglogVisable = true">新增指纹</el-button>
          </el-col>
      </el-row>

      <!-- 资产列表区 -->
      <el-table :data="iplist" border stripe>
        <el-table-column label="name" prop="name"></el-table-column>
        <el-table-column label="description" prop="description"  width="150px" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="finger" prop="finger" width="200px" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column
          label="创建时间"
          prop="created_time"
        >{{ created_at | dateformat('YYYY-MM-DD HH:mm:ss')}}</el-table-column>
        <el-table-column
          label="更新时间"
          prop="updated_time"
        >{{ updated_time | dateformat('YYYY-MM-DD HH:mm:ss')}}</el-table-column>
        <el-table-column label="操作" width="180px">
          <template slot-scope="scope">

            <el-tooltip effect="dark" content="对数据库资产指纹识别" placement="top" >
              <el-button
                      type="warning"
                      icon="el-icon-s-promotion"
                      size="mini"
                      @click="scanfinger(scope.row.id)"
              ></el-button>
            </el-tooltip>

            <el-tooltip effect="dark" content="删除指纹" placement="top" >
            <el-button
                    type="danger"
                    icon="el-icon-delete"
                    size="mini"
                    @click="deleteById(scope.row.id)"
            ></el-button>
            </el-tooltip>

          </template>
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

      <!-- 添加列表对话框 -->
      <el-dialog
        title="新增指纹"
        :visible.sync="addDiglogVisable"
        width="50%"
        @close="diglogClose('add')"
      >
        <!-- 内容主体区域  :rules="formRules"  -->
        <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="100px">
          <el-form-item label="指纹名称" prop="name">
            <el-col :span="14">
              <el-input v-model="addForm.name" placeholder="指纹名称"></el-input>
            </el-col>
          </el-form-item>

          <el-form-item label="指纹描述" prop="description">
            <el-col :span="14">
              <el-input v-model="addForm.description" placeholder="指纹描述"></el-input>
            </el-col>
          </el-form-item>

          <el-form-item label="测试地址" prop="testurl">
            <el-col :span="14">
              <el-input v-model="addForm.testurl" placeholder="输入一个url测试编写的指纹是否正确"></el-input>
            </el-col>
          </el-form-item>

          <el-form-item label="指纹内容" prop="finger">
            <el-tooltip class="item" effect="dark" content="新指纹一定要测试是否正常" placement="top-end">
            <el-col :span="14">
              <el-input type="textarea" v-model="addForm.finger" :placeholder="placeholderValue" :autosize="{minRows: 15,maxRows: 7}" ></el-input>
            </el-col>
            </el-tooltip>
          </el-form-item>

        </el-form>
        <!-- 底部区域 -->
        <span slot="footer" class="dialog-footer">
          <el-col :span="5">
          <el-link type="primary" href="https://github.com/awake1t/linglong/blob/master/img/Finger.md" target="_blank" ><i class="el-icon-view el-icon--right"></i>   如何编写指纹</el-link>
          </el-col>
          <el-button type="danger" @click="testFinger">测试指纹</el-button>
          <el-button @click="addDiglogVisable = false">取 消</el-button>
          <el-button type="primary" @click="addList">保存</el-button>
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
      placeholderValue: '' +
              '"Shiro" : {\n' +
              '  "cookies" : {\n' +
              '      "rememberMe": ""\n' +
              '   },\n' +
              '   "html": "[test]some html word"\n' +
              '}' +
              '',

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
        testurl: "",
        source: "1",
        finger: '' +
                '"Shiro" : {\n' +
                '  "cookies" : {\n' +
                '  "rememberMe": ""\n' +
                '  },\n' +
                '  "html": "[test]some html word"\n' +
                ' }' +
                '',
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
        const { data: res } = await this.$http.get("/finger", {
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
      const { data: res } = await this.$http.get("/finger", {
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
          "/addfinger",
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

    // 测试指纹
    testFinger() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return;
        const { data: resp } = await this.$http.post(
          "/testfinger",
          qs.stringify(this.addForm),
          { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
        );

        if (resp.code == 401) {
          this.$message.success("cookie失效");
          return this.$message.error(resp.msg);
        }
        if (resp.code == 402) {
          return this.$message.error("指纹识别失败");
        }
        // this.$refs.addFormRef.resetFields();
        this.$message.success("指纹识别成功,点击保存指纹");
        // this.getIpList();
      });
    },

    async scanfinger(id) {
      try {
        const { data: res } = await this.$http.get("/scanfinger/" + id);
        if (res.code !== 200) {
          return this.$message.error("获取列表失败");
        }
        this.iplist = res.data.lists;
        this.total = res.data.total;
      } catch (error) {
        return this.$message.error("cookie失效，请点击右上角退出重新登陆");
        error.message; // "Oops!"
      }
      this.$message.success("开始进行指纹识别,设置中可查看运行日志");
      this.getIpList();
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
      const { data: resp } = await this.$http.delete("/delfinger/" + id);
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
