<template>
  <div>
    <div class="form-show-product">
      <router-link to="/user" class="back">Regresar</router-link>
      <div class="form-background">
        <div class="form">
          <h3 class="product-name">{{ product.name }}</h3>
          <table class="product-table">
            <tr>
              <td>
                <img :src="product.image" alt class="product-image" />
              </td>
              <td>
                <div class="product-information">
                  <h2>Descripción</h2>
                  <br />
                  <p>{{ product.description }}</p>
                  <br />
                  <span class="back-button" @click="decrementQuantityToBuy()">-</span>
                  &nbsp;
                  <em>Cantidad:</em>
                  {{quantityToBuy}}
                  &nbsp;
                  <span
                    class="back-button"
                    @click="incrementQuantityToBuy()"
                  >+</span>
                  <br />
                  <br />
                  <em>Precio:</em>
                  <p>${{ product.price }}</p>
                </div>
                <br />
                <br />
                <span
                  class="confirmation-button"
                  @click="addToBasket()"
                  v-if="!disabled"
                >Añadir a la canasta</span>
              </td>
            </tr>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { basketStore } from '@/basketStore/basketStore.js'
import { notify } from '@/notify.js'
import { getSubcategoryProductsList } from '@/api/api.js'

export default {
  data: function () {
    return {
      product: {},
      disabled: false,
      quantityToBuy: 1
    }
  },

  methods: {
    addToBasket () {
      if (this.product.quantity < 1) {
        notify({
          body: 'Lo sentimos, no nos quedan unidades disponibles',
          type: 'alert'
        })

        return
      }

      this.product.quantityToBuy = this.quantityToBuy
      basketStore.addProduct(this.product)

      this.product.quantity -= this.quantityToBuy
      if (this.product.quantity > 0) {
        this.quantityToBuy = 1
        return
      }
      this.quantityToBuy = 0

      notify({
        body: 'El producto se encuentra en su canasta de compras',
        type: 'success'
      })
    },

    loadActualProduct (products) {
      for (let i = 0; i < products.length; i++) {
        if (products[i].id === this.$route.params.product_id) {
          return products[i]
        }
      }

      notify({
        body: 'No se pudo cargar el producto indicado',
        type: 'error'
      })
    },

    incrementQuantityToBuy () {
      if (this.product.quantity > this.quantityToBuy) this.quantityToBuy++
    },

    decrementQuantityToBuy () {
      if (this.quantityToBuy > 1) this.quantityToBuy--
    }
  },

  mounted () {
    getSubcategoryProductsList(this.$route.params.subcategory_id)
      .then(response => {
        this.product = this.loadActualProduct(response.data)
        if (this.product.quantity < 1) {
          notify({
            body: 'Lo sentimos, no nos quedan unidades disponibles',
            type: 'alert'
          })
          this.disabled = true
          this.quantityToBuy = 0
        }
      })
      .catch(error => {
        console.log(error)
        notify({
          body: 'No se pudo cargar el producto indicado',
          type: 'error'
        })
      })
  }
}
</script>

<style scoped>
.form-show-product {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 125vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.product-name {
  color: white;
  margin: 10px;
}

.product-image {
  max-width: 400px;
  max-height: 400px;

  float: left;
}

.product-information {
  color: white;
  display: block;
  width: 400px;
}

.product-table {
  margin: 20px;
}
</style>
