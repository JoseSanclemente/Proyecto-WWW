<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Client Details</h4>
      </md-card-header>
      <md-card-content>
        <h4>Client id</h4>
        {{ clientID }}
        <br />
        <h4>Document ID</h4>
        {{ documentID }}
        <br />
        <h4>Email</h4>
        {{ email }}
        <br />
        <h4>Phone number</h4>
        {{ phone }}
        <br />
        <h4>Address</h4>
        {{ address }}
        <br />
        <h4>Date of birth</h4>
        {{ dateOfBirth }}
        <br />
        <h4>Active</h4>
        {{ clientActive }}
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "substation-details",

  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    },

    clientID: {
      type: String,
      default: "default"
    }
  },

  methods: {
    getFromClients(property) {
      if (this.clients === undefined) return "";

      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (this.clientID === client.client_id) {
          return client[property];
        }
      }
      return "";
    }
  },

  computed: {
    ...mapState({
      clients: state => state.clientsManagement.clients
    }),

    documentID: function() {
      if (this.clients === undefined) return "";

      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (this.clientID === client.client_id) {
          return client.document_id;
        }
      }
      return "";
    },
    email: function() {
      if (this.clients === undefined) return "";

      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (this.clientID === client.client_id) {
          return client.email;
        }
      }
      return "";
    },
    phone: function() {
      if (this.clients === undefined) return "";

      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (this.clientID === client.client_id) {
          return client.phone_number;
        }
      }
      return "";
    },
    address: function() {
      if (this.clients === undefined) return "";

      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (this.clientID === client.client_id) {
          return client.address;
        }
      }
      return "";
    },
    dateOfBirth: function() {
      if (this.clients === undefined) return "";

      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (this.clientID === client.client_id) {
          return client.date_of_birth;
        }
      }
      return "";
    },
    clientActive: function() {
      if (this.clients === undefined) return false;

      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (this.clientID === client.client_id) {
          return client.active;
        }
      }
      return "";
    }
  }
};
</script>

<style scoped>
.search-card {
  display: flex;
  margin: 2rem 0 2rem 1rem;
}

.search-button {
  margin: 1.5rem 0 0 2rem;
}
</style>
