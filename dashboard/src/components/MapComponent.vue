<template>
  <div>
    <GmapMap
      :center="center"
      :map-type-id="mapTypeId"
      :zoom="15"
      :options="{
   scrollwheel: false,
 }"
    >
      <GmapMarker
        v-for="(item, index) in markers"
        :key="index"
        :position="item.position"
        @click="getMarkerInfo(item.position)"
      />
    </GmapMap>
  </div>
</template>

<script>
export default {
  data() {
    return {
      center: { lat: 3.375588, lng: -76.533017 },
      mapTypeId: "roadmap",
      markers: [{ position: { lat: 3.375588, lng: -76.533017 } }]
    };
  },
  methods: {
    getMarkerInfo(position) {
      console.log(position);
    }
  }
};
</script>

<style lang="scss" scoped>
.vue-map-container {
  height: 70vh;
  max-width: 992px;
  width: 100%;
}
</style>