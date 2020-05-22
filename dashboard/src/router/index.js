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
    path: "/employees/login",
    component: () => import(/* webpackChunkName: "employees" */ '../views/UserLogin.vue'),
  },
  {
    path: "/consumer/contracts",
    component: () => import(/* webpackChunkName: "contracts" */ '../views/Consumer/Contracts.vue'),
  },
  {
    path: "/consumer/profile",
    component: () => import(/* webpackChunkName: "contracts" */ '../views/Consumer/Profile.vue'),
  },
  {
    path: "/manager/dashboard",
    component: () => import(/* webpackChunkName: "manager" */ '../views/Manager/Manager.vue'),
  },
  {
    path: "/operator/dashboard",
    component: () => import(/* webpackChunkName: "operator" */ '../views/Operator/Operator.vue'),
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
