import React, { useState } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import { useEffect } from "react";
import { useCookies } from "react-cookie";

const CreateShipment = () => {
  const navigate = useNavigate();
  const [inputValue, setInputValue] = useState({
    user_email: "",
    hall_name: "",
  });
  const { user_email, hall_name } = inputValue;
  const location = useLocation();
  const [cookies, removeCookie] = useCookies([]);
  const [username, setUsername] = useState("");

  useEffect(() => {
    // Parse the query parameters to get the username
    const searchParams = new URLSearchParams(location.search);
    const usernameFromParams = searchParams.get("email");
    setUsername(usernameFromParams);
  }, [location.search]);
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
        `http://localhost:8000/admin/book-ticket/${username}/${inputValue.user_email}/${inputValue.hall_name}`,
        { withCredentials: true }
      );
      const {data} = data1
      if(data.data != null){
        handleSuccess("ticket booked");
        const usernameQueryParam = `?email=${username}`;
        navigate("/admin-home" + usernameQueryParam);
      }
      else if(data.data === null){
        handleError("ticket not booked");
        const usernameQueryParam = `?email=${username}`;
        navigate("/admin-home" + usernameQueryParam);
      }
      console.log(data);
    } catch (error) {
      console.log(error);
    }
    setInputValue({
      ...inputValue,
      user_email: "",
      hall_name: "",
    });
  };

  return (
    <div className="form_container">
      <h2>Book Ticket</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="user_email">user email</label>
          <input
            type="text"
            name="user_email"
            value={user_email}
            placeholder="Enter user email"
            onChange={handleOnChange}
          />
        </div>
        <div>
          <label htmlFor="hall_name">hall name</label>
          <input
            type="text"
            name="hall_name"
            value={hall_name}
            placeholder="Enter hall name"
            onChange={handleOnChange}
          />
        </div>
        <button type="submit">Submit</button>
      </form>
      <ToastContainer />
    </div>
  );
};

export default CreateShipment;