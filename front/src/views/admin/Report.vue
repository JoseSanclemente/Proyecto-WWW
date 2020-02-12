<template>
  <div class="form-profile">
    <router-link to="/admin" class="back">Regresar al menú principal</router-link>
    <div class="form-background">
        <form class="form">
            <div class="form-title">Reportes</div>
            <div class="form-input">
                 <div class="form-input-label">Tipo de Reporte</div>
                <div class="combo-box">
                    <select  @change="showInputs()" v-model="reportType">
                        <option value="top_20_more_sold">Top 20 Productos más vendidos</option>
                        <option value="top_20_less_sold">Top 20 Productos menos vendidos</option>
                        <option value="sales_from_date">Total de ventas</option>
                        <option value="total_sales_product">Total de ventas de un producto</option>
                        <option value="low_quantity_products">Productos con menos existencias</option>
                        <option value="loyal_users">Top mejores clientes</option>
                        <option value="next_users_birthdays">Pròximos cumpleaños de clientes</option>
                    </select>
                </div>
            </div>

            <div class="form-input" v-if="salesFromDate">
                <div class="form-input-label">Desde:</div>
                <datepicker
                placeholder="Seleccionar fecha"
                input-class="form-input-single-line"
                v-model="from"
                ></datepicker>
            </div>

            <div class="form-input" v-if="salesFromDate">
                <div class="form-input-label">Hasta:</div>
                <datepicker
                placeholder="Seleccionar fecha"
                input-class="form-input-single-line"
                v-model="to"
                ></datepicker>
            </div>

            <div class="form-input" v-if="productSales">
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

            <div class="form-input" v-if="productSales">
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

            <div class="form-input" v-if="productSales">
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

            <div class="confirmation-button" @click="generateReport">Generar Reporte</div>

            <div>
              <img :src="this.image" alt />
            </div>

            <div class="form-table" v-if="productTable">
                <table>
                  <tr>
                    <th>Nombre</th>
                    <th>Existencias</th>
                  </tr>
                  <tr
                  v-for="item in list"
                  v-bind:key="item.id"
                  :value="item.id"
                  >
                    <th>{{ item.label }} </th>
                    <th>{{ item.value }} </th>
                  </tr>
                </table>
            </div>

            <div class="form-table" v-if="birthdayTable">
                <table>
                  <tr>
                    <th>Correo electrónico</th>
                    <th>Fecha de cumpleaños</th>
                  </tr>
                  <tr
                  v-for="item in list"
                  v-bind:key="item.id"
                  :value="item.id"
                  >
                    <th>{{ item.label }} </th>
                    <th>{{ item.value }} </th>
                  </tr>
                </table>
            </div>
        </form>
    </div>
  </div>
</template>

<script>
import {
  generateReport,
  getCategoriesList,
  getSubcategoriesList,
  getSubcategoryProductsList,
  generateDateReport,
  generateSalesProductReport
} from '@/api/api.js'
import Datepicker from 'vuejs-datepicker/dist/vuejs-datepicker.esm.js'
import { notify } from '@/notify.js'

export default {
  components: {
    Datepicker
  },
  data: function () {
    return {
      salesFromDate: false,
      productSales: false,
      productTable: false,
      birthdayTable: false,
      reportType: '',
      to: '',
      from: '',
      categoryID: '',
      subcategoryID: '',
      productID: '',
      categories: [],
      subcategories: [],
      products: [],
      list: [],
      image: ''
    }
  },

  methods: {
    showInputs: function () {
      this.birthdayTable = false
      this.productTable = false
      this.image = ''
      if (this.reportType === 'sales_from_date') {
        this.salesFromDate = true
        this.productSales = false
        return
      }

      if (this.reportType === 'total_sales_product') {
        this.productSales = true
        this.salesFromDate = false
        this.categoryID = ''
        this.subcategoryID = ''
        this.productID = ''
        return
      }

      this.productSales = false
      this.salesFromDate = false
    },

    generateReport () {
      if (this.reportType === 'sales_from_date') {
        if(this.to === '' || this.from === '') {
          notify({
            body: "Debe seleccionar un rango de fechas",
            type: "alert"
          })
          return
        }
        generateDateReport(this.reportType, this.to, this.from)
          .then((response) => {
            if (response.data.info === 'No records found') {
              notify({
                body: 'No se encontraron datos',
                type: 'alert'
              })
              return
            }
            this.image = 'data:image/png;base64,' + response.data.image
          }).catch((error) => {
            console.log(error)
            notify({
              body: 'Ha ocurrido un error :(',
              type: 'error'
            })
          })
        return
      }

      if (this.reportType === 'total_sales_product') {
        if(this.categoryID === '' || this.subcategoryID === '' || this.productID === ''){
          notify({
            body: "Debe seleccionar un producto",
            type: "alert"
          })
          return
        }
        return generateSalesProductReport(this.reportType, this.productID)
          .then((response) => {
            if (response.data.info === 'No records found') {
              notify({
                body: 'No se encontraron datos',
                type: 'alert'
              })
              return
            }
            this.image = 'data:image/png;base64,' + response.data.image
          }).catch((error) => {
            console.log(error)
            notify({
              body: 'Ha ocurrido un error :(',
              type: 'error'
            })
          })
      }

      if (this.reportType === 'low_quantity_products') {
        generateReport(this.reportType)
          .then((response) => {
            if (response.data.info === 'No records found') {
              notify({
                body: 'No se encontraron datos',
                type: 'alert'
              })
              return
            }
            this.list = response.data.result
            this.productTable = true
          }).catch((error) => {
            console.log(error)
            notify({
              body: 'Ha ocurrido un error :(',
              type: 'error'
            })
          })
        return
      }

      if (this.reportType === 'next_users_birthdays') {
        generateReport(this.reportType)
          .then((response) => {
            if (response.data.info === 'No records found') {
              notify({
                body: 'No se encontraron datos',
                type: 'alert'
              })
              return
            }
            this.list = response.data.result
            this.birthdayTable = true
          }).catch((error) => {
            console.log(error)
            notify({
              body: 'Ha ocurrido un error :(',
              type: 'error'
            })
          })
        return
      }

      generateReport(this.reportType)
        .then((response) => {
          if (response.data.info === 'No records found') {
            notify({
              body: 'No se encontraron datos',
              type: 'alert'
            })
            return
          }

          if (response.data.image !== null) {
            this.image = 'data:image/png;base64,' + response.data.image
          }
        }).catch((error) => {
          console.log(error)
          notify({
            body: 'Ha ocurrido un error :(',
            type: 'error'
          })
        })
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
.form-profile {
  display: flex;
  align-items: center;
  flex-direction: column;
  height: 150vh;
  font-size: 18px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.form-container {
  display: flex;
  width: 60vw;
  height: 60vh;
  flex-direction: row;
}

.form-image{
  margin: 20px;
  border: 1px solid white;
  border-radius: 100px;
  height: 150px;
  width: 150px;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.form-data{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 30px 30px 5px 30px;
  width: 45vw;
  -webkit-border-radius: 0 20px 20px 0;
  border-radius: 0 20px 20px 0;
}

.form-table{
    border-radius: 10px;
    border: 1px solid white;
}

th {
  width: 100%;
  color: white;
  border: 1px solid white;
}

tr {
  width: 100%;
  color: white;
  border: 1px solid white;
}
</style>
