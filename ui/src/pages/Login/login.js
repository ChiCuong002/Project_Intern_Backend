import React, { useState } from 'react'
import './/login.css'
import { Link } from 'react-router-dom';
// import google from '../Components/Assets/google.png'
// import facebook from '../Components/Assets/facebook.png'
import axios from 'axios';

const Login = () => {
    const [phone, setPhone] = useState('');
    const [password, setPassword] = useState('');

    const handlePhoneChange = (e) => {
        setPhone(e.target.value)
    };

    const handlePasswordChange = (e) => {
        setPassword(e.target.value)
    };

    const handleLogin = (e) => {
        e.preventDefault();
        const loginData = {
            phone_number: phone,
            password: password,
        };
        axios.post('http://localhost:8080/login', loginData)
            .then(response => {
                console.log(response.data);
                const token = response.data.token;
                localStorage.setItem('token', token)
                alert("Dang nhap thanh cong")
            })
            .catch(error => {
                console.error('Lỗi:', error);
            });
    };

    return (
    <div className="login">
        <div className="login-container">
            <h1>LOGIN</h1>
            <form onSubmit={handleLogin}>
            <div className="login-fields">
                <input type="text" placeholder="Phone" value={phone} onChange={handlePhoneChange} />
                <input type="password" placeholder='Password' value={password} onChange={handlePasswordChange} />
            </div>
            <div className="login-p">
                <Link to="/register"><p>Sign up</p></Link>
                <p>Forgot password?</p>
            </div>
            <button type='submit'>Submit</button>
            </form>
            {/* <div className="loginButton google">
                <img src={google} alt='' />
                Google
            </div>
            <div className="loginButton facebook">
                <img src={facebook} alt='' />
                Facebook
            </div> */}
        </div>
    </div>
  )
}

export default Login