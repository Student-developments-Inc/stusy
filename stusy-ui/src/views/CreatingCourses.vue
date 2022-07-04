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
            <div class="button-save">
              <img src="@/assets/added.svg">
              <input type="submit" value="Создать новый"/>
            </div>
          </div>
          <ul class="createdCoursesContent">
            <li class="redactorModule">
              <img class="moving-icon" src="@/assets/moving.svg">
              <a href="">Введение</a></li>
            <li class="redactorModule">
              <img class="moving-icon" src="@/assets/moving.svg">
              <AccordionMenu title="Модуль 1" class="accordionMenu"></AccordionMenu>
              <img class="redactor-icon" src="@/assets/addedV2.svg">
              <img class="redactor-icon" src="@/assets/pen.svg">
              <img class="trash-icon" src="@/assets/trash.svg">
            </li>
            <li class="redactorModule">
              <img class="moving-icon" src="@/assets/moving.svg">
              <AccordionMenu title="Модуль 2" class="accordionMenu">
                <ol>
                  <li><a href="#">Как сжать звуковые файлы</a></li>
                  <li><a href="#">Как сжать графические файлы</a></li>
                  <li><a href="#">Какими возможностями обладают архиваторы</a></li>
                  <li><a href="#">&#10071; Контрольная точка 2 - Тестирование</a></li>
                </ol>
              </AccordionMenu>
              <img class="redactor-icon" src="@/assets/addedV2.svg">
              <img class="redactor-icon" src="@/assets/pen.svg">
              <img class="trash-icon" src="@/assets/trash.svg">
            </li>
            <li class="redactorModule">
              <img class="moving-icon" src="@/assets/moving.svg">
              <AccordionMenu title="Модуль 3" class="accordionMenu"></AccordionMenu>
              <img class="redactor-icon" src="@/assets/addedV2.svg">
              <img class="redactor-icon" src="@/assets/pen.svg">
              <img class="trash-icon" src="@/assets/trash.svg">
            </li>
          </ul>
        </div>
        <div class="coursesAbout">
          <h1>О курсе</h1>
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
  components: {AccordionMenu, TopMenu, AsideMenu}
};
</script>

<style scoped>
h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

a {
  text-decoration: none;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  color: var(--blue);
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

.button-save:hover {
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

.redactor-icon {
  margin-right: 12px;
}

.trash-icon {
  margin-left: 5px;
}

.createdCoursesContent {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.accordionMenu {
  margin-right: 9px;
}

.accordionMenu ol {
  display: flex;
  flex-direction: column;
  gap: 30px;
  margin-top: 30px;
}

ul {
  list-style-type: none;
}

li {
  display: flex;
  flex-direction: row;
}
</style>