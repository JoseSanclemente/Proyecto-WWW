<template>
  <default-layout>
    <template v-slot:header>
      <span class="md-title">User Management</span>
    </template>
    <div>
      <section>
        <user-form v-model="modalOpen"></user-form>
        <md-button class="md-primary md-raised" @click="openModal" :disabled="sending">Create user</md-button>
      </section>
      <section>
        <user-table></user-table>
      </section>
    </div>
  </default-layout>
</template>

<script>
import DefaultLayout from "@/layout/DefaultLayout.vue";
import UserTable from "@/components/UserTable.vue";
import UserForm from "@/components/UserForm.vue";

export default {
  name: "Users",
  data() {
    return {
      modalOpen: false
    };
  },
  components: {
    DefaultLayout,
    UserForm,
    UserTable
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