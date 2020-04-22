import Vue from 'vue'
import App from './App.vue'
import router from './router'
// import './plugins/element.js'

import axios from 'axios'
// 导入样式表
import './assets/css/global.css'
// import 'element-ui/lib/theme-chalk/index.css'

import TreeTable from 'vue-table-with-tree-grid'

import Nprogress from 'nprogress'
// import 'nprogress/nprogress.css'

Vue.config.productionTip = false

axios.defaults.baseURL = 'http://vue.golango.cn/api'
axios.interceptors.request.use(config => {
  Nprogress.start()
  config.headers.Authorization = window.sessionStorage.getItem('token')
  return config
})

axios.interceptors.response.use(config => {
  Nprogress.done()
  return config
})

Vue.prototype.$http = axios

Vue.component('tree-table', TreeTable)
Vue.filter('dateFromat', function(timestamp) {

  const dt = new Date(timestamp * 1000)

  const y = dt.getFullYear()
  const m = ((dt.getMonth() + 1) + '').padStart(2, '0')
  const d = ((dt.getDay()) + '').padStart(2, '0')

  const hh = ((dt.getHours()) + '').padStart(2, '0')
  const mm = ((dt.getMinutes()) + '').padStart(2, '0')
  const ss = ((dt.getSeconds()) + '').padStart(2, '0')

  return `${y}-${m}-${d} ${hh}:${mm}:${ss}`

})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
