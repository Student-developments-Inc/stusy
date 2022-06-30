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
          <ol>
            <li><a href="">Введение</a></li>
            <AccordionMenu title="Модуль 1">

            </AccordionMenu>
            <AccordionMenu title="Модуль 2">
              <ol>
                <li><a href="#">1</a></li>
                <li><a href="#">1</a></li>
                <li><a href="#">1</a></li>
                <li><a href="#">1</a></li>
              </ol>
            </AccordionMenu>
            <AccordionMenu title="Модуль 3">

            </AccordionMenu>
          </ol>
        </div>
        <div class="coursesAbout">
          <article>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore
            et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip
            ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu
            fugiat nulla pariatur.
          </article>
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
  justify-content: space-between;
  gap: 16px;
}

.coursesMain {
  max-width: 70%;
  padding: 27px 32px;
  background: var(--light);
  border-radius: 25px;
}

.coursesAbout {
  max-width: 22%;
  padding: 20px;
  background: var(--light);
  border-radius: 25px;
}

.advertisement {
  padding: 17px 19px;
  border: 3px solid var(--blue);
  border-radius: 25px;
}
</style>