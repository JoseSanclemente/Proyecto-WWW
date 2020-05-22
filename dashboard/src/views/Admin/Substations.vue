<template>
  <div>
    <section>
      <substation-form :modify="substationToModify" v-on:edit-cancelled="addView()"></substation-form>
    </section>
    <section>
      <substation-table v-on:loadModifyForm="loadModifyForm" v-if="showTable"></substation-table>
    </section>
  </div>
</template>

<script>
import SubstationTable from "@/components/substation/SubstationTable.vue";
import SubstationForm from "@/components/substation/SubstationForm.vue";
import { mapActions, mapState } from "vuex";

export default {
  name: "Substations",
  components: {
    SubstationForm,
    SubstationTable
  },
  data() {
    return {
      substationToModify: null,
      showTable: true,
    }
  },
  created() {
    this.listSubstations();
  },
  computed: {
    ...mapState("substation", ["substations"])
  },
  methods: {
    ...mapActions("substation", ["listSubstations"]),
    loadModifyForm(payload) {
      this.substationToModify = payload
      this.showTable = false
    },

    addView() {
      this.substationToModify = null
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