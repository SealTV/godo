// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import VueI18n from 'vue-i18n'
import vueConfig from 'vue-config'
import VueCookie from 'vue-cookie'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueAuth from '@websanova/vue-auth'
import Bootstrap from 'bootstrap-vue'

import App from './App'
import router from './router'

import messages from './i18n/messages'
import constants from './assets/js/constants'

const configs = {
  API: constants.API_SERVER_HOST
}

Vue.use(VueI18n)
Vue.use(vueConfig, configs)
Vue.use(VueCookie)
Vue.use(VueAxios, axios)
Vue.use(Bootstrap)

Vue.config.productionTip = false
Vue.axios.defaults.baseURL = configs.API

Vue.router = router

Vue.use(VueAuth, {
  auth: {
    request: function (req, token) {
      req.headers['Authorization'] = 'Bearer ' + token
    },
    response: function (res) {
      if (res.data.code === 200) {
        return res.data.token
      }
    }
  },
  http: require('@websanova/vue-auth/drivers/http/axios.1.x.js'),
  router: require('@websanova/vue-auth/drivers/router/vue-router.2.x.js'),
  parseUserData: function (data) {
    console.log('USER DATA: ' + JSON.stringify(data))
    return data.token
  },
  fetchData: { url: 'auth/user', method: 'GET', authType: 'bearer' }
})

const i18n = new VueI18n({
  locale: 'ru',
  messages
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  i18n
})
