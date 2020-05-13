<template>
  <div>
    <md-field v-if="showAddressInput">
      <label for="address">Address</label>
      <md-input name="address" v-model="address" />
      <md-button class="md-icon-button md-primary" @click="getLatLng">
        <md-icon>search</md-icon>
      </md-button>
    </md-field>
    <div id="map" ref="map"></div>
  </div>
</template>

<script>
export default {
  props: {
    showAddressInput: {
      required: true
    },
    mapType: {
      required: true
    }
  },
  data() {
    return {
      map: null,
      geocoder: null,
      address: null,
      transMarker: null,
      marker: {
        lat: null,
        lng: null
      },
      options: {
        center: { lat: 3.4516, lng: -76.532 },
        zoom: 14
      }
    };
  },
  mounted() {
    this.map = new window.google.maps.Map(this.$refs["map"], this.options);
    this.geocoder = new window.google.maps.Geocoder();
  },
  methods: {
    getMap(callback) {
      let vm = this;
      function checkForMap() {
        if (vm.map) callback(vm.map);
        else setTimeout(checkForMap, 200);
      }
      checkForMap();
    },
    getLabel() {
      if (this.mapType == "substation") {
        return "S";
      } else {
        return "T";
      }
    },
    getLatLng() {
      let _this = this;
      let formattedAddres = this.address + " Cali Colombia";
      this.geocoder.geocode({ address: formattedAddres }, function(
        results,
        status
      ) {
        if (status === "OK") {
          _this.marker.lat = results[0].geometry.location.lat();
          _this.marker.lng = results[0].geometry.location.lng();

          let newLatLng = { lat: _this.marker.lat, lng: _this.marker.lng };

          _this.map.setZoom(16);
          _this.map.setCenter(newLatLng);

          if (_this.transMarker == null) {
            _this.transMarker = new window.google.maps.Marker({
              position: newLatLng,
              map: _this.map,
              label: _this.getLabel()
            });
          } else {
            _this.transMarker.setPosition(newLatLng);
          }
        } else {
          console.log("failed");
        }
      });

      this.returnLatLng();
    },
    returnLatLng() {
      this.$emit("latlng", this.marker);
    }
  }
};
</script>

<style lang="scss" scoped>
#map {
  height: 50vh;
  width: 100%;
  width: 100%;
}
</style>