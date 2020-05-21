<template>
  <div>
    <!--user-form v-model="modalOpen"></user-form-->
    <md-table v-model="substations" md-card @md-selected="onSelect">
      <md-table-toolbar>
        <span class="md-title">{{$t("Substations")}}</span>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
        <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { getStatusLabel } from "@/helpers/helpers.js";

export default {
  name: "substation-table",
  components: {},
  props: {
    substations: { required: true }
  },
  data: () => ({
    modalOpen: false
  }),
  methods: {
    getStatus(status) {
      return getStatusLabel(status);
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