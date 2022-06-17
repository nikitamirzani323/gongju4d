<script>
    import { link } from "svelte-spa-router";
</script>
<header class="sticky top-0 z-30 flex w-full justify-center bg-opacity-90 backdrop-blur transition-all duration-100 bg-base-100 text-base-content lg:shadow-sm">
  <main class="w-full max-w-6xl mx-auto">
    <nav class="navbar  w-full ">
      <div class="flex-1">
        <a class="btn btn-ghost  normal-case text-xl" use:link href="/">
          <img class="w-full lg:w-[135px]" src="images/logo.png" alt="결공주 Lottery Pools / 집">
        </a>
      </div>
      <div class="flex-none">
        <ul class="hidden lg:menu lg:menu-horizontal p-0">
          <li><a class="text-lg text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/">집</a></li>
          <li><a class="text-lg text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/result">결과</a></li>
          <li><a class="text-lg text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/payment">게임 방법</a></li>
          <li><a class="text-lg text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/responsibility">책임감 있게 플레이</a></li>
          <li><a class="text-lg text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/about">우리에 대해</a></li>
        </ul>
      </div>
      <div class="dropdown dropdown-end lg:hidden">
        <label tabindex="0" class="btn btn-ghost btn-circle">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7" /></svg>
        </label>
        <ul tabindex="0" class="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52">
          <li><a class="text-xs text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/">집</a></li>
          <li><a class="text-xs text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/result">결과</a></li>
          <li><a class="text-xs text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/payment">게임 방법</a></li>
          <li><a class="text-xs text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/responsibility">책임감 있게 플레이</a></li>
          <li><a class="text-xs text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/about">우리에 대해</a></li>
        </ul>
      </div>
    </nav>
  </main>
</header>