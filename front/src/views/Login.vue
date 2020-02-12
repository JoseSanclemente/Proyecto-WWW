<template>
  <div class="form-login">
    <div class="form-background">
      <form class="form">
        <div class="form-title">Iniciar sesion</div>
        <div class="form-input">
          <label class="form-input-label" for="email">Email</label>
          <input class="form-input-single-line" type="text" id="email" v-model="email" />
        </div>
        <div class="form-input">
          <label class="form-input-label" for="password">Contraseña</label>
          <input class="form-input-single-line" type="password" id="password" v-model="password" />
        </div>
        <div class="confirmation-button" @click="authenticateUser">Entrar</div>
      </form>
    </div>
  </div>
</template>

<script>
import { logInClient, logInAdmin } from '@/api/api.js'

export default {
  name: 'login',

  data: function () {
    return {
      email: '',
      password: ''
    }
  },

  methods: {
    routeAutentication () {
      if (this.$router.currentRoute.name === 'login-client') {
        return logInClient
      }
      return logInAdmin
    },

    redirectToHome () {
      if (this.$router.currentRoute.name === 'login-client') {
        console.log('going home')
        this.$router.push({ name: 'user-home' }) // TODO:go to user home here
        return
      }
      this.$router.push({ name: 'admin-home' })
    },

    authenticateUser: function () {
      this.routeAutentication()(this.email, this.password)
        .then(() => this.redirectToHome()) // Arrow function let me use 'this'
        // TODO: notify the error
        .catch(error => console.log(error))
    }
  }
}
</script>

<style scoped>
.form-login {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 110vh;
  font-size: 18px;
  background-image: url("../assets/images/background_palm_tree.jpeg");
}

.forget-password {
  color: #505050;
}
</style>
