<template>
  <div class="map">
    <map-component :showAddressInput="false"></map-component>
  </div>
</template>
<script>
import MapComponent from "@/components/map/MapComponent.vue";
export default {
  components: {
    MapComponent
  }
};
</script>
<style scoped>
</style>