import React, { useState } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import { useEffect } from "react";
import { useCookies } from "react-cookie";
import 'react-datetime/css/react-datetime.css';

const CreateShipment = () => {
  const navigate = useNavigate();
  const [inputValue, setInputValue] = useState({
    movie_name: "",
    hall_name: "",
    date: "",
    time: "",
  });
  const { movie_name, hall_name, date, time } = inputValue;
  const location = useLocation();
  const [cookies, removeCookie] = useCookies([]);
  const [username, setUsername] = useState("");
  const [shipments, setShipments] = useState([]);
  const [halls, setHalls] = useState([]);

  useEffect(() => {
    // Parse the query parameters to get the username
    const searchParams = new URLSearchParams(location.search);
    const usernameFromParams = searchParams.get("email");
    setUsername(usernameFromParams);
    const fetchData = async () => {
        try {
            const response = await axios.get(
              `http://localhost:8000/movie`);
            const shipments1 = response.data.data;
            setShipments(shipments1);
            console.log(shipments1);
        } catch (error) {
            console.error(error);
            toast.error("Error fetching movies", { position: "bottom-left" });
        } 
    };
    fetchData();

    const fetchData1 = async () => {
        try {
            const response = await axios.get(
              `http://localhost:8000/my-halls/${username}`);
            const shipments1 = response.data.data;
            setHalls(shipments1);
            console.log(shipments1);
        } catch (error) {
            console.error(error);
            toast.error("Error fetching HALLS", { position: "bottom-left" });
        } 
    };
    fetchData1();
  }, [location.search, username]);
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
        `http://localhost:8000/add-show/${username}/${inputValue.movie_name}/${inputValue.hall_name}/${inputValue.date}/${inputValue.time}`,
        { withCredentials: true }
      );
      const {data} = data1
      if(data.data != null){
        handleSuccess("movie show added");
        const usernameQueryParam = `?email=${username}`;
        navigate("/admin-home" + usernameQueryParam);
      }
      else if(data.data === null){
        handleError("movie show not added");
        const usernameQueryParam = `?email=${username}`;
        navigate("/admin-home" + usernameQueryParam);
      }
      console.log(data);
    } catch (error) {
      console.log(error);
    }
    setInputValue({
      ...inputValue,
    movie_name: "",
    hall_name: "",
    date: "",
    time: "",
    });
  };

  return (
    <div className="form_container">
      <h2>Add Show</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="movie_name">movie name</label>
          <select
            name="movie_name"
            value={movie_name}
            onChange={handleOnChange}
          >
            <option></option>
            {shipments.map((shipment) => (
            <option key={shipment.id} value={shipment.movie_name}>
                {shipment.movie_name}
            </option>
            ))}
          
            </select>
        </div>
        <div>
          <label htmlFor="hall_name">hall name</label>
          <select
            name="hall_name"
            value={hall_name}
            onChange={handleOnChange}
          >
            <option></option>
            {halls.map((shipment) => (
            <option key={shipment.id} value={shipment.name}>
                {shipment.name}
            </option>
            ))}
          
            </select>
        </div>
        <div>
          <label htmlFor="date">date</label>
          <input
                type="date"
                name="date"
                value={inputValue.date}
                onChange={handleOnChange}
                placeholder="Enter date"
            />
        </div>
        <div>
          <label htmlFor="time">time</label>
          <input
            type="time"
            name="time"
            value={time}
            placeholder="Enter time"
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