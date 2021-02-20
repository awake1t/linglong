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

                <el-form-item label="用户名" prop="username">
                    <el-col :span="10">
                        <el-input v-model="addForm.username"></el-input>
                    </el-col>
                </el-form-item>

                <el-form-item label="旧密码" prop="oldpass">
                    <el-col :span="10">
                        <el-input type="password" v-model="addForm.oldpass"></el-input>
                    </el-col>
                </el-form-item>

                <el-form-item label="新密码" prop="newpass">
                        <el-col :span="10">
                            <el-input type="password" v-model="addForm.newpass"></el-input>
                        </el-col>
                </el-form-item>
                <el-form-item label="确认密码" prop="newpass2">
                        <el-col :span="10">
                            <el-input type="password" v-model="addForm.newpass2"></el-input>
                        </el-col>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click="editListInfo">更新</el-button>
                </el-form-item>
            </el-form>


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
        },
        methods: {
            editListInfo() {
                this.$refs.addFormRef.validate(async valid => {
                    if (!valid) return;
                    const {data: resp} = await this.$http.put(
                        `/modpass`,
                        qs.stringify(this.addForm),
                        {headers: {"Content-Type": "application/x-www-form-urlencoded"}}
                    );
                    if (resp.code !== 200) {
                        return this.$message.error(respeta.msg);
                    }
                    if (resp.code == 402) {
                        return this.$message.error(respeta.msg);
                    }
                    if (resp.code == 403) {
                        return this.$message.error(respeta.msg);
                    }
                    this.$message.success("修改成功");
                });
            }
        }
    };
</script>

<style lang="less" scoped>
</style>
