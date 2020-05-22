<template>
  <div>
    <section id="top-transformers">
      <transformer-form
        :substations="substations"
        :modify="transformerToModify"
        v-on:edit-cancelled="addView()"
      ></transformer-form>
    </section>
    <section v-if="showTable">
      <transformer-table v-on:loadModifyForm="loadModifyForm"></transformer-table>
    </section>
    <!--md-button v-scroll-to="'transformer-form'">
      scroll
    </md-button-->
  </div>
</template>

<script>
import TransformerForm from "@/components/transformer/TransformerForm.vue";
import TransformerTable from "@/components/transformer/TransformerTable.vue";
import { mapActions, mapState } from "vuex";

export default {
  name: "Transformer",
  components: {
    TransformerForm,
    TransformerTable,
  },

  data() {
    return {
      transformerToModify: null,
      showTable: true
    }
  },

  beforeMount() {
    this.listSubstations();
  },
  computed: {
    ...mapState("substation", ["substations"])
  },
  methods: {
    ...mapActions("substation", ["listSubstations"]),

    loadModifyForm(payload) {
      this.transformerToModify = payload
      this.showTable = false
    },

    addView() {
      this.transformerToModify = null
      this.showTable = true
    }
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>