<template>
  <div class="content">
    <div class="user-management">
      <div class="management-form">
        <md-button
          class="md-info"
          @click="showClientsOverview"
          v-if="clientSelected"
        >
          <md-icon>navigate_before</md-icon>
          Clients
          <md-tooltip md-direction="top">Back to clients overview</md-tooltip>
        </md-button>
        <create-client-form
          v-if="!clientSelected"
          data-background-color="purple"
        ></create-client-form>
        <search-client-form
          v-if="!clientSelected"
          v-on:client-details="showClientDetails"
          data-background-color="purple"
        ></search-client-form>
        <client-details
          v-if="clientSelected"
          data-background-color="purple"
          :clientID="clientID"
        >
        </client-details>
        <search-contract-form
          v-if="clientSelected"
          data-background-color="purple"
          :clientlientID="clientID"
        ></search-contract-form>
      </div>
    </div>
  </div>
</template>

<script>
import { CreateClientForm, SearchClientForm, ClientDetails } from "@/pages";

export default {
  components: {
    CreateClientForm,
    SearchClientForm,
    ClientDetails
  },

  data() {
    return {
      clientSelected: false,
      clientID: ""
    };
  },

  mounted: function() {
    this.$store.dispatch("clientsManagement/loadAllClients");
  },

  methods: {
    showClientDetails: function(clientID) {
      this.clientSelected = true;
      this.clientID = clientID;
    },

    showClientsOverview: function() {
      this.clientSelected = false;
    }
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
