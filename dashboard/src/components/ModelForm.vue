<template>
  <div class="modal" v-show="value">
    <form novalidate class="md-layout" @submit.prevent="validateUser">
      <md-card class="md-layout-item md-small-size-100">
        <md-card-header>
          <div class="md-title">Add Transformer Model</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('model')">
                <label for="model">Model</label>
                <md-input name="model" id="model" v-model="form.model" :disabled="sending" />
                <span class="md-error" v-if="!$v.form.model.required">The model is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('manufacturer')">
                <label for="manufacturer">Manufacturer</label>
                <md-input
                  name="manufacturer"
                  id="manufacturer"
                  v-model="form.manufacturer"
                  :disabled="sending"
                />
                <span
                  class="md-error"
                  v-if="!$v.form.manufacturer.required"
                >The manufacturer is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('serialNum')">
                <label for="serialNum">Serial Number</label>
                <md-input
                  name="serialNum"
                  id="serialNum"
                  v-model="form.serialNum"
                  :disabled="sending"
                />
                <span
                  class="md-error"
                  v-if="!$v.form.serialNum.required"
                >The serial number is required</span>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('primaryVolt')">
                <label for="primaryVolt">Primary Voltage</label>
                <md-input
                  type="primaryVolt"
                  id="primaryVolt"
                  name="primaryVolt"
                  autocomplete="primaryVolt"
                  v-model="form.primaryVolt"
                  :disabled="sending"
                />
                <span
                  class="md-error"
                  v-if="!$v.form.primaryVolt.required"
                >The primary voltage is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('secondaryVolt')">
                <label for="secondaryVolt">Secondary Voltage</label>
                <md-input
                  type="secondaryVolt"
                  id="secondaryVolt"
                  name="secondaryVolt"
                  autocomplete="secondaryVolt"
                  v-model="form.secondaryVolt"
                  :disabled="sending"
                />
                <span
                  class="md-error"
                  v-if="!$v.form.secondaryVolt.required"
                >The secondary voltage is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('outputCurrent')">
                <label for="outputCurrent">Output current</label>
                <md-input
                  type="outputCurrent"
                  id="outputCurrent"
                  name="outputCurrent"
                  autocomplete="outputCurrent"
                  v-model="form.outputCurrent"
                  :disabled="sending"
                />
                <span
                  class="md-error"
                  v-if="!$v.form.outputCurrent.required"
                >The output current is required</span>
              </md-field>
            </div>
          </div>

          <div class="md-layout md-gutter">
            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('weight')">
                <label for="weight">Weight</label>
                <md-input
                  type="weight"
                  id="weight"
                  name="weight"
                  autocomplete="weight"
                  v-model="form.weight"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.weight.required">The weight is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('height')">
                <label for="height">Height</label>
                <md-input
                  type="height"
                  id="height"
                  name="height"
                  autocomplete="height"
                  v-model="form.height"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.height.required">The height is required</span>
              </md-field>
            </div>

            <div class="md-layout-item md-small-size-100">
              <md-field :class="getValidationClass('width')">
                <label for="width">Width</label>
                <md-input
                  type="width"
                  id="width"
                  name="width"
                  autocomplete="width"
                  v-model="form.width"
                  :disabled="sending"
                />
                <span class="md-error" v-if="!$v.form.width.required">The width is required</span>
              </md-field>
            </div>
          </div>

          <md-field :class="getValidationClass('depth')">
            <label for="depth">Depth</label>
            <md-input
              type="depth"
              name="depth"
              id="depth"
              autocomplete="depth"
              v-model="form.depth"
              :disabled="sending"
            />
            <span class="md-error" v-if="!$v.form.depth.required">The depth is required</span>
          </md-field>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button @click.prevent="close">Cancel</md-button>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">Add model</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="userSaved">The user {{ lastUser }} was saved with success!</md-snackbar>
    </form>
  </div>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, minLength, maxLength } from "vuelidate/lib/validators";

export default {
  name: "model-form",
  mixins: [validationMixin],
  props: {
    value: {
      required: true
    }
  },
  data: () => ({
    form: {
      model: null,
      manufacturer: null,
      serialNum: null,
      primaryVolt: null,
      secondaryVolt: null,
      outputCurrent: null,
      weight: null,
      height: null,
      width: null,
      depth: null
    },
    userSaved: false,
    sending: false,
    lastUser: null
  }),
  validations: {
    form: {
      model: {
        required,
        minLength: minLength(3)
      },
      manufacturer: {
        required,
        minLength: minLength(3)
      },
      serialNum: {
        required,
        maxLength: maxLength(3)
      },
      primaryVolt: {
        required
      },
      secondaryVolt: {
        required
      },
      outputCurrent: {
        required
      },
      weight: {
        required
      },
      height: {
        required
      },
      width: {
        required
      },
      depth: {
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
  height: 100vh;
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