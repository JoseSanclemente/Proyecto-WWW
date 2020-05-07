<template>
  <div>
    <div>
      <div class="section">
        <user-create-dialog></user-create-dialog>
      </div>
      <div class="section">
        <users-table data-background-color="purple"></users-table>
      </div>
    </div>
  </div>
</template>

<script>
import UserCreateDialog from "@/components/UserManagement/UserCreateDialog.vue";
import UsersTable from "@/components/UserManagement/UsersTable.vue";

export default {
  data() {
    return {
      showDialog: false
    };
  },
  components: {
    UsersTable,
    UserCreateDialog
  }
};
</script>
<style scoped>
.section {
  margin: 1em;
}
</style>