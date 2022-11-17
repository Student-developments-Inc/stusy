<template>
  <div class='block' id='weather' v-if="weather !== null">
    <h1>{{ localeHours }}{{ name }}</h1>
    <span id='wicon'>
      {{ weatherEmoji }}
    </span>
    <p id='dateNowDay'>Сегодня {{ nowDay }}, {{ weather.weatherText }}, {{ weather.weather }}&deg;</p>
    <p>К концу занятий температура {{ getDiff }} {{ weather.lastLessonWeather }}°</p>
    <p>У вас {{ CountLesson }}</p>
    <p>Контрольных точек не запланировано</p>
  </div>
  <div class="block" id="weather" v-else>
    <ScreenLoader/>
  </div>
</template>

<script setup>
import {computed, onMounted, ref} from "vue";
import {getCookie, logout, url, weatherCodes, weatherSymbols} from "@/global";
import ScreenLoader from "@/components/ScreenLoader";
import TimetableWidget from "@/components/TimetableWidget";

const temperaturePromise = getTemperature();
onMounted(() => {
  setTemperature(temperaturePromise);
});

const weather = ref(null);
let lastLessonTime = 1705;

const weatherEmoji = computed(() => {
  return weatherSymbols[weatherCodes[weather.value.weatherCode]];
});

const nowDay = computed(() => {
  return new Date().toLocaleString("ru", {weekday: "long"});
});

const getDiff = computed(() => {
  if (weather.value.weather === weather.value.lastLessonWeather) return "останется";
  if (weather.value.weather > weather.value.lastLessonWeather) return "понизится до";
  if (weather.value.weather < weather.value.lastLessonWeather) return "повысится до";
  throw Error("weather check error");
});

const CountLesson = computed(() => {
  const valueCountLesson = TimetableWidget.methods.getCountLessons();
  let lessonDeclination = (valueCountLesson > 4) ? "пар" : "пары";
  return valueCountLesson + " " + lessonDeclination
});

async function setTemperature(temperaturePromise) {
  weather.value = await temperaturePromise;
}

async function getTemperature() {
  const storedWeather = JSON.parse(localStorage.getItem("weather"));
  if (storedWeather !== null && !isWeatherExpired(storedWeather)) {
    return storedWeather;
  }
  const fetchedWeather = await fetchTemperature();
  localStorage.setItem("weather", JSON.stringify(fetchedWeather));
  return fetchedWeather;
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
  const city = "Волгодонск";
  return fetch(
      `https://wttr.in/${city}?format=j1`
  )
      .then((response) => response.json())
      .then((data) => ({
        weather: parseInt(data.current_condition[0].temp_C),
        weatherText: data.current_condition[0].lang_ru[0].value.toLowerCase(),
        lastLessonWeather: parseInt(data.weather[0].hourly[timeIndex()].tempC),
        weatherCode: data.weather[0].hourly[timeIndex()]["weatherCode"],
        fetchDate: Math.round(Date.now() / 1000)
      }))
      .catch((err) => {
        return ["Невозможно отправить запрос", err];
      });
}

onMounted(() => {
  getUserData();
})

const userData = ref(null)

const localeHours = computed(() => {
  let localeHours = new Date().getHours();
  if (localeHours > 3 && localeHours < 12) return "Доброе утро";
  else if (localeHours > 11 && localeHours < 19) return "Добрый день";
  else if (localeHours > 18 && localeHours < 24) return "Добрый вечер";
  else if (localeHours > 23 || localeHours < 4) return "Привет";
  throw Error('localHours error')
})

const name = computed(() => {
  if (userData.value === null) {
    return '';
  }
  return `, ${userData.value.first_name}`;
})

function getUserData() {
  if (getCookie("ID") === undefined) logout();
  fetch(`${url}/users/${getCookie("ID")}`, {
    headers: {
      "Authorization": `Bearer ${getCookie("TOKEN")}`
    }
  }).then(response => {
    if (response.ok) return response.json();
    if (response.status === 401) {
      logout();
    }
  }).then(data => {
    if (data !== undefined) {
      userData.value = data;
    }
  }).catch(err => {
    console.error("Cannot fetch", err);
  });
}

</script>

<style scoped>
.lds-ripple {
  align-self: center;
}


#weather {
  min-height: 220px;
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
