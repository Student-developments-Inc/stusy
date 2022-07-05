<template>
  <div class="creatingCourses">
    <AsideMenu/>
    <div class="main">
      <header>
        <TopMenu/>
      </header>
      <div class="mainContent">
        <div class="coursesContentCreating">
          <div class="coursesContentCreating-topRow">
            <h1>Название:</h1>
            <div class="button-save">
              <img src="@/assets/diskette.svg">
              <input type="submit" value="Сохранить"/>
            </div>
          </div>
          <label class="input-form">
            <input type="text" alt="coursesName"/>
          </label>
          <div class="sections">
            <h1>Разделы</h1>
            <div class="createNew">
              <div class="button-save" v-bind:class="{ active: menu }" @click="menuAction">
                <img src="@/assets/added.svg">
                <input type="submit" value="Создать новый"/>
              </div>
              <transition name="slide-fade">
                <ul class="sub-menu" v-if="menu">
                  <li><a>C подпунктами</a></li>
                  <li><a>Ссылка</a></li>
                  <li><a>Объявление</a></li>
                </ul>
              </transition>
            </div>
          </div>
          <ul class="createdCoursesContent">
            <li class="redactorModule">
              <img class="moving-icon" src="@/assets/moving.svg">
              <a href="">Введение</a></li>
            <li class="redactorModule">
              <AccordionMenu title="Модуль 1" v-bind:role="this.role" class="accordionMenu"></AccordionMenu>
            </li>
            <li class="redactorModule">
              <AccordionMenu title="Модуль 2" v-bind:role="this.role" class="accordionMenu">
                <ol>
                  <li><a href="#">Как сжать звуковые файлы</a></li>
                  <li><a href="#">Как сжать графические файлы</a></li>
                  <li><a href="#">Какими возможностями обладают архиваторы</a></li>
                  <li><a href="#">&#10071; Контрольная точка 2 - Тестирование</a></li>
                </ol>
              </AccordionMenu>
            </li>
            <li class="redactorModule">
              <AccordionMenu title="Модуль 3" v-bind:role="this.role" class="accordionMenu"></AccordionMenu>
            </li>
          </ul>
        </div>
        <div class="coursesAbout">
          <h1>О курсе</h1>
          <label class="input-form">
            <input type="text" alt="coursesName"/>
          </label>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AsideMenu from "@/components/AsideMenu";
import TopMenu from "@/components/TopMenu";
import AccordionMenu from "@/components/AccordionMenu";

export default {
  name: "CreatedCourses",
  components: {AccordionMenu, TopMenu, AsideMenu},
  data() {
    return {
      menu: false,
      role: "teacher"
    };
  },
  methods: {
    menuAction() {
      this.menu = !this.menu;
    }
  }
};
</script>

<style scoped>
ul {
  list-style-type: none;
}

li {
  display: flex;
  flex-direction: row;
}

h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

a {
  cursor: pointer;
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  color: var(--blue);
}

input {
  font-style: normal;
  font-weight: 500;
  font-size: 18px;
  line-height: 21px;
}

.creatingCourses {
  display: flex !important;
  background: var(--dark-mate-mate);
  min-height: 100vh;
}

.main {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  justify-content: flex-start;
  gap: 10px;
  width: 100%;
  padding: 15px 24px 24px;
  height: 100vh;
  overflow: auto;
}

.main header {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  margin-bottom: 29px;
}

.mainContent {
  height: 100%;
  display: flex;
  align-items: flex-start;
  flex-direction: row;
  justify-content: center;
  gap: 16px;
}

.coursesContentCreating {
  display: flex;
  flex: 3;
  flex-direction: column;
  justify-content: flex-start;
  gap: 21px;
  padding: 21px 26px;
  max-width: 70%;
  min-height: 80%;
  background: var(--light);
  border-radius: 25px;
}

.coursesAbout {
  flex: 1;
  max-width: 22%;
  min-height: 70%;
  padding: 20px;
  background: var(--light);
  border-radius: 25px;
}

.coursesAbout h1 {
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  line-height: 28px;
  margin-bottom: 46px;
}

.coursesContentCreating-topRow {
  display: flex;
  justify-content: space-between;
}

.input-form, .input-form input {
  margin: 0;
}

.sections {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  gap: 26px;
}

.button-save {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  align-content: center;
  padding: 17px 15px;
  gap: 13px;
  background: var(--blue);
  border-radius: 15px;
  cursor: pointer;
}

.button-save input {
  cursor: pointer;
  color: #ffffff;
  border: none;
  background: none;
}

.button-save:hover, .sub-menu li:hover {
  background: var(--blue-mate);
}

.redactorModule {
  display: flex;
  flex-direction: row;
}

.moving-icon {
  margin-right: 28px;
  cursor: pointer !important;
}

.createdCoursesContent {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.accordionMenu ol {
  display: flex;
  flex-direction: column;
  gap: 30px;
  margin-top: 30px;
  margin-left: 51px;
}

.sub-menu {
  position: absolute;
  display: flex;
  flex-direction: column;
  border-radius: 0 0 15px 15px;
  background-color: var(--blue);
}

.active {
  padding: 17px 15px 16px;
  border-radius: 15px 15px 0 0;
  border-bottom: 1px solid #1E42BB;
}

.sub-menu li a{
  padding: 15px 15.52px 21px 45px;
  border-radius: 15px;
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 18px;
  line-height: 21px;
  color: var(--light);
}

.sub-menu > li:last-child {
  border-radius: 0 0 15px 15px;
}

.createNew {
  background-color: var(--blue);
  border-radius: 15px;
}
</style>