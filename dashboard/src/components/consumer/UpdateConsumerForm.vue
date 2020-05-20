<template>
  <div class="modal" v-show="show">
    <form novalidate class="md-layout" @submit.prevent="update">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">{{ $t("Modify consumer") }}</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout">
            <div class="md-layout-item md-size-50 md-small-size-100">
              <md-field>
                <label for="email">Email</label>
                <md-input v-model="form.email" :disabled="sending" />
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
  name: "update-consumer-form",
  props: {
    show: { required: true },
    consumer: { required: true }
  },
  data() {
    return {
      form: {
        id: null,
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
    this.form.id = this.consumer.id;
    this.form.email = this.consumer.email;
    this.form.type = this.consumer.type;
    this.form.deleted = this.consumer.deleted;
  },
  methods: {
    ...mapActions("consumer", ["updateConsumer"]),
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    clearForm() {
      this.form.email = null;
      this.form.password = null;
      this.confirmPass = null;
    },
    update() {
      this.sending = true;

      this.updateConsumer(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
            this.showNotification("The consumer was successfully updated!");

            this.clearForm();
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification(
            "An error had occurred while updating consumer"
          );

          this.clearForm();
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