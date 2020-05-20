<template>
  <div>
    <consumer-form modalType="modify" :inputUser="selected" v-model="modalOpen"></consumer-form>
    <md-empty-state
      v-if="tableConsumers == null || tableConsumers.length == 0"
      class="md-primary"
      md-icon="remove_circle_outline"
      :md-label="$t('There is nothing here yet')"
      :md-description="$t('Consumers added will be showed here.')"
    ></md-empty-state>
    <md-table
      v-if="tableConsumers!= null && tableConsumers.length > 0"
      v-model="tableConsumers"
      md-card
      @md-selected="onSelect"
    >
      <md-table-toolbar>
        <span class="md-title">{{$t('Consumers')}}</span>
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
        <md-table-cell :md-label="$t('Email')">{{ item.email }}</md-table-cell>
        <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import ConsumerForm from "@/components/consumer/ConsumerForm.vue";
import { getStatusLabel } from "@/helpers/helpers.js";
export default {
  name: "consumer-table",
  components: {
    ConsumerForm
  },
  data: () => ({
    modalOpen: false,
    showSnackBar: false,
    message: "",
    selected: [],
    tableConsumers: []
  }),
  computed: {
    ...mapState("consumer", ["consumers"])
  },
  watch: {
    table: function() {
      this.tableConsumers = this.consumers;
    }
  },
  created() {
    this.listConsumers();
    this.tableConsumers = this.consumers;
  },
  methods: {
    ...mapActions("consumer", ["listConsumers"]),
    onSelect(items) {
      this.selected = items;
    },
    getStatus(status) {
      return getStatusLabel(status);
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