<template>
    <div class="form-add-discount">
        <router-link to="/admin" class="back">Regresar al menú principal</router-link>
        <div class="form-background">
            <form class="form">
                <div class="form-title">Crear descuento</div>
                <div class="form-input">
                    <div class="form-input-label">Nombre</div>
                    <input class="form-input-single-line" type="text" id="name" v-model="name">
                </div>

                <div class="form-input">
                    <div class="form-input-label">Email del usuario</div>
                    <input class="form-input-single-line" type="text" id="email" v-model="userEmail">
                </div>

                <div class="form-input">
                    <div class="form-input-label">Porcentaje de descuento</div>
                    <input class="form-input-single-line" type="text" id="percentage" v-model="percentage">
                </div>
                <div class="confirmation-button" @click="createDiscount">Crear</div>
            </form>
        </div>
    </div>
</template>

<script>
import { createDiscount } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      userEmail: '',
      percentage: '',
      name: ''
    }
  },

  methods: {
    validateInput () {
      if (this.name === '') {
        notify({
          body: 'Debe ingresar el nombre del descuento',
          type: 'alert'
        })
        return true
      }

      if (this.userEmail === '') {
        notify({
          body: 'Debe ingresar el usuario destino',
          type: 'alert'
        })
        return true
      }

      if (this.percentage === '') {
        notify({
          body: 'Debe ingresar el porcentaje de descuento',
          type: 'alert'
        })
        return true
      }

      return false
    },

    createDiscount () {
      if (this.validateInput()) {
        return
      }

      createDiscount(this.name, this.userEmail, this.percentage).then((response) => {
        notify({
          body: 'El descuento ha sido enviado',
          type: 'success'
        })
        this.name = ''
        this.userEmail = ''
        this.percentage = ''
      })
        .catch((error) => {
          console.log(error)
          notify({
            body: 'Ha ocurrido un error :(',
            type: 'error'
          })
        })
    }
  }
}
</script>

<style scoped>
.form-add-discount {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 125vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
