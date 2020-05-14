<template>
  <div>
    <section>
      <transformer-form :substations="substations"></transformer-form>
    </section>
    <section></section>
  </div>
</template>

<script>
import TransformerForm from "@/components/transformer/TransformerForm.vue";
import { mapActions, mapState } from "vuex";

export default {
  name: "Transformer",
  components: {
    TransformerForm
  },
  beforeMount() {
    this.listSubstations();
  },
  computed: {
    ...mapState("substation", ["substations"])
  },
  methods: {
    ...mapActions("substation", ["listSubstations"])
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>