<script>
  import { createEventDispatcher } from "svelte";
  import Cookies from "../Cookies.js";

  let c = new Cookies();
  let jwt = c.getCookie("jwt");
  jwt = JSON.parse(jwt);

  const dispatch = createEventDispatcher();

  let contact = {
    name: "",
    phone: ""
  };

  const createContact = async contact => {
    let res = await fetch("http://localhost:8000/api/contacts/new", {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      body: JSON.stringify(contact), // Coordinate the body type with 'Content-Type'
      headers: {
        Authorization: `Bearer ${jwt.token}`
      }
    });
    let data = await res.json();
    return data;
  };

  const onSubmit = e => {
    e.preventDefault();
    if (contact.name != "") {
      createContact(contact).then(data => {
        (contact = {
          id: data.contact.id,
          phone: data.contact.phone,
          name: data.contact.name
        }),
 
          dispatch("addcontact", contact);
        contact = {
          name: "",
          phone: ""
        };
      });
    }
  };
</script>

<form class="grid-1" on:submit={onSubmit}>
  <input type="text" placeholder="Contact Name" bind:value={contact.name} />
  <input type="text" placeholder="Contact Phone" bind:value={contact.phone} />

  <input type="submit" value="Add Contact" class="btn btn-primary btn-block" />
</form>
