<template>
  <div>
    <md-field v-if="showAddressInput">
      <label for="address">{{$t("Address")}}</label>
      <md-input name="address" v-model="address" />
      <md-button class="md-icon-button md-primary" @click="getLatLng">
        <md-icon>search</md-icon>
      </md-button>
    </md-field>
    <div id="map" ref="map">
      <map-marker
        v-for="sub in  substations"
        v-bind:key="sub.id"
        v-bind:lat="parseFloat(sub.latitude)"
        v-bind:lng="parseFloat(sub.longitude)"
        type="substation"
      ></map-marker>

      <map-marker
        v-for="trans in  transformers"
        v-bind:key="trans.id"
        v-bind:lat="parseFloat(trans.latitude)"
        v-bind:lng="parseFloat(trans.longitude)"
        type="transformer"
      ></map-marker>
    </div>
  </div>
</template>

<script>
import MapMarker from "@/components/map/MapMarker.vue";
import { markersWithPivots } from "@/helpers/map.js";
export default {
  components: {
    MapMarker
  },
  props: {
    showAddressInput: {
      required: true
    },
    mapType: {
      required: true
    },
    substations: {
      required: false
    },
    transformers: {
      required: false
    }
  },
  data() {
    return {
      map: null,
      geocoder: null,
      address: null,
      marker: null,
      pos: {
        lat: null,
        lng: null
      },
      options: {
        center: { lat: 3.4516, lng: -76.532 },
        zoom: 12
      }
    };
  },
  mounted() {
    this.map = new window.google.maps.Map(this.$refs["map"], this.options);
    this.geocoder = new window.google.maps.Geocoder();
    this.drawLines();
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
          _this.pos.lat = results[0].geometry.location.lat();
          _this.pos.lng = results[0].geometry.location.lng();

          let newLatLng = { lat: _this.pos.lat, lng: _this.pos.lng };

          _this.map.setZoom(16);
          _this.map.setCenter(newLatLng);

          if (_this.marker == null) {
            _this.marker = new window.google.maps.Marker({
              position: newLatLng,
              map: _this.map,
              label: _this.getLabel()
            });
          } else {
            _this.marker.setPosition(newLatLng);
          }
        } else {
          console.log("failed");
        }
      });

      this.returnLatLng();
    },
    returnLatLng() {
      setTimeout(() => {
        this.$emit("latlng", this.pos);
      }, 1000);
    },
    drawLines() {
      if (this.substations != null) {
        let polyLines = markersWithPivots(this.substations);
        for (let i = 0; i < polyLines.length; i++) {
          var line = new window.google.maps.Polyline({
            path: polyLines[i],
            geodesic: true,
            strokeColor: "#FF0000",
            strokeOpacity: 1.0,
            strokeWeight: 2
          });

          line.setMap(this.map);
        }
      }
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