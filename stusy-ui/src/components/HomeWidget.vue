<template>
  <div class='block' id='weather' v-if="weather.fetchDate !== 0">
    <span id='wicon'>
      {{ weatherEmoji }}
    </span>
    <p id='dateNowDay'>Сегодня {{ nowDay }}, сейчас
      {{ weather.weather }}&deg;
    </p>
    <p>У вас 4 пары</p>
    <p>Температура к концу занятий {{ getDiff }} {{ weather.lastLessonWeather }}°</p>
    <p>Контрольных точек не запланировано</p>
  </div>
  <div class="block" id="weather" v-else>
    <ScreenLoader/>
  </div>
</template>

<script setup>
import {ref, computed, onMounted} from 'vue'
import {weatherCodes, weatherSymbols} from '@/global';
import ScreenLoader from "@/components/ScreenLoader";

onMounted(() => {
  getTemperature()
})

const weather = ref({
  weather: 0,
  lastLessonWeather: 0,
  weatherCode: 0,
  fetchDate: 0,
})
const city = 'Волгодонск'
let lastLessonTime = 1705

const weatherEmoji = computed(() => {
  return weatherSymbols[weatherCodes[weather.value.weatherCode]]
})

const nowDay = computed(() => {
  return new Date().toLocaleString('ru', {weekday: 'long'});
})

const getDiff = computed(() => {
  if (weather.value.weather === weather.value.lastLessonWeather) return 'останется';
  if (weather.value.weather > weather.value.lastLessonWeather) return 'понизится до';
  if (weather.value.weather < weather.value.lastLessonWeather) return 'повысится до';
  throw Error('weather check error')
})

async function getTemperature() {
  const storedWeather = JSON.parse(localStorage.getItem('weather'));
  if (storedWeather !== null && !isWeatherExpired(storedWeather)) {
    console.log('Есть погода')
    weather.value = storedWeather;
  } else {
    console.log('Нет погоды')
    weather.value = await fetchTemperature();
    localStorage.setItem('weather', JSON.stringify(weather.value));
  }
}

function isWeatherExpired(weather) {
  const hoursExpired = 2;
  const secondsExpired = hoursExpired * 60 * 60;
  const nowSeconds = Date.now() / 1000;
  const secondsPassed = nowSeconds - weather.fetchDate;
  return secondsPassed > secondsExpired;
}

function timeIndex() {
  let closestTime;
  let j = Math.floor(lastLessonTime / 100);
  if (j % 3 > 1) {
    closestTime = j + 1;
  } else {
    closestTime = j - j % 3;
  }
  if (closestTime === 24) {
    closestTime = 0;
  }
  return closestTime / 3;
}

function fetchTemperature() {
  return fetch(
      `https://wttr.in/${city}?format=j1`
  )
      .then((response) => response.json())
      .then((data) => ({
        weather: parseInt(data.current_condition[0].FeelsLikeC),
        lastLessonWeather: parseInt(data.weather[0].hourly[timeIndex()].FeelsLikeC),
        weatherCode: data.weather[0].hourly[timeIndex()]['weatherCode'],
        fetchDate: Math.round(Date.now() / 1000),
      }))
      .catch((err) => {
        return ['Невозможно отправить запрос', err];
      });
}

</script>

<style scoped>
.lds-ripple {
  align-self: center;
}


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
