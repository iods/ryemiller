<template>
  <!--
    HEADER
  -->
  <header id="header-wrap" class="relative">
    <!--
      NAVBAR
    -->
    <div class="navigation fixed top-0 left-0 w-full duration-300 z-30">

      <div class="container">
        <nav class="navbar navbar-expand-lg py-1 flex justify-between items-center relative duration-300">
          <a class="navbar-brand" href="/"><img src="../assets/img/logo.svg" alt="Logo" height="40" width="40"></a>
          <button class="navbar-toggler focus:outline-none block lg:hidden" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="toggler-icon"></span>
            <span class="toggler-icon"></span>
            <span class="toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse hidden lg:block duration-300 shadow absolute top-10 left-0 mt-full bg-white z-20 px-5 py-3 w-full lg:static lg:bg-transparent lg:shadow-none" id="navbarSupportedContent">
            <ul class="navbar-nav mr-auto justify-center items-center lg:flex">
              <li class="nav-item">
                <a class="page-scroll active" href="#hero">Home</a>
              </li>
              <li class="nav-item">
                <a class="page-scroll" href="#profile">Profile</a>
              </li>
              <li class="nav-item">
                <a class="page-scroll" href="#projects">Projects</a>
              </li>
              <li class="nav-item">
                <a class="page-scroll" href="#offerings">Offerings</a>
              </li>
              <li class="nav-item">
                <a class="page-scroll" href="/curriculum-vitae">CV</a>
              </li>
              <li class="nav-item">
                <a class="page-scroll" href="/health">Site Health</a>
              </li>
            </ul>
          </div>
          <div class="header-btn hidden sm:block sm:absolute sm:right-0 sm:mr-16 lg:static lg:mr-0">
            <a class="px-3 py-1" href="https://github.com/iods?tab=repositories" target="_blank">
              <img src="../assets/img/github-square.svg" height="40" width="40">
            </a>
          </div>
        </nav>
      </div>

    </div>
    <!--
      NAVBAR END
    -->
  </header>
  <!--
    HEADER END
  -->
</template>


<script>

export default {
  name: 'header',

  mounted() {
    var header_navbar = document.querySelector(".navigation");
    var sticky = header_navbar.offsetTop;
    var backToTop = document.querySelector(".back-to-top");
    var pageLink = document.querySelectorAll('.page-scroll');
    var navbarCollapse = document.querySelector(".navbar-collapse");
    let navbarToggler = document.querySelector(".navbar-toggler");

    window.onscroll = function () {

      if (window.pageYOffset > sticky) {
        header_navbar.classList.add("sticky");
      } else {
        header_navbar.classList.remove("sticky");
      }

      if (document.body.scrollTop > 50 || document.documentElement.scrollTop > 50) {
        backToTop.style.display = "flex";
      } else {
        backToTop.style.display = "none";
      }
    }

    pageLink.forEach(elem => {
      elem.addEventListener('click', e => {
        e.preventDefault();
        document.querySelector(elem.getAttribute('href')).scrollIntoView({
          behavior: 'smooth',
          offsetTop: 1 - 60,
        });
      });
    });

    // section menu active
    function onScroll(event) {// eslint-disable-line no-unused-vars
      var sections = document.querySelectorAll('.page-scroll');
      var scrollPos = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop;

      for (var i = 0; i < sections.length; i++) {
        var currLink = sections[i];
        var val = currLink.getAttribute('href');
        var refElement = document.querySelector(val);
        var scrollTopMinus = scrollPos + 73;
        if (refElement.offsetTop <= scrollTopMinus && (refElement.offsetTop + refElement.offsetHeight > scrollTopMinus)) {
          document.querySelector('.page-scroll').classList.remove('active');
          currLink.classList.add('active');
        } else {
          currLink.classList.remove('active');
        }
      }
    }

    window.document.addEventListener('scroll', onScroll);

    document.querySelectorAll(".page-scroll").forEach(e =>
      e.addEventListener("click", () =>{
        navbarCollapse.classList.remove("show")
        navbarToggler.classList.remove("active")
      })
    )

    navbarToggler.addEventListener('click', function () {
      navbarCollapse.classList.toggle("show");
      navbarToggler.classList.toggle("active")
    })
  }
}
</script>