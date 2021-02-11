<template>
  <!--
    CONTACTS
  -->
  <section class="py-14 px-4">

    <div class="container">
      <h2 class="text-3xl leading-tight mb-6 font-heading iods text-center">Send me a message:</h2>
      <div class="w-full max-w-2xl mx-auto mb-8">
        <Form @submit.prevent="sendEmail" id="contacts">
          <div class="flex mb-4 -mx-2">
            <div class="w-1/2 px-2">
              <Field class="appearance-none block w-full py-3 px-4 leading-tight text-gray-700 bg-gray-50 focus:bg-white border border-gray-200 focus:border-gray-500 rounded focus:outline-none" type="text" name="user_name" placeholder="Name" />
            </div>
            <div class="w-1/2 px-2">
              <Field class="appearance-none block w-full py-3 px-4 leading-tight text-gray-700 bg-gray-50 focus:bg-white border border-gray-200 focus:border-gray-500 rounded focus:outline-none" type="email" name="user_email" :rules="validateContacts" placeholder="Email" />
            </div>
          </div>
          <div class="mb-4"><Field as="textarea" class="appearance-none block w-full py-3 px-4 leading-tight text-gray-700 bg-gray-50 focus:bg-white border border-gray-200 focus:border-gray-500 rounded focus:outline-none" placeholder="Wise men speak because they have something to say; fools because they have to say something..." rows="5" name="message"></Field></div>
          <div>
            <button class="inline-block w-full py-4 px-8 leading-none text-white bg-gray-800 hover:bg-gray-500 iods rounded shadow">Submit</button>
          </div>
          <ErrorMessage as="div" name="user_email" v-slot="{ message }">
            <p class="error-message">{{ message }}</p>
          </ErrorMessage>
        </Form>
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

import { Form, Field, ErrorMessage } from 'vee-validate';

export default {

  components: {
    ErrorMessage,
    Form,
    Field,
  },

  methods: {

    validateContacts(value) {
      if (!value) {
        return 'Your Email is required.';
      }

      if (!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(value)) {
        return 'Your email must be valid, silly.';
      }

      return true;
    },


    post(path, params, method='post') {

      const form = document.createElement('form');
      form.method = method;
      form.action = path;

      for (const key in params) {

        const hiddenField = document.createElement('input');
        hiddenField.type = 'hidden';
        hiddenField.name = key;
        hiddenField.value = params[key];

        form.appendChild(hiddenField);

      }

      document.body.appendChild(form);
      form.submit();
    },

    sendEmail() {
      const form = document.forms.contacts;
      const formData = new FormData(form);
      const name = formData.get('user_name');
      const email = formData.get('user_email');
      const message = formData.get('message');

      this.post('/sendMail', {email: email, name: name, message: message})
    }
  }
}

</script>