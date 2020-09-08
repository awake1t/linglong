<template>
  <div>
    <!-- 面包屑导航 -->
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>设置</el-breadcrumb-item>
      <el-breadcrumb-item>扫描设置</el-breadcrumb-item>
    </el-breadcrumb>
    <el-card>
      <el-form :data="addForm" :model="addForm" ref="addFormRef" label-width="200px">
        <el-form-item label="敏感后台关键字" prop="login_word">
          <el-col :span="10">
            <el-input type="textarea" v-model="addForm.login_word" value="d1"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="敏感后台url" prop="login_url">
          <el-col :span="10">
            <el-input type="textarea" v-model="addForm.login_url"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="masscan线程" prop="masscan_thred">
          <el-col :span="10">
            <el-input  v-model="addForm.masscan_thred"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="masscan删除周期" prop="masscan_deltime">
          
          <el-tooltip class="item" effect="dark" content="n个周期没扫到就会删除" placement="top-end">
            <el-col :span="10">
            <el-input v-model="addForm.masscan_deltime"></el-input>
          </el-col>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="masscan要扫描的列表" prop="masscan_ip">
          <el-tooltip class="item" effect="dark" content="格式同masscan格式" placement="top-end">
          <el-col :span="10">
            <el-input type="textarea" v-model="addForm.masscan_ip"></el-input>
          </el-col>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="masscan要扫描的端口" prop="masscan_port">
          
          <el-tooltip class="item" effect="dark" content="格式同masscan格式" placement="top-end">
            <el-col :span="10">
            <el-input type="textarea" v-model="addForm.masscan_port"></el-input>
          </el-col>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="masscan不扫描的列表" prop="masscan_white">
          <el-col :span="10">
            <el-input type="textarea" v-model="addForm.masscan_white"></el-input>
          </el-col>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="editListInfo">更新</el-button>
          <!-- <el-button>取消</el-button> -->
        </el-form-item>
      </el-form>

      <!-- 底部区域 -->
      <!-- <span slot="footer" class="dialog-footer">
          <el-button @click="addDiglogVisable = false">取 消</el-button>
          <el-button type="primary" @click="addList">确 定</el-button>
      </span>-->
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
        login_word: [""],
        login_url: ""
      },

      addressForm: {
        address1: [],
        address2: "",
        masscan_white: ""
      }
    };
  },
  created() {
    this.getIpList();
  },
  methods: {
    async getIpList() {
      try {
        const { data: res } = await this.$http.get("/setting", {
          params: this.queryInfo
        });
        if (res.code !== 200) {
          return this.$message.error("获取列表失败");
        }
        this.addForm = res.data.lists[0];
        this.total = res.data.total;
      } catch (error) {
        return this.$message.error("cookie失效，请点击右上角退出重新登陆");
        error.message; // "Oops!"
      }
    },

    // this.editForm = JSON.parse(JSON.stringify(userInfo))
    editListInfo() {
      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return;
        const { data: resp } = await this.$http.put(
          `/setting`,
          qs.stringify(this.addForm),
          { headers: { "Content-Type": "application/x-www-form-urlencoded" } }
        );
        if (resp.code !== 200) {
          return this.$message.error(respeta.msg);
        }
        this.$message.success("更新成功,下一周期扫描使用此配置");
        this.$refs.addForm.resetFields();
        this.editDiglogVisable = false;
        // this.$message.success("修改成功");
        this.getIpList();
      });
    }
  }
};
</script>

<style lang="less" scoped>
</style>
