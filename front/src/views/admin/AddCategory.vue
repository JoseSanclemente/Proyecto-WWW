<template>
  <div class="form-add-category">
    <router-link to="/admin" class="back">Regresar al menú principal</router-link>
    <div class="form-background">
      <form class="form">
        <div class="form-title">Añadir categoría</div>
        <div class="form-input">
          <label class="form-input-label" for="name">Nombre</label>
          <input class="form-input-single-line" type="text" id="name" v-model="name" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="description">Descripción</label>
          <textarea class="form-input-multiline description" type="text" id="description" v-model="description" />
        </div>

        <div v-show="false" id="imageBase64"></div>
        <div class="form-input">
          <label class="form-input-label" for="image">Imagen</label>
          <input class="form-input-single-line" type="file" ref="fileupload" @change="onFileChange" />
        </div>

        <!-- TODO: make reactive this button -->
        <div class="confirmation-button" @click="createCategory">Añadir</div>>
      </form>
    </div>
  </div>
</template>

<script>
import { createCategory } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      name: '',
      description: ''
    }
  },

  methods: {
    validateInputs () {
      if (this.name === '') {
        return true
      }

      if (this.description === '') {
        return true
      }

      return false
    },

    clear () {
      const input = this.$refs.fileupload
      input.type = 'text'
      input.type = 'file'
    },

    createCategory () {
      if (this.validateInputs()) {
        notify({
          body: 'Debe ingresar todos los datos',
          type: 'alert'
        })
        return
      }

      let base64Image = document
        .getElementById('imageBase64')
        .innerHTML.split('base64,')[1]
        .split('">')[0]

      base64Image = 'data:image/png;base64,' + base64Image
      createCategory(this.name, this.description, base64Image)
        .then((response) => {
          notify({
            body: 'La categoría ha sido añadida',
            type: 'success'
          })
          this.name = ''
          this.description = ''
          this.clear()
        })
        .catch((error) => {
          console.log(error)
          notify({
            body: 'Ha ocurrido un error :(',
            type: 'error'
          })
        })
    },

    onFileChange (e) {
      let files = e.target.files || e.dataTransfer.files
      if (!files.length) return

      let fileReader = new FileReader()

      fileReader.onload = function (fileLoadedEvent) {
        let srcData = fileLoadedEvent.target.result

        let newImage = document.createElement('img')
        newImage.src = srcData

        document.getElementById('imageBase64').innerHTML = newImage.outerHTML
      }

      fileReader.readAsDataURL(e.target.files[0])
    }
  }
}
</script>

<style scoped>
.form-add-category {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
