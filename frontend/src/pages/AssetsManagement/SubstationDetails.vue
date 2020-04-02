<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Substation Details</h4>
      </md-card-header>
      <md-card-content>
        <h4>Substation id</h4>
        {{ substationID }}
        <br />
        <h4>Name</h4>
        {{ substationName }}
        <br />
        <h4>Address</h4>
        {{ substationAddress }}
        <br />
        <h4>Active</h4>
        {{ substationActive }}
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "substation-details",

  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    },

    substationID: {
      type: String,
      default: "default"
    }
  },

  computed: {
    ...mapState({
      substations: state => state.assetsManagement.substations
    }),

    substationName: function() {
      if (this.substations === undefined) return "";

      for (let i = 0; i < this.substations.length; i++) {
        let substation = this.substations[i];

        if (this.substationID === substation.sub_id) {
          return substation.name;
        }
      }
      return "";
    },

    substationAddress: function() {
      if (this.substations === undefined) return "";

      for (let i = 0; i < this.substations.length; i++) {
        let substation = this.substations[i];

        if (this.substationID === substation.sub_id) {
          return substation.address;
        }
      }
      return "";
    },

    substationActive: function() {
      if (this.substations === undefined) return "";

      for (let i = 0; i < this.substations.length; i++) {
        let substation = this.substations[i];

        if (this.substationID === substation.sub_id) {
          return substation.active;
        }
      }
      return "";
    }
  }
};
</script>

<style scoped>
.search-card {
  display: flex;
  margin: 2rem 0 2rem 1rem;
}

.search-button {
  margin: 1.5rem 0 0 2rem;
}
</style>
