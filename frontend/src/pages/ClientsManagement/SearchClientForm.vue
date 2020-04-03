<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Search Client</h4>
      </md-card-header>
      <div class="md-layout-item md-small-size-100 md-size-33 search-card">
        <md-field>
          <label>ID number</label>
          <md-input v-model="idNumber" type="text"></md-input>
        </md-field>
      </div>
      <md-card-content>
        <clients-table
          :clientsList="clientsFiltered"
          table-header-color="purple"
          v-on:client-details="client_id => $emit('client-details', client_id)"
        ></clients-table>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { ClientsTable } from "@/components";
import { mapState } from "vuex";

export default {
  name: "search-client-form",

  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    }
  },

  components: {
    ClientsTable
  },

  data() {
    return {
      idNumber: ""
    };
  },

  computed: {
    ...mapState({
      clients: state => state.clientsManagement.clients
    }),

    clientsFiltered: function() {
      if (this.idNumber === "") {
        return this.clients;
      }

      let newClientsList = [];
      for (let i = 0; i < this.clients.length; i++) {
        let client = this.clients[i];

        if (client.document_id.includes(this.idNumber)) {
          newClientsList.push(client);
        }
      }

      return newClientsList;
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
