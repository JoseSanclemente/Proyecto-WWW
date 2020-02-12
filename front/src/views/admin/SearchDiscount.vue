<template>
    <div class="form-search-discount">
        <router-link to="/admin" class="back">Regresar al menú principal</router-link>
        <div class="form-background">
            <form class="form">
                <div class="form-title">Consultar descuentos</div>
                <div class="form-input">
                    <div class="form-input-label">Descuentos</div>
                    <div class="combo-box">
                        <select v-model="discountID" @change="getInfo">
                            <option
                            v-for="discount in discounts"
                            v-bind:key="discount.discount_id"
                            :value="discount.discount_id"
                            >{{ discount.name }}</option>
                        </select>
                    </div>
                </div>

                <div class="form-input">
                    <div class="form-input-label">Porcentaje de descuento</div>
                    <input class="form-input-single-line" type="text" id="percentage" v-model="percentage" disabled/>
                </div>

                <div class="form-input">
                    <div class="form-input-label">Usuario</div>
                    <input class="form-input-single-line" type="text" id="user" v-model="user" disabled/>
                </div>
                <div class="dangerous-button" @click="deleteDiscount">Eliminar</div>
            </form>
        </div>
    </div>
</template>

<script>
import { getDiscountList, deleteDiscount } from '@/api/api.js'
import { notify } from '@/notify.js'
export default {
  data: function () {
    return {
      discounts: [],
      discountID: '',
      percentage: '',
      user: ''
    }
  },

  methods: {
    validateInput () {
      if (this.discountID === '') {
        notify({
          body: 'Debe seleccionar un descuento',
          type: 'alert'
        })
        return true
      }

      return false
    },

    deleteDiscount () {
      if (this.validateInput()) {
        return
      }

      deleteDiscount(this.discountID).then((response) => {
        notify({
          body: 'El descuento ha sido eliminado',
          type: 'success'
        })
        this.discountID = ''
        this.user = ''
        this.percentage = ''
        getDiscountList().then((response) => {
          this.discounts = response.data
        })
      })
        .catch((error) => {
          console.log(error)
          notify({
            body: 'Ha ocurrido un error :(',
            type: 'error'
          })
        })
    },

    getInfo () {
      for (let i = 0; i < this.discounts.length; i++) {
        if (this.discounts[i].discount_id === this.discountID) {
          this.percentage = this.discounts[i].percentage
          this.user = this.discounts[i].user_id
        }
      }
    }
  },

  mounted () {
    getDiscountList().then((response) => {
      this.discounts = response.data
    })
  }
}
</script>

<style scoped>
.form-search-discount {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 125vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
