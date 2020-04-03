<template>
  <div class="content">
    <div class="user-management">
      <div class="management-form">
        <create-client-form data-background-color="purple"></create-client-form>
        <search-client-form data-background-color="purple"></search-client-form>
      </div>
    </div>
  </div>
</template>

<script>
import { CreateClientForm, SearchClientForm } from "@/pages";

export default {
  components: {
    CreateClientForm,
    SearchClientForm
  },

  mounted: function() {
    this.$store.dispatch("clientsManagement/loadAllClients");
  }
};
</script>

<style scoped>
.management-form {
  display: block;
  margin: 0 auto;
  width: 90%;
}
</style>
