<template>
  <div>
    <form novalidate class="md-layout" @submit.prevent="saveSubstation">
      <md-card class="md-layout-item md-size-60 md-small-size-100">
        <md-card-header>
          <div class="md-title">Add Substation</div>
        </md-card-header>

        <md-card-content>
          <map-component mapType="substation" :showAddressInput="true" @latlng="setLatLng"></map-component>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">Add</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="showSnackBar">{{ message }}</md-snackbar>
    </form>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import MapComponent from "@/components/map/MapComponent.vue";

export default {
  name: "substation-form",
  components: {
    MapComponent
  },
  data() {
    return {
      form: {
        latitude: null,
        longitude: null,
        deleted: false
      },
      showSnackBar: false,
      message: null,
      sending: false
    };
  },
  methods: {
    ...mapActions("substation", ["createSubstation"]),
    setLatLng(marker) {
      this.form.latitude = marker.lat;
      this.form.longitude = marker.lng;
      console.log(this.form);
    },
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    saveSubstation() {
      this.sending = true;

      this.createSubstation(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
            this.showNotification("The substation was successfully added!");
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification("An error had occured");
          console.log(error);
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.md-progress-bar {
  position: absolute;
  top: 0;
  right: 0;
  left: 0;
}

.md-card-actions {
  margin: 0.5em;
}
</style>