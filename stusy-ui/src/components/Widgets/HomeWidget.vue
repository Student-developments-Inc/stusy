<template>
  <div id='home-widget' v-if="home-widget !== null">
    <h1>{{ localeHours }} {{userData.first_name}}</h1>
    <p>У вас {{ CountLesson }}</p>
    <p>Контрольных точек не запланировано</p>
  </div>
</template>

<script setup>
import {computed, onMounted} from "vue";
// import {getCookie, logout, url} from "@/global";
import TimetableWidget from "@/components/Widgets/TimetableWidget.vue";
import {getUserData} from "@/composables/getUserData";
const CountLesson = computed(() => {
  const valueCountLesson = TimetableWidget.methods.getCountLessons();
  if (valueCountLesson) {
    let lessonDeclination = (valueCountLesson > 4) ? "пар" : "пары";
    return valueCountLesson + " " + lessonDeclination;
  } else {
    return "нету пар";
  }
});

const {userData} = getUserData();

onMounted(() => {
  // getUserData();
})

// const userData = ref(null)

const localeHours = computed(() => {
  let localeHours = new Date().getHours();
  if (localeHours > 3 && localeHours < 12) return "Доброе утро";
  else if (localeHours > 11 && localeHours < 19) return "Добрый день";
  else if (localeHours > 18 && localeHours < 24) return "Добрый вечер";
  else if (localeHours > 23 || localeHours < 4) return "Привет";
  throw Error('localHours error')
})

// const name = computed(() => {
//   if (userData.value === null) {
//     return '';
//   }
//   return `, ${userData.value.first_name}`;
// })

// function getUserData() {
//   if (getCookie("ID") === undefined) logout();
//   fetch(`${url}/users/${getCookie("ID")}`, {
//     headers: {
//       "Authorization": `Bearer ${getCookie("TOKEN")}`
//     }
//   }).then(response => {
//     if (response.ok) return response.json();
//     if (response.status === 401) {
//       logout();
//     }
//   }).then(data => {
//     if (data !== undefined) {
//       userData.value = data;
//     }
//   }).catch(err => {
//     console.error("Cannot fetch", err);
//   });
// }

</script>

<style scoped>
.lds-ripple {
  align-self: center;
}


#home-widget {
  min-height: 220px;
  flex-direction: column;
  justify-content: flex-start;
  gap: 30px;
}

#home-widget h1 {
  padding-top: 27px;
}

#home-widget p, #home-widget h1 {
  font-size: 24px;
  margin: 0;
  z-index: 1;
}
</style>
