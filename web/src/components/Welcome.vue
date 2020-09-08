<template>
  <el-row :gutter="40" class="panel-group">
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('newVisitis')">
        <div class="card-panel-icon-wrapper icon-message">
          <img src="../assets/all.png" alt style="width:75px" />
        </div>
        <div class="card-panel-description">
          <count-to :start-val="0" :end-val="iplist" :duration="3000" class="card-panel-num" />
          <div class="card-panel-text">资产总数</div>
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('messages')">
        <div class="card-panel-icon-wrapper icon-message">
          <img src="../assets/ipall.png" alt style="width:75px" />
        </div>
        <div class="card-panel-description">
          <count-to :start-val="0" :end-val="ip" :duration="3000" class="card-panel-num" />
          <div class="card-panel-text">IP总数</div>
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('purchases')">
        <div class="card-panel-icon-wrapper icon-message">
          <img src="../assets/login.png" alt style="width:75px" />
        </div>
        <div class="card-panel-description">
          <count-to :start-val="0" :end-val="weblogin" :duration="3200" class="card-panel-num" />
          <div class="card-panel-text">管理后台</div>
        </div>
      </div>
    </el-col>
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('shoppings')">
        <div class="card-panel-icon-wrapper icon-message">
          <img src="../assets/vuln.png" alt style="width:75px" />
        </div>
        <div class="card-panel-description">
          <count-to :start-val="0" :end-val="vuln" :duration="3200" class="card-panel-num" />
          <div class="card-panel-text">漏洞</div>
        </div>
      </div>
    </el-col>

    <el-col :xs="12" :sm="12" :lg="8">
      <div class="Echarts">
        <div id="main" style="width: 500px;height:300px;"></div>
      </div>
    </el-col>

    <el-col :xs="12" :sm="12" :lg="8" style="width: 500px;height:300px;">
      <div class="Echarts">
        <div id="main1" style="width: 500px;height:300px;"></div>
      </div>
    </el-col>

        <el-col :xs="12" :sm="12" :lg="6" style="position: absolute;right: -7px;top: 148px;">
      <el-card class="box-card" >
        <div slot="header" class="clearfix">
          <span>系统动态</span>
          <!-- <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button> -->
        </div>
        <!-- <div v-for="o in 4" :key="o" class="text item">{{'新增端口 ' + o }}</div> -->
        <div style="margin-left: -45px;">
          <el-timeline>
            <el-timeline-item
              v-for="(activity, index) in activities"
              :key="index"
              :icon="activity.icon"
              :timestamp="activity.timestamp"
            >{{activity.content}}</el-timeline-item>
          </el-timeline>
        </div>
      </el-card>
    </el-col>

    <el-col :xs="12" :sm="12" :lg="6" style="width: 900px;height:300px">
      <div class="Echarts">
        <div id="main2" style="width: 900px;height:300px;"></div>
      </div>
    </el-col>

    <el-col :xs="12" :sm="12" :lg="6" style="width: 900px;height:300px">
      <div class="Echarts">
        <div id="main3" style="width: 900px;height:300px;"></div>
      </div>
    </el-col>
  </el-row>
</template>


<script>
import CountTo from "vue-count-to";

export default {
  data() {
    return {
      iplist: 0,
      ip: 0,
      weblogin: 0,
      vuln: 0,
      test: 199,
      vulnratio: [{}],
      option: {
        title: {
          text: "漏洞占比"
        },
        color:['#3F9EFF', '#53CB75','#FBD43F','#2EC7C9','#B6A2DE','#D87A80','#E7FAF0','#FFEDED'],
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)"
        },
        series: {
          name: "漏洞数量",
          type: "pie",
          radius: "85%",
          data: [
                    {name: 'MYSQL', value: 1212},
                ]
        }
      },

      option1: {
        title: {
          text: "7日资产新增"
        },
        color: ["#53CB75"],
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "shadow" // 默认为直线，可选为：'line' | 'shadow'
          }
        },
        
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true
        },
        // toolbox: { 
        //   feature: {
        //     saveAsImage: {}
        //   }
        // },
        xAxis: {
          type: "category",
          boundaryGap: false,
          data: ["周一", "周二", "周三", "周四", "周五", "周六", "周日"],
          axisTick: {
              alignWithLabel: true
            },
            axisLabel: {
              interval: 0,
              rotate: 45
            }
        },
        yAxis: {
          type: "value"
        },
        series: [
          {
            name: "新增资产",
            type: "line",
            stack: "总量",
            data: [120, 132, 101, 134, 90, 230, 210]
          }
        ]
      },

      option2: {
        title: {
          text: "端口占比"
        },
        color: ["#3F9EFF"],
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "shadow" // 默认为直线，可选为：'line' | 'shadow'
          }
        },
        xAxis: [
          {
            type: "category",
            data: [{ name: 0, value: "Mon" }],
            axisTick: {
              alignWithLabel: true
            },
            axisLabel: {
              interval: 0
            }
          }
        ],
        yAxis: [
          {
            type: "value"
          }
        ],
        series: [
          {
            name: "端口数量",
            type: "bar",
            barWidth: "60%",
            data: [10, 52, 200, 334, 390, 330, 220]
          }
        ]
      },
      option3: {
        title: {
          text: "服务占比"
        },
        color:['#3F9EFF'],
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "shadow" // 默认为直线，可选为：'line' | 'shadow'
          }
        },
        xAxis: [
          {
            type: "category",
            data: [{ name: 0, value: "Mon" }],
            axisTick: {
              alignWithLabel: true
            },
            axisLabel: {
              interval: 0,
              rotate: 45
            }
          }
        ],
        yAxis: [
          {
            type: "value"
          }
        ],
        series: [
          {
            name: "数量",
            type: "bar",
            barWidth: "60%",
            data: [10, 52, 200, 334, 390, 330, 220]
          }
        ]
      },

      activities: [
        {
          content: "新增漏洞 10.123.154.121",
          timestamp: "2018-04-12 20:46",
          color: "#0bbd87"
        }
      ]
    };
  },
  components: {
    CountTo
  },

  name: "Echarts",
  methods: {
    myEcharts() {
      // 基于准备好的dom，初始化echarts实例
      // var myChart2 = this.$echarts.init(document.getElementById("main2"));
      // 指定图表的配置项和数据
    },

    async getIpList() {
      try {
        const { data: res } = await this.$http.get("/dashboard", {});
        if (res.code !== 200) {
          return this.$message.error("获取列表失败");
        }
        this.iplist = res.data.iplist;
        this.ip = res.data.ip;
        this.weblogin = res.data.weblogin;
        this.vuln = res.data.vuln;
      } catch (error) {
        return this.$message.error("cookie失效，请点击右上角退出重新登陆");
        error.message; // "Oops!"
      }
    }
  },

  handleSetLineChartData(type) {
    this.$emit("handleSetLineChartData", type);
  },

  async mounted() {
    var myChart = this.$echarts.init(document.getElementById("main"));
    var myChart1 = this.$echarts.init(document.getElementById("main1"));
    var myChart2 = this.$echarts.init(document.getElementById("main2"));
    var myChart3 = this.$echarts.init(document.getElementById("main3"));

    const { data: res } = await this.$http.get("/dashboard", {});
    if (res.code !== 200) {
      return this.$message.error("获取列表失败");
    }
    //系统动态
    this.activities = res.data.dynamics;
    // if (res.data.vulnratio == null){
    //   console.log('kong')
    // }else{
    //   console.log('yesyes')
    // }

    //饼图数据
    this.option.series.data = res.data.vulnratio;
    myChart.setOption(this.option);
    myChart1.setOption(this.option1);
    myChart2.setOption(this.option2);

  //  近7天新增数据
    this.option1.xAxis.data = res.data.timelinex;
    this.option1.series[0].data = res.data.timeliney;

    this.option2.xAxis[0].data = res.data.portlinex;
    this.option2.series[0].data = res.data.portliney;

    this.option3.xAxis[0].data = res.data.protocolinex;
    this.option3.series[0].data = res.data.protocoliney;

    myChart1.setOption(this.option1);
    myChart2.setOption(this.option2);
    myChart3.setOption(this.option3);
  },
  created() {
    this.getIpList();
  }
};
</script>
<style lang="scss" scoped>
.panel-group {
  margin-top: 18px;

  .card-panel-col {
    margin-bottom: 32px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);

    .card-panel-icon-wrapper {
      color: #fff;
    }

    .icon-people {
      background: #40c9c6;
    }

    .icon-money {
      background: #f4516c;
    }

    .icon-shopping {
      background: #34bfa3;
    }

    .icon-people {
      color: #40c9c6;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3;
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 25px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
      width: 50px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px;
      margin-left: 0px;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
      }
    }
  }
}

@media (max-width: 550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    .svg-icon {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}

//卡片样式
.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}

.box-card {
  width: 265px;
  height: 900px;
}
</style>
