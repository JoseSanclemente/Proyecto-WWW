<template>
  <div class="form-profile">
    <router-link to="/" class="back">Regresar</router-link>
    <div class="form-background">
      <form class="form">
        <div class="form-title">Registro</div>
        <div class="form-input">
          <label class="form-input-label" for="name">Nombre</label>
          <input class="form-input-single-line" type="text" id="name" v-model="name" />
        </div>

        <div class="form-input">
          <label class="form-input-label">Tipo de ID</label>
          <div class="combo-box">
            <select v-model="documentType">
              <option>C.C</option>
              <option>Passport</option>
              <option>PEP</option>
            </select>
          </div>
        </div>

        <div class="form-input">
          <label class="form-input-label" for="email">Identificación</label>
          <input class="form-input-single-line" type="text" id="id" v-model="id" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="age">Fecha de nacimiento</label>
          <datepicker
            placeholder="Seleccionar fecha"
            input-class="form-input-single-line"
            v-model="birthDate"
          ></datepicker>
        </div>

        <div class="form-input">
          <label class="form-input-label" for="email">Correo electrónico</label>
          <input class="form-input-single-line" type="text" id="email" v-model="email" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="address">Dirección</label>
          <input class="form-input-single-line" type="text" id="address" v-model="address" />
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

        <!-- TODO: make reactive this button -->
        <div class="confirmation-button" @click="postNewClientUser">Registrarse</div>
      </form>
    </div>
  </div>
</template>

<script>
import { signUpClient } from '@/api/api.js'
import { notify } from '@/notify.js'
import Datepicker from 'vuejs-datepicker/dist/vuejs-datepicker.esm.js'
export default {
  components: {
    Datepicker
  },

  data: function () {
    return {
      name: '',
      email: '',
      birthDate: '',
      address: '',
      password: '',
      passwordConfirmation: '',
      id: '',
      documentType: ''
    }
  },

  methods: {
    validateInputs () {
      if (this.name === '') {
        notify({
          body: 'Ingresar un nombre de usuario',
          type: 'alert'
        })
        return true
      }

      if (this.email === '') {
        notify({
          body: 'Ingresar un correo electrónico',
          type: 'alert'
        })
        return true
      }

      if (this.birthDate === '') {
        notify({
          body: 'Ingresar la fecha de nacimiento',
          type: 'alert'
        })
        return true
      }

      if (this.address === '') {
        notify({
          body: 'Ingresar la dirección de residencia',
          type: 'alert'
        })
        return true
      }

      if (this.password === '') {
        notify({
          body: 'Ingresar una contraseña',
          type: 'alert'
        })
        return true
      }

      if (this.id === '') {
        notify({
          body: 'Ingresar un docunento de identidad',
          type: 'alert'
        })
        return true
      }

      if (this.documentType === '') {
        notify({
          body: 'Ingresar el tipo de documento de eidentidad',
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
    },
    postNewClientUser () {
      if (this.validateInputs()) {
        return
      }

      signUpClient(
        this.email,
        this.name,
        this.birthDate,
        this.address,
        this.password,
        this.id,
        this.documentType
      )
        .then((response) => {
          notify({
            body: 'Usuario registrado correctamente',
            type: 'success'
          })
          this.email = ''
          this.name = ''
          this.birthDate = ''
          this.address = ''
          this.password = ''
          this.id = ''
          this.documentType = ''
          this.passwordConfirmation = ''
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

<!-- WARNING: making this style scoped breaks the datepicker -->
<!-- TODO: fix the overflow -->
<style scoped>
.form-profile {
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
