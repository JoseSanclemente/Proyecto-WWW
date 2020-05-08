<template>
  <div>
    <user-form v-model="modalOpen"></user-form>
    <md-table v-model="substations" md-card @md-selected="onSelect">
      <md-table-toolbar>
        <span class="md-title">Substations List</span>
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
        <md-table-cell md-label="Substation name" md-sort-by="name">{{ item.name }}</md-table-cell>
        <md-table-cell md-label="Address">{{ item.address }}</md-table-cell>
        <md-table-cell md-label="Status" md-sort-by="activeStatus">{{ item.activeStatus }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import UserForm from "@/components/user/UserForm.vue";
export default {
  name: "substation-table",
  components: {
    UserForm
  },
  data: () => ({
    modalOpen: false,
    selected: [],
    substations: [
      {
        id: "SUB72332db984166914e",
        name: "Substation A",
        address: "Cra 64 #14-105",
        activeStatus: "Active"
      },
      {
        id: "SUBbd7e04c617ad5bef",
        name: "Substation B",
        address: "Cra 32 #1a-198",
        activeStatus: "Active"
      },
      {
        id: "SUB72332dbbd7e04c614",
        name: "Substation Center",
        address: "Cra 67 #23-105",
        activeStatus: "Inactive"
      },
      {
        id: "SUB72332db98b412314a",
        name: "External Substation",
        address: "Cra 90 #14-105",
        activeStatus: "Active"
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

      return `${count} substation${plural} selected`;
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