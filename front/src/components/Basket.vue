<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h3>Mi canasta de compras</h3>
          </div>

          <div class="modal-body">
            <table class="basket-item" align="center">
              <tr v-for="product in productsInBasket" :key="product.id">
                <td>
                  <img :src="product.image" alt="item image" class="item-image" />
                </td>
                <td>
                  <div class="item-name">{{ product.name }}</div>
                  <div class="item-price">${{ product.price }} x {{ product.quantityToBuy }}</div>
                </td>
                <td>
                  <div class="controls-container">
                    <div class="dangerous-button" @click="removeUnit(product.id)">X</div>
                  </div>
                </td>
              </tr>
            </table>
          </div>

          <div class="modal-footer">
            <div>Total: ${{ getTotalPrice() }}</div>
            <br />
            <slot name="footer">
              <span
                class="confirmation-button"
                @click="goToCheckout(); $emit('buy')"
              >&nbsp;Comprar&nbsp;</span>
              &nbsp;
              <span
                class="dangerous-button"
                @click="$emit('close')"
              >&nbsp;Cerrar&nbsp;</span>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
import { basketStore } from '@/basketStore/basketStore.js'
import { notify } from '@/notify.js'
export default {
  data: function () {
    return {
      productsInBasket: []
    }
  },

  methods: {
    getTotalPrice () {
      return basketStore.getTotalPrice()
    },

    removeUnit (productID) {
      basketStore.removeUnit(productID)
      this.productsInBasket = basketStore.getProductsInBasket()
    },

    goToCheckout () {
      if (this.productsInBasket.length < 1) {
        notify({
          body: 'Tu canasta de compras se encuentra vacía',
          type: 'alert'
        })
        return
      }
      this.$emit('close')
      this.$router.push({
        name: 'checkout'
      })
    }
  },

  mounted () {
    this.productsInBasket = basketStore.getProductsInBasket()
  }
}
</script>

<style scoped>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: black;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
  color: white;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
  color: white;
  max-height: 200px;
  overflow: auto;
}

.modal-default-button {
  float: right;
}

.basket-item {
  width: 100%;
}

.item-image {
  max-width: 68px;
  max-height: 68px;
}

.controls-container {
  width: 30px;
  margin: auto;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>
