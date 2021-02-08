<template>
  <!--
    HERO
  -->
  <section id="hero" class="bg-gray-100 pt-28 pb-10 text-center">

    <div class="container">
      <div class="w-full max-w-3xl mx-auto">
        <span class="text-sm iods">Welcome, thanks for stopping by.</span>
        <h2 class="text-5xl mt-4 mb-6 pt-10 leading-tight font-semibold font-heading">I am <span id="companionMethods"></span></h2>
        <p class="mb-8 text-gray-700 leading-relaxed">I have a passion for bringing clarity and simplicity to software and the teams behind it. It begins with a healthy obsession for the
          technologies and processes from a project's discovery to it's delivery, focusing on being agile not fragile. My goal is to move the web forward with each character
          considered, each pixel defined, every pipeline a success.</p>
        <div>
          <a class="inline-block hover:underline pb-8 page-scroll" href="#projects">My vision is in the code »</a>
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