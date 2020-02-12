<template>
  <div>
    <div class="form-check-out">
      <div class="form-background">
        <div class="form">
          <h3 class="title">Finalizar compra</h3>
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

          <div class="form-input">
            <label class="form-input-label">Descuento</label>
            <div class="combo-box">
              <select v-model="selectedDiscount">
                <option
                  v-for="discount in discounts"
                  v-bind:key="discount.discount_id"
                  :value="discount"
                >{{ discount.name + ' - ' + discount.percentage + '%'}}</option>
              </select>
            </div>
          </div>

          <div class="form-input">
            <label class="form-input-label">Añadir método de pago</label>
            <div class="combo-box">
              <select v-model="selectedPayMethod" @change="addPayMethod(selectedPayMethod)">
                <option
                  v-for="payMethod in payMethods"
                  v-bind:key="payMethod.card_id"
                  :value="payMethod"
                >{{ payMethod.name }}</option>
              </select>
            </div>
          </div>

          <div v-for="payMethod in selectedPayMethods" :key="payMethod.card_id">
            <div class="form-input">
              <label class="form-input-label" :for="payMethod.card_id">{{ payMethod.name }} (Porcentaje)</label>
              <input
                class="form-input-single-line"
                :id="payMethod.card_id"
                :ref="payMethod.card_id"
              />
            </div>
            <div v-if="payMethod.type === 'credit'" class="form-input">
              <label class="form-input-label" for="creditQuotas">Cuotas</label>
              <input class="form-input-single-line" :ref="payMethod.card_id+'-quotas'" />
            </div>
          </div>

          <div class="confirmation-button" @click="pay()">Pagar</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { basketStore } from '@/basketStore/basketStore.js'
import { notify } from '@/notify.js'
import {
  makePurchase,
  getPayMethods,
  getUserDiscounts,
  getUserEmail
} from '@/api/api.js'

export default {
  data: function () {
    return {
      productsInBasket: [],
      address: '',
      payMethods: [],
      selectedPayMethods: [],
      selectedPayMethod: {},
      discounts: [],
      selectedDiscount: {}
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

    addPayMethod (payMethod) {
      if (payMethod === undefined) return

      this.selectedPayMethods.push(JSON.parse(JSON.stringify(payMethod)))

      for (let i = 0; i < this.payMethods.length; i++) {
        if (this.payMethods[i].card_id === payMethod.card_id) {
          this.payMethods.splice(i, 1)
          break
        }
      }
    },

    validatePorcentaje (str) {
      if (!/^\d+\.*\d*$/.test(str)) {
        notify({
          body: 'La cantidad a pagar debe ser ingresada como porcentaje.',
          type: 'alert'
        })
        return false
      }

      return true
    },

    validatedQuotas (str) {
      let number = Number(str)

      if (isNaN(number) || !Number.isInteger(number)) {
        notify({
          body: 'Digite el numero de quotas como un numero entero',
          type: 'alert'
        })

        return false
      }

      if (number < 1 || number > 36) {
        notify({
          body: 'El número de quotas debe estar entre 1 y 36',
          type: 'alert'
        })

        return false
      }

      return true
    },

    validatePayment (paymentInfo) {
      if (paymentInfo.length < 1) {
        notify({
          body: 'Indique al menos un medio de pago',
          type: 'alert'
        })
        return false
      }

      let totalProportion = 0
      for (let i = 0; i < paymentInfo.length; i++) {
        totalProportion += paymentInfo[i].percentage
      }

      if (totalProportion !== 100.0) {
        notify({
          body: 'La distribucion de los porcentajes debe sumar 100',
          type: 'alert'
        })
        return false
      }

      return true
    },

    pay () {
      if (this.productsInBasket.length < 1) {
        notify({
          body: 'No hay productos en la canasta',
          type: 'alert'
        })

        return
      }

      let paymentInfo = []
      for (let i = 0; i < this.selectedPayMethods.length; i++) {
        if (
          !this.validatePorcentaje(
            this.$refs[this.selectedPayMethods[i].card_id][0].value
          )
        ) {
          return
        }

        let quotas = ''
        if (
          this.$refs[this.selectedPayMethods[i].card_id + '-quotas'] !==
          undefined
        ) {
          if (
            this.validatedQuotas(
              this.$refs[this.selectedPayMethods[i].card_id + '-quotas'][0]
                .value
            )
          ) {
            quotas = this.$refs[this.selectedPayMethods[i].card_id + '-quotas'][0].value
          } else {
            return
          }
        }

        paymentInfo.push({
          type: this.selectedPayMethods[i].type,
          card_id: this.selectedPayMethods[i].card_id,
          percentage: parseFloat(
            this.$refs[this.selectedPayMethods[i].card_id][0].value
          ),
          quotas: parseInt(quotas)
        })
      }

      if (!this.validatePayment(paymentInfo)) {
        return
      }

      makePurchase({
        discount: this.selectedDiscount.discount_id,
        paymentMethods: paymentInfo
      })
        .then(response => {
          if (response.status === 204) {
            notify({
              body: 'Alguien ha comprado el artículo antes que tú',
              type: 'alert'
            })
          } else {
            notify({
              body: 'La compra ha sido exitosa',
              type: 'success'
            })
          }
          basketStore.state.products = []
          this.$router.push({
            name: 'user-home'
          })
        })
        .catch(error => {
          console.log(error)
          notify({
            body: 'Ha ocurrido un error :(',
            type: 'error'
          })
        })
    }
  },

  computed: {
    payWithCreditCard: function () {
      for (let i = 0; i < this.selectedPayMethods.length; i++) {
        if (this.selectedPayMethods[i].type === 'credit') {
          return true
        }
      }

      return false
    }
  },

  mounted () {
    this.productsInBasket = basketStore.getProductsInBasket()

    getPayMethods()
      .then(response => {
        this.payMethods = response.data
        this.payMethods.push({
          name: 'Efectivo',
          card_id: '123',
          type: 'cash'
        })
      })
      .catch(() => {
        console.log('ERROR GETTING PAYMENT METHODS')
      })

    getUserDiscounts()
      .then(response => {
        this.discounts = response.data
      })
      .catch(err => {
        console.log('ERROR ON LOAD DISCOUNT' + err)
      })
  }
}
</script>

<style scoped>
.form-check-out {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 125vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.title {
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
</style>
