import Vue from 'vue'
import VueRouter from 'vue-router'

// import Login from '../components/Login.vue'
// import Home from '../components/Home.vue'
// import Welcome from '../components/Welcome.vue'

const Login = () => import(/* webpackChunkName: "login-home-welcome" */ '../components/Login.vue')
const Home = () => import(/* webpackChunkName: "login-home-welcome" */ '../components/Home.vue')
const Welcome = () => import(/* webpackChunkName: "login-home-welcome" */ '../components/Welcome.vue')

// import Users from '../components/user/Users.vue'
// import Rights from '../components/power/Rights.vue'
// import Role from '../components/role/Roles'

const Users = () => import(/* webpackChunkName: "user-right-role" */ '../components/user/Users.vue')
const Rights = () => import(/* webpackChunkName: "user-right-role" */ '../components/power/Rights.vue')
const Role = () => import(/* webpackChunkName: "user-right-role" */ '../components/role/Roles.vue')

// import Cates from '../components/good/Cates'
// import Params from '../components/good/Params'
// import Goods from '../components/good/Goods'
// import EditGoods from '../components/good/EditGoods'

const Cates = () => import(/* webpackChunkName: "good-cate-param" */ '../components/good/Cates.vue')
const Params = () => import(/* webpackChunkName: "good-cate-param" */ '../components/good/Params.vue')
const Goods = () => import(/* webpackChunkName: "good-cate-param" */ '../components/good/Goods.vue')
const EditGoods = () => import(/* webpackChunkName: "good-cate-param" */ '../components/good/EditGoods.vue')

// import Common from '../components/common/Common'
// import Reports from '../components/common/Reports'

const Common = () => import(/* webpackChunkName: "common-report" */ '../components/common/Common.vue')
const Reports = () => import(/* webpackChunkName: "common-report" */ '../components/common/Reports.vue')

Vue.use(VueRouter)

const routes = [
  { path: '/login', name: 'Login', component: Login },
  { path: '/', redirect: '/login' },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    redirect: '/welcome',
    children: [
      { path: '/welcome', component: Welcome },
      { path: '/users', component: Users },
      { path: '/rights', component: Rights },
      { path: '/roles', component: Role },
      { path: '/cates', component: Cates },
      { path: '/params', component: Params },
      { path: '/goods', component: Goods },
      { path: '/goods/edit', component: EditGoods },
      { path: '/common', component: Common },
      { path: '/reports', component: Reports }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  if (to.path === '/login') {
    return next()
  }

  const item = sessionStorage.getItem('data')
  if (!item) {
    return next('/login')
  }
  return next()

})

export default router
