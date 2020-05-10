<template>
  <div class="modal" v-show="value">
    <form novalidate class="md-layout" @submit.prevent="validateUser">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">Add user</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
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
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('role')">
                <label for="role">Role</label>
                <md-select name="role" id="role" v-model="form.role" md-dense :disabled="sending">
                  <md-option value="admin">Admin</md-option>
                  <md-option value="manager">Manager</md-option>
                  <md-option value="operator">Operator</md-option>
                </md-select>
                <span class="md-error">The role is required</span>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('pass')">
                <label for="pass">Password</label>
                <md-input
                  type="pass"
                  name="pass"
                  id="pass"
                  v-model="form.pass"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.pass.required">The password is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('pass')">
                <label for="confirmPass">Confirm password</label>
                <md-input
                  type="confirmPass"
                  name="confirmPass"
                  id="confirmPass"
                  v-model="form.confirmPass"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.confirmPass.required">This field is required</span>
              </md-field>
            </div>
          </div>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button @click.prevent="close">Cancel</md-button>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">Create</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="userSaved">The user {{ lastUser }} was saved with success!</md-snackbar>
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
        role: null,
        pass: null,
        confirmPass: null
      },
      userSaved: false,
      sending: false,
      lastUser: null
    };
  },
  validations: {
    form: {
      role: {
        required
      },
      email: {
        required,
        email
      },
      pass: {
        required
      },
      confirmPass: {
        required
      }
    }
  },
  methods: {
    ...mapActions("user", ["createUser"]),
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
      this.form.pass = null;
      this.form.confirmPass = null;
    },
    saveUser() {
      this.sending = true;

      console.log(this.form);
      this.createUser(this.form)
        .then(response => {
          this.sending = false;
          console.log(response);
        })
        .catch(error => {
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