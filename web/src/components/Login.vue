<template>
  <div class="login_container">
    <div class="login_box">
      <!-- 头部区域 -->
      <div class="avatar_box">
        <img src="../assets/logo.png" alt />
      </div>
      <!-- 表单区域 -->
      <el-form :model="loginForm" :rules="loginRules" ref="loginFormRef" label-width="0px" class="login_form">
        <!-- 用户名 -->
        <el-form-item prop="username">
          <el-input v-model="loginForm.username" prefix-icon="iconfont icon-user"></el-input>
        </el-form-item>

        <!-- 密码 -->
        <el-form-item prop="password">
          <!-- <el-input v-model="loginForm.password" prefix-icon="iconfont icon-3702mima" :type="pwdType"> -->
          <el-input v-model="loginForm.password" prefix-icon="iconfont icon-3702mima" type="password">
            <i slot="suffix" class="iconfont icon-showpassword" @click="showPwd"></i>
          </el-input>
        </el-form-item>

        <!-- 按钮 -->
        <el-form-item class="btns">
          <el-button type="primary" @click="login">登录</el-button>
          <el-button type="info" @click="resetLoginForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 登陆表单数据对象
      loginForm: {
        username: '',
        password: ''
      },
      loginFormRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 3, max: 10, message: '长度在 3 到 10 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 3, max: 16, message: '长度在 3 到 16 个字符', trigger: 'blur' }
        ],
        pwdType: 'password'
      }
    }
  },
  methods: {
    showPwd() {
      if (this.pwdType === 'password') {
        this.pwdType = ''
      } else {
        this.pwdType = 'password'
      }
    },

    resetLoginForm() {
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      this.$refs.loginFormRef.validate(async valid => {
        if (!valid) {
          return
        }
        const { data: resp } = await this.$http.get('/auth?username=' + this.loginForm.username + '&password=' + this.loginForm.password)
        // console.log(resp);
        if (resp.code !== 200) {
          return this.$message.error('登录失败')
        } else {
          this.$message.success('登录成功')
          window.sessionStorage.setItem('token', resp.data.token)
          this.$router.push('/home')
        }        
        // http://127.0.0.1:8000/auth?username=drunk&password=@plin1998.
        // if (resp.meta.status !== 200) {
        //   return this.$message.error('登录失败')
        // } else {
        //   // 1.将登录成功的token保存到客户端的sessionStorage中
        //   //  1.1 项目中除了登录之外的API接口，必须在登录之后才能访问
        //   //  1.2 token 只应当在当前网站打开期间生效，所以将token保存在sessionStorage中
        //   this.$message.success('登录成功')
        // }
      })
    }
  }
}
</script>

<style lang="less" scoped>
.login_container {
  background: url(../assets/img/bg.png);
  // background-color: #2b4b6b;
  height: 100%;
}

.login_box {
  width: 450px;
  height: 300px;
  background-color: #fff;
  border-radius: 10px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.avatar_box {
  width: 130px;
  height: 130px;
  border-radius: 50%;
  border: 1px solid #eee;
  padding: 10px;
  box-shadow: 0 0 10px #ddd;
  position: absolute;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: #fff;
  img {
    width: 100%;
    height: 100%;
    border-radius: 50%;
    background-color: #eee;
  }
}

.login_form {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
  .icon-showpassword {
    margin-right: 8px;
  }
}

.btns {
  display: flex;
  justify-content: flex-end;
}
</style>
