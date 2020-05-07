<template>
  <default-layout>
    <template v-slot:header>
      <span class="md-title">Substations</span>
    </template>
    <div>
      <section>
        <substation-form></substation-form>
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
import SubstationForm from "@/components/SubstationForm.vue";

export default {
  name: "Substations",
  components: {
    DefaultLayout,
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