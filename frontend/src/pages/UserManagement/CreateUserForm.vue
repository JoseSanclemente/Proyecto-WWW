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
          <div class="md-layout-item md-small-size-100 md-size-33">
            <md-field>
              <label>Name</label>
              <md-input v-model="name" type="email"></md-input>
            </md-field>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-50">
            <md-field>
              <label>Date of birth</label>
              <md-input v-model="dateOfBirth" type="text"></md-input>
            </md-field>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-50">
            <md-field>
              <label>Email</label>
              <md-input v-model="email" type="text"></md-input>
            </md-field>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-100">
            <md-field>
              <label>Adress</label>
              <md-input v-model="address" type="text"></md-input>
            </md-field>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-33">
            <md-checkbox v-model="active">Active</md-checkbox>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-33">
            <md-field>
              <label for="userType">Role</label>
              <md-select v-model="role" name="movie" id="movie">
                <md-option value="admin">Administrator</md-option>
                <md-option value="manager">Manager</md-option>
                <md-option value="operator">Operator</md-option>
              </md-select>
            </md-field>
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
import { USER_CREATE } from "@/store/user";
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
      name: "luis",
      dateOfBirth: "1988-01-01",
      email: "lamnrique@truora.com",
      address: "f",
      active: false,
      role: "f"
    };
  },
  methods: {
    createUser() {
      let user = {
        national_id: this.nationalID,
        name: this.name,
        date_of_birth: this.dateOfBirth,
        email: this.email,
        address: this.address,
        active: this.active,
        role: this.role
      };

      this.$store.dispatch(USER_CREATE, user);
    }
  }
};
</script>
<style></style>
