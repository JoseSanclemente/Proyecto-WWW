<template>
  <div class="modal" v-show="value">
    <form novalidate class="md-layout" @submit.prevent="validateConsumer">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">{{ title }}</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-80 md-small-size-100">
              <md-field :class="getValidationClass('id')">
                <label for="email">Cédula de ciudanía</label>
                <md-input name="id" v-model="form.id" :disabled="sending" />
                <span class="md-error" v-if="!$v.form.id.required">The email is required</span>
                <span class="md-error" v-else-if="!$v.form.minLength">Invalid id</span>
              </md-field>
            </div>
            <div class="md-layout-item md-size-80 md-small-size-100">
              <md-field :class="getValidationClass('email')">
                <label for="email">Email</label>
                <md-input
                  type="email"
                  name="email"
                  autocomplete="email"
                  v-model="form.email"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.email.required">The email is required</span>
                <span class="md-error" v-else-if="!$v.form.email.email">Invalid email</span>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('password')">
                <label for="password">Password</label>
                <md-input type="password" v-model="form.password" :disabled="sending" />
                <span class="md-error" v-if="!$v.form.password.required">The password is required</span>
                <span class="md-error" v-if="!$v.form.password.minLength">The password is too short!</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('confirmPass')">
                <label for="confirmPass">Confirm password</label>
                <md-input type="password" v-model="confirmPass" :disabled="sending" />
                <span class="md-error" v-if="!$v.confirmPass.required">This field is required</span>
              </md-field>
            </div>
          </div>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button @click.prevent="close">Cancel</md-button>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">Add</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="showSnackBar">{{ message }}</md-snackbar>
    </form>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, email, minLength } from "vuelidate/lib/validators";
import { mapActions } from "vuex";

export default {
  name: "consumer-form",
  mixins: [validationMixin],
  props: {
    value: {
      required: true
    },
    modalType: {
      required: true
    },
    inputUser: null
  },
  data() {
    return {
      form: {
        id: null,
        email: null,
        password: null,
        deleted: false
      },
      confirmPass: null,
      showSnackBar: false,
      sending: false,
      message: null,
      title: null
    };
  },
  validations: {
    form: {
      id: {
        required,
        minLength: minLength(6)
      },
      email: {
        required,
        email
      },
      password: {
        required,
        minLength: minLength(6)
      }
    },
    confirmPass: {
      required
    }
  },
  mounted() {
    this.setTitle();
  },
  methods: {
    ...mapActions("consumer", ["createConsumer"]),
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    setTitle() {
      if (this.modalType == "create") {
        this.title = "Add consumer";
      } else {
        this.title = "Modify consumer";
        this.email = this.inputUser.email;
      }
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
      this.form.password = null;
      this.form.confirmPass = null;
      this.form.id = null;
    },
    saveConsumer() {
      this.sending = true;

      this.createConsumer(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
            this.showNotification("The consumer was successfully added!");
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification("An error had occured");
          console.log(error);
        });
    },
    validateConsumer() {
      this.$v.$touch();
      if (!this.$v.$invalid) {
        console.log(this.form);
        this.saveConsumer();
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