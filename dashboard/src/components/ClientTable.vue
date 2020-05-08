<template>
  <div>
    <client-form v-model="modalOpen"></client-form>
    <md-table v-model="clients" md-card @md-selected="onSelect">
      <md-table-toolbar>
        <span class="md-title">Clients List</span>
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
        <md-table-cell md-label="ID" md-sort-by="id">{{ item.id }}</md-table-cell>
        <md-table-cell md-label="Document number" md-sort-by="email">{{ item.docNumber }}</md-table-cell>
        <md-table-cell md-label="Document type" md-sort-by="gender">{{ item.docType }}</md-table-cell>
        <md-table-cell md-label="Email" md-sort-by="title">{{ item.email }}</md-table-cell>
        <md-table-cell md-label="Phone number" md-sort-by="title">{{ item.phoneNumber }}</md-table-cell>
        <md-table-cell md-label="Address" md-sort-by="title">{{ item.address }}</md-table-cell>
        <md-table-cell md-label="Status" md-sort-by="title">{{ item.status }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import ClientForm from "@/components/ClientForm.vue";
export default {
  name: "client-table",
  components: {
    ClientForm
  },
  data: () => ({
    modalOpen: false,
    selected: [],
    clients: [
      {
        id: "CLIeb3ec3ed0c184551",
        docNumber: "1144712362",
        docType: "Cédula",
        email: "dummy@gmail.com",
        phoneNumber: "314652346",
        address: "Cra 84 #18-165",
        status: "Active"
      },
      {
        id: "CLIeb3ec3eB12451s",
        docNumber: "665531462",
        docType: "Cédula",
        email: "sdubbin0@geocities.com",
        phoneNumber: "316327654281",
        address: "Cll. 75 #45-90",
        status: "Active"
      },
      {
        id: "CLI12413e41234bcd",
        docNumber: "94123141-3",
        docType: "NIT",
        email: "contact@acities.comr",
        phoneNumber: "31642341242",
        address: "Cra 56 #58-10",
        status: "Inactive"
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