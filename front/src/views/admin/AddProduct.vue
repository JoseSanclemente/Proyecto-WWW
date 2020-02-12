<template>
  <div class="form-add-category">
    <router-link to="/admin" class="back">Regresar al menú principal</router-link>
    <div class="form-background">
      <form class="form">
        <div class="form-title">Añadir producto</div>
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
          <label class="form-input-label" for="subcategory">Subcategoría</label>
          <div class="combo-box">
            <select v-model="subcategoryID">
              <option
                v-for="subcategory in subcategories"
                v-bind:key="subcategory.id"
                :value="subcategory.id"
              >{{ subcategory.name }}</option>
            </select>
          </div>
        </div>

        <div class="form-input">
          <label class="form-input-label" for="name">Nombre</label>
          <input class="form-input-single-line" type="text" id="name" v-model="name" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="description">Descripción</label>
          <textarea
            class="form-input-multiline"
            type="text"
            id="description"
            v-model="description"
          />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="price">Precio</label>
          <textarea class="form-input-single-line" type="text" id="price" v-model="price" />
        </div>

        <div class="form-input">
          <label class="form-input-label" for="quantity">Cantidad</label>
          <textarea class="form-input-single-line" type="text" id="quantity" v-model="quantity" />
        </div>

        <div v-show="false" id="imageBase64"></div>
        <div class="form-input">
          <label class="form-input-label" for="image">Imagen</label>
          <input class="form-input-single-line" type="file" ref="fileupload" @change="onFileChange" />
        </div>

        <div class="confirmation-button" @click="createProduct">Añadir</div>
      </form>
    </div>
  </div>
</template>

<script>
import {
  createProduct,
  getCategoriesList,
  getSubcategoriesList
} from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      name: '',
      description: '',
      price: '',
      categories: [],
      subcategories: [],
      categoryID: '',
      subcategoryID: '',
      quantity: ''
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

      if (this.subcategoryID === '') {
        notify({
          body: 'Debe seleccionar una sub-categoría',
          type: 'alert'
        })
        return true
      }

      if (this.name === '') {
        notify({
          body: 'El producto debe tener nombre',
          type: 'alert'
        })
        return true
      }

      if (this.description === '') {
        notify({
          body: 'El producto debe descripción',
          type: 'alert'
        })
        return true
      }

      if (this.price === '') {
        notify({
          body: 'El producto debe tener precio',
          type: 'alert'
        })
        return true
      }

      if (!/^\d+\.*\d*$/.test(this.price)) {
        notify({
          body: 'El precio solo puede ser un numero entero o decimal',
          type: 'alert'
        })
        return true
      }

      if (this.quantity === '') {
        notify({
          body: 'El producto debe tener cantidad',
          type: 'alert'
        })
        return true
      }

      if (!/^\d+$/.test(this.quantity)) {
        notify({
          body: 'La cantidad solo puede ser un numero entero',
          type: 'alert'
        })
        return true
      }

      if (
        document.getElementById('imageBase64').innerHTML.split('base64,')[1] ===
        undefined
      ) {
        notify({
          body: 'El producto debe tener imagen',
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

    createProduct () {
      let badInput = this.validateInput()
      if (badInput) {
        return
      }

      let base64Image = document
        .getElementById('imageBase64')
        .innerHTML.split('base64,')[1]
        .split('">')[0]

      base64Image = 'data:image/png;base64,' + base64Image
      createProduct({
        subcategoryID: this.subcategoryID,
        name: this.name,
        price: parseFloat(this.price),
        description: this.description,
        image: base64Image,
        quantity: this.quantity
      })
        .then((response) => {
          notify({
            body: 'El producto ha sido añadido',
            type: 'success'
          })
          this.categoryID = ''
          this.subcategoryID = ''
          this.name = ''
          this.price = ''
          this.description = ''
          this.image = ''
          this.quantity = ''
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

  watch: {
    categoryID: function (val) {
      return getSubcategoriesList(this.categoryID).then(response => {
        this.subcategories = response.data
      })
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
