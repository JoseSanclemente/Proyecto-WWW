<template>
  <div>
    <!--user-form v-model="modalOpen"></user-form-->
    <md-table v-model="searchedSubstations" md-card>
      <md-table-toolbar>
        <span class="md-title">{{$t("Substations")}}</span>
        <md-field md-clearable class="md-layout-item md-size-30 md-toolbar-section-end">
          <md-input :placeholder="$t('Search by name')" v-model="searchedSubstation" @input="searchOnTable" />
        </md-field>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
        <md-table-cell :md-label="$t('Name')" md-sort-by="id">{{ item.name }}</md-table-cell>
        <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
        <md-table-cell>
          <md-button class="md-icon-button md-primary" @click="$emit('loadModifyForm', item)">
            <md-icon>edit</md-icon>
          </md-button>
        </md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { getStatusLabel } from "@/helpers/helpers.js";
import {mapMutations, mapState} from "vuex";

export default {
  name: "substation-table",
  components: {},
  data: () => ({
    modalOpen: false,
    searchedSubstation: "",
  }),
  computed: {
    ...mapState("substation", ["searchedSubstations"])
  },
  methods: {
    ...mapMutations("substation", ["searchSubstation"]),
    getStatus(status) {
      return getStatusLabel(status);
    },
    openModal() {
      this.modalOpen = !this.modalOpen;
    },
    searchOnTable() {
      this.searchSubstation(this.searchedSubstation);
    },
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>