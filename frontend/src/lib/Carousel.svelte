<script>
  import Carousel from 'svelte-carousel'
  let slider = [
    {img:"images/banner-1.jpg"},
    {img:"images/banner-2.jpg"},
    {img:"images/banner-0.jpg"}
  ]
</script>
<Carousel
  arrows={false}
  autoplay
  autoplayDuration={5000}
  autoplayProgressVisible
  pauseOnFocus>
  {#each slider as rec}
  <img src="{rec.img}" class="w-full" />
  {/each}
</Carousel>