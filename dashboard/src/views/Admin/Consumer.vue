<template>
  <div>
    <section>
      <consumer-form modalType="create" :transformers="transformers" v-model="modalOpen"></consumer-form>
      <md-button class="md-primary md-raised" @click="openModal">{{$t('Add Consumer')}}</md-button>
    </section>
    <section>
      <consumer-table></consumer-table>
    </section>
  </div>
</template>

<script>
import ConsumerTable from "@/components/consumer/ConsumerTable.vue";
import ConsumerForm from "@/components/consumer/ConsumerForm.vue";
import { mapState } from "vuex";

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
  computed: {
    ...mapState("transformer", ["transformers"])
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