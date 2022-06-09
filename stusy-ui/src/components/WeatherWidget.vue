<template>
  <div class="block" id="weather">
    <img
        id="wicon"
        v-bind:src="[`http://openweathermap.org/img/wn/01d@2x.png`]"
        alt="weathericon"
    />
    <p id="dateNowDay">Сегодня {{ getNowDay() }}, сейчас {{ weather }}&deg;</p>
    <p>У вас 4 пары</p>
    <p>К концу занятий температура поднимется до {{ lastLessonWeather }}°</p>
    <p>Контрольных точек не запланировано</p>
  </div>
</template>

<script>
export default {
  name: "WeatherWidget",
  data() {
    return {
      city: "Волгодонск",
      weather: '',
      lastLessonTime: 1705,
      lastLessonWeather: ''
    };
  },
  methods: {
    getNowDay() {
      return new Date().toLocaleString("ru", {
        weekday: "long"
      });
    },
    getTemp() {
      fetch(
          `https://wttr.in/${this.city}?format=j1`
      )
          .then((response) => response.json())
          .then((data) => {
            this.weather = data.current_condition[0].FeelsLikeC;
            this.lastLessonWeather = data.weather[0].hourly[this.timeIndex()].FeelsLikeC;
          })
          .catch((err) => {
            console.error("Невозможно отправить запрос", err);
          });
    },
    timeIndex() {
      let closestTime;
      let j = Math.floor(this.lastLessonTime / 100);
      if (j % 3 > 1) {
        closestTime = j + 1
      } else {
        closestTime = j - j % 3
      }
      if (closestTime === 24) {
        closestTime = 0;
      }
      return closestTime / 3
    }
  },
  mounted() {
    this.getTemp();
  }
};
</script>

<style scoped>
#weather {
  height: 220px;
}

#weather p {
  font-size: 24px;
  margin: 0;
  z-index: 1;
}

#weather #wicon {
  position: absolute;
  right: -10px;
  top: -10px;
}
</style>
