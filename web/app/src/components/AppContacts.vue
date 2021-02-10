<template>
  <!--
    CONTACTS
  -->
  <section class="py-14 px-4">

    <div class="container">
      <h2 class="text-3xl leading-tight mb-6 font-heading iods text-center">Send me a message:</h2>
      <div class="w-full max-w-2xl mx-auto mb-8">
        <form @submit.prevent="sendEmail">
          <div class="flex mb-4 -mx-2">
            <div class="w-1/2 px-2">
              <input class="appearance-none block w-full py-3 px-4 leading-tight text-gray-700 bg-gray-50 focus:bg-white border border-gray-200 focus:border-gray-500 rounded focus:outline-none" type="text" name="user_name" placeholder="Name">
            </div>
            <div class="w-1/2 px-2">
              <input class="appearance-none block w-full py-3 px-4 leading-tight text-gray-700 bg-gray-50 focus:bg-white border border-gray-200 focus:border-gray-500 rounded focus:outline-none" type="email" name="user_email" placeholder="Email">
            </div>
          </div>
          <div class="mb-4"><textarea class="appearance-none block w-full py-3 px-4 leading-tight text-gray-700 bg-gray-50 focus:bg-white border border-gray-200 focus:border-gray-500 rounded focus:outline-none" placeholder="Wise men speak because they have something to say; fools because they have to say something..." rows="5" name="message"></textarea></div>
          <div>
            <button class="inline-block w-full py-4 px-8 leading-none text-white bg-gray-800 hover:bg-gray-500 iods rounded shadow">Submit</button>
          </div>
        </form>
      </div>
      <div class="text-center">
        <p class="mb-2">Interested in collaborating? Let's chat, I will even buy the coffee.</p>
        <a class="text-gray-600 hover:underline" href="#">rye@drkstr.dev</a>
      </div>


    </div>
  </section>
  <!--
    CONTACTS
  -->
</template>


<script>

import emailjs from 'emailjs-com';

export default {
  methods: {

    sendEmail: (e) => {
      emailjs.sendForm('service_vhw1qcu', 'template_dgpsrlc', e.target, 'user_0WfdVkJzFB4xuv6Iua1yf')
        .then((result) => {
          console.log('SUCCESS!', result.status, result.text);
        }, (error) => {
          console.log('FAILED...', error);
        });
    }
  }
}

</script>