<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>资产管理</el-breadcrumb-item>
      <el-breadcrumb-item>ip资产</el-breadcrumb-item>
    </el-breadcrumb>
    <el-card>
      <!-- 搜索 添加区域 -->
      <el-row :gutter="30">
        <el-col :span="6">
          <el-input
            placeholder="请输入Ip"
            class="input-with-select"
            v-model="queryInfo.ip"
            clearable
            @clear="getIpList"
            @keyup.enter.native="getIpListSingle"
          ></el-input>
        </el-col>
        <el-col :span="4">
          <el-input
            placeholder="请输入title"
            class="input-with-select"
            v-model="queryInfo.title"
            clearable
            @clear="getIpList"
            @keyup.enter.native="getIpListSingle"
          >
          </el-input>
        </el-col>
                <el-col :span="4">
          <el-input
            placeholder="请输入Port"
            class="input-with-select"
            v-model="queryInfo.port"
            clearable
            @clear="getIpList"
            @keyup.enter.native="getIpListSingle"
          >
          </el-input>
        </el-col>
        <el-col :span="8">
          <el-button type="primary" @click="getIpListSingle">搜索</el-button>
        </el-col>
      </el-row>

      <!-- 资产列表区 -->
      <el-table :data="iplist" border stripe>
        <el-table-column label="ip" prop="ip">
        <template slot-scope="scope">
            <a
              :href="'http://'+scope.row.ip+':'+scope.row.port"
              target="_blank"
              prop
              style="color:#606266;"
            >{{ scope.row.ip }}</a>
          </template>
          </el-table-column>
        <el-table-column label="port" prop="port"></el-table-column>
        <el-table-column label="protocol" prop="protocol"></el-table-column>
        <el-table-column label="title" prop="title"></el-table-column>
        <el-table-column label="server" prop="server"></el-table-column>
        <el-table-column label="指纹" prop="cms"></el-table-column>
        <!-- <el-table-column label="测试布尔" prop="fandomain">
          <template slot-scope="scope">
            {{scope.row.fandomain}}
            <el-switch v-model="scope.row.fandomain"></el-switch>
          </template>
        </el-table-column>-->
        <!-- <el-table-column label="操作" width="180px">
          <template slot-scope="scope">
            <el-button
              type="primary"
              icon="el-icon-edit"
              size="mini"
              @click="editDiglogInit(scope.row)"
            ></el-button>
            <el-button
              type="danger"
              icon="el-icon-delete"
              size="mini"
              @click="deleteById(scope.row.id)"
            ></el-button>
            <el-tooltip effect="dark" content="再想想" placement="top" :enterable="false">
              <el-button
                type="warning"
                icon="el-icon-setting"
                size="mini"
                @click="setRole(scope.row)"
              ></el-button>
            </el-tooltip>
          </template>
        </el-table-column> -->
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
        title="添加信息"
        :visible.sync="addDiglogVisable"
        width="50%"
        @close="diglogClose('add')"
      >
        <!-- 内容主体区域  :rules="formRules"  -->
        <el-form :model="addForm" ref="addFormRef" label-width="70px">
          <el-form-item label="ip" prop="ip">
            <el-input v-model="addForm.ip"></el-input>
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
      // 参数列表
      queryInfo: {
        ip: "",
        title: "",
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
      },

      editForm: {},
      editDiglogVisable: false,

      formRules: {
        // ip: [
        //   { required: true, message: '请输入ip', trigger: 'blur' },
        //   { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        // ]
        // password: [
        //   { required: true, message: '请输入密码', trigger: 'blur' },
        //   { min: 3, max: 16, message: '长度在 3 到 16 个字符', trigger: 'blur' }
        // ],
        // email: [
        //   { required: true, message: '请输入邮箱', trigger: 'blur' },
        //   { validator: checkEmail, trigger: 'blur' }
        // ],
        // mobile: [
        //   { required: true, message: '请输入手机号', trigger: 'blur' },
        //   { validator: checkMobile, trigger: 'blur' }
        // ]
      }
    };
  },
  created() {
    this.getIpList();
  },
  methods: {
    async getIpList() {
      try {
        const { data: res } = await this.$http.get("/masstasks", {
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
      const { data: res } = await this.$http.get("/masstask", {
        params: this.queryInfo
      });
      if (res.code !== 200) {
        return this.$message.error("获取列表失败");
      }
      this.iplist = res.data.lists;
      this.total = res.data.total;
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
    // 监听switch开关状态的改变
    // async userStateChange(userInfo) {
    //   const { data: resp } = await this.$http.put(`users/${userInfo.id}/state/${userInfo.state}`)
    //   if (resp.meta.status !== 200) {
    //     this.userInfo.state = !this.userInfo.state
    //     this.$message.error('更新状态失败')
    //   }
    //   this.$message.success('更新状态成功')
    // },
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
        console.log(this.addForm.ip);
        // const { data: resp } = await this.$http.post('/jobiplist', { params: this.addForm })
        const { data: resp } = await this.$http.post(
          "/jobiplist",
          qs.stringify(this.addForm),
          { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
        );
        if (resp.code !== 200) {
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
      //   const confirmResult = await this.$confirm('此操作将永久删除, 是否继续?', '提示', {
      //     confirmButtonText: '确定',
      //     cancelButtonText: '取消',
      //     type: 'warning'
      //   }).catch(err => err)
      //   if (confirmResult !== 'confirm') {
      //     return this.$message.info('已取消删除')
      //   }
      const { data: resp } = await this.$http.delete("/jobiplist?id=" + id);
      if (resp.code !== 200) {
        return this.$message.error(resp.msg);
      }
      this.$message.success("删除成功");
      this.getIpList();
    }
  }
};
</script>

<style lang="less" scoped>
a{text-decoration:none}
</style>
