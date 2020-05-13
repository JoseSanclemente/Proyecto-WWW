<template>
  <div class="map">
    <map-component mapType="list" :showAddressInput="false"></map-component>
  </div>
</template>
<script>
import MapComponent from "@/components/map/MapComponent.vue";
import { mapActions } from "vuex";
export default {
  components: {
    MapComponent
  },
  data() {
    return {
      substationList: []
    };
  },
  beforeMount() {
    this.listSubstations();
  },
  methods: {
    ...mapActions("substation", ["listSubstations"])
  }
};
</script>
<style scoped>
</style>