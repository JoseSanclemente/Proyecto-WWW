<template>
  <default-layout>
    <template v-slot:header>
      <span class="md-title">Substations</span>
    </template>
    <div>
      <section>
        <transformer-form></transformer-form>
      </section>
      <section>
        <substation-table></substation-table>
      </section>
    </div>
  </default-layout>
</template>

<script>
import DefaultLayout from "@/layout/DefaultLayout.vue";
import SubstationTable from "@/components/SubstationTable.vue";
import TransformerForm from "@/components/TransformerForm.vue";

export default {
  name: "Transformer",
  components: {
    DefaultLayout,
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