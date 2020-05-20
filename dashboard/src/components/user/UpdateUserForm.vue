<template>
  <div class="modal" v-show="show">
    <form novalidate class="md-layout" @submit.prevent="update">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">{{ $t("Modify user") }}</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout">
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field>
                <label for="type">{{$t("Role")}}</label>
                <md-select name="type" v-model="form.type" md-dense :disabled="sending">
                  <md-option value="admin">{{$t("Admin")}}</md-option>
                  <md-option value="manager">{{$t("Manager")}}</md-option>
                  <md-option value="operator">{{$t("Operator")}}</md-option>
                </md-select>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field>
                <label for="password">{{$t("New password")}}</label>
                <md-input type="password" v-model="form.password" :disabled="sending" />
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field>
                <label for="confirmPass">{{$t("Confirm password")}}</label>
                <md-input type="password" v-model="confirmPass" :disabled="sending" />
              </md-field>
            </div>
          </div>

          <md-checkbox v-model="form.deleted" class="md-primary">{{$t("Deactivate")}}</md-checkbox>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button @click.prevent="close">{{$t("Close")}}</md-button>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">{{$t("Modify")}}</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
    </form>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "update-user-form",
  props: {
    show: { required: true },
    user: { required: true }
  },
  data() {
    return {
      form: {
        email: null,
        type: null,
        password: null,
        deleted: false
      },
      confirmPass: null,
      showSnackBar: false,
      sending: false,
      message: null
    };
  },
  mounted() {
    this.form.email = this.user.email;
    this.form.type = this.user.type;
    this.form.deleted = this.user.deleted;
  },
  methods: {
    ...mapActions("user", ["updateUser"]),
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    clearForm() {
      this.form.password = null;
      this.confirmPass = null;
    },
    update() {
      this.sending = true;

      this.updateUser(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;

            let message = "The user was successfully updated!";
            this.showNotification(message);

            this.clearForm();
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.clearForm();
          this.showNotification(
            this.$t("An error had occurred while updating user")
          );
          console.log(error);
        });
    },
    close() {
      this.clearForm();
      this.$emit("destroyModal", !this.value);
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