<template>
  <div>
    <consumer-login></consumer-login>
  </div>
</template>

<script>
import ConsumerLogin from "@/components/consumer/ConsumerLogin.vue";
import { mapActions } from "vuex";

export default {
  name: "login",
  components: {
    ConsumerLogin
  },
  beforeMount() {
    this.listConsumers();
  },
  methods: {
    ...mapActions("consumer", ["listConsumers"])
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>