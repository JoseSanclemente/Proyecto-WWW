<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">Contracts</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <router-link :to="{ name: 'consumer-contracts', params: {id: consumer_id } }">
            <md-list-item>
              <md-icon>description</md-icon>
              <span class="md-list-item-text">Contracts</span>
            </md-list-item>
          </router-link>

          <router-link :to="{ name: 'consumer-profile', params: {id: consumer_id } }">
            <md-list-item>
              <md-icon>account_circle</md-icon>
              <span class="md-list-item-text">Profile</span>
            </md-list-item>
          </router-link>

          <router-link to="/login">
            <md-list-item>
              <md-icon>exit_to_app</md-icon>
              <span class="md-list-item-text">Sign out</span>
            </md-list-item>
          </router-link>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-table
          class="table"
          v-if="contracts!= null && contracts.length > 0"
          v-model="contracts"
          md-card
        >
          <md-table-toolbar class="table">
            <span class="md-title">{{$t('Contracts')}}</span>
          </md-table-toolbar>

          <md-table-row slot="md-table-row" slot-scope="{ item }">
            <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
            <md-table-cell :md-label="$t('Address')" md-sort-by="id">{{ item.address }}</md-table-cell>
            <md-table-cell
              :md-label="$t('Notification type')"
              md-sort-by="id"
            >{{ item.notification_type }}</md-table-cell>
            <md-table-cell>
              <md-button
                class="md-icon-button md-primary md-raised"
                @click="getConsumerPDF(item.id)"
              >
                <md-icon>save_alt</md-icon>
              </md-button>
            </md-table-cell>
          </md-table-row>
        </md-table>

        <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
export default {
  name: "Operator",
  data() {
    return {
      consumer_id: null,
      sending: false,
      showSnackBar: false,
      message: null
    };
  },
  computed: {
    ...mapState("consumer", ["contracts", "loggedConsumer"])
  },
  created() {
    this.consumer_id = this.loggedConsumer[0].id;
  },
  beforeMount() {
    this.getConsumerContracts();
  },
  methods: {
    ...mapActions("consumer", ["getPDF", "listContracts"]),
    getConsumerContracts() {
      this.sending = true;

      this.listContracts(this.consumer_id)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
          }, 1000);
        })
        .catch(error => {
          this.sending = false;
          console.log(error);
        });
    },
    getConsumerPDF(contract_id) {
      this.sending = true;
      this.getPDF(contract_id)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
          }, 1000);
        })
        .catch(error => {
          if (error == "not_found") {
            this.showNotification("Bill not generated yet for this month");
          }
          this.showNotification(
            "An error has occured while downloading the bill"
          );
          console.log(error);
        });
    },
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    }
  }
};
</script>

<style lang="scss" scoped>
.card-background {
  background-color: #333333;
}

.table {
  border: 1px solid #6e6e6e;
}

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

