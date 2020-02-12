<template>
    <div>
        <div class='form-search-user' v-if="searchScreen">
          <router-link to="/admin" class="back">Regresar al menú principal</router-link>
            <div class="form-background">
                <form class="form">
                    <div class="form-title">Consultar usuario</div>
                    <div class="form-input">
                        <label class="form-input-label">Email</label>
                        <input class="form-input-single-line" type="text" id="email" v-model="userEmail" />
                    </div>
                    <div class="form-input">
                        <div class="form-input-label">Tipo de usuario</div>
                        <div class="combo-box">
                            <select v-model="userType">
                                <option>Cliente</option>
                                <option>Administrador</option>
                            </select>
                        </div>
                    </div>
                    <div class="confirmation-button" @click="getUser">Consultar</div>
                </form>
            </div>
        </div>

        <div class="form-search-user" v-if="adminScreen">
            <div class="form-background">
                <form class="form">
                    <div class="form-back">
                        <div class="back-button back"  @click="back">Regresar</div>
                    </div>
                    <div class="form-title">Datos del administrador</div>
                    <div class="form-input">
                        <div class="form-input-label">Usuario</div>
                        <input class="form-input-single-line" type="text" id="user" v-model="username" />
                    </div>
                    <div class="form-input">
                        <div class="dangerous-button" @click="deleteAdminUser">Eliminar</div>
                    </div>
                </form>
            </div>
        </div>

        <div class="form-search-user" v-if="clientScreen">
          <router-link to="/admin">Back to Home</router-link>
            <div class="form-background">
                <form class="form">
                    <div class="form-back">
                      <div class="back-button" @click="back">Regresar</div>
                    </div>
                    <div class="form-title">Datos del cliente</div>
                    <div class="form-input">
                        <div class="form-input-label">Usuario</div>
                        <input class="form-input-single-line" type="text" id="user" v-model="username" />
                    </div>

                    <div class="form-input">
                        <div class="form-input-label">Fecha de nacimiento</div>
                        <input class="form-input-single-line" type="text" id="birthdate" v-model="birthdate" disabled/>
                    </div>

                    <div class="form-input">
                        <div class="form-input-label">Dirección</div>
                        <input class="form-input-single-line" type="text" id="address" v-model="address" disabled/>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import { getAdminUser, getClientUser, deleteAdminUser } from '@/api/api.js'
import { notify } from '@/notify.js'
export default {
  data: function () {
    return {
      searchScreen: true,
      adminScreen: false,
      clientScreen: false,
      userType: '',
      userEmail: '',
      username: '',
      birthdate: '',
      address: '',
      image: ''
    }
  },

  methods: {
    back () {
      this.searchScreen = true
      this.adminScreen = false
      this.clientScreen = false
      this.userEmail = ''
      this.userType = ''
    },

    validateInput () {
      if (this.userEmail === '') {
        notify({
          body: 'Debe ingresar un correo electrónico',
          type: 'alert'
        })
        return true
      }

      if (this.userType === '') {
        notify({
          body: 'Debe seleccionar un tipo de usuario',
          type: 'alert'
        })
        return true
      }

      return false
    },

    getUser () {
      if (this.validateInput()) {
        return
      }

      if (this.userType === 'Administrador') {
        getAdminUser(this.userEmail).then((response) => {
          this.username = response.data.name
          this.searchScreen = false
          this.adminScreen = true
        })
          .catch((error) => {
            console.log(error)
            this.userEmail = ''
            notify({
              body: 'Administrador no encontrado',
              type: 'alert'
            })
          })
      } else {
        getClientUser(this.userEmail).then((response) => {
          this.username = response.data.name
          this.birthdate = response.data.birth_date
          this.address = response.data.address
          this.searchScreen = false
          this.clientScreen = true
        })
          .catch((error) => {
            console.log(error)
            this.userEmail = ''
            notify({
              body: 'Cliente no encontrado',
              type: 'alert'
            })
          })
      }
    },

    deleteAdminUser () {
      deleteAdminUser(this.userEmail).then((response) => {
        notify({
          body: 'El administrador ha sido eliminado correctamente',
          type: 'success'
        })
        this.searchScreen = true
        this.adminScreen = false
        this.userEmail = ''
        this.userType = ''
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
.form-search-user {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 110vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.back {
  margin-bottom: 30px;
}
</style>
