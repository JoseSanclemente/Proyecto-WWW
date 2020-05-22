<template>
  <div>
    <form novalidate class="md-layout" @submit.prevent="saveSubstation">
      <md-card class="md-layout-item">
        <md-card-header>
          <div class="md-title">{{$t(title)}}</div>
        </md-card-header>

        <md-card-content>
          <div class="md-layout-item md-size-50">
            <md-field>
              <label for>{{$t("Name")}}</label>
              <md-input v-model="form.name" />
            </md-field>
          </div>

          <div class="md-layout md-gutter" v-show="showCancel">
            <div class="md-layout-item md-size-50 md-small-size-50">
              <md-field>
                <label for="deleted">{{$t("Active")}}</label>
                <md-select v-model="form.deleted" md-dense :disabled="sending">
                  <md-option
                          value=false
                  >{{ $t("Yes") }}</md-option>
                  <md-option
                          value=true
                  >{{ $t("No") }}</md-option>
                </md-select>
              </md-field>
            </div>
          </div>

          <map-component mapType="substation" :showAddressInput="!showCancel" @latlng="setLatLng" :substations="substationMap"></map-component>
        </md-card-content>

        <md-progress-bar md-mode="indeterminate" v-if="sending" />

        <md-card-actions>
          <md-button class="md-default md-raised" v-show="showCancel" @click="$emit('edit-cancelled')">{{$t("Cancel")}}</md-button>
          <md-button type="submit" class="md-primary md-raised" :disabled="sending">{{$t(confirmButtonText)}}</md-button>
        </md-card-actions>
      </md-card>

      <md-snackbar :md-active.sync="showSnackBar">{{ $t(message) }}</md-snackbar>
    </form>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import MapComponent from "@/components/map/MapComponent.vue";

export default {
  name: "substation-form",
  components: {
    MapComponent
  },
  props:{
    modify: {type: Object, required: false},
  },
  data() {
    return {
      title: "Add Substation",
      showCancel: false,
      substationMap: [],
      confirmButtonText: "Add",
      form: {
        name: null,
        latitude: null,
        longitude: null,
        deleted: null,
        id: null,
//        transformers: null,
      },
      showSnackBar: false,
      message: null,
      sending: false
    };
  },
  watch: {
    modify: function (modify) {
      if (modify == null) {
        this.title = "Add Substation"
        this.showCancel= false
        this.substationMap= []
        this.form.name= null
        this.form.latitude= null
        this.form.longitude= null
        this.form.deleted= null
        this.confirmButtonText="Add"
        this.form.id= null
        //this.form.transformers= null

        return

      }

      this.title = "Modify Substation"
      this.showCancel= true
      this.substationMap.push(modify)
      this.form.name= modify.name
      this.form.latitude= modify.latitude
      this.form.longitude= modify.longitude
      this.form.deleted= modify.deleted
      this.confirmButtonText="Modify"
      this.form.id= modify.id
     // this.form.transformers= modify.transformers
    }

  },
  methods: {
    ...mapActions("substation", ["createSubstation", "updateSubstation", "listSubstations"]),
    setLatLng(marker) {
      this.form.latitude = marker.lat;
      this.form.longitude = marker.lng;
      console.log(this.form);
    },
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    saveSubstation() {
      this.sending = true;

      if (this.title == "Modify Substation") {
        this.updateSubstation(this.form)
                .then(() => {
                  setTimeout(() => {
                    this.sending = false;
                    this.showNotification("The substation was successfully added!");
                  }, 2000);
                  this.listSubstations()
                  this.$emit('edit-cancelled')
                })
                .catch(error => {
                  this.sending = false;
                  this.showNotification("An error had occured");
                  console.log(error);
                  this.$emit('edit-cancelled')
                });
        return
      }

      this.createSubstation(this.form)
        .then(() => {
          setTimeout(() => {
            this.sending = false;
            this.showNotification("The substation was successfully added!");
            this.listSubstations()
          }, 2000);
        })
        .catch(error => {
          this.sending = false;
          this.showNotification("An error had occured");
          console.log(error);
        });
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
</style>