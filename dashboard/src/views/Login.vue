<template>
  <div>
    <user-login></user-login>
  </div>
</template>

<script>
// import ConsumerLogin from "@/components/consumer/ConsumerLogin.vue";
import UserLogin from "@/components/user/UserLogin.vue";
//import { mapActions } from "vuex";

export default {
  name: "login",
  data() {
    return {

    };
  },
  components: {
    //ConsumerLogin,
    UserLogin
  },
  beforeMount() {

  },
  methods: {

  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>