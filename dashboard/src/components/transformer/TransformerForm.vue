<template>
  <div>
    <form novalidate class="md-layout" @submit.prevent="validateTrans">
      <md-card class="md-layout-item md-size-60 md-small-size-100">
        <md-card-header>
          <div class="md-title">Add Transformer</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field :class="getValidationClass('substation_id')">
                <label for="substation_id">Select substation</label>
                <md-select v-model="form.substation_id" md-dense :disabled="sending">
                  <md-option value="national_id">Sub A</md-option>
                  <md-option value="tax_id">Sub B</md-option>
                </md-select>
              </md-field>
            </div>
          </div>
          <map-component mapType="transformer" :showAddressInput="true" @latlng="setLatLng"></map-component>
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
import { validationMixin } from "vuelidate";
import { required } from "vuelidate/lib/validators";
import { mapActions } from "vuex";
import MapComponent from "@/components/map/MapComponent.vue";

export default {
  name: "transformer-form",
  mixins: [validationMixin],
  components: {
    MapComponent
  },
  data: () => ({
    form: {
      substation_id: null,
      latitude: null,
      longitude: null,
      //address: null,
      deleted: false
    },
    showSnackBar: false,
    message: null,
    sending: false
  }),
  validations: {
    form: {
      substation_id: {
        required
      }
    }
  },
  methods: {
    ...mapActions("transformer", ["createTransformer"]),
    setLatLng(marker) {
      this.form.latitude = marker.lat;
      this.form.longitude = marker.lng;
    },
    getValidationClass(fieldName) {
      const field = this.$v.form[fieldName];

      if (field) {
        return {
          "md-invalid": field.$invalid && field.$dirty
        };
      }
    },
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    clearForm() {
      this.$v.$reset();
      this.form.substation_id = null;
    },
    saveTransformer() {
      this.sending = true;

      this.createTransformer(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
            this.clearForm();
            this.showNotification("The transform was successfully added!");
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification("An error had occured");
          console.log(error);
        });
    },
    validateTrans() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.saveTransformer();
      }
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