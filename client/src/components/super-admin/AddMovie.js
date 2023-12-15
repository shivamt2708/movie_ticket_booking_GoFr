import React, { useState } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import { useEffect } from "react";
import { useCookies } from "react-cookie";

const CreateShipment = () => {
  const navigate = useNavigate();
  const [inputValue, setInputValue] = useState({
    movie_name: "",
  });
  const { movie_name } = inputValue;
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
        `http://localhost:8000/add-movie/${inputValue.movie_name}`,
        { withCredentials: true }
      );
      const {data} = data1
      if(data.data != null){
        handleSuccess("movie added");
        const usernameQueryParam = `?email=${username}`;
        navigate("/super-admin-home" + usernameQueryParam);
      }
      else if(data.data === null){
        handleError("movie not added");
        const usernameQueryParam = `?email=${username}`;
        navigate("/super-admin-home" + usernameQueryParam);
      }
      console.log(data);
    } catch (error) {
      console.log(error);
    }
    setInputValue({
      ...inputValue,
      movie_name: "",
    });
  };

  return (
    <div className="form_container">
      <h2>Add Movie Hall</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="movie_name">Movie's name</label>
          <input
            type="text"
            name="movie_name"
            value={movie_name}
            placeholder="Enter movie's name"
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