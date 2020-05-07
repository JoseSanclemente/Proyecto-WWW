<template>
  <default-layout>
    <template v-slot:header>
      <span class="md-title">Client Management</span>
    </template>
    <div>
      <section>
        <client-form v-model="modalOpen"></client-form>
        <md-button class="md-primary md-raised" @click="openModal" :disabled="sending">Add Client</md-button>
      </section>
      <section>
        <client-table></client-table>
      </section>
    </div>
  </default-layout>
</template>

<script>
import DefaultLayout from "@/layout/DefaultLayout.vue";
import ClientTable from "@/components/ClientTable.vue";
import ClientForm from "@/components/ClientForm.vue";

export default {
  name: "Clients",
  data() {
    return {
      modalOpen: false
    };
  },
  components: {
    DefaultLayout,
    ClientForm,
    ClientTable
  },
  methods: {
    openModal() {
      this.modalOpen = !this.modalOpen;
    }
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>