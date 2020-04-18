import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Home from '../components/Home.vue'
import Welcome from '../components/Welcome.vue'
import Users from '../components/user/Users.vue'
import Rights from '../components/power/Rights.vue'
import Role from '../components/role/Roles'
import Cates from '../components/good/Cates'
import Params from '../components/good/Params'


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
      { path: '/params', component: Params }
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
