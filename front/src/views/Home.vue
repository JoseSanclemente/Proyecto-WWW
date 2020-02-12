<template>
  <div class="home">
    <!-- TODO: make this image responsive -->
    <img alt="logo" src="../assets/images/logo2.png">
  </div>
</template>

<script>
export default {
  name: 'home'
}
</script>

<style>
.home {
  background-image: url("../assets/images/background_palm_tree.jpeg");
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  font-size: 18px;
}
</style>
