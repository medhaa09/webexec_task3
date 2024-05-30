
import './App.css'
import SignupForm from './signup.js';
import LoginForm from './LoginForm.js'
import { BrowserRouter, Routes, Route } from 'react-router-dom';

function App() {
  return (
    <div className="App">
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginForm/>} />
        <Route path="/user/signup" element={<SignupForm/>} />
      </Routes>
    </BrowserRouter>
    </div>
  );
}

export default App;
