<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Employees Stats</h4>
        <div class="md-layout-item md-small-size-100 md-size-33 search-header">
          <md-field>
            <label>User Name</label>
            <md-input v-model="username" type="text"></md-input>
          </md-field>
          <md-button class="md-success" @click="searchUser">Search</md-button>
        </div>
      </md-card-header>
      <md-card-content>
        <ordered-table
          table-header-color="orange"
          :userList="users"
        ></ordered-table>
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
          salary: "$36,738",
          country: "Niger",
          city: "Oud-Turnhout"
        },
        {
          id: 2,
          name: "Camilo",
          salary: "$36,738",
          country: "Niger",
          city: "Oud-Turnhout"
        },
        {
          id: 3,
          name: "Sofía",
          salary: "$36,738",
          country: "Niger",
          city: "Oud-Turnhout"
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
.search-header {
  display: block;
}
</style>
