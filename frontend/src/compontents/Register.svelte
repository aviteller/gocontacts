<script>

  let user = {
    email:"",
    password:""
  }

  const createUser = async(user) => {
   let res = await fetch("http://localhost:8000/api/user/new", {
      method: "POST", // 'GET', 'PUT', 'DELETE', etc.
      body: JSON.stringify(user), // Coordinate the body type with 'Content-Type'

    });
    let data = await res.json();
    return data;
  }


  const onSubmit = (e) => {
    e.preventDefault();
    createUser(user).then(data => console.log(data))
 
  }

</script>




<form class="grid-1" on:submit={onSubmit}>

  <input type="email" placeholder="Email" bind:value={user.email} />
  <input type="password" placeholder="Enter password" bind:value={user.password} />


  <input type="submit" value="Register User" class="btn btn-primary btn-block" />
</form>