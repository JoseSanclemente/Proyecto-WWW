<template>
  <div>
    <form novalidate class="md-layout" @submit.prevent="validateTrans">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">{{$t("Add Transformer")}}</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-50 md-small-size-50">
              <md-field :class="getValidationClass('substation_id')">
                <label for="substation_id">{{$t("Select substation")}}</label>
                <md-select v-model="transformerForm.substation_id" md-dense :disabled="sending">
                  <md-option
                    v-for="sub in substations"
                    v-bind:key="sub.id"
                    v-bind:value="sub.id"
                  >{{ sub.id }}</md-option>
                </md-select>
              </md-field>
            </div>

            <div class="md-layout-item md-size-50">
              <md-field>
                <label for>{{$t("Name")}}</label>
                <md-input v-model="transformerForm.name" />
              </md-field>
            </div>
          </div>
          <map-component mapType="transformer" :showAddressInput="true" @latlng="setLatLng"></map-component>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">{{$t("Add")}}</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
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
  props: {
    substations: { type: Array, required: true }
  },
  data() {
    return {
      transformerForm: {
        substation_id: null,
        name: null,
        latitude: null,
        longitude: null
      },
      showSnackBar: false,
      message: null,
      sending: false
    };
  },
  validations: {
    transformerForm: {
      substation_id: {
        required
      }
    }
  },
  methods: {
    ...mapActions("substation", ["createTransformer"]),
    setLatLng(marker) {
      this.transformerForm.latitude = marker.lat;
      this.transformerForm.longitude = marker.lng;
    },
    getValidationClass(fieldName) {
      const field = this.$v.transformerForm[fieldName];

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
      this.transformerForm.substation_id = null;
      this.transformerForm.name = null;
    },
    saveTransformer() {
      this.sending = true;

      this.createTransformer(this.transformerForm)
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