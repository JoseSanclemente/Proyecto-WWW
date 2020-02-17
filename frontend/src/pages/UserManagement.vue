<template>
  <div class="content">
    <div class="user-management">
      <div class="management-form">
        <create-user-form data-background-color="purple"></create-user-form>
        <search-user-form data-background-color="purple"></search-user-form>
      </div>
    </div>
  </div>
</template>

<script>
import { CreateUserForm, SearchUserForm } from "@/pages";

export default {
  components: {
    CreateUserForm,
    SearchUserForm
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
