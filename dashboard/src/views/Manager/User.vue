<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">{{$t("Users")}}</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <router-link to="/manager/reports">
            <md-list-item>
              <md-icon>assessment</md-icon>
              <span class="md-list-item-text">{{$t("Reports")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/manager/users">
            <md-list-item>
              <md-icon>account_circle</md-icon>
              <span class="md-list-item-text">{{$t("Users")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/login">
            <md-list-item>
              <md-icon>exit_to_app</md-icon>
              <span class="md-list-item-text">{{$t("Sign out")}}</span>
            </md-list-item>
          </router-link>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <md-table v-model="searched" md-card md-fixed-header>
          <md-table-toolbar class="md-primary" slot="md-table-alternate-header">
            <md-field md-clearable class="md-layout-item md-size-30 md-toolbar-section-end">
              <md-input
                placeholder="Search by email"
                v-model="searchedUser"
                @input="searchOnTable"
              />
            </md-field>
          </md-table-toolbar>

          <md-table-row slot="md-table-row" slot-scope="{ item }">
            <md-table-cell :md-label="$t('Email')" md-sort-by="email">{{ item.email }}</md-table-cell>
            <md-table-cell :md-label="$t('Role')" md-sort-by="type">{{ $t(getRole(item.type)) }}</md-table-cell>
            <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
          </md-table-row>
        </md-table>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import { mapState, mapMutations, mapActions } from "vuex";
import { getStatusLabel, getRoleLabel } from "@/helpers/helpers.js";

export default {
  name: "user-table",
  data: () => ({
    searchedUser: null
  }),
  computed: {
    ...mapState("user", ["searched"])
  },
  created() {
    this.listUsers();
  },
  methods: {
    ...mapMutations("user", ["searchUser"]),
    ...mapActions("user", ["listUsers"]),
    searchOnTable() {
      this.searchUser(this.searchedUser);
    },
    getRole(role) {
      return getRoleLabel(role);
    },
    getStatus(status) {
      return getStatusLabel(status);
    }
  }
};
</script>

<style lang="scss" scoped>
.md-app {
  height: 100vh;
  border: 1px solid rgba(#000, 0.12);
}

.chart {
  margin: 2em 0.2em;
  width: 40vw;
}

.table {
  margin: 2em 0.2em;
}

// Demo purposes only
.md-drawer {
  width: 200px;
  max-width: calc(100vw - 125px);
}
</style>

