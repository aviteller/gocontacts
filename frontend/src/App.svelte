<script>
  import Navbar from "./layout/Navbar.svelte";
  import Register from "./compontents/Register.svelte";
  import Login from "./compontents/Login.svelte";

  let register = false;

  const toggle = () => (register = !register);
</script>

<style>

</style>

<Navbar />
<div class="container">
  <button class="btn btn-dark" on:click={toggle}>
    {register ? 'Login' : 'Register'}
  </button>
  {#if register}
    <Register />
  {:else}
    <Login />
  {/if}

</div>
