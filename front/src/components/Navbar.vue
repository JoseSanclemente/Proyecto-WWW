<template>
  <div id="nav">
    <div class="store-related">
      <router-link class="navbar-element" to="/" v-if="!isLoged">Inicio</router-link>
      <router-link class="navbar-element" to="/admin" v-if="isLoged && isAdmin">Inicio</router-link>
      <router-link class="navbar-element" to="/user" v-if="isLoged && isClient">Inicio</router-link>
      <router-link class="navbar-element" to="/about">About</router-link>
    </div>
    <div class="user-related">
      <div class="basket-container" v-if="isLoged && isClient" @click="openBasket()">
        <img src="@/assets/images/shopping-basket.png" class="basket-icon" />
      </div>
      <router-link class="navbar-element" to="/user-profile" v-if="isLoged && isClient">Perfil</router-link>
      <router-link class="navbar-element" to="/history" v-if="isLoged && isClient">Historial</router-link>
      <router-link class="navbar-element" to="/admin-profile" v-if="isLoged && isAdmin">Perfil</router-link>
      <a class="navbar-element" href='http://localhost:8081/' v-if="isLoged">Log Out</a>
      <router-link class="navbar-element" to="/login-client" v-if="!isLoged">Log In</router-link>
      <router-link class="navbar-element" to="/sign-up" v-if="!isLoged">Sign Up</router-link>
    </div>
    <basket v-if="showBasket" @close="showBasket = false"/>
  </div>
</template>

<script>
import { isLoged } from '@/api/api.js'
import { isClientRoute, isAdminRoute } from '@/routes.js'
import Basket from '@/components/Basket.vue'

export default {
  name: 'NavBar',

  components: {
    Basket
  },

  data: function () {
    return {
      isClient: false,
      isAdmin: false,
      isLoged: false,
      showBasket: false
    }
  },

  watch: {
    $route (to, from) {
      this.isClient = this.isClientCheck()
      this.isAdmin = this.isAdminCheck()
      this.isLoged = this.isLogedCheck()
    }
  },

  methods: {
    isClientCheck: function () {
      if (
        this.$router.currentRoute.name === null ||
        this.$router.currentRoute.name === undefined
      ) {
        return false
      }

      return isClientRoute(this.$router.currentRoute.name)
    },

    isAdminCheck: function () {
      if (
        this.$router.currentRoute.name === null ||
        this.$router.currentRoute.name === undefined
      ) {
        return false
      }

      return isAdminRoute(this.$router.currentRoute.name)
    },

    isLogedCheck: function () {
      return isLoged()
    },

    openBasket: function () {
      this.showBasket = true
    },

    logOut: function () {
      console.log(this.isLoged)
      this.isLoged = false
    }
  }
}
</script>

<style scoped>
#nav {
  position: -webkit-sticky;
  position: sticky;
  top: 0;
  display: flex;
  padding: 30px;
  background-color: #000;
  justify-content: space-between;
  height: 78px;
}

#nav a {
  font-weight: bold;
  color: #fff;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

.store-related {
  display: inline;
}

.user-related {
  display: flex;
}

.navbar-element {
  margin: 0px 5px;
}

.basket-container {
  height: 32px;
  width: 100px;
}

.basket-icon {
  height: 32px;
  width: 32px;
  position: absolute;
  top: 50%;
  transform: translate(0, -50%);
}
</style>
