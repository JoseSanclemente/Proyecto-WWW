<template>
  <div class="form-profile">
    <div class="form-background">
      <form class="form-container">
        <div class="form-options">
          <div class="form-image">
          </div>
          <div class="option" @click="infoSelect"> Datos de la cuenta</div>
          <div class="option" @click="passwordSelect"> Cambiar contraseña</div>
          <div class="option" @click="payMethodSelect">Métodos de Pago</div>
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

          <div class="form-input">
            <div class="form-input-label"> Fecha de nacimiento </div>
            <input class="form-input-single-line" type="text" v-model="birthDate" disabled/>
          </div>

          <div class="form-input">
            <div class="form-input-label"> Dirección </div>
            <input class="form-input-single-line" type="text" v-model="address" />
          </div>

          <div class="confirmation-button" @click="updateClient">Modificar datos</div>
          <div class="dangerous-button" @click="deleteClientUser">Eliminar cuenta</div>
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

        <div class="form-data" v-if="cardSelected">
          <div class="form-title">Agregar método de pago</div>

          <div class="form-input">
            <div class="form-input-label"> Tipo de tarjeta </div>
            <div class="combo-box">
                <select v-model="card_type">
                    <option>Crédito</option>
                    <option>Débito</option>
                </select>
            </div>
          </div>
          <div class="form-input">
            <div class="form-input-label">Número de la tarjeta</div>
            <input class="form-input-single-line" type="text" v-model="card_id"/>
          </div>
          <div class="form-input">
            <div class="form-input-label">Fecha de expiración</div>
            <input class="form-input-single-line" type="text" v-model="expiration_date" />
          </div>
          <div class="form-input">
            <div class="form-input-label">CVV</div>
            <input class="form-input-single-line" type="text" v-model="card_pin" />
          </div>
          <div class="form-input">
            <div class="form-input-label">Nombre de la tarjeta</div>
            <input class="form-input-single-line" type="text" v-model="card_name" />
          </div>

          <div class="confirmation-button" @click="addPayMethod">Agregar tarjeta</div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { getClientProfile, deleteClientUser, updatePassword, updateClient, addCard } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      infoSelected: true,
      passwordSelected: false,
      cardSelected: false,
      email: '',
      name: '',
      birthDate: '',
      address: '',
      password: '',
      currentPassword: '',
      newPassword: '',
      verifyNewPass: '',
      card_id: '',
      card_type: '',
      expiration_date: '',
      card_pin: '',
      card_name: '',
      image: ''
    }
  },

  methods: {
    infoSelect () {
      this.infoSelected = true
      this.passwordSelected = false
      this.cardSelected = false
    },

    passwordSelect () {
      this.infoSelected = false
      this.passwordSelected = true
      this.cardSelected = false
    },

    payMethodSelect () {
      this.infoSelected = false
      this.passwordSelected = false
      this.cardSelected = true
    },

    validatePayment () {
      if (this.card_id === '') {
        notify({
          body: 'El campo del número de la tarjeta está vacío',
          type: 'alert'
        })
        return true
      }
      if (this.card_type === '') {
        notify({
          body: 'Seleccione el tipo de pago',
          type: 'alert'
        })
        return true
      }
      if (this.expiration_date === '') {
        notify({
          body: 'El campo de fecha de expiración está vacío',
          type: 'alert'
        })
        return true
      }
      if (this.card_pin === '') {
        notify({
          body: 'El campo del código de seguridad está vacío',
          type: 'alert'
        })
        return true
      }
      if (this.card_name === '') {
        notify({
          body: 'El campo del nombre del tarjeta está vacío',
          type: 'alert'
        })
        return true
      }
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

      updatePassword(this.email, this.newPassword).then((response) => {
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

    updateClient () {
      if (this.validateInputs()) {
        return
      }

      updateClient(this.email, this.name, this.address).then((response) => {
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
    },

    deleteClientUser () {
      deleteClientUser(this.email).then(() => {
        notify({
          body: 'La cuenta ha sido eliminada correctamente',
          type: 'success'
        })
        this.$router.push({ name: 'home' })
      }).catch((error) => {
        notify({
          body: 'No se pudo eliminar el usuario',
          type: 'error'
        })
        console.log(error)
      })
    },

    addPayMethod () {
      if (this.validatePayment()) {
        return
      }

      let paymentType
      if (this.card_type === 'Crédito') {
        paymentType = 'credit'
      } else {
        paymentType = 'debit'
      }

      addCard(this.card_id, this.email, paymentType, parseInt(this.card_pin), this.expiration_date, this.card_name)
        .then(() => {
          notify({
            body: 'Se ha agregado el método de pago',
            type: 'success'
          })
          this.card_id = ''
          this.card_type = ''
          this.expiration_date = ''
          this.card_pin = ''
          this.card_name = ''
        }).catch((error) => {
          notify({
            body: 'No se pudo agregar el método de pago',
            type: 'error'
          })
          console.log(error)
        })
    }
  },

  mounted () {
    getClientProfile()
      .then((response) => {
        this.email = response.data.email
        this.name = response.data.name
        this.birthDate = response.data['birth_date']
        this.address = response.data.address
        this.password = response.data.password
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
  content: url("../../assets/images/client_pic.jpg");
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
