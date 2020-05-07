<template>
  <div>
    <md-dialog :md-active.sync="showDialog">
      <md-dialog-title>Add User</md-dialog-title>
      <div class="md-layout">
        <div class="md-layout-item md-small-size-100 md-size-33">
          <md-field>
            <label>National ID</label>
            <md-input v-model="nationalID" type="text"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item md-small-size-100 md-size-50">
          <md-field>
            <label>Name</label>
            <md-input v-model="name" type="text"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item md-small-size-100 md-size-33">
          <md-datepicker v-model="dateOfBirth">
            <label>Select date</label>
          </md-datepicker>
        </div>
        <div class="md-layout-item md-small-size-100 md-size-50">
          <md-field>
            <label>Email</label>
            <md-input v-model="email" type="email"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item md-small-size-100 md-size-33">
          <md-field>
            <label>Address</label>
            <md-input v-model="address" type="text"></md-input>
          </md-field>
        </div>
        <div class="md-layout-item md-small-size-100 md-size-33">
          <md-field>
            <label for="userType">Role</label>
            <md-select v-model="role" name="movie" id="movie">
              <md-option value="Administrator">Administrator</md-option>
              <md-option value="Manager">Manager</md-option>
              <md-option value="Operator">Operator</md-option>
            </md-select>
          </md-field>
        </div>
        <div class="md-layout-item md-small-size-100 md-size-50">
          <md-checkbox v-model="active">Active</md-checkbox>
        </div>
      </div>
      <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        <md-button class="md-primary" @click="showDialog = false">Create</md-button>
      </md-dialog-actions>
    </md-dialog>

    <md-button class="md-primary md-raised" @click="showDialog = true">Add User</md-button>
  </div>
</template>
<script>
export default {
  name: "user-create-dialog",
  data() {
    return {
      showDialog: false
    };
  }
};
</script>

<style lang="scss" scoped>
.md-dialog /deep/.md-dialog-container {
  max-width: 768px;
}
</style>