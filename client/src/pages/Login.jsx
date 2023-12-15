import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";

const Login = () => {
  const navigate = useNavigate();
  const [inputValue, setInputValue] = useState({
    email: "",
    password: "",
  });
  const { email, password } = inputValue;
  const handleOnChange = (e) => {
    const { name, value } = e.target;
    setInputValue({
      ...inputValue,
      [name]: value,
    });
  };

  const handleError = (err) =>
    toast.error(err, {
      position: "bottom-left",
    });
  const handleSuccess = (msg) =>
    toast.success(msg, {
      position: "bottom-left",
    });

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data1 = await axios.post(
        `http://localhost:8000/login/${inputValue.email}/${inputValue.password}`,
        { withCredentials: true }
      );
      const usernameQueryParam = `?email=${inputValue.email}`;
      const {data} = data1
      console.log(data);
      if(data.data === "user"){
        console.log("user");
        setTimeout(() => {
          navigate("/user-home" + usernameQueryParam);
        }, 1000);
      }
      else if(data.data === "admin"){
        console.log("admin");
        setTimeout(() => {
          navigate("/admin-home" + usernameQueryParam);
        }, 1000);
      }
      else if(data.data === "super-admin"){
        console.log("super-admin");
        setTimeout(() => {
          navigate("/super-admin-home" + usernameQueryParam);
        }, 1000);
      }
      if (data.data != null) {
        handleSuccess("login successful");
      } else {
        handleError("login unsuccessful");
      }
    } catch (error) {
      console.log(error);
    }
    setInputValue({
      ...inputValue,
      email: "",
      password: "",
    });
  };

  return (
    <div className="form_container">
      <h2>Login Account</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email</label>
          <input
            type="email"
            name="email"
            value={email}
            placeholder="Enter your email"
            onChange={handleOnChange}
          />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input
            type="password"
            name="password"
            value={password}
            placeholder="Enter your password"
            onChange={handleOnChange}
          />
        </div>
        <button type="submit">Submit</button>
        <span>
          Already have an account? <Link to={"/signup"}>Signup</Link>
        </span>
      </form>
      <ToastContainer />
    </div>
  );
};

export default Login;