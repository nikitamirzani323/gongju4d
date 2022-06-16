<script>
  import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
  import Navbar from './lib/Navbar.svelte'
  import Footer from './lib/Footer.svelte'
  import Home from "./pages/home.svelte";
  import Result from "./pages/result.svelte";
  import Payment from "./pages/payment.svelte";
  import About from "./pages/about.svelte";
  import Responsibility from "./pages/responsibility.svelte";
  import NotFound from "./pages/NotFound.svelte";
  export let path_api = "";
  let routes = "";
  routes = {
			"/": wrap({
        props: {
						path_api: path_api,
				},
				component: Home,
			}),
      "/result": wrap({
        props: {
						path_api: path_api,
				},
				component: Result,
			}),
      "/payment": wrap({
				component: Payment,
			}),
      "/about": wrap({
				component: About,
			}),
      "/responsibility": wrap({
				component: Responsibility,
			}),
			"*": NotFound,
	};
</script>

<main class="p-2 lg:p-3 max-w-6xl mx-auto  mt-2 mb-2 ">
  <Navbar />
  <Router {routes} />
  
  <Footer />
</main>

<style global lang="postcss">
  @import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;500;700&display=swap");
  
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  .divider{
    display: flex;
    flex-direction: row;
    align-items: center;
    align-self: stretch;
    margin-top: 1rem/* 16px */;
    margin-bottom: 1rem/* 16px */;
    height: 1rem/* 16px */;
    white-space: nowrap;
  }
</style>
