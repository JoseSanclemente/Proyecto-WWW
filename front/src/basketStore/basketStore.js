import { helpers } from './helpers.js'
import { addProductToCart, deleteFromCart } from '../api/api.js'

var basketStore = {
  debug: true,

  state: {
    products: []
  },

  // Please don't modify the state directly, instead add methods, this make it easier to debug
  /* SETTERS */
  addProduct (product) {
    this.state.products.push(JSON.parse(JSON.stringify(product)))
    if (this.debug) console.log('PRODUCTS', this.state.products)
    this.updateServerCart()
  },

  removeUnit (productID) {
    for (let i = 0; i < this.state.products.length; i++) {
      if (this.state.products[i].id === productID) {
        this.state.products[i].quantityToBuy--

        if (this.state.products[i].quantityToBuy < 1) {
          this.state.products.splice(i, 1)
          this.updateServerCart(productID)
          return
        }
        this.updateServerCart()
        return
      }
    }
  },

  // Provide a productID to detele it from the server cart
  updateServerCart (productID) {
    if (productID !== undefined) {
      deleteFromCart(productID)
      return
    }

    let productsInBasket = this.getProductsInBasket()
    for (let i = 0; i < productsInBasket.length; i++) {
      // If the product already exists in the cart it gets updated
      addProductToCart({
        productID: productsInBasket[i].id,
        quantity: productsInBasket[i].quantityToBuy
      })
        .then(() => {
          if (this.debug) console.log('UPDATING CART')
        })
        .catch(() => {
          if (this.debug) console.log('ERROR UPDATING CART')
        })
    }
  },

  /* GETTERS */
  getProductsInBasket () {
    let productsInBasket = []

    for (let i = 0; i < this.state.products.length; i++) {
      let product = helpers.searchProductByID(
        productsInBasket,
        this.state.products[i].id
      )
      if (product === undefined) {
        productsInBasket.push(
          JSON.parse(JSON.stringify(this.state.products[i]))
        )
        continue
      }
      product.quantityToBuy += this.state.products[i].quantityToBuy
    }

    return productsInBasket
  },

  getTotalPrice () {
    let totalPrice = 0

    for (let i = 0; i < this.state.products.length; i++) {
      totalPrice +=
        this.state.products[i].price * this.state.products[i].quantityToBuy
    }

    return totalPrice
  }
}

export { basketStore }
