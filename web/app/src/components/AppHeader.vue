<template>

<!-- Header & Nav Template -->
<header class="relative">

  <div class="navigation fixed top-0 left-0 w-full bg-gray-100 z-30"><!-- NAV z-30 and duration-300 -->
    <div class="container"><!-- mx-auto, p-5 "h-8 w-8 mr-2"-->

      <nav class="flex items-center justify-between flex-wrap p-6">
        <div class="flex items-center flex-no-shrink mr-6">
          <img src="../assets/logo.svg" alt="logo" class="h-8 w-8 mr-2"/>
          <span class="font-bold text-black text-2xl tracking-tight iods px-2">Rye Miller</span>
        </div>
        <div class="block lg:hidden">
          <button class="flex items-center px-3 py-2 border rounded text-gray-lighter border-black hover:text-gray hover:border-gray">
            <svg class="h-3 w-3" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><title>Menu</title><path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"/></svg>
          </button>
        </div>
        <div class="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
          <div class="text-sm lg:flex-grow">
            <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-gray-lighter hover:text-white mr-5">Home</a>
            <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-gray-lighter hover:text-white mr-5">Profile</a>
            <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-gray-lighter hover:text-white mr-5">Process</a>
            <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-gray-lighter hover:text-white mr-5">Projects</a>
            <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-gray-lighter hover:text-white mr-5">CV</a>
            <a href="#responsive-header" class="block mt-4 lg:inline-block lg:mt-0 text-gray-lighter hover:text-white">Site Health</a>
          </div>
          <div>
            <a href="#" class="inline-block text-sm leading-none hover:border-transparent hover:text-gray hover:bg-transparent mt-4 lg:mt-0">Get In Touch!</a>
          </div>
        </div>
      </nav>


    </div>
  </div>
</header>

</template>
