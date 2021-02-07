<template>
  <AppHeader />
  <AppHero />
  <AppContacts />
  <AppFooter />
</template>

<script>
import AppHeader from '@/components/AppHeader'
import AppHero from '@/components/AppHero'
import AppContacts from '@/components/AppContacts'
import AppFooter from "@/components/AppFooter"

export default {
  name: 'App',
  components: {
    AppHeader,
    AppHero,
    AppContacts,
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