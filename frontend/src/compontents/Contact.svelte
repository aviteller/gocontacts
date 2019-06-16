<script>
  import { createEventDispatcher } from "svelte";
  import Cookies from "../Cookies.js";

  let c = new Cookies();
  let jwt = c.getCookie("jwt");
  jwt = JSON.parse(jwt);

  export let name;
  export let phone;
  export let id;
  const dispatch = createEventDispatcher();

  const removeContact = () => {
    if (confirm("Are you sure ?")) {
      fetch(`http://localhost:8000/api/contacts/delete/${id}`, {
        headers: {
          Authorization: `Bearer ${jwt.token}`
        }
      }).then(response => response.json(), dispatch("removecontact", id));
    }
  };
</script>

<div class="card">
  <h4>{name}</h4>
  <h4>{phone}</h4>
  <button class="btn btn-danger btn-sm" on:click={removeContact}>x</button>
</div>
