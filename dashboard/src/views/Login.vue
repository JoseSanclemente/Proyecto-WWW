<template>
  <div class="centered-container">
    <div class="md-layout">
      <div class="login-image"></div>

      <md-content class="md-elevation-3">
        <div class="title">
          <img src="../assets/logo.png" />
          <div class="md-title">{{$t("Frens Co.")}}</div>
        </div>

        <form novalidate class="md-layout" @submit.prevent="validateLoginData">
          <md-field :class="getValidationClass('id')">
            <label>{{$t("User ID")}}</label>
            <md-input v-model="form.id" name="id" autocomplete="id" />
            <span class="md-error" v-if="!$v.form.id.required">{{$t("The user id is required")}}</span>
            <span class="md-error" v-else-if="!$v.form.minLength">{{$t("Invalid id")}}</span>
          </md-field>

          <md-field md-has-password :class="getValidationClass('password')">
            <label>{{$t('Password')}}</label>
            <md-input v-model="form.password" type="password" name="password"></md-input>
            <span
              class="md-error"
              v-if="!$v.form.password.required"
            >{{$t("The password is required")}}</span>
          </md-field>

          <img class="captcha-image" v-bind:src="'data:image/jpeg;base64,'+ captcha.captcha" />

          <md-field :class="getValidationClass('captcha')">
            <label>Enter Captcha</label>
            <md-input v-model="form.captcha"></md-input>
          </md-field>
        </form>

        <md-card-actions class="actions md-layout md-alignment-center-space-between">
          <md-button
            type="submit"
            class="md-raised md-accent"
            @click="validateLoginData"
            :disable="sending"
          >{{$t("Log in")}}</md-button>
          <md-progress-spinner md-mode="indeterminate" v-if="sending" />
        </md-card-actions>

        <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
      </md-content>
    </div>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, minLength } from "vuelidate/lib/validators";
import { mapActions, mapState, mapMutations } from "vuex";

export default {
  name: "consumer-login",
  mixins: [validationMixin],
  data() {
    return {
      form: {
        id: "",
        password: "",
        captcha: ""
      },
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
      password: {
        required
      },
      captcha: {
        required
      }
    }
  },
  methods: {
    ...mapActions("consumer", [
      "loadCaptcha",
      "sendLoginData",
      "listConsumers"
    ]),
    ...mapMutations("consumer", ["setLoggedConsumer"]),

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
      this.form["captcha_id"] = this.captcha.captcha_id;
      this.sending = true;

      this.sendLoginData(this.form)
        .then(() => {
          this.setLoggedConsumer(this.form.id);

          setTimeout(() => {
            this.sending = false;
            this.showNotification("Login successful!");
            this.$router.push("/consumer/contracts");
            this.$router.replace({
              name: "consumer-contracts",
              params: { id: this.form.id }
            });
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.loadCaptcha();
          this.showNotification("The information sent is incorrect");
          console.log(error);
          this.form.captcha = null;
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
    ...mapState("consumer", ["captcha"])
  },
  beforeMount() {
    this.listConsumers();
    this.loadCaptcha();
  }
};
</script>

<style lang="scss" scoped>
.captcha-image {
  background-color: white;
}

.login-image {
  background-size: cover;
  background-image: url("../assets/login_image.jpg");
  width: 30vw;
}

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