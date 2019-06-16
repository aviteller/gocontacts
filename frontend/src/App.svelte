<script>
  import Navbar from "./layout/Navbar.svelte";
  import Register from "./compontents/Register.svelte";
  import Login from "./compontents/Login.svelte";
  import Contact from "./compontents/Contact.svelte";
  import CreateContact from "./compontents/CreateContact.svelte";
  import Cookies from "./Cookies.js";

  import { onMount } from "svelte";

  let c = new Cookies();

  onMount(() => {
    let jwt = "";
    if ((jwt = c.getCookie("jwt"))) {
      jwt = JSON.parse(jwt);
      user = {
        id: jwt.id,
        email: jwt.email,
        token: jwt.token
      };
      user.loggedIn = true;
      getContacts(user);
    }
  });

  let register = false;

  let user = {};

  let contacts = [];

  const toggle = () => (register = !register);
  const logout = () => {
    c.eraseCookie("jwt");
    user = {};
  };

  const getContacts = async user => {
    let res = "";

    res = await fetch(`http://localhost:8000/api/user/${user.id}/contacts`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${user.token}`
      }
    });

    contacts = await res.json();
  };

  const userLoggedIn = e => {
    user = e.detail;

    c.setCookie("jwt", JSON.stringify(user), 1);

    getContacts(user);
  };

  const addContact = e => {
    const newContact = e.detail;

    contacts.data = [...contacts.data, newContact];
  }

   const removeContact = e => {
    contacts.data = contacts.data.filter(contact => contact.ID !== e.detail);
  };
</script>

<style>

</style>

<Navbar />
<div class="container">
  {#if user.loggedIn}
    <button class="btn btn-danger" on:click={logout}>Logout</button>
    <h1>{user.email}</h1>
    <CreateContact on:addcontact={addContact}/>
    {#if contacts.data !== undefined}
      {#each contacts.data as contact}
        <Contact name={contact.name} phone={contact.phone} id={contact.ID} on:removecontact={removeContact}/>
      {/each}
    {:else}
      <h3>please add contacts</h3>
    {/if}
  {:else}
    <button class="btn btn-dark" on:click={toggle}>
       {register ? 'Login' : 'Register'}
    </button>
    {#if register}
      <Register />
    {:else}
      <Login on:userloggedin={userLoggedIn} />
    {/if}
  {/if}

</div>
