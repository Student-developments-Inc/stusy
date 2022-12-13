<template>
  <div class="block" id="timetable">
    <h1>{{ localeDate }}</h1>
    <div class="block-timetable" v-for="item in localTimetable['ИС-20-Д'][localeDate.split(',')[0]]" :key="item.id">
      <div>
        <ul class="block-timetable_topLeft-content">
          <li>{{ item[1] }}</li>
          <li>{{ item[2] }}</li>
          <li>{{ item[5] }}</li>
        </ul>
        <ul class="block-timetable_topRight-content">
          <li class="right-post">{{ item[3] }}</li>
          <li>{{ item[4] }}</li>
        </ul>
      </div>
      <p>{{ item[0] }}</p>
    </div>
  </div>
</template>

<!--<script setup>-->
<!--import {timetable} from "@/global";-->
<!--import {onMounted, ref} from "vue";-->

<!--let localeDate = ref("")-->
<!--let localTimetable = ref(timetable)-->

<!--function getDate() {-->
<!--  const options = {weekday: "long", month: "numeric", day: "numeric"};-->
<!--  localeDate.value = new Date().toLocaleDateString(undefined, options);-->
<!--  localeDate.value = localeDate.value.charAt(0).toUpperCase() + localeDate.value.slice(1);-->
<!--}-->
<!-- -->
<!--function getCountLessons() {-->
<!--  getDate();-->
<!--  const weekDay = timetable['ИС-20-Д'][localeDate.value.split(",")[0]]-->
<!--  if (weekDay) {-->
<!--    return Object.keys(weekDay).length;-->
<!--  } else {-->
<!--    return null-->
<!--  }-->
<!--}-->

<!--onMounted(() => {-->
<!--  getDate();-->
<!--});-->

<!--</script>-->

<script>
import {timetable} from "@/global";

export default {
  name: "TimetableWidget",
  data() {
    return {
      localeDate: "",
      localTimetable: timetable
    };
  },
  methods: {
    getDate() {
      const options = {weekday: "long", month: "numeric", day: "numeric"};
      this.localeDate = new Date().toLocaleDateString(undefined, options);
      this.localeDate = this.localeDate.charAt(0).toUpperCase() + this.localeDate.slice(1);
    },
    getCountLessons() {
      this.getDate();
      const weekDay = timetable['ИС-20-Д'][this.localeDate.split(",")[0]]
      if (weekDay) {
        return Object.keys(weekDay).length;
      } else {
        return null
      }
    }
  },
  mounted() {
    this.getDate();
  }
};
</script>

<style scoped>
#timetable {
  display: flex;
  flex-direction: column;
  flex-wrap: nowrap;
  overflow: auto;
}

.block-timetable {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 7px 10px;
  margin-top: 15px;
  margin-bottom: -5px;
  box-shadow: 0 0 4px rgba(0, 0, 0, 0.25);
  border-radius: 15px;
}

.block-timetable div {
  display: flex;
  flex: 1 1 auto;
  flex-wrap: wrap;
  justify-content: center;
  width: 100%;
}

.block-timetable_topLeft-content, .block-timetable_topRight-content {
  display: flex;
  align-items: center;
  margin: 0;
  float: left;
  padding: 0;
  list-style: none;
}

.block-timetable_topLeft-content {
  float: left;
}

.block-timetable_topLeft-content li {
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  padding: 4px 20px;
  font-style: normal;
  font-weight: 400;
  font-size: 14px;
  line-height: 16px;
  margin: 0 8px 0 0;
  border: 1px solid var(--dark);
  border-radius: 25px;
}

.block-timetable_topRight-content {
  float: right;
}

.right-post {
  text-align: left;
  margin: 0 -20px 0 0;
  padding: 4px 35px 4px 20px;
}

.block-timetable_topRight-content li {
  flex: 1 1 auto;
  margin: 0 0 0 4px;
  text-align: center;
  padding: 4px 20px;
  font-style: normal;
  font-weight: 400;
  font-size: 14px;
  line-height: 16px;
  box-shadow: 0 0 4px rgba(0, 0, 0, 0.25);
  border-radius: 25px;
  background: var(--light);
}

.block-timetable p {
  margin: 15px 0;
  text-align: center;
}

@media (max-width: 1600px) {
  .block-timetable div {
    width: auto;
  }
}

@media (max-width: 1039px) {
  .block-timetable ul li {
    margin: 0 0 8px;
  }

  .block-timetable div {

  }

  .block-timetable ul {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }
}

@media (max-width: 670px) {
  .block-timetable li {
    width: 100%;
  }
}
</style>