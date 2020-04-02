<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Search transformer</h4>
      </md-card-header>
      <div class="md-layout-item md-small-size-100 md-size-33 search-card">
        <md-field>
          <label>Transformer name</label>
          <md-input v-model="transformerName" type="text"></md-input>
        </md-field>
      </div>
      <md-card-content>
        <transformers-table
          :transformersList="transformersFiltered"
          table-header-color="purple"
          v-on:transformer-details="
            transformerID => $emit('transformer-details', transformerID)
          "
        ></transformers-table>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { TransformersTable } from "@/components";
import { mapState } from "vuex";

export default {
  name: "search-transformer-form",

  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    },

    selectedSubstation: {
      type: String,
      default: ""
    }
  },

  components: {
    TransformersTable
  },

  data() {
    return {
      transformerName: ""
    };
  },

  computed: {
    ...mapState({
      transformers: state => state.assetsManagement.transformers
    }),

    transformersFiltered: function() {
      let newTransformersList = [];
      for (let i = 0; i < this.transformers.length; i++) {
        let transformer = this.transformers[i];

        let normTransformerName = transformer.name
          .split(" ")
          .join("")
          .toUpperCase();

        let normSearchTerm = this.transformerName
          .split(" ")
          .join("")
          .toUpperCase();

        if (normTransformerName.includes(normSearchTerm) && transformer.subestation === this.selectedSubstation) {
          newTransformersList.push(transformer);
        }
      }

      return newTransformersList;
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
