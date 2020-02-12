<template>
  <div class="form-add-admin">
    <router-link to="/admin" class="back">Regresar al menú principal</router-link>
    <div class="form-background">
      <form class="form">
        <div class="form-title">Registrar administrador</div>
        <div class="form-input">
          <label class="form-input-label" for="name">Nombre</label>
          <input class="form-input-single-line" type="text" id="name" v-model="name" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="email">Correo electrónico</label>
          <input class="form-input-single-line" type="text" id="email" v-model="email" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="password">Contraseña</label>
          <input class="form-input-single-line" type="password" id="password" v-model="password" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="password-confirmation">Confirmar contraseña</label>
          <input
            class="form-input-single-line"
            type="password"
            id="password-confirmation"
            v-model="passwordConfirmation"
          />
        </div>
        <div class="confirmation-button create" @click="postNewAdmin">Crear</div>
      </form>
    </div>
  </div>
</template>

<script>
import { createAdmin } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      name: '',
      email: '',
      password: '',
      passwordConfirmation: ''
    }
  },

  methods: {
    validateInputs () {
      if (this.name === '') {
        notify({
          body: 'Debe ingresar un nombre de usuario',
          type: 'alert'
        })
        return true
      }

      if (this.email === '') {
        notify({
          body: 'Debe ingresar un correo',
          type: 'alert'
        })
        return true
      }

      if (this.password === '') {
        notify({
          body: 'Debe ingresar una contraseña',
          type: 'alert'
        })
        return true
      }

      if (this.passwordConfirmation === '') {
        notify({
          body: 'Debe ingresar la verificación de contraseña',
          type: 'alert'
        })
        return true
      }

      if (this.password !== this.passwordConfirmation) {
        notify({
          body: 'Las contraseñas no coinciden',
          type: 'alert'
        })
        return true
      }

      return false
    },

    postNewAdmin () {
      if (this.validateInputs()) {
        return
      }

      createAdmin(this.email, this.name, this.password)
        .then((response) => {
          notify({
            body: 'Se ha agregado un nuevo administrador',
            type: 'success'
          })
          this.name = ''
          this.email = ''
          this.password = ''
          this.passwordConfirmation = ''
        })
        .catch(function (error) {
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
<style>
.form-add-admin {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 110vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

</style>
