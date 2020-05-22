<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">Dashboard</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">{{$t("Navigation")}}</md-toolbar>

        <md-list>
          <router-link to="/manager/reports">
            <md-list-item>
              <md-icon>assessment</md-icon>
              <span class="md-list-item-text">{{$t("Reports")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/manager/users">
            <md-list-item>
              <md-icon>account_circle</md-icon>
              <span class="md-list-item-text">{{$t("Users")}}</span>
            </md-list-item>
          </router-link>

          <router-link to="/employee/login">
            <md-list-item>
              <md-icon>exit_to_app</md-icon>
              <span class="md-list-item-text">{{$t("Sign out")}}</span>
            </md-list-item>
          </router-link>

          <md-list-item>
            <md-field>
              <label>{{$t("Language")}}</label>
              <md-select v-model="$i18n.locale">
                <md-option
                  v-for="(lang, i) in langs"
                  :key="`Lang${i}`"
                  :value="lang.code"
                >{{ lang.name }}</md-option>
              </md-select>
            </md-field>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content class="content">
        <div class="md-layout md-glutter">
          <div class="md-layout-item">
            <md-card class="chart">
              <md-card-header>
                <div class="md-title">{{$t("Monthly Consumption")}}</div>
              </md-card-header>
              <md-card-content>
                <div class="md-layout md-gutter">
                  <div class="md-layout-item md-size-60">
                    <md-field>
                      <label for="startYear">{{$t("Select year")}}</label>
                      <md-select v-model="inputMonthly" md-dense>
                        <md-option value="2020">2020</md-option>
                      </md-select>
                    </md-field>
                  </div>
                </div>
                <div class="md-layout">
                  <div class="md-layout-item chart-layout">
                    <bar-chart :data="monthlyData"></bar-chart>
                  </div>
                </div>
              </md-card-content>
            </md-card>
          </div>

          <div class="md-layout-item">
            <md-card class="chart">
              <md-card-header>
                <div class="md-title">{{$t("Annual Consumption")}}</div>
              </md-card-header>
              <md-card-content>
                <div class="md-layout md-gutter">
                  <div class="md-layout-item">
                    <md-field>
                      <label for="startYear">{{$t("Select start year")}}</label>
                      <md-select v-model="inputStartYear" md-dense>
                        <md-option value="2020">2020</md-option>
                      </md-select>
                    </md-field>
                  </div>

                  <div class="md-layout-item">
                    <md-field>
                      <label for="endYear">{{$t("Select end year")}}</label>
                      <md-select v-model="inputEndYear" md-dense>
                        <md-option value="2020">2020</md-option>
                      </md-select>
                    </md-field>
                  </div>
                </div>
                <div class="md-layout">
                  <div class="md-layout-item chart-layout">
                    <bar-chart :data="yearlyData"></bar-chart>
                  </div>
                </div>
              </md-card-content>
            </md-card>
          </div>
        </div>

        <div class="md-layout md-size-50">
          <md-table v-model="topConsumersList" md-card md-fixed-header>
            <md-table-toolbar class="md-primary" slot="md-table-alternate-header">
              <span class="md-title">{{$t("Top Consumers")}}</span>
            </md-table-toolbar>

            <md-table-row slot="md-table-row" slot-scope="{ item }">
              <md-table-cell :md-label="$t('ID')" md-sort-by="id">{{ item.id }}</md-table-cell>
              <md-table-cell :md-label="$t('Value')">{{ item.value }}</md-table-cell>
            </md-table-row>
          </md-table>
        </div>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import BarChart from "@/components/charts/BarChart.vue";
import { mapActions, mapState } from "vuex";

export default {
  name: "Manager",
  components: {
    BarChart
  },
  data() {
    return {
      inputMonthly: "2020",
      inputStartYear: "2020",
      inputEndYear: "2020",
      monthlyData: {
        labels: [
          "Jan",
          "Feb",
          "Mar",
          "Apr",
          "May",
          "Jun",
          "Jul",
          "Aug",
          "Sep",
          "Oct",
          "Nov",
          "Dec"
        ],
        datasets: [
          {
            label: "Total consumption",
            data: [163, 250, 353, 380, 204],
            backgroundColor: [
              "rgba(255, 206, 86, 0.2)",
              "rgba(255, 206, 86, 0.2)",
              "rgba(255, 206, 86, 0.2)",
              "rgba(255, 206, 86, 0.2)",
              "rgba(255, 206, 86, 0.2)",
              "rgba(255, 206, 86, 0.2)"
            ],
            borderColor: [
              "rgba(255, 206, 86, 0.9)",
              "rgba(255, 206, 86, 0.9)",
              "rgba(255, 206, 86, 0.9)",
              "rgba(255, 206, 86, 0.9)",
              "rgba(255, 206, 86, 0.9)",
              "rgba(255, 206, 86, 0.9)"
            ],
            borderWidth: 1
          }
        ]
      },
      yearlyData: {
        labels: ["2020"],
        datasets: [
          {
            label: "Total consumption",
            data: [1350],
            backgroundColor: ["rgba(90, 86, 204, 0.5)"],
            borderColor: ["rgba(90, 86, 204, 1)"],
            borderWidth: 1
          }
        ]
      },
      langs: [
        { name: "English", code: "en" },
        { name: "Español", code: "es" },
        { name: "Português", code: "pt" }
      ]
    };
  },
  computed: {
    ...mapState("report", ["topConsumersList"])
  },
  created() {
    this.topConsumers();
  },
  methods: {
    ...mapActions("report", ["topConsumers"]),
    getMonthlyData() {
      if (this.monthYear == 2020) {
        this.monthlyConsumption.year = 2020;
        this.monthlyConsumption = ["Jan", "Feb", "Mar", "Apr", "May"];
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.md-app {
  height: 100vh;
  border: 1px solid rgba(#000, 0.12);
}

.content {
  background-color: #333333;
}

.chart {
  margin: 2em 0.2em;
  width: 40vw;
}

.table {
  margin: 2em 0.2em;
}

// Demo purposes only
.md-drawer {
  width: 200px;
  max-width: calc(100vw - 125px);
}
</style>

