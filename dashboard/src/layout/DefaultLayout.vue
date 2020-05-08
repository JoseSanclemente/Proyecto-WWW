<template>
  <div class="page-container">
    <md-app>
      <md-app-toolbar class="md-primary" md-elevation="0">
        <span class="md-title">{{title}}</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <router-link to="/admin/users">
            <md-list-item @click="changeTitle('users')">
              <md-icon>supervisor_account</md-icon>
              <span class="md-list-item-text">User Management</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/clients">
            <md-list-item @click="changeTitle('clients')">
              <md-icon>account_box</md-icon>
              <span class="md-list-item-text">Client Management</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/transformers">
            <md-list-item @click="changeTitle('trans')">
              <md-icon>offline_bolt</md-icon>
              <span class="md-list-item-text">Transformers</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/substations">
            <md-list-item @click="changeTitle('subs')">
              <md-icon>apartment</md-icon>
              <span class="md-list-item-text">Substations</span>
            </md-list-item>
          </router-link>

          <router-link to="/admin/map">
            <md-list-item @click="changeTitle('map')">
              <md-icon>location_on</md-icon>
              <span class="md-list-item-text">Map</span>
            </md-list-item>
          </router-link>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <slot />
        <router-view />
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
export default {
  name: "PersistentMini",
  data: () => ({
    menuVisible: false,
    title: "User Management"
  }),
  methods: {
    toggleMenu() {
      this.menuVisible = !this.menuVisible;
    },
    changeTitle(page) {
      if (page == "clients") {
        this.title = "Clients Management";
      }

      if (page == "users") {
        this.title = "Users Management";
      }

      if (page == "trans") {
        this.title = "Transformers Management";
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

// Demo purposes only
.md-drawer {
  width: 230px;
  max-width: calc(100vw - 125px);
}
</style>