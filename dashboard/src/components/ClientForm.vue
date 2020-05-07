<template>
  <div class="modal" v-show="value">
    <form novalidate class="md-layout" @submit.prevent="validateUser">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">Add Client</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('docNumber')">
                <label for="docNumber">Document number</label>
                <md-input
                  name="docNumber"
                  id="docNumber"
                  v-model="form.docNumber"
                  :disabled="sending"
                />
                <span
                  class="md-error"
                  v-if="!$v.form.docNumber.required"
                >The document number is required</span>
                <span
                  class="md-error"
                  v-else-if="!$v.form.docNumber.minlength"
                >Invalid document number</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('docType')">
                <label for="docType">Document Type</label>
                <md-select
                  name="docType"
                  id="docType"
                  v-model="form.docType"
                  md-dense
                  :disabled="sending"
                >
                  <md-option></md-option>
                  <md-option value="national_id">Cédula</md-option>
                  <md-option value="tax_id">NIT</md-option>
                </md-select>
                <span
                  class="md-error"
                  v-if="!$v.form.docType.required"
                >The document type is required</span>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('phoneNumber')">
                <label for="phoneNumber">Phone number</label>
                <md-input
                  type="phoneNumber"
                  name="phoneNumber"
                  id="phoneNumber"
                  autocomplete="phoneNumber"
                  v-model="form.phoneNumber"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.phoneNumber.required">The email is required</span>
                <span class="md-error" v-if="!$v.form.phoneNumber.minLength">Invalid phone number</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('address')">
                <label for="address">Address</label>
                <md-input
                  type="address"
                  name="address"
                  id="address"
                  autocomplete="address"
                  v-model="form.address"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.address.required">The address is required</span>
              </md-field>
            </div>
          </div>

          <md-field :class="getValidationClass('email')">
            <label for="email">Email</label>
            <md-input
              type="email"
              name="email"
              id="email"
              autocomplete="email"
              v-model="form.email"
              :disabled="sending"
            />
            <span class="md-error" v-if="!$v.form.email.required">The email is required</span>
            <span class="md-error" v-else-if="!$v.form.email.email">Invalid email</span>
          </md-field>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-checkbox class="md-primary" v-model="form.active">Active</md-checkbox>
            </div>
          </div>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button @click.prevent="close">Cancel</md-button>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">Add</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="userSaved">The user {{ lastUser }} was saved with success!</md-snackbar>
    </form>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, email, minLength } from "vuelidate/lib/validators";

export default {
  name: "client-form",
  mixins: [validationMixin],
  props: {
    value: {
      required: true
    }
  },
  data: () => ({
    form: {
      docNumber: null,
      docType: null,
      email: null,
      phoneNumber: null,
      address: null
    },
    userSaved: false,
    sending: false,
    lastUser: null
  }),
  validations: {
    form: {
      docNumber: {
        required,
        minLength: minLength(3)
      },
      docType: {
        required,
        minLength: minLength(3)
      },
      email: {
        required,
        email
      },
      phoneNumber: {
        required,
        minLength: minLength(7)
      },
      address: {
        required
      }
    }
  },
  methods: {
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
      this.form.firstName = null;
      this.form.lastName = null;
      this.form.age = null;
      this.form.gender = null;
      this.form.email = null;
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
    },
    close() {
      this.$emit("input", !this.value);
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

.modal {
  position: fixed;
  right: 0;
  left: 0;
  bottom: 0;
  top: 0;
  z-index: 100;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.7);
}
</style>