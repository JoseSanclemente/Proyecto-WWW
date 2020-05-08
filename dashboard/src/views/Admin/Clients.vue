<template>
  <div>
    <section>
      <client-form v-model="modalOpen"></client-form>
      <md-button class="md-primary md-raised" @click="openModal" :disabled="sending">Add Client</md-button>
    </section>
    <section>
      <client-table></client-table>
    </section>
  </div>
</template>

<script>
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