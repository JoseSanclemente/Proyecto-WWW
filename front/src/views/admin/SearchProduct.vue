<template>
  <div>
    <div class="form-add-category" v-if="selectionScreen">
      <router-link to="/admin" class="back">Regresar al menú principal</router-link>
      <div class="form-background">
        <form class="form">
          <div class="form-title">Consultar producto</div>
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
            <label class="form-input-label" for="subcategory">Producto</label>
            <div class="combo-box">
              <select v-model="productID">
                <option
                  v-for="product in products"
                  v-bind:key="product.id"
                  :value="product.id"
                >{{ product.name }}</option>
              </select>
            </div>
          </div>

          <div class="confirmation-button" @click="selectProduct">Seleccionar</div>
        </form>
      </div>
    </div>

    <div class="form-add-category" v-if="!selectionScreen">
      <div class="form-background">
        <form class="form">
          <div class="form-back">
            <div class="back-button" @click="back">Regresar</div>
          </div>
          <div class="form-title">Detalles</div>
          <div class="category-container">
            <img class="category-pic" :src="this.image" alt />
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
            <input class="form-input-single-line" type="file" ref="fileupload" @change="onFileChange"/>
          </div>

          <div class="form-input">
            <div class="confirmation-button" @click="updateProduct">Modificar datos</div>
            <div class="dangerous-button" @click="deleteProduct">Eliminar</div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import {
  getCategoriesList,
  getSubcategoriesList,
  getSubcategoryProductsList,
  deleteProduct,
  updateProduct
} from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      selectionScreen: true,
      name: '',
      description: '',
      price: '',
      categories: [],
      subcategories: [],
      products: [],
      categoryID: '',
      subcategoryID: '',
      productID: '',
      quantity: '',
      image: ''
    }
  },

  methods: {
    back () {
      this.selectionScreen = true
      this.categoryID = ''
      this.subcategoryID = ''
      this.productID = ''
    },

    selectProduct () {
      if (this.validateInputSelectProduct()) {
        return
      }

      this.selectionScreen = false

      for (let i = 0; i < this.products.length; i++) {
        if (this.products[i].id === this.productID) {
          this.name = this.products[i].name
          this.price = this.products[i].price
          this.description = this.products[i].description
          this.quantity = this.products[i].quantity
          this.image = this.products[i].image
        }
      }
    },

    validateInputSelectProduct () {
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

      if (this.productID === '') {
        notify({
          body: 'Debe seleccionar un producto',
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

    updateProduct () {
      let base64Image = document
        .getElementById('imageBase64').innerHTML

      if (base64Image !== '') {
        base64Image = base64Image.split('base64,')[1]
          .split('">')[0]
        this.image = 'data:image/png;base64,' + base64Image
      }

      updateProduct({
        id: this.productID,
        subcategory: this.subcategoryID,
        name: this.name,
        description: this.description,
        price: this.price,
        quantity: this.quantity,
        image: this.image
      })
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

    deleteProduct () {
      deleteProduct(this.productID).then(() => {
        notify({
          body: 'El producto ha sido eliminado',
          type: 'success'
        })
        this.selectionScreen = true
        this.name = ''
        this.description = ''
        this.price = ''
        this.quantity = ''
        this.image = ''
        this.productID = ''
        getSubcategoryProductsList(this.subcategoryID)
          .then(response => {
            this.products = response.data
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
    categoryID: function (val) {
      return getSubcategoriesList(this.categoryID).then(response => {
        this.subcategories = response.data
      })
    },

    subcategoryID: function (val) {
      return getSubcategoryProductsList(this.subcategoryID)
        .then(response => {
          this.products = response.data
        })
        .catch(error => {
          console.log(error)
          notify({
            body:
              'No se pueden cargar los productos de la subcategoría indicada',
            type: 'error'
          })
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
  height: 125vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}
</style>
