<template>
  <div class="homepage">
    <AsideMenu/>
    <div class="home">
      <div class="content">
        <header>
          <TopMenu></TopMenu>
        </header>
        <router-view/>
      </div>
    </div>
  </div>
</template>

<script setup>
import AsideMenu from "@/components/AsideMenu";
import TopMenu from "@/components/TopMenu";

</script>

<style scoped>
.homepage {
  display: flex !important;
}

.home {
  flex-grow: 1;
  background: var(--dark-mate-mate);
}

.content header {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.content header h1 {
  font-weight: 500;
  font-size: 36px;
  margin: 0;
}

.content {
  padding: 15px 24px 24px;
  height: 100vh;
  overflow: auto;
}

header {
  margin: 12px 0;
}
</style>