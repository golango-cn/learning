import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'

import axios from 'axios'
// 导入样式表
import './assets/css/global.css'
import 'element-ui/lib/theme-chalk/index.css'

import TreeTable from 'vue-table-with-tree-grid'

Vue.config.productionTip = false

axios.defaults.baseURL = '/api'
axios.interceptors.request.use(config => {
  config.headers.Authorization = window.sessionStorage.getItem('token')
  return config
})

Vue.prototype.$http = axios

Vue.component('tree-table', TreeTable)

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
