<template>
  <div>
    <update-consumer-form
      v-if="render"
      :show="modalOpen"
      :consumer="inputUser"
      @destroyModal="destroyModal"
    ></update-consumer-form>

    <h1 class="md-title">{{ $t("Consumers") }}</h1>
    <md-table v-if="consumers!= null && consumers.length > 0" v-model="consumers" md-card>
      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
        <md-table-cell :md-label="$t('Email')">{{ item.email }}</md-table-cell>
        <md-table-cell :md-label="$t('Status')">{{ $t(getStatus(item.deleted)) }}</md-table-cell>
        <md-table-cell>
          <md-button class="md-icon-button md-primary" @click="openUpdateModal(item)">
            <md-icon>edit</md-icon>
          </md-button>
        </md-table-cell>
      </md-table-row>
    </md-table>
  </div>
</template>

<script>
import { mapState } from "vuex";
import UpdateConsumerForm from "@/components/consumer/UpdateConsumerForm.vue";
import { getStatusLabel } from "@/helpers/helpers.js";
export default {
  name: "consumer-table",
  components: {
    UpdateConsumerForm
  },
  data: () => ({
    modalOpen: false,
    showSnackBar: false,
    message: "",
    render: false
  }),
  computed: {
    ...mapState("consumer", ["consumers"])
  },
  methods: {
    getStatus(status) {
      return getStatusLabel(status);
    },
    openUpdateModal(item) {
      this.render = true;
      this.inputUser = item;
      this.modalOpen = true;
    },
    showNotification(input) {
      this.message = input;
      this.showSnackBar = true;
    },
    destroyModal(value) {
      this.modalOpen = value;
      this.render = false;
    }
  }
};
</script>

<style scoped>
section {
  margin: 1.2em 0;
}
</style>