refactor:<template>
  <aside id="menuVertical">
    <h1>StudentSystem</h1>
    <ol>
      <li id="li_1"><a href="/">
        <div class="LeftMenu_icon">
          <img src="@/assets/home.svg">
        </div>
        <span style="color: var(--light)">Домашняя страница</span></a></li>
      <li><a href="/courses">
        <div class="LeftMenu_icon">
          <img src="@/assets/book.svg">
        </div>
        <span>Мои курсы</span></a></li>
      <li><a href="#">
        <div class="LeftMenu_icon">
          <img src="@/assets/calendar.svg">
        </div>
        <span>Расписание</span></a></li>
    </ol>
  </aside>
</template>

<script>
export default {
  name: "AsideMenu"
}
</script>

<style scoped>
#menuVertical {
  display: block;
  text-align: center;
  padding: 39px 24px;
  width: 300px;
  min-height: 100%;
  box-sizing: border-box;
  border-radius: 0 25px 0 0;
  background: var(--light);
}

#menuVertical h1 {
  font-style: normal;
  font-weight: 500;
  font-size: 28px;
  color: var(--blue);
  margin-top: 0;
  margin-bottom: 30px;
}

#menuVertical ol {
  display: block;
  margin: 0;
  padding: 0;
  list-style: none;
  position: relative;
}

#menuVertical ol li {
  display: list-item;
  box-sizing: border-box;
  padding: 16px 21px;
  width: 252px;
  height: 54px;
  border-radius: 15px;
}

#li_1 {
  background: var(--blue);
}

#menuVertical ol li a {
  display: flex;
  color: var(--dark);
  text-decoration: none;
  font-size: 18px;
}
.LeftMenu_icon {
  padding: 0 13px 0 0;
}
</style>