<template>
  <div class="map">
    <map-component
      mapType="list"
      :substations="substations"
      :transformers="transformers"
      :showAddressInput="false"
    ></map-component>
  </div>
</template>
<script>
import MapComponent from "@/components/map/MapComponent.vue";
import { mapActions, mapState } from "vuex";
export default {
  components: {
    MapComponent
  },
  data() {
    return {};
  },
  created() {
    this.listSubstations();
  },
  computed: {
    ...mapState("substation", ["substations", "transformers"])
  },
  methods: {
    ...mapActions("substation", ["listSubstations"])
  }
};
</script>
<style scoped>
</style>