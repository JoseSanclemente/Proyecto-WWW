<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">Profile</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <router-link :to="{ name: 'consumer-contracts', params: {id: consumer.id } }">
            <md-list-item>
              <md-icon>description</md-icon>
              <span class="md-list-item-text">Contracts</span>
            </md-list-item>
          </router-link>

          <router-link :to="{ name: 'consumer-profile', params: {id: consumer.id } }">
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
              @click="update"
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
        id: null,
        email: null,
        password: null,
        delete: null
      },
      consumer_id: null,
      confirmPassword: null,
      sending: false,
      showSnackBar: false,
      message: null
    };
  },
  created() {
    this.consumer.id = this.loggedConsumer[0].id;
    this.consumer.email = this.loggedConsumer[0].email;
  },
  computed: {
    ...mapState("consumer", ["loggedConsumer"])
  },
  methods: {
    ...mapActions("consumer", ["updateConsumer"]),
    update() {
      this.sending = true;

      this.updateConsumer(this.consumer).then(() => {
        setTimeout(() => {
          this.sending = false;
          this.showNotification("Information successfully updated!");
        }, 2000);
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

