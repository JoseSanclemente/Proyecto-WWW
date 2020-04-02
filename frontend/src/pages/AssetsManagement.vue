<template>
  <div class="content">
    <div class="assets-management">
      <div class="management-form">
        <search-substation-form
          data-background-color="purple"
        ></search-substation-form>
      </div>
    </div>
  </div>
</template>

<script>
import { SearchSubstationForm } from "@/pages";
import { mapActions } from "vuex";
export default {
  components: {
    SearchSubstationForm
  },

  mounted: function() {
    this.$store.dispatch("assetsManagement/loadAllSubstations");
  }
};
</script>

<style scoped>
.management-form {
  display: block;
  margin: 0 auto;
  width: 90%;
}
</style>
