<template>
  <div class="modal" v-show="value">
    <form novalidate class="md-layout" @submit.prevent="validateUser">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">{{ $t("Add user") }}</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-80 md-small-size-100">
              <md-field :class="getValidationClass('email')">
                <label for="email">{{$t("Email")}}</label>
                <md-input
                  type="email"
                  name="email"
                  autocomplete="email"
                  v-model="form.email"
                  :disabled="sending"
                />
                <span
                  class="md-error"
                  v-if="!$v.form.email.required"
                >{{$t("The email is required")}}</span>
                <span class="md-error" v-else-if="!$v.form.email.email">{{$t("Invalid email")}}</span>
              </md-field>
            </div>
          </div>

          <div class="md-layout">
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field :class="getValidationClass('type')">
                <label for="type">{{$t("Role")}}</label>
                <md-select name="type" v-model="form.type" md-dense :disabled="sending">
                  <md-option value="admin">{{$t("Admin")}}</md-option>
                  <md-option value="manager">{{$t("Manager")}}</md-option>
                  <md-option value="operator">{{$t("Operator")}}</md-option>
                </md-select>
                <span class="md-error">{{$t("The role is required")}}</span>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('password')">
                <label for="password">{{$t("Password")}}</label>
                <md-input type="password" v-model="form.password" :disabled="sending" />
                <span
                  class="md-error"
                  v-if="!$v.form.password.required"
                >{{$t("The password is required")}}</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('password')">
                <label for="confirmPass">{{$t("Confirm password")}}</label>
                <md-input type="password" v-model="confirmPass" :disabled="sending" />
                <span
                  class="md-error"
                  v-if="!$v.confirmPass.required"
                >{{$t("This field is required")}}</span>
              </md-field>
            </div>
          </div>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button @click.prevent="close">{{$t("Close")}}</md-button>
          <md-button type="submit" class="md-accent md-raised" :disabled="sending">{{$t("Add")}}</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
    </form>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, email } from "vuelidate/lib/validators";
import { mapActions } from "vuex";

export default {
  name: "user-form",
  mixins: [validationMixin],
  props: {
    value: {
      required: true
    }
  },
  data() {
    return {
      form: {
        email: null,
        type: null,
        password: null
      },
      confirmPass: null,
      showSnackBar: false,
      sending: false,
      message: null
    };
  },
  validations: {
    form: {
      type: {
        required
      },
      email: {
        required,
        email
      },
      password: {
        required
      }
    },
    confirmPass: {
      required
    }
  },
  methods: {
    ...mapActions("user", ["createUser"]),
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
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
      this.form.email = null;
      this.form.role = null;
      this.form.password = null;
      this.form.type = null;
      this.confirmPass = null;
    },
    saveUser() {
      this.sending = true;

      this.createUser(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;

            let message =
              "The user " + this.form.email + " was successfully added!";
            this.showNotification(message);

            this.clearForm();
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.clearForm();
          this.showNotification(this.$t("An error had occurred"));
          console.log(error);
        });
    },
    validateUser() {
      this.$v.$touch();

      if (!this.$v.$invalid) {
        this.saveUser();
      }
    },
    close() {
      this.clearForm();
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