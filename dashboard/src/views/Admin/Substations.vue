<template>
  <div>
    <section>
      <substation-form></substation-form>
    </section>
    <section>
      <substation-table :substations="substations"></substation-table>
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
  created() {
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