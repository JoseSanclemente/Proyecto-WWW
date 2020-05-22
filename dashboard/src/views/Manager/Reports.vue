<template>
  <div>
    <md-app md-waterfall md-mode="fixed">
      <md-app-toolbar class="md-primary">
        <span class="md-title">Dashboard</span>
      </md-app-toolbar>

      <md-app-drawer md-permanent="full">
        <md-toolbar class="md-transparent" md-elevation="0">Navigation</md-toolbar>

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

          <router-link to="/login">
            <md-list-item>
              <md-icon>exit_to_app</md-icon>
              <span class="md-list-item-text">{{$t("Sign out")}}</span>
            </md-list-item>
          </router-link>
        </md-list>
      </md-app-drawer>

      <md-app-content>
        <div class="md-layout md-glutter">
          <div class="md-layout-item">
            <md-card class="chart">
              <md-card-header>
                <div class="md-title">Monthly Consumption</div>
              </md-card-header>
              <md-card-content>
                <div class="md-layout md-gutter">
                  <div class="md-layout-item md-size-60">
                    <md-field>
                      <label for="startYear">Select year</label>
                      <md-select md-dense>
                        <md-option value="2018">2018</md-option>
                        <md-option value="2019">2019</md-option>
                        <md-option value="2020">2020</md-option>
                      </md-select>
                    </md-field>
                  </div>
                </div>
                <div class="md-layout">
                  <div class="md-layout-item chart-layout">
                    <bar-chart :data="chartData"></bar-chart>
                  </div>
                </div>
              </md-card-content>
            </md-card>
          </div>

          <div class="md-layout-item">
            <md-card class="chart">
              <md-card-header>
                <div class="md-title">Year Consumption</div>
              </md-card-header>
              <md-card-content>
                <div class="md-layout md-gutter">
                  <div class="md-layout-item">
                    <md-field>
                      <label for="startYear">Select start year</label>
                      <md-select md-dense>
                        <md-option value="2018">2018</md-option>
                        <md-option value="2019">2019</md-option>
                        <md-option value="2020">2020</md-option>
                      </md-select>
                    </md-field>
                  </div>

                  <div class="md-layout-item">
                    <md-field>
                      <label for="endYear">Select end year</label>
                      <md-select md-dense>
                        <md-option value="2018">2018</md-option>
                        <md-option value="2019">2019</md-option>
                        <md-option value="2020">2020</md-option>
                      </md-select>
                    </md-field>
                  </div>
                </div>
                <div class="md-layout">
                  <div class="md-layout-item chart-layout">
                    <bar-chart :data="chartData"></bar-chart>
                  </div>
                </div>
              </md-card-content>
            </md-card>
          </div>
        </div>

        <div class="md-layout md-size-50">
          <md-table md-card class="table">
            <md-table-toolbar>
              <h1 class="md-title">Top Consumers</h1>
            </md-table-toolbar>
            <md-table-row>
              <md-table-head md-numeric>ID</md-table-head>
              <md-table-head>Total consumption</md-table-head>
            </md-table-row>

            <md-table-row>
              <md-table-cell md-numeric>1</md-table-cell>
              <md-table-cell>310.020</md-table-cell>
            </md-table-row>

            <md-table-row>
              <md-table-cell md-numeric>2</md-table-cell>
              <md-table-cell>298.098</md-table-cell>
            </md-table-row>

            <md-table-row>
              <md-table-cell md-numeric>3</md-table-cell>
              <md-table-cell>Vera Taleworth</md-table-cell>
            </md-table-row>
          </md-table>
        </div>
      </md-app-content>
    </md-app>
  </div>
</template>

<script>
import BarChart from "@/components/charts/BarChart.vue";

export default {
  name: "Manager",
  components: {
    BarChart
  },
  data() {
    return {
      chartData: {
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
            data: [12, 19, 3, 5, 2, 3],
            backgroundColor: [
              "rgba(255, 99, 132, 0.2)",
              "rgba(54, 162, 235, 0.2)",
              "rgba(255, 206, 86, 0.2)",
              "rgba(75, 192, 192, 0.2)",
              "rgba(153, 102, 255, 0.2)",
              "rgba(255, 159, 64, 0.2)"
            ],
            borderColor: [
              "rgba(255, 99, 132, 1)",
              "rgba(54, 162, 235, 1)",
              "rgba(255, 206, 86, 1)",
              "rgba(75, 192, 192, 1)",
              "rgba(153, 102, 255, 1)",
              "rgba(255, 159, 64, 1)"
            ],
            borderWidth: 1
          }
        ]
      }
    };
  }
};
</script>

<style lang="scss" scoped>
.md-app {
  height: 100vh;
  border: 1px solid rgba(#000, 0.12);
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

