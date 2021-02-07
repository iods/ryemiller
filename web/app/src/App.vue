<template>
  <AppHeader />
  <AppHero />
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <AppFooter />
</template>

<script>
import AppHeader from '@/components/AppHeader'
import AppHero from '@/components/AppHero'
import HelloWorld from './components/HelloWorld.vue'
import AppFooter from "@/components/AppFooter"

export default {
  name: 'App',
  components: {
    AppHeader,
    AppHero,
    HelloWorld,
    AppFooter
  }
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  color: #2c3e50;
}
</style>
