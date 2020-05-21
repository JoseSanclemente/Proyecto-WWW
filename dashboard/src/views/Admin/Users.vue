<template>
  <div>
    <section>
      <user-form v-model="modalOpen"></user-form>
      <md-button class="md-primary md-raised" @click="openModal">{{ $t('Add user') }}</md-button>
    </section>
    <section>
      <user-table></user-table>
    </section>
  </div>
</template>

<script>
import UserTable from "@/components/user/UserTable.vue";
import UserForm from "@/components/user/UserForm.vue";
import { mapActions, mapState } from "vuex";

export default {
  name: "Users",
  data() {
    return {
      modalOpen: false
    };
  },
  components: {
    UserForm,
    UserTable
  },
  computed: {
    ...mapState("user", ["users"])
  },
  created() {
    this.listUsers();
  },
  methods: {
    ...mapActions("user", ["listUsers"]),
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