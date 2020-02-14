<template>
  <div class="content">
    <div class="form">
      <div>
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
.form {
  width: 100%;
}
</style>
