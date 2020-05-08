<template>
  <default-layout>
    <div class="map">
      <map-component></map-component>
    </div>
  </default-layout>
</template>
<script>
import DefaultLayout from "@/layout/DefaultLayout.vue";
import MapComponent from "@/components/MapComponent.vue";
export default {
  components: {
    DefaultLayout,
    MapComponent
  }
};
</script>
<style scoped>
</style>