<script>
  import { createEventDispatcher } from "svelte";

  let dispatch = createEventDispatcher();


  let user = {
    email: "",
    password: "",
    token: ""
  };

  const loginUser = async user => {
    let res = await fetch("http://localhost:8000/api/user/login", {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      body: JSON.stringify(user) // Coordinate the body type with 'Content-Type'
    });
    let data = await res.json();
    return data;
  };

  const onSubmit = e => {
    e.preventDefault();
    loginUser(user).then(data => {
      (user = {
        id: data.account.ID,
        email: data.account.email,
        token: data.account.token,
        loggedIn: true
      }),
        //console.log(data.account.ID),
        dispatch("userloggedin", user);
         user = {
          email: "",
          password: ""
        };
    });
  };
</script>

<form class="grid-1" on:submit={onSubmit}>

  <input type="email" placeholder="Email" bind:value={user.email} />
  <input
    type="password"
    placeholder="Enter password"
    bind:value={user.password} />

  <input type="submit" value="Login User" class="btn btn-primary btn-block" />
</form>
