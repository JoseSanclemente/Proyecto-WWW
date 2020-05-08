<template>
  <div>
    <section>
      <transformer-form></transformer-form>
    </section>
    <section>
      <substation-table></substation-table>
    </section>
  </div>
</template>

<script>
import SubstationTable from "@/components/SubstationTable.vue";
import TransformerForm from "@/components/TransformerForm.vue";

export default {
  name: "Transformer",
  components: {
    TransformerForm,
    SubstationTable
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>