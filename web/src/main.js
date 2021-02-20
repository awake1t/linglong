import Vue from 'vue'
import App from './App.vue'
import router from './router'

import './plugins/element.js'
// 导入字体图标
import './assets/fonts/iconfont.css'
// 全局样式表
import './assets/css/global.css'

import echarts from 'echarts'
Vue.prototype.$echarts = echarts

import ElementUI from 'element-ui' //element-ui的全部组件
import 'element-ui/lib/theme-chalk/index.css'//element-ui的css
Vue.use(ElementUI) //使用elementUI

import moment from 'moment'
Vue.filter('dateformat', function(dataStr, pattern = 'YYYY-MM-DD HH:mm:ss') {
  return moment(dataStr).format(pattern)
})
// // 导入进度条
// import NProgress from 'nprogress'
// import 'nprogress/nprogress.css'


import axios from 'axios'
import serverConfig from '../server-config'

axios.defaults.baseURL = process.env.VUE_APP_BASE_API,
// axios.defaults.baseURL = 'http://127.0.0.1:18000/api/v1/'
Vue.prototype.$http = axios



// 新方式
axios.interceptors.request.use(
  config => {
    config.headers.Authorization = window.sessionStorage.getItem('token');
    return config;
},
error => {
  console.log('nonon')
    return Promise.error(error);
});

// 默认方式
// axios.interceptors.request.use(config => {
//   // NProgress.start();
//   config.headers.Authorization = window.sessionStorage.getItem('token');
//   // console.log(config)
//   // config.headers['Content-Type'] = 'application/x-www-form-urlencoded';
//   // Content-Type: ; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW
//   // if (config.method === 'post' || config.method === 'put') {

//   // }
//   return config;
// })


Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
