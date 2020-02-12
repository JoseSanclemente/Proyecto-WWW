var helpers = {
  searchProductByID (products, productID) {
    for (let i = 0; i < products.length; i++) {
      if (products[i].id === productID) {
        return products[i]
      }
    }
  }
}

export {
  helpers
}
