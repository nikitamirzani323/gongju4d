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
				component: Home,
			}),
      "/result": wrap({
        props: {
            props: {
                path_api: path_api,
            },
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

<Navbar />
<main class="p-2 lg:p-3 max-w-6xl mx-auto  mt-2 mb-2 ">
  <Router {routes} />
</main>
<Footer />


<style global lang="postcss">
  @import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;500;700&display=swap");
  
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
