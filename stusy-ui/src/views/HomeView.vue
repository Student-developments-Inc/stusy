<template>
  <div class="mainContent">
    <HomeWidget class="block TwoFor"/>
    <div class="block TwoTwo" id="courses">
      <div class="blockHeader">
        <h1>Ваши курсы</h1>
        <a href="#">Смотреть все</a>
      </div>
      <div class="block-courses">
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
    <TimetableWidget class="TwoFor" v-if="checkCountLesson()"/>

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

.block {
  position: relative;
  overflow: hidden;
  background: var(--light);
  border-radius: 25px;

  display: flex;
  flex-wrap: wrap;
  align-content: space-between;
  justify-content: space-between;
  padding: 15px;
  margin: 5px;

}

.blockHeader {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.block > .block {
  padding: 0;
}

.OneOne {
  height: 145px;
  width: 145px;
  background: #000;
}

.TwoTwo {
  height: 310px;
  width: 310px;
}

.TwoFor {
  height: 310px;
  max-height: 310px;
  width: 640px;
  max-width: 640px;
}

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

@media (max-width: 1393px) {
  .TwoFor {
    height: 310px;
    width: 310px;
  }
}
</style>
