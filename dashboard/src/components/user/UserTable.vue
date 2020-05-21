<template>
  <div>
    <user-form modalType="modify" :inputUser="selected" v-model="modalOpen"></user-form>
    <md-empty-state
      v-if="users == null || users.length == 0"
      class="md-primary"
      md-icon="remove_circle_outline"
      :md-label="$t('There is nothing here yet')"
      :md-description="$t('Users added will be showed here.')"
    ></md-empty-state>
    <md-table
      v-if="users != null && users.length > 0"
      v-model="users"
      md-card
      @md-selected="onSelect"
    >
      <md-table-toolbar>
        <span class="md-title">{{ $t("Users") }}</span>
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
          <md-button @click="changeUserStatus" class="md-icon-button md-raised md-accent">
            <md-icon>voice_over_off</md-icon>
          </md-button>

          <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
        </div>
      </md-table-toolbar>

      <md-table-row
        slot="md-table-row"
        slot-scope="{ item }"
        md-selectable="multiple"
        md-auto-select
      >
        <md-table-cell :md-label="$t('Email')" md-sort-by="email">{{ item.email }}</md-table-cell>
        <md-table-cell :md-label="$t('Role')" md-sort-by="type">{{ item.type }}</md-table-cell>
        <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import UserForm from "@/components/user/UserForm.vue";
import { getStatusLabel } from "@/helpers/helpers.js";

export default {
  name: "user-table",
  components: {
    UserForm
  },
  data: () => ({
    modalOpen: false,
    showSnackBar: false,
    message: "",
    selected: []
  }),
  computed: {
    ...mapState("user", ["users"])
  },
  methods: {
    ...mapActions("user", ["listUsers", "updateUser"]),
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
    getStatus(status) {
      return getStatusLabel(status);
    },
    openModal() {
      this.modalOpen = !this.modalOpen;
    },
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    changeUserStatus() {
      for (let i = 0; i < this.selected.length; i++) {
        let item = this.selected[i];
        item.deleted = !item.deleted;
        this.updateUser(item)
          .then(() => {
            this.listUsers();
            this.showNotification("Users updated successfully!");
          })
          .catch(error => {
            console.log(error);
          });
      }
    }
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>