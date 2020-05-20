<template>
  <div>
    <md-empty-state
      v-if="transformers == null || transformers.length == 0"
      class="md-primary"
      md-icon="remove_circle_outline"
      md-label="There is nothing here yet"
      md-description="Transformers added will be showed here."
    ></md-empty-state>
    <md-table
      v-if="transformers != null && transformers.length > 0"
      v-model="transformers"
      md-card
      @md-selected="onSelect"
    >
      <md-table-toolbar>
        <span class="md-title">{{$t("Transformers")}}</span>
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
            <md-icon>delete</md-icon>
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
        <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
        <md-table-cell :md-label="$t('Latitude')">{{ item.latitude }}</md-table-cell>
        <md-table-cell :md-label="$t('Longitude')">{{ item.longitude }}</md-table-cell>
        <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { getStatusLabel } from "@/helpers/helpers.js";

export default {
  name: "transformer-table",
  props: {
    transformers: { type: Array, required: true }
  },
  data: () => ({
    modalOpen: false,
    showSnackBar: false,
    message: "",
    selected: []
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
    },
    getStatus(status) {
      return getStatusLabel(status);
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
            this.showNotification("Transformers updated successfully!");
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