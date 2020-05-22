<template>
  <div>
    <md-empty-state
      v-if="searchedTransformers == null || searchedTransformers.length == 0"
      class="md-primary"
      md-icon="remove_circle_outline"
      md-label="There is nothing here yet"
      md-description="Transformers added will be showed here."
    ></md-empty-state>
    <md-table
      v-if="searchedTransformers != null && searchedTransformers.length > 0"
      v-model="searchedTransformers"
      md-card
      @md-selected="onSelect"
    >
      <md-table-toolbar class="md-primary" slot="md-table-alternate-header">
        <span class="md-title">{{$t("Transformers")}}</span>
        <md-field md-clearable class="md-layout-item md-size-30 md-toolbar-section-end">
          <md-input placeholder="Search by name" v-model="searchedTransformer" @input="searchOnTable" />
        </md-field>
      </md-table-toolbar>

      <md-table-row
        slot="md-table-row"
        slot-scope="{ item }"
      >
        <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
        <md-table-cell :md-label="$t('Name')">{{ item.name }}</md-table-cell>
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
  name: "transformer-table",
  computed: {
    ...mapState("substation", ["searchedTransformers"])
  },
  data: () => ({
    modalOpen: false,
    showSnackBar: false,
    message: "",
    selected: [],
    searchedTransformer: null
  }),
  methods: {
    ...mapMutations("substation", ["searchTransformer"]),
    onSelect(items) {
      this.selected = items;
    },
    searchOnTable() {
      this.searchTransformer(this.searchedTransformer);
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
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>