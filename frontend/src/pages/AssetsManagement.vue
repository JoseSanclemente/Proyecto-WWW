<template>
  <div class="content">
    <div class="assets-management">
      <div class="management-form">
        <md-button class="md-info" @click="showSubstationsList" v-if="selectedSubstation">
          <md-icon>navigate_before</md-icon>Substations list
          <md-tooltip md-direction="top">Back to substations list</md-tooltip>
        </md-button>
        <create-substation-form data-background-color="purple" v-if="!selectedSubstation"></create-substation-form>
        <search-substation-form
          data-background-color="purple"
          v-on:substation-details="showSubstationDetails"
          v-if="!selectedSubstation"
        ></search-substation-form>
        <substation-details
          v-if="selectedSubstation"
          data-background-color="purple"
          :substationID="substationID"
        ></substation-details>
        <search-transformer-form
          v-if="selectedSubstation"
          data-background-color="purple"
          v-on:substation-details="showSubstationDetails"
          :selectedSubstation="substationID"
        ></search-transformer-form>
      </div>
    </div>
  </div>
</template>

<script>
import {
  SearchSubstationForm,
  SubstationDetails,
  SearchTransformerForm,
  CreateSubstationForm
} from "@/pages";
import { mapActions } from "vuex";
export default {
  components: {
    SearchSubstationForm,
    SubstationDetails,
    SearchTransformerForm,
    CreateSubstationForm
  },

  data: function() {
    return {
      substationID: "",
      selectedSubstation: false
    };
  },

  mounted: function() {
    this.$store.dispatch("assetsManagement/loadAllSubstations");
  },

  methods: {
    showSubstationDetails: function(substationID) {
      this.substationID = substationID;
      this.selectedSubstation = true;
      this.$store.dispatch("assetsManagement/loadAllTransformers");
    },

    showSubstationsList: function() {
      this.selectedSubstation = false;
    }
  }
};
</script>