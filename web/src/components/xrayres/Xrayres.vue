<template>
    <div>
        <!-- 面包屑导航 -->
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>XrayPoc扫描结果</el-breadcrumb-item>
        </el-breadcrumb>
        <el-card>
            <!-- 搜索 添加区域 -->
            <el-row :gutter="30">
                <el-col :span="3">
                    <el-input
                            placeholder="漏洞url"
                            class="input-with-select"
                            v-model="queryInfo.url"
                            clearable
                            @clear="getIpList"
                            @keyup.enter.native="getIpListSingle"
                    ></el-input>
                </el-col>
                <el-col :span="3">
                    <el-input
                            placeholder="poc"
                            class="input-with-select"
                            v-model="queryInfo.poc"
                            clearable
                            @clear="getIpList"
                            @keyup.enter.native="getIpListSingle"
                    ></el-input>
                </el-col>

<!--                <el-col :span="3">-->
<!--                    <el-input-->
<!--                            placeholder="detail"-->
<!--                            class="input-with-select"-->
<!--                            v-model="queryInfo.snapshot"-->
<!--                            clearable-->
<!--                            @clear="getIpList"-->
<!--                            @keyup.enter.native="getIpListSingle"-->
<!--                    >-->
<!--                    </el-input>-->
<!--                </el-col>-->

                <el-col :span="8">
                    <el-button type="primary" @click="getIpListSingle">搜索</el-button>
                </el-col>
            </el-row>

            <!-- 资产列表区 -->
            <el-table :data="iplist" border stripe>
                <el-table-column label="url" prop="url">
                    <template slot-scope="scope">
                        <a
                                :href="scope.row.url"
                                target="_blank"
                                prop
                                style="color:#606266;"
                        >{{ scope.row.url }}</a>
                    </template>
                </el-table-column>

                <el-table-column label="poc" prop="poc"></el-table-column>

                <el-table-column
                        label="创建时间"
                        prop="created_time"
                >{{ created_at | dateformat('YYYY-MM-DD HH:mm:ss')}}</el-table-column>

                <el-table-column label="操作" width="180px">
                    <template slot-scope="scope">

                        <el-tooltip effect="dark" content="删除结果" placement="top" >
                            <el-button
                                    type="danger"
                                    icon="el-icon-delete"
                                    size="mini"
                                    @click="deleteById(scope.row.id)"
                            ></el-button>
                        </el-tooltip>

                    </template>
                </el-table-column>

<!--                <el-table-column label="snapshot" prop="snapshot"></el-table-column>-->
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
                // 参数列表
                queryInfo: {
                    url: "",
                    poc: "",
                    snapshot: "",
                    pagenum: 1,
                    // 当前每页显示多少条数据
                    pagesize: 10
                },
                iplist: [],
                total: 0,

            };
        },
        created() {
            this.getIpList();
        },
        methods: {
            async getIpList() {
                try {
                    const {data: res} = await this.$http.get("/gerXrayRes", {
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
                const {data: res} = await this.$http.get("/gerXrayRes", {
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

            //监听对话框关闭事件
            diglogClose(operation) {
                if (operation === "add") {
                    this.$refs.addFormRef.resetFields();
                } else if (operation === "edit") {
                    this.$refs.editFormRef.resetFields();
                }
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
                const { data: resp } = await this.$http.delete("/delxrayres/" + id);
                if (resp.code !== 200) {
                    return this.$message.error(resp.msg);
                }
                this.$message.success("删除成功");
                this.getIpList();
            },

            editListInfo() {
                this.$refs.editFormRef.validate(async valid => {
                    if (!valid) return;
                    const {
                        data: resp
                    } = await this.$http.put(
                        `/jobiplist/${this.editForm.id}`,
                        qs.stringify(this.editForm),
                        {headers: {"Content-Type": "application/x-www-form-urlencoded"}}
                    );
                    if (resp.code !== 200) {
                        return this.$message.error(respeta.msg);
                    }
                    this.$refs.editFormRef.resetFields();
                    this.editDiglogVisable = false;
                    this.$message.success("修改成功");
                    this.getIpList();
                });
            }

        }
    };
</script>

<style lang="less" scoped>
    a {
        text-decoration: none
    }
</style>
