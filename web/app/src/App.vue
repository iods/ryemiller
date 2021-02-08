<template>
  <AppHeader />
  <AppHero />
  <AppProfile />
  <AppProjects />
  <AppCollab />
  <AppOfferings />
  <AppContacts />
  <AppFooter />
</template>

<script>
import AppHeader from '@/components/AppHeader'
import AppHero from '@/components/AppHero'
import AppProfile from '@/components/AppProfile'
import AppProjects from '@/components/AppProjects'
import AppCollab from '@/components/AppCollab'
import AppOfferings from '@/components/AppOfferings'
import AppContacts from '@/components/AppContacts'
import AppFooter from "@/components/AppFooter"

export default {
  name: 'App',
  components: {
    AppHeader,
    AppHero,
    AppProfile,
    AppProjects,
    AppCollab,
    AppOfferings,
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