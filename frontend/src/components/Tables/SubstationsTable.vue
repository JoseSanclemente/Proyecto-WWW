<template>
  <div>
    <md-table v-model="substationsList" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <div v-if="showNotFound">No substations found</div>
        <md-table-cell md-label="ID">{{ item.sub_id }}</md-table-cell>
        <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Addres">{{ item.address }}</md-table-cell>
        <md-table-cell md-label="Active">{{ item.active }}</md-table-cell>
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
  name: "ordered-table",
  props: {
    tableHeaderColor: {
      type: String,
      default: ""
    },
    substationsList: {
      type: Array,
      default: () => [{ name: "default" }]
    }
  },
  data() {
    return {
      showNotFound: false
    };
  },
  methods: {}
};
</script>

<style scoped>
.icon {
  font-size: 36px;
}
</style>
