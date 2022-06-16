<script>
  import banner_0 from './assets/images/banner-0.jpg'
  import banner_1 from './assets/images/banner-1.jpg'
  import banner_2 from './assets/images/banner-2.jpg'
  import Navbar from './lib/Navbar.svelte'
</script>

<main class="p-2 lg:p-3 max-w-6xl mx-auto  mt-2 mb-2 font-['Verdana']">
  <Navbar />
  <section class="flex gap-2">
    <article class="carousel w-full">
      <div id="slide1" class="carousel-item relative w-full">
        <img src="{banner_0}" class="w-full" />
        <div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
          <a href="#slide4" class="btn btn-circle">❮</a> 
          <a href="#slide2" class="btn btn-circle">❯</a>
        </div>
      </div> 
      <div id="slide2" class="carousel-item relative w-full">
        <img src="{banner_1}" class="w-full" />
        <div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
          <a href="#slide1" class="btn btn-circle">❮</a> 
          <a href="#slide3" class="btn btn-circle">❯</a>
        </div>
      </div> 
      <div id="slide3" class="carousel-item relative w-full">
        <img src="{banner_2}" class="w-full" />
        <div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
          <a href="#slide2" class="btn btn-circle">❮</a> 
          <a href="#slide4" class="btn btn-circle">❯</a>
        </div>
      </div> 
    </article>
  </section>

  <div class="titleweb">
    <hr>
    <h4>다음 그림</h4>
  </div>

</main>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  @import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;500;700&display=swap");
</style>
