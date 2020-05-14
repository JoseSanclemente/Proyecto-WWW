<script>
export default {
  name: "map-marker",
  props: {
    lat: { type: Number, required: true },
    lng: { type: Number, required: true },
    type: { type: String, required: true }
  },
  data: () => ({
    marker: null
  }),
  mounted() {
    this.$parent.getMap(map => {
      this.marker = new window.google.maps.Marker({
        position: { lat: this.lat, lng: this.lng },
        map: map,
        icon: {
          url: this.getMarkerIcon()
        }
      });
    });
  },
  beforeDestroy() {
    this.marker.setMap(null);
    window.google.maps.event.clearInstanceListeners(this.marker);
  },
  render() {
    return null;
  },
  methods: {
    getMarkerIcon() {
      if (this.type == "substation") {
        return "http://maps.google.com/mapfiles/kml/paddle/purple-circle.png";
      } else {
        return "http://maps.google.com/mapfiles/kml/paddle/blu-circle.png";
      }
    }
  }
};
</script>