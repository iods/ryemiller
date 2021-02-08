<template>
  <!--
    HERO
  -->
  <section id="hero" class="bg-gray-100 pt-36 pb-10 text-center">

    <div class="container">
      <div class="w-full max-w-2xl mx-auto">
        <span class="text-sm font-semibold iods">Welcome, thanks for stopping by.</span>
        <h2 class="text-5xl mt-2 mb-6 leading-tight font-semibold font-heading">I am <span id="companionMethods"></span></h2>
        <p class="mb-8 text-gray-700 leading-relaxed">There will be a bunch of text here that I will add and put in here to make myself sound really cool and really important
          I just do not know what to add here yet since I am building the site.</p>
        <div>
          <a class="inline-block hover:underline pb-8 page-scroll" href="#projects">View Some Projects »</a>
        </div>
      </div>
    </div>

  </section>
  <!--
    HERO END
  -->
</template>

<script>

import TypeIt from "typeit";

export default {


  mounted() {

    new TypeIt("#companionMethods", {
      speed: 70,
      waitUntilVisible: true
    })
    .type("a software developer...", {delay: 1000})
    .delete(10, {delay: 1000, speed: 100})
    .type("bugger...", {delay: 600})
    .delete(20, {delay: 1000, speed: 100})
    .type(" caped crusader...", {delay: 600})
    .delete(16, {delay: 1000, speed: 100})
    .type("loud architect...", {delay: 600})
    .delete(26, {delay: 1000, speed: 100})
    .type("Rye Miller.")
    .go();

  }
}
</script>