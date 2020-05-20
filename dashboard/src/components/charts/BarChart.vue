<script>
import { Bar } from "vue-chartjs";

export default {
  extends: Bar,
  name: "bar-chart",
  props: {
    data: { required: true }
  },
  data() {
    return {
      options: {
        maintainAspectRatio: false,
        scales: {
          yAxes: [
            {
              ticks: {
                beginAtZero: true
              }
            }
          ]
        }
      }
    };
  },
  mounted() {
    this.renderChart(this.data, this.options);
  }
};
</script>