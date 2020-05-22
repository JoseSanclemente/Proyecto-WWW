<template>
  <div>
    <section>
      <consumer-form v-model="modalOpen"></consumer-form>
      <md-button class="md-accent md-raised" @click="openModal">{{$t('Add consumer')}}</md-button>
    </section>
    <section>
      <consumer-table></consumer-table>
    </section>
  </div>
</template>

<script>
import ConsumerTable from "@/components/consumer/ConsumerTable.vue";
import ConsumerForm from "@/components/consumer/ConsumerForm.vue";
import { mapActions } from "vuex";

export default {
  name: "Consumer",
  data() {
    return {
      modalOpen: false
    };
  },
  components: {
    ConsumerForm,
    ConsumerTable
  },
  beforeMount() {
    this.listConsumers();
    this.listSubstations();
  },
  methods: {
    ...mapActions("consumer", ["listConsumers"]),
    ...mapActions("substation", ["listSubstations"]),
    openModal() {
      this.modalOpen = !this.modalOpen;
    }
  }
};
</script>

<style src="../../assets/styles/main.scss" lang="scss">
section {
  margin: 1.2em 0;
}
</style>