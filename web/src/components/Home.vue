<template>
  <!-- 页面主题区域 -->
  <el-container>
    <!-- 左边侧边栏 -->
    <el-aside :width="isCollapse ? '64px' : '200px'">
      <!-- 左边侧边菜单栏 -->
      <el-menu
        background-color="#293846"
        text-color="#fff"
        active-text-color="#ffd04b"
        unique-opened
        :collapse="isCollapse"
        :collapse-transition="isTransition"
        router
        default-active="./jobips"
      >
        <!-- 一级菜单 -->
        <!-- 首页 -->
        <el-menu-item index="/welcome">
          <i class="el-icon-s-home"></i>
          <span slot="title">首页</span>
        </el-menu-item>

        <el-menu-item index="/jobips">
          <i class="el-icon-menu"></i>
          <span slot="title">ip资产</span>
        </el-menu-item>

        <el-menu-item index="/webloginlist">
          <i class="el-icon-s-promotion"></i>
          <span slot="title">敏感后台</span>
        </el-menu-item>

        <el-menu-item index="/task">
          <i class="el-icon-s-order"></i>
          <span slot="title">任务列表</span>
        </el-menu-item>

        <el-submenu index="/zichan">
          <template slot="title">
            <!-- 图标 -->
            <i class="el-icon-s-tools"></i>
            <!-- 文本 -->
            <span slot="title">设置</span>
          </template>
          <!-- 二级菜单 -->
          <el-menu-item-group>
            <el-menu-item index="/log" @click="saveNavState('/log')">
              <i class="el-icon-s-help"></i>
              <span>运行日志</span>
            </el-menu-item>
            <el-menu-item index="/setting" @click="saveNavState('/setting')">
              <i class="el-icon-s-opportunity"></i>
              <span>扫描设置</span>
            </el-menu-item>
          </el-menu-item-group>
        </el-submenu>
      </el-menu>
    </el-aside>

    <el-container>
      <el-header>
        <el-button
          size="mini"
          class="toggle-button"
          type="primary"
          icon="el-icon-s-fold"
          @click="toggleCollapse"
        ></el-button>
        <el-button type="primary" @click="logout">退出</el-button>
      </el-header>
      <!-- 右侧主题区域 -->
      <el-main>
        <router-view>nih</router-view>
      </el-main>
    </el-container>
  </el-container>
</template>


<script>
export default {
  data() {
    return {
      menuList: ["dsa ß"],
      iconsObj: {
        "125": "iconfont icon-user",
        "103": "iconfont icon-tijikongjian",
        "101": "iconfont icon-shangpin",
        "102": "iconfont icon-danju",
        "145": "iconfont icon-baobiao"
      },
      isCollapse: false,
      isTransition: false,
      currentPath: ""
    };
  },

  created() {
    // this.getMenuList()
    this.currentPath = window.sessionStorage.getItem("currentPath");
  },

  methods: {
    logout() {
      window.sessionStorage.clear();
      this.$router.push("/login");
    },
    // async getMenuList() {
    //   // http://127.0.0.1:8000/api/v1/srclist?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRydW5rIiwicGFzc3dvcmQiOiJAcGxpbjE5OTguIiwiZXhwIjoxNTkyNTA3MTkxLCJpc3MiOiJnaW4tYmxvZyJ9.VRSpjwrNFHwrGDWqP4cDOlbxlBQJiy2vpaz3TMP0dDw
    //   //   const { data: resp } = await this.$http.get('menus')
    //   const { data: resp } = await this.$http.get('api/v1/srclist')
    //   if (resp.code !== 200) return this.$message.error(resp.meta.msg)
    //   console.log(resp)
    //   this.menuList = resp.data
    // },
    // 点击按钮折叠左侧菜单
    toggleCollapse() {
      this.isCollapse = !this.isCollapse;
    },
    saveNavState(currentPath) {
      window.sessionStorage.setItem("currentPath", currentPath);
      this.currentPath = currentPath;
    }
  }
};
</script>
<style lang="less" scoped>
.el-container {
  height: 100%;
}
.el-header {
  background-color: #fff;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #ddd;
  // color: #fff;
  .toggle-button {
    cursor: pointer;
  }
}

.el-aside {
  background-color: #2f4050;
  .el-menu {
    border-right: none;
  }
}

.el-main {
  background-color: #ffffff;
}

.iconfont {
  margin-right: 8px;
}
</style>

