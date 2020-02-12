<template>
  <div class="page">
    <div class="container">
      <div v-for="product in products" :key="product.id" class="category-container">
        <h2>{{ product.name }}</h2>
        <img class="category-images" :src="product.image" :alt="product.name" @click="goToProduct(product.id)"/>
      </div>
    </div>
  </div>
</template>

<script>
import { getSubcategoryProductsList } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      products: []
    }
  },

  methods: {
    goToProduct (productID) {
      this.$router.push({
        name: 'product',
        params: {
          category_id: this.$route.params.category_id,
          subcategory_id: this.$route.params.subcategory_id,
          product_id: productID
        }
      })
    }
  },

  mounted () {
    getSubcategoryProductsList(this.$route.params.subcategory_id)
      .then(response => {
        console.log(response)
        this.products = response.data
      })
      .catch(error => {
        console.log(error)
        notify({
          body:
            'No se pueden cargar los productos de la sub-categoría indicada',
          type: 'error'
        })
      })
  }
}
</script>

<style scoped>
.page {
  display: flex;
  height: 120vh;
  background-image: url("../../assets/images/background_palm_tree.jpeg");
}

.sidebar {
  background: rgba(224, 224, 224, 0.2);
  color: white;
  display: flex;
  flex-direction: column;
  width: 10vw;
}

.container {
  background: rgba(0, 0, 0, 0.4);
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
  align-items: flex-start;
  width: 80vw;
  margin: 50px auto;
  border-radius: 20px;
  border: 1px solid darkgray;
}

.category-container {
  color: white;
  margin: 40px 20px 20px 20px;
  border-radius: 20px;
  width: 20vw;
  height: 30vh;
}

.category-images {
  display: flex;
  width: 20vw;
  height: 30vh;
  border-radius: 10px;
  border: 1px solid white;
}
</style>
