<template>
  <div class="form-add-subcategory">
    <router-link to="/admin" class="back">Regresar al menú principal</router-link>
    <div class="form-background">
      <form class="form">
        <div class="form-title">Añadir subcategoría</div>
        <div class="form-input">
          <label class="form-input-label">Categoría</label>
          <div class="combo-box">
            <select v-model="categoryID">
              <option
                v-for="category in categories"
                v-bind:key="category.category_id"
                :value="category.category_id"
              >{{ category.name }}</option>
            </select>
          </div>
        </div>

        <div class="form-input">
          <label class="form-input-label" for="name">Nombre</label>
          <input class="form-input-single-line" type="text" id="name" v-model="name" />
        </div>

        <div v-show="false" id="imageBase64"></div>
        <div class="form-input">
          <label class="form-input-label" for="image">Imagen</label>
          <input class="form-input-single-line" type="file" ref="fileupload" @change="onFileChange" />
        </div>

        <div class="confirmation-button" @click="createSubcategory">Añadir</div>
      </form>
    </div>
  </div>
</template>

<script>
import { createSubcategory, getCategoriesList } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      name: '',
      categories: [],
      categoryID: ''
    }
  },

  methods: {
    validateInput () {
      if (this.categoryID === '') {
        notify({
          body: 'Debe seleccionar una categoría',
          type: 'alert'
        })
        return true
      }

      if (this.name === '') {
        notify({
          body: 'La subcategoría debe tener nombre',
          type: 'alert'
        })
        return true
      }

      if (
        document.getElementById('imageBase64').innerHTML.split('base64,')[1] ===
        undefined
      ) {
        notify({
          body: 'La subcategoría debe tener imagen',
          type: 'alert'
        })
        return true
      }

      return false
    },

    clear () {
      const input = this.$refs.fileupload
      input.type = 'text'
      input.type = 'file'
    },

    createSubcategory () {
      let badInput = this.validateInput()
      if (badInput) {
        return
      }

      let base64Image = document
        .getElementById('imageBase64')
        .innerHTML.split('base64,')[1]
        .split('">')[0]

      base64Image = 'data:image/png;base64,' + base64Image
      createSubcategory({
        categoryID: this.categoryID,
        name: this.name,
        image: base64Image
      })
        .then((response) => {
          notify({
            body: 'La subcategoría ha sido añadida',
            type: 'success'
          })
          this.categoryID = ''
          this.name = ''
          this.image = ''
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
  },

  mounted () {
    getCategoriesList().then(response => {
      this.categories = response.data
    })
  }
}
</script>

<style scoped>
.form-add-subcategory {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
