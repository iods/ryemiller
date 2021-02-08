<template>
  <!--
    PROJECTS
  -->
  <section class="bg-gray-100 pattern py-14">

    <div class="max-w-5xl px-6 mx-auto text-center">
      <h2 class="text-3xl leading-tight mb-6 font-heading iods">My projects and recent work:</h2>
      <div class="flex items-center justify-center mt-10">
        <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
          <div class="max-w-xs w-full">
            <div class="flex items-center justify-center h-48 bg-white border-b-8 border-gray-400 rounded-md overflow-hidden">
              <img class="object-cover h-32" src="../assets/img/web-design.svg" alt="">
            </div>
            <a href="https://github.com/iods/ryemiller.io" class="block mt-5 rounded-md overflow-hidden transition duration-500 ease-in-out transform hover:-translate-y-1 hover:scale-110" target="_blank">
              <div class="py-2 px-3 text-center text-sm">
                <p class="text-gray-800">My site is a Vue web application built and running inside of a Go binary.</p>
                <span class="block text-gray-500 mt-2">View On Github</span>
              </div>
            </a>
          </div>
          <div class="max-w-xs w-full">
            <div class="flex items-center justify-center h-48 bg-white border-b-8 border-gray-400 rounded-md overflow-hidden">
              <img class="object-cover h-32" src="../assets/img/code_3.svg" alt="">
            </div>
            <a href="https://github.com/search?q=user%3Aiods+bones" target="_blank" class="block mt-5 rounded-md overflow-hidden transition duration-500 ease-in-out transform hover:-translate-y-1 hover:scale-110">
              <div class="py-2 px-3 text-center text-sm">
                <p class="text-gray-800">Bones creates a standard for projects in PHP, JavaScript, Python and Go.</p>
                <span class="block text-gray-500 mt-2">View On Github</span>
              </div>
            </a>
          </div>
          <div class="max-w-xs w-full">
            <div class="flex items-center justify-center h-48 bg-white border-b-8 border-gray-400 rounded-md overflow-hidden">
              <img class="object-cover h-32" src="../assets/img/search.svg" alt="">
            </div>

            <a href="https://github.com/iods/go-eddie" target="_blank" class="block mt-5 rounded-md overflow-hidden transition duration-500 ease-in-out transform hover:-translate-y-1 hover:scale-110">
              <div class="py-2 px-3 text-center text-sm">
                <p class="text-gray-800">Eddie is a command line tool for tracking your routines, built with Go.</p>
                <span class="block text-gray-500 mt-2">View On Github</span>
              </div>
            </a>
          </div>
        </div>
      </div>
      <div class="flex items-center justify-center mt-12">
        <a class="flex items-center text-black hover:underline hover:text-gray-800" href="https://github.com/iods?tab=repositories" target="_blank">
          <span>View More On Github</span>
          <svg class="h-5 w-5 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"/>
          </svg>
        </a>
      </div>
    </div>

  </section>
  <!--
    PROJECTS
  -->
</template>