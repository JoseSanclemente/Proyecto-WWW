<template>
  <div>
    <md-table v-model="userList" :table-header-color="tableHeaderColor">
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="ID">{{ item.id }}</md-table-cell>
        <md-table-cell md-label="Name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Salary">{{ item.salary }}</md-table-cell>
        <md-table-cell md-label="Country">{{ item.country }}</md-table-cell>
        <md-table-cell md-label="City">{{ item.city }}</md-table-cell>
        <md-table-cell md-label>
          <md-button class="md-danger" @click="deleteUser(item.id)"
            >Delete</md-button
          >
        </md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
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
  methods: {
    deleteUser(id) {
      console.log("Hola:", id);
      for (let i = 0; i < this.userList.length; i++) {
        let user = this.userList[i];

        if (user.id === id) {
          this.userList.splice(this.userList.indexOf(user), 1);
          return;
        }
      }
    }
  }
};
</script>
