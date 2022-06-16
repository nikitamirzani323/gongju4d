<script>
    import { link } from "svelte-spa-router";
    import logo from '../assets/images/logo.png'
</script>
<div class="navbar bg-base-100">
    <div class="flex-1">
      <a class="btn btn-ghost normal-case text-xl" use:link href="/">
        <img src="{logo}" alt="결공주 Lottery Pools / 집">
      </a>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal p-0">
        <li><a use:link href="/">집</a></li>
        <li><a use:link href="/result">결과</a></li>
        <li><a use:link href="/payment">게임 방법</a></li>
        <li><a use:link href="/responsibility">책임감 있게 플레이</a></li>
        <li><a use:link href="/about">우리에 대해</a></li>
      </ul>
    </div>
  </div>