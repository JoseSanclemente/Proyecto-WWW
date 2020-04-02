<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Search substation</h4>
      </md-card-header>
      <div class="md-layout-item md-small-size-100 md-size-33 search-card">
        <md-field>
          <label>substation name</label>
          <md-input v-model="substationName" type="text"></md-input>
        </md-field>
      </div>
      <md-card-content>
        <substations-table
          :substationsList="substationsFiltered"
          table-header-color="purple"
          v-on:substation-details="
            substationID => $emit('substation-details', substationID)
          "
        ></substations-table>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { SubstationsTable } from "@/components";
import { mapState } from "vuex";

export default {
  name: "search-substation-form",

  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    }
  },

  components: {
    SubstationsTable
  },

  data() {
    return {
      substationName: ""
    };
  },

  computed: {
    ...mapState({
      substations: state => state.assetsManagement.substations
    }),

    substationsFiltered: function() {
      if (this.substationName === "") {
        return this.substations;
      }

      let newSubstationList = [];
      for (let i = 0; i < this.substations.length; i++) {
        let substation = this.substations[i];

        let normSubstationName = substation.name
          .split(" ")
          .join("")
          .toUpperCase();

        let normSearchTerm = this.substationName
          .split(" ")
          .join("")
          .toUpperCase();

        if (normSubstationName.includes(normSearchTerm)) {
          newSubstationList.push(substation);
        }
      }

      return newSubstationList;
    }
  }
};
</script>

<style scoped>
.search-card {
  display: flex;
  margin: 2rem 0 2rem 1rem;
}

.search-button {
  margin: 1.5rem 0 0 2rem;
}
</style>
