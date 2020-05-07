<template>
  <div class="page-container">
    <md-app>
      <md-app-toolbar class="md-primary" md-elevation="0">
        <slot name="header" />
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

        <md-list>
          <router-link to="/users">
            <md-list-item>
              <md-icon>supervisor_account</md-icon>
              <span class="md-list-item-text">User Management</span>
            </md-list-item>
          </router-link>

          <router-link to="/clients">
            <md-list-item>
              <md-icon>account_box</md-icon>
              <span class="md-list-item-text">Client Management</span>
            </md-list-item>
          </router-link>

          <router-link to="/substations">
            <md-list-item>
              <md-icon>apartment</md-icon>
              <span class="md-list-item-text">Substations</span>
            </md-list-item>
          </router-link>

          <router-link to="/about">
            <md-list-item>
              <md-icon>location_on</md-icon>
              <span class="md-list-item-text">Maps</span>
            </md-list-item>
          </router-link>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <slot />
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
export default {
  name: "PersistentMini",
  data: () => ({
    menuVisible: false
  }),
  methods: {
    toggleMenu() {
      this.menuVisible = !this.menuVisible;
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