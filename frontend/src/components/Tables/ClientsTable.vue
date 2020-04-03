<template>
  <div>
    <md-table v-model="clientsList" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <div v-if="showNotFound">No clients found</div>
        <md-table-cell md-label="Client ID">{{ item.client_id }}</md-table-cell>
        <md-table-cell md-label="Document ID">{{
          item.document_id
        }}</md-table-cell>
        <md-table-cell md-label="Email">{{ item.email }}</md-table-cell>
        <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Phone number">{{
          item.phone_number
        }}</md-table-cell>
        <md-table-cell md-label="Address">{{ item.address }}</md-table-cell>
        <md-table-cell md-label="Date of birth">{{
          item.date_of_birth
        }}</md-table-cell>
        <md-table-cell md-label="Active">{{ item.active }}</md-table-cell>
        <md-table-cell class="icon">
          <md-button
            class="md-just-icon md-simple md-info"
            @click="$emit('client-details', item.client_id)"
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
import Vue from "vue";

export default {
  name: "clients-table",
  props: {
    tableHeaderColor: {
      type: String,
      default: ""
    },
    clientsList: {
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
