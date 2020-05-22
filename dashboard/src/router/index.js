import Vue from 'vue'
import VueRouter from 'vue-router'


Vue.use(VueRouter)

const routes = [
  {
    path: "/",
    redirect: "/login"
  },
  {
    path: "/home",
    component: () => import(/* webpackChunkName: "landing" */ '../views/Landing.vue'),
  },
  {
    path: "/employee/login",
    component: () => import(/* webpackChunkName: "employees" */ '../views/UserLogin.vue'),
  },
  {
    name: "consumer-contracts",
    path: "/consumer/contracts/:id",
    component: () => import(/* webpackChunkName: "contracts" */ '../views/Consumer/Contracts.vue'),
  },
  {

    name: "consumer-profile",
    path: "/consumer/profile/:id",
    component: () => import(/* webpackChunkName: "contracts" */ '../views/Consumer/Profile.vue'),
  },
  {
    path: "/manager/reports",
    component: () => import(/* webpackChunkName: "manager" */ '../views/Manager/Reports.vue'),
  },
  {
    path: "/manager/users",
    component: () => import(/* webpackChunkName: "manager" */ '../views/Manager/User.vue'),
  },
  {
    path: "/operator/payment",
    component: () => import(/* webpackChunkName: "operator" */ '../views/Operator/Operator.vue'),
  },
  {
    path: "/operator/consumers",
    component: () => import(/* webpackChunkName: "operator" */ '../views/Operator/Consumer.vue'),
  },
  {
    path: "/admin",
    redirect: "/admin/users",
    component: () => import(/* webpackChunkName: "default" */ '../layout/DefaultLayout.vue'),
    children: [
      {
        path: 'users',
        name: 'Users',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Admin/Users.vue')
      },
      {
        path: 'consumer',
        name: 'Consumer',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Admin/Consumer.vue')
      },
      {
        path: 'substations',
        name: 'Substations',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Admin/Substations.vue')
      },
      {
        path: 'transformers',
        name: 'Transformers',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Admin/Transformer.vue')
      },
      {
        path: 'map',
        name: 'Map',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Map.vue')
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
