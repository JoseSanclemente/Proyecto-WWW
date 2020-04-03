<template>
  <form>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Create user</h4>
      </md-card-header>

      <md-card-content>
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
          <div class="md-layout-item md-size-100 text-right">
            <md-button class="md-raised md-primary" @click="createUser"
              >Create User</md-button
            >
          </div>
        </div>
      </md-card-content>
    </md-card>
  </form>
</template>
<script>
import axios from "axios";

export default {
  name: "create-user-form",
  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      nationalID: "",
      name: "",
      dateOfBirth: "",
      email: "",
      address: "",
      active: false,
      role: ""
    };
  },

  // TODO: Create a function to validate and notiy events
  methods: {
    createUser() {
      const url = "http://127.0.0.1:8000/employee/";
      let payload = {
        national_id: this.nationalID,
        name: this.name,
        date_of_birth: this.dateOfBirth,
        email: this.email,
        address: this.address,
        active: this.active,
        role: this.role
      };

      axios // TODO: Move this to a separate file
        .post(url, payload)
        .then(response => {
          this.nationalID = "";
          this.name = "";
          this.dateOfBirth = "";
          this.email = "";
          this.address = "";
          this.active = "";
          this.role = "";

          this.$notify({
            message: "User created",
            icon: "add_alert",
            horizontalAlign: "right",
            verticalAlign: "top",
            type: "success"
          });
        })
        .catch(error => {
          console.log(error.response.data);
          this.$notify({
            message: "Please enter all fields",
            icon: "error_outline",
            horizontalAlign: "right",
            verticalAlign: "top",
            type: "danger"
          });
        });
    }
  }
};
</script>
<style></style>
