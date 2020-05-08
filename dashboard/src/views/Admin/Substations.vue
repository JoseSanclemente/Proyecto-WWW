<template>
  <div>
    <section>
      <substation-form></substation-form>
    </section>
    <section>
      <substation-table></substation-table>
    </section>
  </div>
</template>

<script>
import SubstationTable from "@/components/SubstationTable.vue";
import SubstationForm from "@/components/SubstationForm.vue";

export default {
  name: "Substations",
  components: {
    SubstationForm,
    SubstationTable
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>