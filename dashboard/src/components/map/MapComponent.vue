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
//import { mapState } from "vuex";
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
      },
      substations: [
        {
          id: "SUB438c141b5a787a6701d21b6a821698df",
          longitude: -76.5264907,
          latitude: 3.4031701,
          deleted: false,
          transformers: [
            {
              id: "TRA41f20a10133f2eb4779f4296b8e4f9d6",
              substation_id: "SUB438c141b5a787a6701d21b6a821698df",
              longitude: -76.530981,
              latitude: 3.3828656,
              deleted: false
            },
            {
              id: "TRAedacf364e6eb8c7391abd09002ec0114",
              substation_id: "SUB438c141b5a787a6701d21b6a821698df",
              longitude: -76.5308582,
              latitude: 3.3837602,
              deleted: false
            }
          ]
        }
      ]
    };
  },
  computed: {
    //...mapState("substation", ["substation"])
  },
  mounted() {
    this.map = new window.google.maps.Map(this.$refs["map"], this.options);
    this.geocoder = new window.google.maps.Geocoder();
    this.setMarkers();
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
    setMarkers() {
      for (let i = 0; i < this.substations.length; i++) {
        let sub = this.substations[i];
        let subPos = { lat: sub.latitude, lng: sub.longitude };
        new window.google.maps.Marker({
          position: subPos,
          map: this.map,
          icon: {
            url: "http://maps.google.com/mapfiles/kml/paddle/purple-circle.png"
          }
        });

        for (let j = 0; j < sub.transformers.length; j++) {
          let trans = sub.transformers[j];
          let pos = { lat: trans.latitude, lng: trans.longitude };
          new window.google.maps.Marker({
            position: pos,
            map: this.map,
            icon: {
              url: "http://maps.google.com/mapfiles/kml/paddle/blu-circle.png"
            }
          });
        }
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
      setTimeout(() => {
        this.$emit("latlng", this.marker);
      }, 1000);
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