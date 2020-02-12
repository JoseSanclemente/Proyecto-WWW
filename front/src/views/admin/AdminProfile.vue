<template>
  <div class="form-profile">
    <div class="form-background">
      <form class="form-container">
        <div class="form-options">
          <div class="form-image">
            
          </div>
          <div class="option" @click="infoSelect"> Datos de la cuenta</div>
          <div class="option" @click="passwordSelect"> Cambiar contraseña</div>
        </div>

        <div class="form-data" v-if="infoSelected">
          <div class="form-title">Información del usuario</div>
          <div class="form-input">
            <div class="form-input-label"> Nombre </div>
            <input class="form-input-single-line" type="text" v-model="name" />
          </div>

          <div class="form-input">
            <div class="form-input-label"> Correo electrónico </div>
            <input class="form-input-single-line" type="text" v-model="email" disabled />
          </div>

          <div class="confirmation-button" @click="updateAdmin">Modificar datos</div>
        </div>

        <div class="form-data" v-if="passwordSelected">
          <div class="form-title">Cambiar contraseña</div>

          <div class="form-input">
            <div class="form-input-label"> Nueva contraseña </div>
            <input class="form-input-single-line" type="password" v-model="newPassword"/>
          </div>

          <div class="form-input">
            <div class="form-input-label"> Verificar contraseña </div>
            <input class="form-input-single-line" type="password" v-model="verifyNewPass" />
          </div>

          <div class="confirmation-button button" @click="changePassword">Actualizar</div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { getAdminProfile, updateAdminPassword, updateAdmin } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      infoSelected: true,
      passwordSelected: false,
      email: '',
      name: '',
      password: '',
      currentPassword: '',
      newPassword: '',
      verifyNewPass: '',

      image: ''
    }
  },

  methods: {
    infoSelect () {
      this.infoSelected = true
      this.passwordSelected = false
    },

    passwordSelect () {
      this.infoSelected = false
      this.passwordSelected = true
    },

    validateInputs () {
      if (this.name === '') {
        notify({
          body: 'El campo de nombre está vacío',
          type: 'alert'
        })
        return true
      }

      if (this.address === '') {
        notify({
          body: 'El campo de dirección está vacío',
          type: 'alert'
        })
        return true
      }

      return false
    },

    changePassword () {
      if (this.newPassword !== this.verifyNewPass) {
        notify({
          body: 'Las contraseñan no coinciden',
          type: 'alert'
        })
        return
      }

      updateAdminPassword(this.email, this.newPassword).then((response) => {
        notify({
          body: 'La contraseña ha sido actualizada',
          type: 'success'
        })
        this.newPassword = ''
        this.verifyNewPass = ''
      })
        .catch((error) => {
          console.log(error)
          notify({
            body: 'Ha ocurrido un error :(',
            type: 'error'
          })
        })
    },

    updateAdmin () {
      if (this.validateInputs()) {
        return
      }

      updateAdmin(this.email, this.name, this.address).then((response) => {
        notify({
          body: 'Los datos han sido actualizados',
          type: 'success'
        })
      })
        .catch((error) => {
          console.log(error)
          notify({
            body: 'Ha ocurrido un erro :(',
            type: 'error'
          })
        })
    }
  },

  mounted () {
    getAdminProfile()
      .then((response) => {
        this.email = response.data.email
        this.name = response.data.name
        this.birthDate = response.data['birth_date']
        this.address = response.data.address
      })
      .catch(function (error) {
        console.log(error)
      })
  }
}
</script>

<style scoped>
.form-profile {
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.form-container {
  display: flex;
  width: 60vw;
  height: 60vh;
  flex-direction: row;
  border: 1px solid blue;
}

.form-image{
  margin: 20px;
  border: 1px solid white;
  border-radius: 100px;
  height: 150px;
  width: 150px;
  content: url("../../assets/images/admin-pic.jpg");
}

.form-options{
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0px;
  width: 15vw;
  height: 60vh;
  color: white;
  -webkit-border-radius: 20px 0 0 20px;
  border-radius: 20px 0 0 20px;
  border: 1px solid white;
}

.option {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 20px;
  width: 15vw;
  border: 1px solid white;
  height: 8vh;
}

.option:hover {
  color: lightgreen;
}

.option:active {
  background: lightgreen;
  color: black;
}

.form-data{
  display: flex;
  flex-direction: column;
  padding: 30px 30px 5px 30px;
  height: 60vh;
  width: 45vw;
  -webkit-border-radius: 0 20px 20px 0;
  border-radius: 0 20px 20px 0;
}

.button{
  margin: 40px;
}
</style>
