<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Search user</h4>
      </md-card-header>
      <div class="md-layout-item md-small-size-100 md-size-33 search-card">
        <md-field>
          <label>Username</label>
          <md-input v-model="username" type="text"></md-input>
        </md-field>

        <md-button
          class="md-just-icon md-raised md-primary search-button"
          @click="searchUser"
        >
          <md-icon class="md-size-3x">search</md-icon>
        </md-button>
      </div>
      <md-card-content>
        <ordered-table
          :userList="users"
          table-header-color="purple"
        ></ordered-table>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { OrderedTable } from "@/components";
import axios from "axios";

export default {
  name: "search-user-form",
  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    }
  },
  components: {
    OrderedTable
  },
  data() {
    return {
      username: null,
      users: []
    };
  },
  methods: {
    getDummy() {
      const url = "http://127.0.0.1:8000/api/v1.0/dummy/";

      axios
        .get(url)
        .then(response => {
          console.log(response.data);
          this.user = response.data;
        })
        .catch(error => {
          console.log(error);
        });
    },
    searchUser() {
      let newUserList = [];
      for (let i = 0; i < this.users.length; i++) {
        let user = this.users[i];

        if (user.name === this.username) {
          newUserList.push(user);
          this.users = newUserList;
          return;
        }
      }
    }
  },

  created() {
    this.getDummy();
  }
};
</script>

<style scoped>
.search-card {
  display: flex;
  margin: 0 0 2rem 1rem;
}

.search-button {
  margin: 1.5rem 0 0 2rem;
}
</style>
