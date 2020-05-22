<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">Profile</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <router-link to="/consumer/contracts">
            <md-list-item>
              <md-icon>description</md-icon>
              <span class="md-list-item-text">Contracts</span>
            </md-list-item>
          </router-link>

          <router-link to="/consumer/profile">
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
        <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
        <md-card class="md-layout-item md-size-50 md-small-size-100">
          <md-card-media>
            <img src="../../assets/banner.jpg" alt="banner" />
          </md-card-media>
          <md-card-header class="card-background">
            <div class="md-title">{{ $t("Edit information") }}</div>
          </md-card-header>

          <md-card-content class="card-content card-background">
            <div class="md-layout md-gutter">
              <div class="md-layout-item md-size-50 md-small-size-100">
                <md-field>
                  <label for="email">Email</label>
                  <md-input v-model="consumer.email" :disabled="sending" />
                </md-field>
              </div>
            </div>

            <div class="md-layout md-gutter">
              <div class="md-layout-item md-small-size-100">
                <md-field>
                  <label for="password">{{$t("Password")}}</label>
                  <md-input type="password" v-model="consumer.password" :disabled="sending" />
                </md-field>
              </div>

              <div class="md-layout-item md-small-size-100">
                <md-field>
                  <label for="confirmPass">{{$t("Confirm password")}}</label>
                  <md-input type="password" v-model="confirmPassword" :disabled="sending" />
                </md-field>
              </div>
            </div>
          </md-card-content>

          <md-progress-bar md-mode="indeterminate" v-if="sending" />

          <md-card-actions class="card-actions card-background">
            <md-button
              type="submit"
              class="md-primary md-raised"
              :disabled="sending"
            >{{$t("Modify")}}</md-button>
          </md-card-actions>
        </md-card>
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
      consumer: {
        email: "dummy@gmail.com",
        notification_type: "email",
        password: null
      },
      confirmPassword: null,
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
img {
  height: 350px;
}

.card-background {
  background-color: #333333;
}

.card-content {
  padding: 0 2em 2em 2em;
}

.card-actions {
  padding-right: 20px;
  padding-bottom: 20px;
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

