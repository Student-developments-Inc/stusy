<template>
  <div class="courses">
    <AsideMenu/>
    <div class="main">
      <header>
        <TopMenu/>
      </header>
      <div class="mainContent">
        <div class="coursesMain">
          <h1>Архивация и сжатие данных</h1>
          <section class="advertisement">
            <h1>Объявление</h1>
            <article>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore
              et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip
              ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
              fugiat nulla pariatur.
            </article>
          </section>
          <ol class="courseContent">
            <li><a href="">Введение</a></li>
            <AccordionMenu title="Модуль 1" class="accordionMenu">

            </AccordionMenu>
            <AccordionMenu title="Модуль 2" class="accordionMenu">
              <ol>
                <li><a href="#">Как сжать звуковые файлы</a></li>
                <li><a href="#">Как сжать графические файлы</a></li>
                <li><a href="#">Какими возможностями обладают архиваторы</a></li>
                <li><a href="#">&#10071; Контрольная точка 2 - Тестирование</a></li>
              </ol>
            </AccordionMenu>
            <AccordionMenu title="Модуль 3" class="accordionMenu">

            </AccordionMenu>
          </ol>
        </div>
        <div class="coursesAbout">
          <h1>О курсе</h1>
          <div class="coursesAboutInfo">
            <p>3 зачетных единицы</p>
            <p>5 контрольных точек</p>
          </div>
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
  name: "MyCourses",
  components: {AccordionMenu, TopMenu, AsideMenu},
  data() {
    return {
      menu: false
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
h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

p {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

a {
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  color: var(--blue);
}

ol {
  list-style-type: none;
}

.courses {
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

.coursesMain {
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

.coursesAbout {
  flex: 1;
  max-width: 22%;
  min-height: 70%;
  padding: 20px;
  background: var(--light);
  border-radius: 25px;
}

.coursesAbout p {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
}

.coursesAbout h1 {
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  line-height: 28px;
}

.coursesAboutInfo {
  margin-top: 16px;
  display: flex;
  flex-direction: column;
  gap: 11px;
}

.advertisement {
  padding: 17px 19px;
  border: 3px solid var(--blue);
  border-radius: 25px;
  margin-bottom: 30px;
}

.advertisement h1 {
  font-style: normal;
  font-weight: 500;
  font-size: 24px;
  line-height: 28px;
}

.advertisement article {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 160%;
  text-align: justify;
}

.courseContent {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.accordionMenu ol {
  display: flex;
  flex-direction: column;
  gap: 30px;
  margin-top: 30px;
}
</style>