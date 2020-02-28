<template>
  <div>
    <md-table v-model="userList" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <div v-if="showNotFound">No users found.</div>
        <md-table-cell md-label="ID">{{ item.national_id }}</md-table-cell>
        <md-table-cell md-label="Username">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Role">{{ item.role }}</md-table-cell>
        <md-table-cell md-label="Date of birth">{{
          item.date_of_birth
        }}</md-table-cell>
        <md-table-cell md-label="Email">{{ item.email }}</md-table-cell>
        <md-table-cell md-label="Adress">{{ item.address }}</md-table-cell>
        <md-table-cell md-label="Active">{{
          item.active ? "yes" : "no"
        }}</md-table-cell>
        <md-table-cell class="icon">
          <md-button
            class="md-just-icon md-simple md-danger"
            @click="deleteUser(item.user_id)"
          >
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
import { USER_DELETE } from "@/store/user";
export default {
  name: "ordered-table",
  props: {
    tableHeaderColor: {
      type: String,
      default: ""
    },
    userList: {
      type: Array,
      default: null
    }
  },
  data() {
    return {
      showNotFound: false
    };
  },
  methods: {
    deleteUser(id) {
      this.$store.dispatch(USER_DELETE, id);
    }
  }
};
</script>

<style scoped>
.icon {
  font-size: 36px;
}
</style>
