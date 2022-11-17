<template>
  <div className="mainContent">
    <HomeWidget/>
    <TimetableWidget v-if="checkCountLesson()"/>
    <div className="block" id="courses">
      <div>
        <h1>Ваши курсы</h1>
        <a href="#">Смотреть все</a>
      </div>
      <div className="block-courses">
        <div>
          <img src="@/assets/data_archiving_and_compression.svg">
          <p>Архивация и сжатие данных</p>
        </div>
        <div>
          <img src="@/assets/botany_and_plant_physiology.svg">
          <p>Ботаника и физиология растений</p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import {getCookie} from "@/global";
import HomeWidget from "@/components/HomeWidget";
import TimetableWidget from "@/components/TimetableWidget";

export default {
  name: "HomeView",
  components: {HomeWidget, TimetableWidget},
  data() {
    return {}
  },
  methods: {
    menuAction() {
      this.menu = !this.menu;
    },
    checkCountLesson() {
      let value = TimetableWidget.methods.getCountLessons()
      return !!value;
    }
  },
  mounted() {
    if (!getCookie("TOKEN")) {
      this.$router.push("/auth");
    }
  }
};
</script>

<style>

.slide-fade-leave-active {
  transition: all .3s ease;
}

.slide-fade-leave-to {
  transform: translateY(-10px);
  opacity: 0;
}

.mainContent {
  display: flex;
  width: 100%;
  flex-wrap: wrap;
  align-items: flex-start;
  justify-content: center;
  align-content: flex-start;
}
.mainContent {
  display: flex;
  width: 100%;
  flex-wrap: wrap;
  align-items: stretch;
  justify-content: center;
  align-content: flex-start;
}

.block {
  position: relative;
  overflow: hidden;
  padding: 38px 39px;
  background: var(--light);
  border-radius: 25px;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  margin: 0 10px 10px;
}

#courses {
  padding: 18px 41px 23px;
  display: flex;
  flex-direction: column;
}

.mainContent h1 {
  margin: 0;
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  float: left;
}

#courses div a {
  float: right;
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  color: var(--blue);
  text-decoration: none;
  text-align: right;
}

@media (max-width: 1600px) {
  .block {
    width: 100%;
    margin: 0 0 0 10px;
  }
}
</style>
