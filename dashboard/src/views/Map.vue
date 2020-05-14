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
import { mapActions, mapState, mapMutations } from "vuex";
export default {
  components: {
    MapComponent
  },
  data() {
    return {};
  },
  created() {
    this.listSubstations();
    this.setTransformers(this.substations);
  },
  computed: {
    ...mapState("substation", ["substations"]),
    ...mapState("transformer", ["transformers"])
  },
  methods: {
    ...mapMutations("transformer", ["setTransformers"]),
    ...mapActions("substation", ["listSubstations"])
  }
};
</script>
<style scoped>
</style>