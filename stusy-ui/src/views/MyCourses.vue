<template>
  <div class="courses">
    <AsideMenu/>
    <div class="main">
      <header>
        <TopMenu/>
      </header>
      <div class="mainContent">
        <div class="coursesMainLeft" id="coursesMain">
          <h1>Мои курсы</h1>
          <div class="block-courses">
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/data_archiving_and_compression.svg">
              <p>Архивация и сжатие данных</p>
            </div>
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/botany_and_plant_physiology.svg">
              <p>Ботаника и физиология растений</p>
            </div>
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/data_archiving_and_compression.svg">
              <p>Архивация и сжатие данных</p>
            </div>
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/botany_and_plant_physiology.svg">
              <p>Ботаника и физиология растений</p>
            </div>
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/data_archiving_and_compression.svg">
              <p>Архивация и сжатие данных</p>
            </div>
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/botany_and_plant_physiology.svg">
              <p>Ботаника и физиология растений</p>
            </div>
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/data_archiving_and_compression.svg">
              <p>Архивация и сжатие данных</p>
            </div>
            <div v-on:click="goToCourse(1)">
              <img src="@/assets/botany_and_plant_physiology.svg">
              <p>Ботаника и физиология растений</p>
            </div>
          </div>
        </div>
        <div class="coursesMainAbout">
          <h1>Какая-то инфа</h1>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AsideMenu from "@/components/AsideMenu";
import TopMenu from "@/components/TopMenu";

export default {
  name: "MyCourses",
  components: {TopMenu, AsideMenu},
  data() {
    return {
      menu: false
    };
  },
  methods: {
    menuAction() {
      this.menu = !this.menu;
    },
    goToCourse() {
      this.$router.push("/courses");
    }
  }
};
</script>

<style scoped>
h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

#coursesMain {
  display: flex;
  flex: 3;
  flex-direction: column;
  gap: 28px;
  padding: 27px 32px;
  max-width: 70%;
  min-height: 80%;
  background: var(--light);
  border-radius: 25px;
}

.block-courses > div:hover:hover {
  transform: scale(1.01);
}

.block-courses > div {
  cursor: pointer;
  transition: .3s;
}
</style>