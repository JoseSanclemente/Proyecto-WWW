<template>
  <div>
    <update-user-form
      v-if="render"
      :show="modalOpen"
      :user="inputUser"
      @destroyModal="destroyModal"
    ></update-user-form>
    <h1 class="md-title">{{ $t("Users") }}</h1>
    <md-table v-if="users != null && users.length > 0" v-model="users" md-card>
      <md-table-toolbar class="md-primary" slot="md-table-alternate-header" slot-scope="{ count }">
        <div class="md-toolbar-section-start">{{ getAlternateLabel(count) }}</div>
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
import { mapState } from "vuex";
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
    render: false
  }),
  computed: {
    ...mapState("user", ["users"])
  },
  methods: {
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