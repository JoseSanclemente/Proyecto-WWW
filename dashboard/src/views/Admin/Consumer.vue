<template>
  <div>
    <section>
      <consumer-form modalType="create" v-model="modalOpen"></consumer-form>
      <md-button class="md-primary md-raised" @click="openModal">Add Consumer</md-button>
    </section>
    <section>
      <consumer-table></consumer-table>
    </section>
  </div>
</template>

<script>
import ConsumerTable from "@/components/consumer/ConsumerTable.vue";
import ConsumerForm from "@/components/consumer/ConsumerForm.vue";

export default {
  name: "clients",
  data() {
    return {
      modalOpen: false
    };
  },
  components: {
    ConsumerForm,
    ConsumerTable
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