
import React, { useState } from 'react';
import './LoginForm.css'; 
import axios from 'axios';
const LoginForm = () => {
  const [name, setUsername] = useState('');
  const [pass, setPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/user/login',{name, pass})
    .then((res) => {
      const responseData = res.data;
      console.log(res.data)
      // Split the response data into separate JSON objects
      const splitData = responseData.split('}{');
      // Parse each JSON object individually
      const parsedData = splitData.map((item, index) => {
        // Add back the curly braces to the first and last items
        if (index === 0) {
          return JSON.parse(item + '}');
        } else if (index === splitData.length - 1) {
          return JSON.parse('{' + item);
        }
        return JSON.parse('{' + item + '}');
      });
      // Extract tokens from parsed data
      const { token, refreshToken } = parsedData[1];
      console.log(token, refreshToken);
      localStorage.setItem('token', token);
      localStorage.setItem('refreshToken', refreshToken)
     alert('login Successful')
    
     
    })
    .catch((error) => {
      alert('error occured:' + error.message)
    });
  };
  return (
    <div className="login-container">
      <form className="login-form" onSubmit={handleSubmit}>
        <h2>Login</h2>
        <div className="input-group">
          <label htmlFor="Username">Username</label>
          <input
            id="Username"
            required
            value={name}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className="input-group">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            required
            value={pass}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button type="submit" className="login-button">Log In</button>
      </form>
    </div>
  );
};
export default LoginForm;