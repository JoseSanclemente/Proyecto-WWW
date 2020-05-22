<template>
  <div class="page-container">
    <md-app>
      <md-app-toolbar class="md-accent" md-elevation="0">
        <span class="md-title">{{$t(title)}}</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">{{$t("Navigation")}}</md-toolbar>

        <md-list>
          <router-link to="/admin/users">
            <md-list-item @click="changeTitle('users')">
              <md-icon>supervisor_account</md-icon>
              <span class="md-list-item-text">{{$t("User")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/consumer">
            <md-list-item @click="changeTitle('consumer')">
              <md-icon>account_box</md-icon>
              <span class="md-list-item-text">{{$t("Consumer")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/transformers">
            <md-list-item @click="changeTitle('trans')">
              <md-icon>offline_bolt</md-icon>
              <span class="md-list-item-text">{{$t("Transformers")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/substations">
            <md-list-item @click="changeTitle('subs')">
              <md-icon>apartment</md-icon>
              <span class="md-list-item-text">{{$t("Substations")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/map">
            <md-list-item @click="changeTitle('map')">
              <md-icon>location_on</md-icon>
              <span class="md-list-item-text">{{$t("Map")}}</span>
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

      <md-app-content class="content">
        <router-view />
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
export default {
  name: "default-layout",
  data: () => ({
    menuVisible: false,
    title: "User Management",
    langs: [
      { name: "English", code: "en" },
      { name: "Español", code: "es" },
      { name: "Português", code: "pt" }
    ]
  }),
  methods: {
    toggleMenu() {
      this.menuVisible = !this.menuVisible;
    },
    changeTitle(page) {
      if (page == "consumer") {
        this.title = "Consumer Management";
      }

      if (page == "users") {
        this.title = "User Management";
      }

      if (page == "trans") {
        this.title = "Transformer Management";
      }

      if (page == "subs") {
        this.title = "Substations";
      }

      if (page == "map") {
        this.title = "Map";
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.md-app {
  height: 100vh;
  border: 1px solid rgba(#000, 0.12);
}

.content {
  background-color: #333333;
}

/* Demo purposes only */
.md-drawer {
  width: 230px;
  max-width: calc(100vw - 125px);
}
</style>