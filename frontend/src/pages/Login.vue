<template>
  <form class="md-layout-item md-size-33">
    <md-card>
      <md-card-content>
        <div>
          <h4 class="title">Log In</h4>
          <div>
            <md-field>
              <label>Username</label>
              <md-input v-model="username" type="text"></md-input>
            </md-field>
          </div>
          <div>
            <md-field>
              <label>Password</label>
              <md-input v-model="password" type="password"></md-input>
            </md-field>
          </div>
          <div>
            <md-button class="md-raised md-primary" @click="login">Log In</md-button>
          </div>
        </div>
      </md-card-content>
    </md-card>
  </form>
</template>
<script>
import axios from "axios";

export default {
  name: "login",
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
    login() {
      console.log("Hola");
      this.$router.push("user-management");
    }
  }
};
</script>
