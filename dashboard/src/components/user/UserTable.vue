<template>
  <div>
    <user-form v-model="modalOpen"></user-form>
    <md-table v-model="people" md-card @md-selected="onSelect">
      <md-table-toolbar>
        <span class="md-title">Users List</span>
      </md-table-toolbar>

      <md-table-toolbar class="md-primary" slot="md-table-alternate-header" slot-scope="{ count }">
        <div class="md-toolbar-section-start">{{ getAlternateLabel(count) }}</div>

        <div class="md-toolbar-section-end">
          <md-button
            v-if="count == 1"
            class="md-icon-button md-raised md-accent"
            @click="openModal"
          >
            <md-icon>edit</md-icon>
          </md-button>
          <md-button class="md-icon-button md-raised md-accent">
            <md-icon>delete</md-icon>
          </md-button>
        </div>
      </md-table-toolbar>

      <md-table-row
        slot="md-table-row"
        slot-scope="{ item }"
        md-selectable="multiple"
        md-auto-select
      >
        <md-table-cell md-label="ID" md-sort-by="name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Name" md-sort-by="email">{{ item.email }}</md-table-cell>
        <md-table-cell md-label="Role" md-sort-by="gender">{{ item.gender }}</md-table-cell>
        <md-table-cell md-label="Active" md-sort-by="title">{{ item.title }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import UserForm from "@/components/user/UserForm.vue";
export default {
  name: "user-table",
  components: {
    UserForm
  },
  data: () => ({
    modalOpen: false,
    selected: [],
    people: [
      {
        name: "Shawna Dubbin",
        email: "sdubbin0@geocities.com",
        gender: "Male",
        title: "Assistant Media Planner"
      },
      {
        name: "Odette Demageard",
        email: "odemageard1@spotify.com",
        gender: "Female",
        title: "Account Coordinator"
      },
      {
        name: "Lonnie Izkovitz",
        email: "lizkovitz3@youtu.be",
        gender: "Female",
        title: "Operator"
      },
      {
        name: "Thatcher Stave",
        email: "tstave4@reference.com",
        gender: "Male",
        title: "Software Test Engineer III"
      },
      {
        name: "Clarinda Marieton",
        email: "cmarietonh@theatlantic.com",
        gender: "Female",
        title: "Paralegal"
      }
    ]
  }),
  methods: {
    onSelect(items) {
      this.selected = items;
    },
    getAlternateLabel(count) {
      let plural = "";

      if (count > 1) {
        plural = "s";
      }

      return `${count} user${plural} selected`;
    },
    openModal() {
      this.modalOpen = !this.modalOpen;
    }
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>