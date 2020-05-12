<template>
  <div>
    <form novalidate class="md-layout" @submit.prevent="validateSubstation">
      <md-card class="md-layout-item md-size-40 md-small-size-100">
        <md-card-header>
          <div class="md-title">Add Substation</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field :class="getValidationClass('latitude')">
                <label for="latitude">Latitude</label>
                <md-input v-model="form.latitude" :disabled="sending" />
                <span class="md-error" v-if="!$v.form.latitude.required">The latitude is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field :class="getValidationClass('longitude')">
                <label for="longitude">Longitude</label>
                <md-input v-model="form.longitude" :disabled="sending" />
                <span class="md-error" v-if="!$v.form.longitude.required">The longitude is required</span>
              </md-field>
            </div>
          </div>
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

export default {
  name: "substation-form",
  mixins: [validationMixin],
  components: {
    //MapComponent
  },
  data: () => ({
    form: {
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
      latitude: {
        required
      },
      longitude: {
        required
      }
    }
  },
  methods: {
    ...mapActions("substation", ["createSubstation"]),
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
      this.form.latitude = null;
      this.form.longitude = null;
    },
    saveSubstation() {
      this.sending = true;

      this.createSubstation(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
            this.clearForm();
            this.showNotification("The substation was successfully added!");
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification("An error had occured");
          console.log(error);
        });
    },
    validateSubstation() {
      this.$v.$touch();
      console.log(this.form);
      if (!this.$v.$invalid) {
        this.saveSubstation();
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