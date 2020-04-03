<template>
  <div>
    <md-table v-model="contractsList" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <div v-if="contractsList.length <= 0">No contracts found</div>
        <md-table-cell md-label="ID">{{ item.contract_id }}</md-table-cell>
        <md-table-cell md-label="Predial">{{ item.predial }}</md-table-cell>
        <md-table-cell md-label="Address">{{ item.address }}</md-table-cell>
        <md-table-cell md-label="Client ID">{{ item.client_id }}</md-table-cell>
        <md-table-cell md-label="Meter ID">{{ item.meter }}</md-table-cell>
        <md-table-cell class="icon">
          <md-button
            class="md-just-icon md-simple md-info"
            @click="$emit('contract-details', item.contract_id)"
          >
            <md-icon>visibility</md-icon>
            <md-tooltip md-direction="top">Details</md-tooltip>
          </md-button>
        </md-table-cell>
        <md-table-cell class="icon">
          <md-button class="md-just-icon md-simple md-danger">
            <md-icon>clear</md-icon>
            <md-tooltip md-direction="top">Delete</md-tooltip>
          </md-button>
        </md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import axios from "axios";
import Vue from "vue";

export default {
  name: "contracts-table",
  props: {
    tableHeaderColor: {
      type: String,
      default: ""
    },

    contractsList: {
      type: Array,
      default: () => [{ name: "default" }]
    }
  },
  data() {
    return {
      showNotFound: false
    };
  }
};
</script>

<style scoped>
.icon {
  font-size: 36px;
}
</style>
