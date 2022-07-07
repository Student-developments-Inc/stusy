<template>
  <div class="block" id="weather">
    <!--    <img-->
    <!--        id="wicon"-->
    <!--        v-bind:src="require(`../assets/weather/${weatherCode}.svg`)"-->
    <!--        alt="weathericon"-->
    <!--      />-->
    <span id="wicon">
      {{ weatherSymbols[weatherCodes[weatherCode]] }}
    </span>
    <p id="dateNowDay">Сегодня {{ getNowDay() }}, сейчас {{ weather }}&deg;</p>
    <p>У вас 4 пары</p>
    <p>Температура к концу занятий {{ getDiff() + lastLessonWeather }}°</p>
    <p>Контрольных точек не запланировано</p>
  </div>
</template>

<script>
export default {
  name: "WeatherWidget",
  data() {
    return {
      city: "Волгодонск",
      weather: "",
      lastLessonTime: 1705,
      lastLessonWeather: "",
      weatherCode: 0,
      weatherCodes: {
        "0": "Unknown",
        "113": "Sunny",
        "116": "PartlyCloudy",
        "119": "Cloudy",
        "122": "VeryCloudy",
        "143": "Fog",
        "176": "LightShowers",
        "179": "LightSleetShowers",
        "182": "LightSleet",
        "185": "LightSleet",
        "200": "ThunderyShowers",
        "227": "LightSnow",
        "230": "HeavySnow",
        "248": "Fog",
        "260": "Fog",
        "263": "LightShowers",
        "266": "LightRain",
        "281": "LightSleet",
        "284": "LightSleet",
        "293": "LightRain",
        "296": "LightRain",
        "299": "HeavyShowers",
        "302": "HeavyRain",
        "305": "HeavyShowers",
        "308": "HeavyRain",
        "311": "LightSleet",
        "314": "LightSleet",
        "317": "LightSleet",
        "320": "LightSnow",
        "323": "LightSnowShowers",
        "326": "LightSnowShowers",
        "329": "HeavySnow",
        "332": "HeavySnow",
        "335": "HeavySnowShowers",
        "338": "HeavySnow",
        "350": "LightSleet",
        "353": "LightShowers",
        "356": "HeavyShowers",
        "359": "HeavyRain",
        "362": "LightSleetShowers",
        "365": "LightSleetShowers",
        "368": "LightSnowShowers",
        "371": "HeavySnowShowers",
        "374": "LightSleetShowers",
        "377": "LightSleet",
        "386": "ThunderyShowers",
        "389": "ThunderyHeavyRain",
        "392": "ThunderySnowShowers",
        "395": "HeavySnowShowers",
      },
      weatherSymbols: {
        "Unknown": "✨",
        "Cloudy": "☁️",
        "Fog": "🌫",
        "HeavyRain": "🌧",
        "HeavyShowers": "🌧",
        "HeavySnow": "❄️",
        "HeavySnowShowers": "❄️",
        "LightRain": "🌦",
        "LightShowers": "🌦",
        "LightSleet": "🌧",
        "LightSleetShowers": "🌧",
        "LightSnow": "🌨",
        "LightSnowShowers": "🌨",
        "PartlyCloudy": "⛅️",
        "Sunny": "☀️",
        "ThunderyHeavyRain": "🌩",
        "ThunderyShowers": "⛈",
        "ThunderySnowShowers": "⛈",
        "VeryCloudy": "☁️",
      }
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
            this.weatherCode = data.weather[0].hourly[this.timeIndex()]["weatherCode"];
          })
          .catch((err) => {
            console.error("Невозможно отправить запрос", err);
          });
    },
    timeIndex() {
      let closestTime;
      let j = Math.floor(this.lastLessonTime / 100);
      if (j % 3 > 1) {
        closestTime = j + 1;
      } else {
        closestTime = j - j % 3;
      }
      if (closestTime === 24) {
        closestTime = 0;
      }
      return closestTime / 3;
    },
    getDiff() {
      if (this.weather == this.lastLessonWeather) return "останется ";
      if (this.weather > this.lastLessonWeather) return "понизится до ";
      if (this.weather < this.lastLessonWeather) return "повысится до ";
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
  top: -15px;
  font-size: 85pt;
}
</style>
