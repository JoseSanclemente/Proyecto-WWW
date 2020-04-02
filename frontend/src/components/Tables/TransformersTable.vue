<template>
  <div>
    <md-table v-model="transformersList" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <div v-if="transformersList.length <= 0">No transformers found</div>
        <md-table-cell md-label="ID">{{ item.trans_id }}</md-table-cell>
        <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Longitude">{{ item.longitude }}</md-table-cell>
        <md-table-cell md-label="Latitude">{{ item.latitude }}</md-table-cell>
        <md-table-cell md-label="Transformer model">{{
          item.transmodel
        }}</md-table-cell>
        <md-table-cell class="icon">
          <md-button
            class="md-just-icon md-simple md-info"
            @click="$emit('transformer-details', item.trans_id)"
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
  name: "transformers-table",
  props: {
    tableHeaderColor: {
      type: String,
      default: ""
    },

    transformersList: {
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
