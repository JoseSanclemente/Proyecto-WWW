<template>
  <div>
    <section>
      <transformer-form :substations="substations"></transformer-form>
    </section>
    <section>
      <transformer-table :transformers="transformers"></transformer-table>
    </section>
  </div>
</template>

<script>
import TransformerForm from "@/components/transformer/TransformerForm.vue";
import TransformerTable from "@/components/transformer/TransformerTable.vue";
import { mapActions, mapState, mapMutations } from "vuex";

export default {
  name: "Transformer",
  components: {
    TransformerForm,
    TransformerTable
  },
  beforeMount() {
    this.listSubstations();
    this.setTransformers(this.substations);
  },
  computed: {
    ...mapState("substation", ["substations"]),
    ...mapState("transformer", ["transformers"])
  },
  methods: {
    ...mapActions("substation", ["listSubstations"]),
    ...mapMutations("transformer", ["setTransformers"])
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>