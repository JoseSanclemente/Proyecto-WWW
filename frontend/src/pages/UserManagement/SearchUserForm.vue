<template>
  <div>
    <md-card>
      <md-card-header :data-background-color="dataBackgroundColor">
        <h4 class="title">Search user</h4>
      </md-card-header>
      <div class="md-layout-item md-small-size-100 md-size-33 search-card">
        <md-field md-clearable>
          <label>Filter</label>
          <md-input v-model="filter" type="text"></md-input>
        </md-field>
      </div>
      <md-card-content>
        <ordered-table
          :userList="users"
          table-header-color="purple"
        ></ordered-table>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
import { OrderedTable } from "@/components";
import axios from "axios";
import { USERS_FETCH, USER_SET_FILTER } from "@/store/user";
import { mapGetters } from "vuex";
export default {
  name: "search-user-form",
  props: {
    dataBackgroundColor: {
      type: String,
      default: ""
    }
  },
  computed: {
    ...mapGetters(["users"])
  },
  watch: {
    filter: function(val) {
      this.$store.commit(USER_SET_FILTER, val);
    }
  },
  components: {
    OrderedTable
  },
  data() {
    return {
      filter: ""
    };
  },
  methods: {
    getUser() {
      this.$store.dispatch(USERS_FETCH);
    }
  },
  created() {
    this.getUser();
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
