<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">{{$t("Consumers")}}</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">{{$t("Navigation")}}</md-toolbar>

        <md-list>
          <router-link to="/operator/payment">
            <md-list-item>
              <md-icon>description</md-icon>
              <span class="md-list-item-text">{{$t("Bill Payment")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/operator/consumers">
            <md-list-item>
              <md-icon>account_circle</md-icon>
              <span class="md-list-item-text">{{$t("Consumers")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/employee/login">
            <md-list-item>
              <md-icon>exit_to_app</md-icon>
              <span class="md-list-item-text">{{$t("Sign out")}}</span>
            </md-list-item>
          </router-link>

          <md-list-item>
            <md-field>
              <label>{{$t("Language")}}</label>
              <md-select v-model="$i18n.locale">
                <md-option
                  v-for="(lang, i) in langs"
                  :key="`Lang${i}`"
                  :value="lang.code"
                >{{ lang.name }}</md-option>
              </md-select>
            </md-field>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <md-table v-model="searched" md-card md-fixed-header>
          <md-table-toolbar class="md-primary" slot="md-table-alternate-header">
            <md-field md-clearable class="md-layout-item md-size-30 md-toolbar-section-end">
              <label>{{$t("Search by ID")}}</label>
              <md-input v-model="searchedConsumer" @input="searchOnTable" />
            </md-field>
          </md-table-toolbar>

          <md-table-row slot="md-table-row" slot-scope="{ item }">
            <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
            <md-table-cell :md-label="$t('Email')">{{ item.email }}</md-table-cell>
            <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
            <md-table-cell>
              <md-button class="md-icon-button md-primary" @click="openUpdateModal(item)">
                <md-icon>edit</md-icon>
              </md-button>
            </md-table-cell>
          </md-table-row>
        </md-table>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import { mapState, mapMutations, mapActions } from "vuex";
import { getStatusLabel } from "@/helpers/helpers.js";
export default {
  data: () => ({
    searchedConsumer: null,
    langs: [
      { name: "English", code: "en" },
      { name: "Español", code: "es" },
      { name: "Português", code: "pt" }
    ]
  }),
  computed: {
    ...mapState("consumer", ["searched"])
  },
  created() {
    this.listConsumers();
  },
  methods: {
    ...mapMutations("consumer", ["searchConsumer"]),
    ...mapActions("consumer", ["listConsumers"]),
    searchOnTable() {
      this.searchConsumer(this.searchedConsumer);
    },
    getStatus(status) {
      return getStatusLabel(status);
    }
  }
};
</script>

<style lang="scss" scoped>
.md-app {
  height: 100vh;
  border: 1px solid rgba(#000, 0.12);
}

// Demo purposes only
.md-drawer {
  width: 200px;
  max-width: calc(100vw - 125px);
}
</style>

