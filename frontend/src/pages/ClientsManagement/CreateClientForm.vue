<template>
  <form>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Create client</h4>
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
              <label>Date of birth</label>
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
              <label>Phone number</label>
              <md-input v-model="phone" type="number"></md-input>
            </md-field>
          </div>
          <div class="md-layout-item md-small-size-100 md-size-50">
            <md-checkbox v-model="active">Active</md-checkbox>
          </div>
          <div class="md-layout-item md-size-100 text-right">
            <md-button class="md-raised md-primary" @click="createClient"
              >Create Client</md-button
            >
          </div>
        </div>
      </md-card-content>
    </md-card>
  </form>
</template>
<script>
import clients from "@/core/services/clients.js";

export default {
  name: "create-client-form",
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
      phone: ""
    };
  },

  // TODO: Create a function to validate and notiy events
  methods: {
    createClient() {
      let payload = {
        document_id: this.nationalID,
        email: this.email,
        name: this.name,
        phone_number: this.phone,
        address: this.address,
        active: this.active,
        date_of_birth: this.dateOfBirth
      };

      clients
        .create(payload)
        .then(response => {
          console.log(response);
          this.$notify({
            message: "Client created correctly",
            horizontalAlign: "right",
            verticalAlign: "top",
            type: "info"
          });

          this.nationalID = "";
          this.email = "";
          this.name = "";
          this.phone = "";
          this.address = "";
          this.active = false;
          this.dateOfBirth = "";

          this.$store.dispatch("clientsManagement/loadAllClients");
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
