<template>
  <div class="centered-container">
    <md-content class="md-elevation-3">
      <div class="title">
        <img src="https://vuematerial.io/assets/logo-color.png" />
        <div class="md-title">{{$t("Electricaribe")}}</div>
        <div class="md-subtitle">{{$t("Employees")}}</div>
      </div>

      <form novalidate class="md-layout" @submit.prevent="validateLoginData">
        <md-field :class="getValidationClass('email')">
          <label>{{$t("Email")}}</label>
          <md-input v-model="form.email" name="email" autocomplete="email" />
          <span class="md-error" v-if="!$v.form.email.required">{{$t("The user email is required")}}</span>
          <span class="md-error" v-else-if="!$v.form.minLength">{{$t("Invalid email")}}</span>
        </md-field>

        <md-field md-has-password :class="getValidationClass('password')">
          <label>{{$t('Password')}}</label>
          <md-input v-model="form.password" type="password" name="password"></md-input>
          <span
            class="md-error"
            v-if="!$v.form.password.required"
          >{{$t("The password is required")}}</span>
        </md-field>

      </form>

      <md-card-actions class="actions md-layout md-alignment-center-space-between">
        <md-button
          type="submit"
          class="md-raised md-primary"
          @click="validateLoginData"
          :disable="sending"
        >{{$t("Log in")}}</md-button>
      </md-card-actions>

      <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
    </md-content>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, minLength } from "vuelidate/lib/validators";
import { mapActions, mapState } from "vuex";

export default {
  name: "user-login",
  mixins: [validationMixin],
  data() {
    return {
      form: {
        email: "",
        password: "",
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
      email: {
        required,
        minLength: minLength(6)
      },
      password: {
        required
      }
    }
  },
  methods: {
    ...mapActions("user", ["sendLoginData"]),

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
    LoginData() {

      //console.log(JSON.stringify(this.form))
      this.sending = true;
      this.sendLoginData(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
            this.showNotification("Login successful!");
            this.$router.push('')
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification("The email or password is incorrect, Check your email then type your password again");
          console.log(error);
        });
    },
    validateLoginData() {
      this.$v.$touch();
      if (!this.$v.$invalid) {
        this.LoginData();
      }
    }
  },
  computed: {
    ...mapState()
  },
  watch: {},
  mounted() {},
  created() {
    this.loadCaptcha();
  }
};
</script>

<style lang="scss" scoped>
.centered-container {
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  height: 100vh;
  .title {
    text-align: center;
    margin-bottom: 30px;
    img {
      margin-bottom: 16px;
      max-width: 80px;
    }
  }
  .actions {
    .md-button {
      margin: 0;
    }
  }
  .form {
    margin-bottom: 60px;
  }
  .background {
    position: absolute;
    height: 100%;
    width: 100%;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    z-index: 0;
  }
  .md-content {
    z-index: 1;
    padding: 40px;
    width: 100%;
    max-width: 400px;
    position: relative;
  }
  .loading-overlay {
    z-index: 10;
    top: 0;
    left: 0;
    right: 0;
    position: absolute;
    width: 100%;
    height: 100%;
    background: rgba(255, 255, 255, 0.9);
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
</style>