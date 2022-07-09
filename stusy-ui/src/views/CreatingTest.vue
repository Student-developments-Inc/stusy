<template>
  <div class="courses">
    <AsideMenu/>
    <div class="main">
      <header>
        <TopMenu/>
      </header>
      <div class="mainContent">
        <div class="coursesMainLeft">
          <div class="redactorContentCreating-topRow">
            <h1>Название теста:</h1>
            <div class="button-save">
              <img src="@/assets/diskette.svg">
              <input type="submit" value="Сохранить"/>
            </div>
          </div>
          <label class="input-form">
            <input type="text" alt="coursesName"/>
          </label>
          <div class="redactorContentCreating-headerOfSection">
            <h1>Вопросы</h1>
            <img src="@/assets/sorting.svg"/>
            <img src="@/assets/mix.svg"/>
            <div class="button-save">
              <img src="@/assets/added.svg">
              <input type="submit" value="Создать новый"/>
            </div>
          </div>
          <div class="groupForm-row">
            <label>Текст вопроса:</label>
            <label class="input-form input-form-allWidth">
              <input type="text" alt="questionText"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Тип ответа:</label>
            <a>Выбрать</a>
            <a>Повторить</a>
            <label>Количество баллов за вопрос:</label>
            <label class="input-form">
              <input type="text" alt="numberOfPoints" id="numberOfPoints"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Ответы</label>
            <img src="@/assets/addedV2.svg">
          </div>
          <div class="groupForm-row">
            <label>Ответ 1</label>
            <input type="checkbox" class="custom-checkbox" id="answer1" name="answer" value="yes">
            <label for="answer1"></label>
            <label class="input-form input-form-allWidth">
              <input type="text" alt="questionText"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Ответ 2</label>
            <input type="checkbox" class="custom-checkbox" id="answer2" name="answer" value="yes">
            <label for="answer2"></label>
            <label class="input-form input-form-allWidth">
              <input type="text" alt="questionText"/>
            </label>
          </div>
          <div class="button-save-right">
            <div class="button-save">
              <img src="@/assets/diskette.svg">
              <input type="submit" value="Сохранить"/>
            </div>
          </div>
          <ol class="accordionMenu-list">
            <li>
              <AccordionMenu class="accordionMenu" title-after="3. ZIP - это" v-bind:role="this.role"></AccordionMenu>
            </li>
            <li>
              <AccordionMenu class="accordionMenu" title-after="2. ZIP - это" v-bind:role="this.role"></AccordionMenu>
            </li>
            <li>
              <AccordionMenu class="accordionMenu" title-after="1. ZIP - это" v-bind:role="this.role"></AccordionMenu>
            </li>
          </ol>
        </div>
        <div class="coursesMainAbout">
          <h1>Настройки</h1>
          <div>
            <p>Макс.баллов:</p>
            <label class="input-form">
              <input type="text" alt="maxPoints"/>
            </label>
          </div>
          <div>
            <p>Мин. оценка в баллах:</p>
            <label class="input-form">
              <input type="text" alt="minPoints"/>
            </label>
          </div>
          <div>
            <p>Мин. оценка в баллах:</p>
            <label class="input-form">
              <input type="text" alt="minPoints"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Равномерное распределение баллов</label>
            <input type="checkbox" class="custom-checkbox" id="UniformDistributionOfPoints"
                   name="pointsDistribution" value="yes">
            <label for="UniformDistributionOfPoints"></label>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TopMenu from "@/components/TopMenu";
import AsideMenu from "@/components/AsideMenu";
import AccordionMenu from "@/components/AccordionMenu";

export default {
  name: "CreatingTest",
  components: {AccordionMenu, AsideMenu, TopMenu},
  data() {
    return {
      role: "teacher"
    };
  }
};
</script>

<style scoped>
p {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  margin-bottom: 8px;
}

h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

a {
  cursor: pointer;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  color: var(--blue);
}

.input-form, .input-form input {
  margin: 0;
}

#numberOfPoints {
  width: 40%;
  padding: 20px;
}

.custom-checkbox {
  position: absolute;
  display: block;
  z-index: -1;
  opacity: 0;
}

.custom-checkbox + label {
  cursor: pointer;
  display: block;
  align-items: center;
  user-select: none;
}

.custom-checkbox + label::before {
  content: '';
  display: inline-block;
  width: 1em;
  height: 1em;
  flex-shrink: 0;
  flex-grow: 0;
  border: 2px solid var(--blue);
  background-repeat: no-repeat;
  background-position: center center;
  background-size: 100% 100%;
}

.custom-checkbox:checked + label::before {
  display: inline-block;
  background-image: url(../assets/tick.svg);
}

.button-save-right {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}

.accordionMenu-list {
  gap: 20px;
}
</style>