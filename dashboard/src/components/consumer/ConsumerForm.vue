<template>
  <div class="modal" v-show="value">
    <form novalidate class="md-layout" @submit.prevent="validateConsumer">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">{{ $t("Add consumer") }}</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field :class="getValidationClass('id')">
                <label for="email">{{$t("ID number")}}</label>
                <md-input name="id" v-model="form.id" :disabled="sending" />
                <span class="md-error" v-if="!$v.form.id.required">{{$t("The ID is required")}}</span>
                <span class="md-error" v-else-if="!$v.form.minLength">{{$t("Invalid id")}}</span>
              </md-field>
            </div>
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field :class="getValidationClass('email')">
                <label for="email">Email</label>
                <md-input v-model="form.email" :disabled="sending" />
                <span
                  class="md-error"
                  v-if="!$v.form.email.required"
                >{{$t("The email is required")}}</span>
                <span class="md-error" v-else-if="!$v.form.email.email">{{$t("Invalid email")}}</span>
              </md-field>
            </div>
            <div class="md-layout-item">
              <md-field>
                <label for>{{$t("Address")}}</label>
                <md-input v-model="contractForm.address" :disabled="sending" />
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
                <span
                  class="md-error"
                  v-if="!$v.form.password.minLength"
                >{{$t("The password is too short!")}}</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('confirmPass')">
                <label for="confirmPass">{{$t("Confirm password")}}</label>
                <md-input type="password" v-model="confirmPass" :disabled="sending" />
                <span
                  class="md-error"
                  v-if="!$v.confirmPass.required"
                >{{$t("This field is required")}}</span>
              </md-field>
            </div>
          </div>

          <h2 class="md-title">{{$t("Contract")}}</h2>
          <div class="md-layout md-gutter">
            <div class="md-layout-item">
              <md-field>
                <label for>{{$t("Bill notification")}}</label>
                <md-select v-model="contractForm.notification_type">
                  <md-option value="email">{{$t("Email")}}</md-option>
                  <md-option value="sendToAddress">{{$t("Send to address")}}</md-option>
                </md-select>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item">
              <md-field>
                <label for>{{$t("Transformer ID")}}</label>
                <md-select v-model="contractForm.transformer_id">
                  <md-option
                    v-for="trans in transformers"
                    v-bind:key="trans.id"
                    v-bind:value="trans.id"
                  >{{ trans.id }}</md-option>
                </md-select>
              </md-field>
            </div>
          </div>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button @click.prevent="close">{{$t("Close")}}</md-button>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">{{$t("Add")}}</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
    </form>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, email, minLength } from "vuelidate/lib/validators";
import { mapActions, mapState } from "vuex";

export default {
  name: "consumer-form",
  mixins: [validationMixin],
  props: {
    value: { required: true }
  },
  computed: {
    ...mapState("substation", ["transformers"])
  },
  data() {
    return {
      form: {
        id: null,
        email: null,
        password: null
      },
      contractForm: {
        consumer_id: null,
        transformer_id: null,
        address: null,
        notification_type: null,
        email: null
      },
      confirmPass: null,
      showSnackBar: false,
      sending: false,
      message: null
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
  methods: {
    ...mapActions("consumer", ["createConsumer", "createContract"]),
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
            this.saveContract();
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification("An error has occured while creating consumer");
          console.log(error);
        });
    },
    saveContract() {
      this.contractForm.consumer_id = this.form.id;
      this.contractForm.email = this.form.email;

      this.createContract(this.contractForm)
        .then(() => {
          this.showNotification("The consumer was successfully created!");
        })
        .catch(error => {
          this.showNotification(
            "An error has occured while creating the contract"
          );
          console.log(error);
        });
    },
    validateConsumer() {
      this.$v.$touch();
      if (!this.$v.$invalid) {
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