<template>
  <div class="block" id="schedule">
    <h1>{{ localeDate }}</h1>
    <div class="block-schedule" v-for="item in timetable['ИС-20-Д'][localeDate.split(',')[0]]" :key="item.id">
      <div>
        <ul class="block-schedule_topLeft-content">
          <li>{{ item[1] }}</li>
          <li>{{ item[2] }}</li>
          <li>{{ item[5] }}</li>
        </ul>
        <ul class="block-schedule_topRight-content">
          <li class="right-post">{{ item[3] }}</li>
          <li>{{ item[4] }}</li>
        </ul>
      </div>
      <p>{{ item[0] }}</p>
    </div>
  </div>
</template>

<script>
import {timetable} from "@/global";

export default {
  name: "TimetableWidget",
  data() {
    return {
      localeDate: "",
      timetable: timetable
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
      return Object.keys(timetable["ИС-20-Д"][this.localeDate.split(",")[0]]).length;
    }
  },
  mounted() {
    this.getDate();
  }
};
</script>

<style scoped>
#schedule {
  display: flex;

  flex-direction: column;
  padding: 18px 23px;
}

#schedule h1 {
  margin: 0;
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  color: var(--dark);
}

.block-schedule {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 7px 10px;
  margin-top: 15px;
  margin-bottom: -5px;
  box-shadow: 0 0 4px rgba(0, 0, 0, 0.25);
  border-radius: 15px;
}

.block-schedule div {
  display: flex;
  flex: 1 1 auto;
  flex-wrap: wrap;
  justify-content: center;
  width: 100%;
}

.block-schedule_topLeft-content, .block-schedule_topRight-content {
  display: flex;
  align-items: center;
  margin: 0;
  float: left;
  padding: 0;
  list-style: none;
}

.block-schedule_topLeft-content {
  float: left;
}

.block-schedule_topLeft-content li {
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

.block-schedule_topRight-content {
  float: right;
}

.right-post {
  text-align: left;
  margin: 0 -20px 0 0;
  padding: 4px 35px 4px 20px;
}

.block-schedule_topRight-content li {
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

.block-schedule p {
  margin: 15px 0;
  text-align: center;
}

@media (max-width: 1600px) {
  .block-schedule div {
    width: auto;
  }
}

@media (max-width: 1039px) {
  .block-schedule ul li {
    margin: 0 0 8px;
  }

  .block-schedule div {

  }

  .block-schedule ul {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
  }
}

@media (max-width: 670px) {
  .block-schedule li {
    width: 100%;
  }
}
</style>