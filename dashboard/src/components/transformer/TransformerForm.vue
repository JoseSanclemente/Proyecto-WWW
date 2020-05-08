<template>
  <div>
    <form novalidate class="md-layout" @submit.prevent="validateUser">
      <md-card class="md-layout-item md-size-60 md-small-size-100">
        <md-card-header>
          <div class="md-title">Add Transformer</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-40 md-small-size-100">
              <md-field :class="getValidationClass('name')">
                <label for="transModel">Select model</label>
                <md-select
                  name="transModel"
                  id="transModel"
                  v-model="form.transModel"
                  md-dense
                  :disabled="sending"
                >
                  <md-option value="national_id">A</md-option>
                  <md-option value="tax_id">B</md-option>
                </md-select>
              </md-field>
            </div>

            <md-button
              class="md-icon-button md-primary md-raised"
              @click="openModal"
              :disabled="sending"
            >
              <md-icon>add</md-icon>
            </md-button>
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('substationID')">
                <label for="substationID">Select substation</label>
                <md-select
                  name="substationID"
                  id="substationID"
                  v-model="form.substationID"
                  md-dense
                  :disabled="sending"
                >
                  <md-option value="national_id">Sub A</md-option>
                  <md-option value="tax_id">Sub B</md-option>
                </md-select>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field :class="getValidationClass('address')">
                <label for="address">Transformer address</label>
                <md-input name="address" id="address" v-model="form.address" :disabled="sending" />
                <md-button class="md-icon-button md-mini">
                  <md-icon>search</md-icon>
                </md-button>
                <span class="md-error" v-if="!$v.form.address.required">The address is required</span>
                <span class="md-error" v-else-if="!$v.form.address.minlength">Invalid address</span>
              </md-field>
            </div>
          </div>

          <map-component></map-component>
          <model-form v-model="modalOpen"></model-form>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">Add</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="userSaved">The user {{ lastUser }} was saved with success!</md-snackbar>
    </form>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, minLength } from "vuelidate/lib/validators";
import MapComponent from "@/components/MapComponent.vue";
import ModelForm from "@/components/ModelForm.vue";

export default {
  name: "transformer-form",
  mixins: [validationMixin],
  components: {
    MapComponent,
    ModelForm
  },
  data: () => ({
    form: {
      name: null,
      address: null,
      active: false
    },
    userSaved: false,
    sending: false,
    lastUser: null,
    modalOpen: false
  }),
  validations: {
    form: {
      name: {
        required,
        minLength: minLength(3)
      },
      address: {
        required,
        minLength: minLength(3)
      }
    }
  },
  methods: {
    openModal() {
      this.modalOpen = !this.modalOpen;
    },
    getValidationClass(fieldName) {
      const field = this.$v.form[fieldName];

      if (field) {
        return {
          "md-invalid": field.$invalid && field.$dirty
        };
      }
    },
    clearForm() {
      this.$v.$reset();
      this.form.name = null;
      this.form.address = null;
      this.form.active = null;
    },
    saveUser() {
      this.sending = true;

      // Instead of this timeout, here you can call your API
      window.setTimeout(() => {
        this.lastUser = `${this.form.firstName} ${this.form.lastName}`;
        this.userSaved = true;
        this.sending = false;
        this.clearForm();
      }, 1500);
    },
    validateUser() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.saveUser();
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