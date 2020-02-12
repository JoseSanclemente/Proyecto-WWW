<template>
  <div class="page">
    <div class="container">
      <div v-for="subcategory in subcategories" :key="subcategory.id" class="category-container">
        <h2>{{ subcategory.name }}</h2>
        <img
          class="category-images"
          :src="subcategory.image"
          :alt="subcategory.name"
          @click="goToSubcategory(subcategory.id)"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { getCategorySubcategoriesList } from '@/api/api.js'
import { notify } from '@/notify.js'

export default {
  data: function () {
    return {
      subcategories: []
    }
  },

  mounted () {
    getCategorySubcategoriesList(this.$route.params.category_id)
      .then(response => {
        console.log(response)
        this.subcategories = response.data
      })
      .catch(error => {
        console.log(error)
        notify({
          body:
            'No se pueden cargar las subcategorias de la categoría indicada',
          type: 'error'
        })
      })
  },

  methods: {
    goToSubcategory (subcategoryID) {
      this.$router.push({
        name: 'subcategory-products',
        params: {
          category_id: this.$route.params.category_id,
          subcategory_id: subcategoryID
        }
      })
    }
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
