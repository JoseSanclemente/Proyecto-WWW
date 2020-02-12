<template>
  <div>
    <div class="form-search-subcategory" v-if="selectionScreen">
      <router-link to="/admin" class="back">Regresar al menú principal</router-link>
      <div class="form-background">
        <form class="form">
          <div class="form-title">Consultar subcategoría</div>
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

        <div class="form-input">
            <label class="form-input-label" for="subcategory">Subcategoría</label>
            <div class="combo-box">
              <select v-model="searchedSubcategory">
                <option
                  v-for="subcategory in subcategories"
                  v-bind:key="subcategory.id"
                  :value="subcategory.id"
                >{{ subcategory.name }}</option>
              </select>
            </div>
        </div>

          <div class="confirmation-button" @click="selectSubcategory">Seleccionar</div>
        </form>
      </div>
    </div>

    <div class="form-search-subcategory" v-if="!selectionScreen">
      <div class="form-background">
        <form class="form">
          <div class="form-back">
            <div class="back-button" @click="back">Regresar</div>
          </div>
          <div class="form-title">Detalles</div>
          <div class="category-container">
            <img class="category-pic" :src="this.subcategoryImage" alt />
          </div>

          <div class="form-input">
            <label class="form-input-label" for="name">Nombre</label>
            <input class="form-input-single-line" type="text" id="name" v-model="subcategoryName" />
          </div>

          <div v-show="false" id="imageBase64"></div>
          <div class="form-input">
            <label class="form-input-label" for="image">Imagen</label>
            <input class="form-input-single-line" type="file" ref="fileupload" @change="onFileChange"/>
          </div>

          <div class="form-input">
            <div class="confirmation-button" @click="updateSubcategory">Modificar datos</div>
            <div class="dangerous-button" @click="deleteSubCategory">Eliminar</div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { getCategoriesList, getSubcategoriesList, deleteSubCategory, updateSubcategory } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      selectionScreen: true,
      searchedCategory: '',
      searchedSubcategory: '',
      subcategoryName: '',
      subcategoryImage: '',
      categories: [],
      subcategories: []
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
      this.searchedSubcategory = ''
    },

    validateInput () {
      if (this.searchedCategory === '') {
        notify({
          body: 'Debe seleccionar una categoría',
          type: 'alert'
        })
        return true
      }

      if (this.searchedSubcategory === '') {
        notify({
          body: 'Debe seleccionar una subcategoría',
          type: 'alert'
        })
        return true
      }

      return false
    },

    validateNewInfo () {
      if (this.subcategoryName === '') {
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

    selectSubcategory () {
      if (this.validateInput()) {
        return
      }

      this.selectionScreen = false

      for (let i = 0; i < this.subcategories.length; i++) {
        if (this.subcategories[i].id === this.searchedSubcategory) {
          this.subcategoryName = this.subcategories[i].name
          // TODO: store the image in the database with the type from the start to avoid this
          this.subcategoryImage = this.subcategories[i].image
        }
      }
    },

    getSubcategory () {
      getSubcategoriesList(this.searchedCategory).then(response => {
        this.subcategories = response.data
      })
    },

    updateSubcategory () {
      if (this.validateNewInfo()) {
        notify({
          body: 'No deben haber campos vacíos',
          type: 'alert'
        })
        return
      }

      let base64Image = document.getElementById('imageBase64').innerHTML

      if (base64Image !== '') {
        base64Image = base64Image.split('base64,')[1]
          .split('">')[0]
        this.subcategoryImage = 'data:image/png;base64,' + base64Image
      }

      updateSubcategory(this.searchedSubcategory, this.searchedCategory, this.subcategoryName, this.subcategoryImage)
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
            body: 'Ocurrió un error :(',
            type: 'error'
          })
        })
    },

    deleteSubCategory () {
      deleteSubCategory(this.searchedSubcategory).then(() => {
        notify({
          body: 'La subcategoría ha sido eliminada',
          type: 'success'
        })
        this.selectionScreen = true
        this.searchedSubcategory = ''
        this.subcategoryName = ''
        this.subcategoryImage = ''
        getSubcategoriesList(this.searchedCategory).then(response => {
          this.subcategories = response.data
        })
      })
        .catch(function (error) {
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
  watch: {
    searchedCategory: function (val) {
      return getSubcategoriesList(this.searchedCategory).then(response => {
        this.subcategories = response.data
      })
    }
  }
}
</script>

<style scoped>
.form-search-subcategory {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 95vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
