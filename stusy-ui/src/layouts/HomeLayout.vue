<template>
  <p>Домашний слой</p>
  <div class="homepage">
    <AsideMenu/>
    <router-view/>
  </div>
</template>

<script>
import AsideMenu from "@/components/AsideMenu";

export default {
  name: "HomeLayout",
  components: {
    AsideMenu
  }
}
</script>

<style scoped>
.homepage {
  display: flex !important;
}
</style>