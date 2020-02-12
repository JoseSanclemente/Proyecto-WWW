import Home from './views/Home.vue'

// TODO: define better routes

const CLIENT_ROUTES = [
  {
    path: '/login-client',
    name: 'login-client',
    component: () => import(/* webpackChunkName: "login-client" */ './views/Login.vue')
  },
  {
    path: '/user-profile',
    name: 'user-profile',
    component: () => import(/* webpackChunkName: "user-profile" */ './views/client/ClientProfile.vue')
  },
  {
    path: '/user',
    name: 'user-home',
    component: () => import(/* webpackChunkName: "user-home" */ './views/client/ClientHome.vue')
  },
  {
    path: '/category/:category_id/subcategory/:subcategory_id/product/:product_id',
    name: 'product',
    component: () => import(/* webpackChunkName: "subcategory-products" */ './views/client/Product.vue')
  },
  {
    path: '/category/:category_id/subcategory/:subcategory_id',
    name: 'subcategory-products',
    component: () => import(/* webpackChunkName: "subcategory-products" */ './views/client/ClientSubcategoryProducts.vue')
  },
  {
    path: '/category/:category_id/',
    name: 'category-subcategories',
    component: () => import(/* webpackChunkName: "subcategory-products" */ './views/client/ClientCategorySubcategories.vue')
  },
  {
    path: '/checkout',
    name: 'checkout',
    component: () => import(/* webpackChunkName: "checkout" */ './views/client/Checkout.vue')
  },
  {
    path: '/history',
    name: 'hsitory',
    component: () => import(/* webpackChunkName: "history" */ './views/client/History.vue')
  },
  {
    path: '/invoice/:invoice_id/',
    name: 'invoice',
    component: () => import(/* webpackChunkName: "invoice" */ './views/client/Invoice.vue')
  }
]

const ADMIN_ROUTES = [
  {
    path: '/admin/create-admin',
    name: 'create-admin',
    component: () => import(/* webpackChunkName: "create-admin" */ './views/admin/CreateAdmin.vue')
  },
  {
    path: '/admin',
    name: 'admin-home',
    component: () => import(/* webpackChunkName: "admin-home" */ './views/admin/AdminHome.vue')
  },
  {
    path: '/admin-profile',
    name: 'admin-profile',
    component: () => import(/* webpackChunkName: "admin-profile" */ './views/admin/AdminProfile.vue')
  },
  {
    path: '/login-admin',
    name: 'login-admin',
    component: () => import(/* webpackChunkName: "login-admin" */ './views/Login.vue')
  },
  {
    path: '/admin/search-user',
    name: 'search-user',
    component: () => import(/* webpackChunkName: "search-user" */ './views/admin/SearchUser.vue')
  },
  {
    path: '/admin/add-category',
    name: 'add-category',
    component: () => import(/* webpackChunkName: "add-category" */ './views/admin/AddCategory.vue')
  },
  {
    path: '/admin/search-category',
    name: 'search-category',
    component: () => import(/* webpackChunkName: "search-category" */ './views/admin/SearchCategory.vue')
  },
  {
    path: '/admin/search-subcategory',
    name: 'search-subcategory',
    component: () => import(/* webpackChunkName: "search-category" */ './views/admin/SearchSubCategory.vue')
  },
  {
    path: '/admin/add-products',
    name: 'add-products',
    component: () => import(/* webpackChunkName: "add-products" */ './views/admin/AddProduct.vue')
  },
  {
    path: '/admin/search-products',
    name: 'search-products',
    component: () => import(/* webpackChunkName: "search-products" */ './views/admin/SearchProduct.vue')
  },
  {
    path: '/admin/add-subcategory',
    name: 'add-subcategory',
    component: () => import(/* webpackChunkName: "add-subcategoty" */ './views/admin/AddSubcategory.vue')
  },
  {
    path: '/admin/add-discount',
    name: 'add-discount',
    component: () => import(/* webpackChunkName: "add-discount" */ './views/admin/AddDiscount.vue')
  },
  {
    path: '/admin/search-discount',
    name: 'search-discount',
    component: () => import(/* webpackChunkName: "search-subcategoty" */ './views/admin/SearchDiscount.vue')
  },
  {
    path: '/admin/report',
    name: 'report',
    component: () => import(/* webpackChunkName: "report" */ './views/admin/Report.vue')
  }

]

const GUEST_ROUTES = [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/about',
    name: 'about',
    component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
  },
  {
    path: '/sign-up',
    name: 'sign-up',
    component: () => import(/* webpackChunkName: "sign-up" */ './views/client/SignUp.vue')
  }
]

function isAdminRoute (name) {
  for (let i = 0; i < ADMIN_ROUTES.length; i++) {
    if (ADMIN_ROUTES[i].name === name) {
      return true
    }
  }
  return false
}

function isClientRoute (name) {
  for (let i = 0; i < CLIENT_ROUTES.length; i++) {
    if (CLIENT_ROUTES[i].name === name) {
      return true
    }
  }
  return false
}

function isGuestRoute (name) {
  for (let i = 0; i < GUEST_ROUTES.length; i++) {
    if (GUEST_ROUTES[i].name === name) {
      return true
    }
  }
  return false
}

export { CLIENT_ROUTES, ADMIN_ROUTES, GUEST_ROUTES, isAdminRoute, isClientRoute, isGuestRoute }
