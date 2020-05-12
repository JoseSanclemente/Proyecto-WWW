<template>
  <div>
    <consumer-form modalType="modify" :inputUser="selected" v-model="modalOpen"></consumer-form>
    <md-table
      v-if="tableConsumers.length > 0"
      v-model="tableConsumers"
      md-card
      @md-selected="onSelect"
    >
      <md-table-toolbar>
        <span class="md-title">Consumers List</span>
      </md-table-toolbar>
      <md-empty-state
        v-if="tableConsumers.length == 0"
        class="md-primary"
        md-icon="remove_circle_outline"
        md-label="There is nothing here yet"
        md-description="Users added will be showed here."
      ></md-empty-state>
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

          <md-snackbar :md-active.sync="showSnackBar">{{ message }}</md-snackbar>
        </div>
      </md-table-toolbar>

      <md-table-row
        slot="md-table-row"
        slot-scope="{ item }"
        md-selectable="multiple"
        md-auto-select
      >
        <md-table-cell md-label="ID" md-sort-by="id">{{ item.id }}</md-table-cell>
        <md-table-cell md-label="Latitude" md-sort-by="email">{{ item.lat }}</md-table-cell>
        <md-table-cell md-label="Longitude">{{ item.long }}</md-table-cell>
        <md-table-cell md-label="Status">{{ item.deleted }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import TransformerForm from "@/components/transformer/TransformerForm.vue";
export default {
  name: "transformer-table",
  components: {
    TransformerForm
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