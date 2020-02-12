<template>
  <div id="app">
    <navbar />
    <router-view />
    <VueToastify eventHandler="eventBus"/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue'
import VueToastify from 'vue-toastify'

export default {
  components: {
    Navbar,
    VueToastify
  }
}
</script>

<style>
@import url("./assets/css/styles.css");
</style>
