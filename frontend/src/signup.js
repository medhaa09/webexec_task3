import React from 'react';
import './signup.css';
import { useState }from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

export default function SignUpForm() {
  const [name, setName] = useState("");
  const [pass, setPass] = useState("");
  const [handle, setHandle] = useState("");
  
  const handlechange=(c)=>{
    c.preventDefault()  //the default behaviour will be that this function will try to navigate to the route below on click but we dont want that so we use preventDefault()
    axios.post('http://localhost:8080/user/signup',{name, handle, pass})
    .then((res) => {
     alert('Signup Successful')
     console.log(res)
    })
    .catch((error) => {
      alert('error occured:' + error.message)
    });
}
const navigate=useNavigate();
  return (
    <div className="extradiv">
    <div className="signup-form">
      <div className="header">
        <h1>Sign Up</h1>
        <p>Enter your information to create an account</p>
      </div>
      <form className="form"  onSubmit={handlechange}>
      <div className="form-group">
          <label htmlFor="Name">Name</label>
          <input id="cfhandle" placeholder="Enter your name" onChange={(e)=>setHandle(e.target.value)} required />
        </div>
        <div className="form-group">
          <label htmlFor="username">Username</label>
          <input id="username" placeholder="Enter your username" onChange={(e)=>setName(e.target.value)}required />
        </div>
        
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input id="password" type="password" placeholder="Enter your password" onChange={(e)=>setPass(e.target.value)} required />
        </div>
        <button type="submit" className="submit-btn" >Sign Up</button>
        <p>Already Signed up? <span onClick={() => navigate("/login")} id="loginhere">Login here!</span></p>
      </form>
    </div>
    </div>
  );
}
