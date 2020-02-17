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

        <md-button class="md-just-icon md-raised md-primary search-button" @click="searchUser">
          <md-icon class="md-size-3x">search</md-icon>
        </md-button>
      </div>
      <md-card-content>
        <ordered-table :userList="users" table-header-color="purple"></ordered-table>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { OrderedTable } from "@/components";

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
      users: [
        {
          id: 1,
          name: "Jose Manuel",
          role: "Administrator",
          salary: "$36,738"
        },
        {
          id: 2,
          name: "Camilo",
          role: "Administrator",
          salary: "$36,738"
        },
        {
          id: 3,
          name: "Sofía",
          role: "Operator",
          salary: "$36,738"
        }
      ],
      // TODO: Implement a way to reset table's rows
      searchedUsers: []
    };
  },
  methods: {
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