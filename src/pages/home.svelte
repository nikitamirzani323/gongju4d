<script>
    import Carousel from '../lib/Carousel.svelte'
    import bg_day from '../assets/images/day.png'
    import bg_night from '../assets/images/night.png'
  </script>
<section class="flex gap-2 mb-10">
    <Carousel />
</section>
<section class="divider mt-10 mb-10">
    <h4 class="text-[pink] text-2xl">다음 그림</h4>
</section>
<section class="border-2 border-solid border-[#74aa63] mt-10 mb-10 rounded-lg  p-5">
    <div class="flex flex-col">
        <p class="text-[#74aa63] text-2xl">아침(왼쪽)과 밤(오른쪽) 출애굽</p>
        <div class="grid grid-cols-2 gap-2 w-full mt-10 mb-10">
        <div class="w-full relative">
            <img src="{bg_day}" alt="">
            <div class="flex justify-center gap-16 absolute top-10 left-16">
            <span class="text-6xl text-white">09</span>
            <span class="text-6xl text-white mx-5">20</span>
            <span class="text-6xl text-white -mx-5">30</span>
            </div>
        </div>
        <div class="w-full relative">
            <img src="{bg_night}" alt="">
            <div class="flex justify-center gap-16 absolute top-10 left-20">
            <span class="text-6xl text-white">09</span>
            <span class="text-6xl text-white mx-5">20</span>
            <span class="text-6xl text-white -mx-5">30</span>
            </div>
        </div>
        </div>
    </div>
</section>
<section class="divider mt-10">
    <h4 class="text-[pink] text-2xl">오늘의 출력</h4>
</section>