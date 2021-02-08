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
      <div class="flex items-center justify-center mt-12 pb-16">
        <a class="flex items-center text-black hover:underline hover:text-gray-800" href="https://github.com/iods?tab=repositories" target="_blank">
          <span>View More On Github</span>
          <svg class="h-5 w-5 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"/>
          </svg>
        </a>
      </div>
    </div>

  </section>

  <section class="container -mt-28">
    <div class="flex flex-col md:h-56 bg-white rounded-lg shadow-lg overflow-hidden md:flex-row my-10">
      <div class="md:flex items-center justify-center md:w-1/2 md:bg-gray-400">
        <div class="py-6 px-8 md:py-0">
          <h2 class="text-gray-700 text-2xl md:text-white iods">Interested in collaborating?</h2>
          <p class="mt-2 text-gray-600 md:text-gray-900">I am always open to discussing development opportunities and just code in general. I love to work with people who are passionate about what they do and are passionate about their business' success.</p>
        </div>
      </div>
      <div class="flex items-center justify-center pb-6 md:py-0 md:w-1/2">
        <form>
          <div class="flex flex-col rounded overflow-hidden sm:flex-row">
            <input class="appearance-none py-3 px-4 bg-gray-50 text-gray-700 border-gray-200 border focus:outline-none placeholder-gray-500 focus:border-gray-500" type="text" name="email" placeholder="Start a conversation...">
            <button class="py-3 px-4 bg-gray-800 iods text-gray-100 hover:bg-gray-400">Let's Chat!</button>
          </div>
        </form>
      </div>
    </div>
  </section>

  <section class="container pt-10">
    <h2 class="text-3xl leading-tight mb-6 font-heading iods text-center">Testimonials:</h2>
    <div class="text-center pb-14">
      <span class="text-sm text-center">People I have had the chance to work with have said some nice things...</span>
    </div>
    <div class="mx-10 pt-15">
      <div class="grid gap-6 mb-4 md:grid-cols-2">
        <div class="min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800">
          <div class="w-full mx-auto rounded-lg bg-white shadow-lg px-5 pt-5 pb-10 text-gray-800" style="max-width: 600px">
            <div class="w-full pt-1 pb-5">
              <div class="overflow-hidden rounded-full w-24 -mt-16 mx-auto shadow-lg">
                <img src="https://randomuser.me/api/portraits/women/12.jpg" alt="" class="opacity-50">
              </div>
            </div>
            <div class="w-full mb-10">
              <div class="text-5xl text-gray-500 text-left leading-tight h-3">“</div>
              <p class="text-sm text-gray-600 text-center px-5">Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nam obcaecati laudantium recusandae, debitis eum voluptatem ad, illo voluptatibus temporibus odio provident.</p>
              <div class="text-4xl text-gray-800 text-right leading-tight h-3 -mt-3">”</div>
            </div>
            <div class="w-full">
              <p class="text-md iods text-center">Rye's Mom</p>
              <p class="text-xs iods text-center">@butteredtoast</p>
            </div>
          </div>
        </div>
        <div class="min-w-0 p-4 text-white rounded-lg shadow-xs">
          <div class="w-full mx-auto rounded-lg bg-white shadow-lg px-5 pt-5 pb-10 text-gray-800" style="max-width: 600px">
            <div class="w-full pt-1 pb-5">
              <div class="overflow-hidden rounded-full w-24 -mt-16 mx-auto shadow-lg">
                <img src="https://randomuser.me/api/portraits/men/8.jpg" alt="" class="opacity-50">
              </div>
            </div>
            <div class="w-full mb-10">
              <div class="text-5xl text-gray-500 text-left leading-tight h-3">“</div>
              <p class="text-sm text-gray-600 text-center px-5">Lorem, ipsum dolor sit amet consectetur adipisicing elit. Nam obcaecati laudantium recusandae, debitis eum voluptatem ad, illo voluptatibus temporibus odio provident.</p>
              <div class="text-4xl text-gray-800 text-right leading-tight h-3 -mt-3">”</div>
            </div>
            <div class="w-full">
              <p class="text-md iods text-center">Scott Windon</p>
              <p class="text-xs iods text-center">CEO of Red Bull</p>
            </div>
          </div>
        </div>
      </div>
    </div>

  </section>
  <!--
    PROJECTS
  -->
</template>