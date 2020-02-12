import axios from 'axios'
import { helpers } from './helpers.js'

/* API urls */
const BASE_URL = 'http://localhost:8080'
const CLIENT_PATH = BASE_URL + '/user'
const CLIENT_LOGIN_PATH = BASE_URL + '/login/user'
const CLIENT_DELETE_PATH = CLIENT_PATH + '/delete'
const ADMIN_PATH = BASE_URL + '/admin'
const ADMIN_DELETE_PATH = ADMIN_PATH + '/delete'
const ADMIN_LOGIN_PATH = BASE_URL + '/login/admin'
const CATEGORY_PATH = BASE_URL + '/category'
const CATEGORY_DELETE_PATH = CATEGORY_PATH + '/delete'
const PRODUCT_PATH = BASE_URL + '/product'
const PRODUCT_DELETE_PATH = PRODUCT_PATH + '/delete'
const SUBCATEGORY_PATH = BASE_URL + '/subcategory'
const SUBCATEGORY_DELETE_PATH = SUBCATEGORY_PATH + '/delete'
const DISCOUNT_PATH = BASE_URL + '/discount'
const DISCOUNT_DELETE_PATH = DISCOUNT_PATH + '/delete'
const CART_PATH = BASE_URL + '/cart'
const CART_DELETE_PATH = CART_PATH + '/delete'
const CARD_PATH = BASE_URL + '/card'
const PURCHASE_PATH = BASE_URL + '/purchase'
const LIST_CARDS_PATH = BASE_URL + '/card/list'
const DISCOUNTS_PATH = BASE_URL + '/discount'
const REPORT_PATH = BASE_URL + '/report'
const HISTORY_PATH = BASE_URL + '/invoice/list'
const INVOICE_PATH = BASE_URL + '/invoice'
const DISCOUNTS_USER_PATH = BASE_URL + '/discount/user'

/* API token */
let token = ''

let iva = 19.0

let userEmail = ''

/* API methods */
function logInClient (email, password) {
  return axios
    .post(
      CLIENT_LOGIN_PATH,
      { email: email, password: password },
      helpers.generateConfig()
    )
    .then(function (response) {
      token = response.data.token
      userEmail = email
    })
}

function logInAdmin (email, password) {
  return axios
    .post(
      ADMIN_LOGIN_PATH,
      { email: email, password: password },
      helpers.generateConfig()
    )
    .then(function (response) {
      token = response.data.token
      userEmail = email
    })
}

function signUpClient (email, name, birthDate, address, password, id, type) {
  return axios.post(
    CLIENT_PATH,
    {
      email: email,
      name: name,
      birth_date: birthDate,
      address: address,
      password: password,
      document_id: id,
      document_type: type
    },
    helpers.generateConfig()
  )
}

function createAdmin (email, name, password) {
  return axios.post(
    ADMIN_PATH,
    {
      email: email,
      name: name,
      password: password
    },
    helpers.generateConfig()
  )
}

function updatePassword (email, password) {
  return axios.put(
    CLIENT_PATH,
    {
      email: email,
      password: password
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function updateAdminPassword (email, password) {
  return axios.put(
    ADMIN_PATH,
    {
      email: email,
      password: password
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function updateClient (email, name, address) {
  return axios.put(
    CLIENT_PATH,
    {
      email: email,
      name: name,
      address: address
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function updateAdmin (email, name) {
  return axios.put(
    ADMIN_PATH,
    {
      email: email,
      name: name
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getAdminProfile () {
  return axios.get(
    ADMIN_PATH,
    {
      params: {
        email: userEmail
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getClientProfile () {
  return axios.get(
    CLIENT_PATH,
    {
      params: {
        email: userEmail
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getClientUser (clientEmail) {
  return axios.get(
    CLIENT_PATH,
    {
      params: {
        email: clientEmail
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getAdminUser (adminEmail) {
  return axios.get(
    ADMIN_PATH,
    {
      params: {
        email: adminEmail
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function deleteAdminUser (email) {
  return axios.post(
    ADMIN_DELETE_PATH,
    {
      email: email
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function deleteClientUser (email) {
  return axios.post(
    CLIENT_DELETE_PATH,
    {
      email: email
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function createCategory (name, description, image) {
  return axios.post(
    CATEGORY_PATH,
    {
      name: name,
      description: description,
      image: image
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function updateCategory (id, name, description, image) {
  return axios.put(
    CATEGORY_PATH,
    {
      category_id: id,
      name: name,
      description: description,
      image: image
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getCategoriesList () {
  return axios.get(
    CATEGORY_PATH,
    {},
    helpers.generateConfig({ Authorizer: token })
  )
}

function deleteCategory (categoryID) {
  return axios.post(
    CATEGORY_DELETE_PATH,
    {
      category_id: categoryID
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getSubcategoriesList (categoryID) {
  return axios.get(
    SUBCATEGORY_PATH,
    {
      params: {
        category_id: categoryID
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function createSubcategory (payload) {
  return axios.post(
    SUBCATEGORY_PATH,
    {
      category_id: payload.categoryID,
      name: payload.name,
      image: payload.image
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function deleteSubCategory (subcategoryID) {
  return axios.post(
    SUBCATEGORY_DELETE_PATH,
    {
      id: subcategoryID
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function updateSubcategory (id, category, name, image) {
  return axios.put(
    SUBCATEGORY_PATH,
    {
      id: id,
      category_id: category,
      name: name,
      image: image
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function createProduct (payload) {
  return axios.post(PRODUCT_PATH, {
    subcategory_id: payload.subcategoryID,
    name: payload.name,
    price: payload.price,
    description: payload.description,
    image: payload.image,
    quantity: parseInt(payload.quantity),
    iva: iva
  })
}

function getSubcategoryProductsList (subcategoryID) {
  return axios.get(
    PRODUCT_PATH,
    {
      params: {
        subcategory_id: subcategoryID
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getCategorySubcategoriesList (categoryID) {
  return axios.get(
    SUBCATEGORY_PATH,
    {
      params: {
        category_id: categoryID
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function deleteProduct (productID) {
  return axios.post(PRODUCT_DELETE_PATH, {
    id: productID
  })
}

function updateProduct (payload) {
  return axios.put(
    PRODUCT_PATH,
    {
      id: payload.id,
      subcategory_id: payload.subcategory,
      name: payload.name,
      description: payload.description,
      price: payload.price,
      quantity: parseInt(payload.quantity),
      image: payload.image
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function createDiscount (name, email, percentage) {
  return axios.post(
    DISCOUNT_PATH,
    {
      name: name,
      user_id: email,
      percentage: parseFloat(percentage)
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getDiscountList () {
  return axios.get(
    DISCOUNT_PATH,
    {},
    helpers.generateConfig({ Authorizer: token })
  )
}

function deleteDiscount (id) {
  return axios.post(
    DISCOUNT_DELETE_PATH,
    {
      discount_id: id
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function addProductToCart (payload) {
  console.log(userEmail)
  console.log(payload)
  return axios.post(
    CART_PATH,
    {
      user_id: userEmail,
      product_id: payload.productID,
      quantity: payload.quantity
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function deleteFromCart (productID) {
  return axios.post(
    CART_DELETE_PATH,
    {
      product_id: productID
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function addCard (id, userID, type, pin, expDate, name) {
  return axios.post(
    CARD_PATH,
    {
      card_id: id,
      type: type,
      user_id: userID,
      expiration_date: expDate,
      pin: pin,
      name: name
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function makePurchase (payload) {
  console.log(payload)
  return axios.post(
    PURCHASE_PATH,
    {
      user_id: userEmail,
      payment_methods: payload.paymentMethods,
      discount_id: payload.discount
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getPayMethods () {
  return axios.get(
    LIST_CARDS_PATH,
    {
      params: {
        user_id: userEmail
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getDiscounts () {
  return axios.get(
    DISCOUNTS_PATH,
    {},
    helpers.generateConfig({ Authorizer: token })
  )
}

function getUserDiscounts(){
  console.log(userEmail)
  return axios.get(
    DISCOUNTS_USER_PATH,
    {
      params:{
        user_id: userEmail
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getUserEmail () {
  return userEmail
}

// TODO: change this for a cookie or another system, if the user reloads the session gets closed
function isLoged () {
  return token !== ''
}

function generateReport (type) {
  console.log(type)
  return axios.post(
    REPORT_PATH,
    {
      report_type: type
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function generateDateReport (type, to, from) {
  console.log(type, to, from)
  return axios.post(
    REPORT_PATH,
    {
      report_type: type,
      to: to,
      from: from
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function generateSalesProductReport (type, productID) {
  console.log(type, productID)
  return axios.post(
    REPORT_PATH,
    {
      report_type: type,
      id: productID
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getHistory () {
  return axios.get(
    HISTORY_PATH,
    {
      params: {
        user_id: userEmail
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

function getInvoice (invoiceID) {
  return axios.get(
    INVOICE_PATH,
    {
      params: {
        invoice_id: invoiceID
      }
    },
    helpers.generateConfig({ Authorizer: token })
  )
}

export {
  logInClient,
  logInAdmin,
  signUpClient,
  createAdmin,
  isLoged,
  getClientUser,
  getAdminUser,
  getAdminProfile,
  updateAdmin,
  getClientProfile,
  deleteAdminUser,
  deleteClientUser,
  updateClient,
  updatePassword,
  updateAdminPassword,
  createCategory,
  updateCategory,
  getCategoriesList,
  deleteCategory,
  getSubcategoriesList,
  createSubcategory,
  deleteSubCategory,
  updateSubcategory,
  createProduct,
  getSubcategoryProductsList,
  getCategorySubcategoriesList,
  deleteProduct,
  updateProduct,
  createDiscount,
  getDiscountList,
  deleteDiscount,
  addProductToCart,
  deleteFromCart,
  addCard,
  makePurchase,
  getPayMethods,
  getDiscounts,
  getUserEmail,
  generateReport,
  generateDateReport,
  generateSalesProductReport,
  getHistory,
  getInvoice,
  getUserDiscounts
}
