<template>
  <!--
    FOOTER
  -->
  <div class="container text-center">

    <footer class="p-4">
      <div class="flex items-center">
        <div class="w-full text-center iods text-gray-500">&copy; {{ year }} Rye Miller</div>
      </div>
    </footer>

  </div>

  <a href="#" class="back-to-top w-10 h-10 fixed bottom-0 right-0 mb-5 mr-5 flex items-center justify-center rounded-full bg-black text-white text-lg z-20 duration-300 hover:bg-gray-400"></a>
  <!--
    FOOTER END
  -->
</template>

<script>

import { computed } from 'vue'

export default {
  setup () {
    const year = computed(() => new Date().getFullYear())
    return { year }
  }
}
</script>