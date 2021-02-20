import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Home from '../components/Home.vue'
import Welcome from '../components/Welcome.vue'
import Jobips from '../components/jobips/Jobips.vue'
import Task from '../components/task/Task.vue'
import Tasklog from '../components/task/Tasklog.vue'
import Webloginlist from '../components/webloginlist/Webloginlist.vue'
import Log from '../components/log/Log.vue'
import Setting from '../components/setting/Setting.vue'
import Finger from '../components/finger/finger.vue'
import Modpass from '../components/modpass/Modpass.vue'
import Xrayres from '../components/xrayres/Xrayres.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    component: Login
  },
  {
    path: '/home',
    component: Home,
    redirect: '/welcome',
    children: [{ path: '/welcome', component: Welcome },
    {  path: '/jobips', component: Jobips  },
    {  path: '/task', component: Task  },
    {  path: '/webloginlist', component: Webloginlist  },
    { path: '/tasklog/:id', component: Tasklog },
    { path: '/log', component: Log },
    { path: '/setting', component: Setting },
    { path: '/finger', component: Finger },
    { path: '/modpass', component: Modpass },
    { path: '/xrayres', component: Xrayres },

    ]
  }
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]
const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  if (to.path === '/login') return next();
  const tokenStr = window.sessionStorage.getItem('token');
  if (!tokenStr) return next('/login');
  next();
});

// 解决ElementUI导航栏中的vue-router在3.0版本以上重复点菜单报错问题
const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}


export default router
