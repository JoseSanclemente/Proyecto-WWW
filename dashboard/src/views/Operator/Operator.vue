<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">Bill payment</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <router-link to="/operator/payment">
            <md-list-item>
              <md-icon>description</md-icon>
              <span class="md-list-item-text">Bill Payment</span>
            </md-list-item>
          </router-link>

          <router-link to="/operator/consumers">
            <md-list-item>
              <md-icon>account_circle</md-icon>
              <span class="md-list-item-text">Consumers</span>
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
        <h1 class="md-title">Search consumer</h1>

        <div class="md-layout-item md-size-40">
          <md-field>
            <label for>Consumer ID</label>
            <md-input v-model="consumer_id" />
            <md-button class="md-icon-button md-primary" @click="getConsumerContracts">
              <md-icon>search</md-icon>
            </md-button>
          </md-field>
        </div>

        <md-table v-if="contracts!= null && contracts.length > 0" v-model="contracts" md-card>
          <md-table-toolbar>
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
              <md-button
                class="md-icon-button md-primary md-raised"
                @click="payConsumerBill(item.id)"
              >
                <md-icon>attach_money</md-icon>
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
      contract_id: "123456",
      sending: false,
      showSnackBar: false,
      message: null
    };
  },
  computed: {
    ...mapState("consumer", ["contracts"])
  },
  methods: {
    ...mapActions("consumer", ["getPDF", "listContracts", "payBill"]),
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
    getAlternateLabel(count) {
      let plural = "";

      if (count > 1) {
        plural = "s";
      }

      return `${count} contract${plural} selected`;
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
          this.showNotification(
            "An error has occured while downloading the bill"
          );
          console.log(error);
        });
    },
    payConsumerBill(id) {
      let payload = {
        contract_id: id
      };

      this.payBill(payload)
        .then(() => {
          this.showNotification("Bill paid sucessfully!");
        })
        .catch(error => {
          this.showNotification("An error has occured while paying the bill");
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

