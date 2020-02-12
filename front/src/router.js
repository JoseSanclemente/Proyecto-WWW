import Vue from 'vue'
import Router from 'vue-router'
import { CLIENT_ROUTES, ADMIN_ROUTES, GUEST_ROUTES } from './routes.js'

Vue.use(Router)

export default new Router({
  routes: [
    ...CLIENT_ROUTES,
    ...ADMIN_ROUTES,
    ...GUEST_ROUTES
  ]
})
