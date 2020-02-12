<template>
  <div class="page">
    <div class="container">
      <div v-for="category in categories" :key="category.id" class="category-container">
        <h2>{{ category.name }}</h2>
        <img
          class="category-images"
          :src="category.image"
          :alt="category.name"
          @click="goToCategory(category.category_id)"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { getCategoriesList } from '@/api/api.js'

export default {
  data: function () {
    return {
      categories: []
    }
  },

  methods: {
    goToCategory (categoryID) {
      this.$router.push({
        name: 'category-subcategories',
        params: { category_id: categoryID }
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
