// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import VueI18n from 'vue-i18n'
import vueConfig from 'vue-config'
import App from './App'
import router from './router'

import messages from './i18n/messages'
import constants from './assets/js/constants'

const configs = {
  API: constants.API_SERVER_HOST
}

Vue.use(VueI18n)
Vue.use(vueConfig, configs)

Vue.config.productionTip = false

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
