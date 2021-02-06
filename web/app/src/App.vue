<template>
  <AppHeader />
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <AppFooter />
</template>

<script>
import AppHeader from '@/components/AppHeader'
import HelloWorld from './components/HelloWorld.vue'
import AppFooter from "@/components/AppFooter"

export default {
  name: 'App',
  components: {
    AppHeader,
    HelloWorld,
    AppFooter
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
