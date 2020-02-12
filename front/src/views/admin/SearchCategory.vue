<template>
  <div>
    <div class="form-search-category" v-if="selectionScreen">
      <router-link to="/admin" class="back">Regresar al menú principal</router-link>
      <div class="form-background">
        <form class="form">
          <div class="form-title">Consultar categoría</div>
          <div class="form-input">
            <label class="form-input-label">Categoría</label>
            <div class="combo-box">
              <select v-model="searchedCategory">
                <option
                  v-for="category in categories"
                  v-bind:key="category.category_id"
                  :value="category.category_id"
                >{{ category.name }}</option>
              </select>
            </div>
          </div>
          <div class="confirmation-button" @click="selectCategory">Seleccionar</div>
        </form>
      </div>
    </div>

    <div class="form-search-category" v-if="!selectionScreen">
      <div class="form-background">
        <form class="form">
          <div class="form-back">
            <div class="back-button" @click="back">Regresar</div>
          </div>
          <div class="form-title">Detalles</div>
          <div class="category-container">
            <img class="category-pic" :src="this.categoryImage" alt />
          </div>

          <div class="form-input">
            <label class="form-input-label" for="name">Nombre</label>
            <input class="form-input-single-line" type="text" id="name" v-model="categoryName" />
          </div>

          <div class="form-input">
            <label class="form-input-label" for="description">Descripción</label>
            <textarea
              class="form-input-multiline"
              type="text"
              id="description"
              v-model="categoryDescription"
            />
          </div>

          <!--TODO: maybe to change the image in modify-->
          <div v-show="false" id="imageBase64"></div>
          <div class="form-input">
            <label class="form-input-label" for="image">Imagen</label>
            <input class="form-input-single-line" type="file" ref="fileupload" @change="onFileChange" />
          </div>

          <div class="form-input">
            <div class="confirmation-button" @click="updateCategory">Modificar datos</div>
            <div class="dangerous-button" @click="deleteCategory">Eliminar</div>
          </div>
          <!--TODO: maybe add options to modify or delete-->
          <!--div class="confirmation-button" @click="createProduct">Añadir</div-->
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { getCategoriesList, deleteCategory, updateCategory } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      selectionScreen: true,
      searchedCategory: '',
      categoryName: '',
      categoryDescription: '',
      categories: [],
      categoryImage: ''
    }
  },

  mounted () {
    getCategoriesList().then(response => {
      this.categories = response.data
    })
  },

  methods: {
    back () {
      this.selectionScreen = true
      this.searchedCategory = ''
    },

    validateInput () {
      if (this.searchedCategory === '') {
        notify({
          body: 'Debe seleccionar una categoría',
          type: 'alert'
        })
        return true
      }

      return false
    },

    validateNewInfo () {
      if (this.categoryName === '') {
        return true
      }

      if (this.categoryDescription === '') {
        return true
      }

      if (this.categoryImage === '') {
        return true
      }

      return false
    },

    clear () {
      const input = this.$refs.fileupload
      input.type = 'text'
      input.type = 'file'
    },

    selectCategory () {
      if (this.validateInput()) {
        return
      }

      this.selectionScreen = false

      for (let i = 0; i < this.categories.length; i++) {
        if (this.categories[i].category_id === this.searchedCategory) {
          this.categoryName = this.categories[i].name
          this.categoryDescription = this.categories[i].description
          // TODO: store the image in the database with the type from the start to avoid this
          this.categoryImage = this.categories[i].image
        }
      }
    },

    updateCategory () {
      if (this.validateNewInfo()) {
        notify({
          body: 'No deben haber campos vacíos',
          type: 'alert'
        })
        return
      }

      let base64Image = document
        .getElementById('imageBase64').innerHTML

      if (base64Image !== '') {
        base64Image = base64Image.split('base64,')[1]
          .split('">')[0]

        this.categoryImage = 'data:image/png;base64,' + base64Image
      }

      updateCategory(this.searchedCategory, this.categoryName, this.categoryDescription, this.categoryImage)
        .then((response) => {
          notify({
            body: 'Los datos han sido actualizados',
            type: 'success'
          })

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

    deleteCategory () {
      deleteCategory(this.searchedCategory).then(() => {
        notify({
          body: 'La categoría ha sido eliminada',
          type: 'success'
        })
        this.selectionScreen = true
        this.categoryName = ''
        this.searchedCategory = ''
        this.categoryDescription = ''
        this.categoryImage = ''
        getCategoriesList().then(response => {
          this.categories = response.data
        })
      })
        .catch(function (error) {
          console.log(error)
          notify({
            body: 'Ocurrió un error :(',
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
.form-search-category {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 120vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
