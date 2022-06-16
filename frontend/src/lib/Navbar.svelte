<script>
    import { link } from "svelte-spa-router";
</script>
<header class="sticky top-0 z-30 flex w-full justify-center bg-opacity-90 backdrop-blur transition-all duration-100 bg-base-100 text-base-content lg:shadow-sm">
  <nav class="navbar bg-base-100">
    <div class="flex-1">
      <a class="btn btn-ghost normal-case text-xl" use:link href="/">
        <img src="images/logo.png" alt="결공주 Lottery Pools / 집">
      </a>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal p-0">
        <li><a class="text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/">집</a></li>
        <li><a class="text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/result">결과</a></li>
        <li><a class="text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/payment">게임 방법</a></li>
        <li><a class="text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/responsibility">책임감 있게 플레이</a></li>
        <li><a class="text-black font-bold hover:bg-transparent focus:bg-transparent"use:link href="/about">우리에 대해</a></li>
      </ul>
    </div>
  </nav>
</header>