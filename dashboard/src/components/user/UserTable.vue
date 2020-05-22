<template>
  <div>
    <update-user-form
      v-if="render"
      :show="modalOpen"
      :user="inputUser"
      @destroyModal="destroyModal"
    ></update-user-form>
    <h1 class="md-title">{{ $t("Users") }}</h1>
    <md-table v-model="searched" md-card md-fixed-header>
      <md-table-toolbar class="md-primary" slot="md-table-alternate-header">
        <md-field md-clearable class="md-layout-item md-size-30 md-toolbar-section-end">
          <md-input placeholder="Search by email" v-model="searchedUser" @input="searchOnTable" />
        </md-field>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell :md-label="$t('Email')" md-sort-by="email">{{ item.email }}</md-table-cell>
        <md-table-cell :md-label="$t('Role')" md-sort-by="type">{{ $t(getRole(item.type)) }}</md-table-cell>
        <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
        <md-table-cell>
          <md-button class="md-icon-button md-primary" @click="openUpdateModal(item)">
            <md-icon>edit</md-icon>
          </md-button>
        </md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { mapState, mapMutations } from "vuex";
import UpdateUserForm from "@/components/user/UpdateUserForm.vue";
import { getStatusLabel, getRoleLabel } from "@/helpers/helpers.js";

export default {
  name: "user-table",
  components: {
    UpdateUserForm
  },
  data: () => ({
    modalOpen: false,
    showSnackBar: false,
    message: "",
    inputUser: null,
    render: false,
    searchedUser: null
  }),
  computed: {
    ...mapState("user", ["searched"])
  },
  methods: {
    ...mapMutations("user", ["searchUser"]),
    searchOnTable() {
      this.searchUser(this.searchedUser);
    },
    getRole(role) {
      return getRoleLabel(role);
    },
    getStatus(status) {
      return getStatusLabel(status);
    },
    openUpdateModal(item) {
      this.render = true;
      this.inputUser = item;
      this.modalOpen = true;
    },
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    destroyModal(value) {
      this.modalOpen = value;
      this.render = false;
    }
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>