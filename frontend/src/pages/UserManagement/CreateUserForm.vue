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
import moment from "moment";
import { USER_CREATE } from "@/store/user";

const DATE_FORMAT = "YYYY-MM-DD";
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
      nationalID: "121212",
      name: "John Doe",
      dateOfBirth: "1988-12-10",
      email: "john@doe.com",
      address: "f",
      active: false,
      role: "f"
    };
  },

  // TODO: Create a function to validate and notiy events
  methods: {
    reset() {
      this.nationalID = "";
      this.name = "";
      this.dateOfBirth = moment().format(DATE_FORMAT);
      this.email = "";
      this.address = "";
      this.active = "";
      this.role = "";
    },

    createUser() {
      let user = {
        national_id: this.nationalID,
        name: this.name,
        date_of_birth: moment(this.dateOfBirth).format(DATE_FORMAT),
        email: this.email,
        address: this.address,
        active: this.active,
        role: this.role
      };
      this.$store
        .dispatch(USER_CREATE, user)
        .then(response => {
          this.reset();
          this.$notify({
            message: "User created",
            icon: "add_alert",
            horizontalAlign: "right",
            verticalAlign: "top",
            type: "success"
          });
        })
        .catch(error => {
          console.log(error.response);
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
