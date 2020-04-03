<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Search Contract</h4>
      </md-card-header>
      <div class="md-layout-item md-small-size-100 md-size-33 search-card">
        <md-field>
          <label>Predial</label>
          <md-input v-model="predialNumber" type="text"></md-input>
        </md-field>
      </div>
      <md-card-content>
        <contracts-table
          :contractsList="contractsFiltered"
          table-header-color="purple"
          v-on:contract-details="
            contractID => $emit('transformer-details', contractID)
          "
        ></contracts-table>
        <md-button
          class="md-info"
          @click="showClientsOverview"
          v-if="clientSelected"
        >
          <md-icon>navigate_before</md-icon>Clients
          <md-tooltip md-direction="top">Back to clients overview</md-tooltip>
        </md-button>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { ContractsTable } from "@/components";
import { mapState } from "vuex";

export default {
  name: "search-contract-form",

  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    },

    selectedClientID: {
      type: String,
      default: ""
    }
  },

  components: {
    ContractsTable
  },

  data() {
    return {
      predialNumber: ""
    };
  },

  computed: {
    ...mapState({
      contracts: state => state.clientsManagement.contracts
    }),

    contractsFiltered: function() {
      let newContractsList = [];
      for (let i = 0; i < this.contracts.length; i++) {
        let contract = this.contracts[i];

        let normPredialNumber = contract.predial
          .toString()
          .split(" ")
          .join("")
          .toUpperCase();

        let normSearchTerm = this.predialNumber
          .split(" ")
          .join("")
          .toUpperCase();

        if (
          normPredialNumber.includes(normSearchTerm) &&
          contract.client_id === this.selectedClientID
        ) {
          newContractsList.push(contract);
        }
      }

      return newContractsList;
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
