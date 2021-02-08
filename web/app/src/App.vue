<template>
  <AppHeader />
  <AppHero />
  <AppProfile />
  <AppCollab />
  <AppFooter />
</template>

<script>
import AppHeader from '@/components/AppHeader'
import AppHero from '@/components/AppHero'
import AppProfile from '@/components/AppProfile'
import AppCollab from '@/components/AppCollab'
import AppFooter from "@/components/AppFooter"

export default {
  name: 'App',
  components: {
    AppHeader,
    AppHero,
    AppProfile,
    AppCollab,
    AppFooter
  }
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>