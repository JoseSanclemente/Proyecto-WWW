import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'
Vue.use(VueMaterial)


import * as VueGoogleMaps from "vue2-google-maps";
import i18n from './i18n'
Vue.use(VueGoogleMaps, {
  load: { key: "AIzaSyBbAIAW4XQ_nJWC43SBDVAQVI4uDewyVPY" }
});

Vue.config.productionTip = false

new Vue({
  router,
  store,
  i18n,
  render: h => h(App)
}).$mount('#app')
