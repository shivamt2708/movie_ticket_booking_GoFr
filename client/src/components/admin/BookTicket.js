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
    user_email: "",
  });
  const { movie_name, hall_name, date, time, user_email } = inputValue;
  const location = useLocation();
  const [cookies, removeCookie] = useCookies([]);
  const [username, setUsername] = useState("");
  const [show_id, setshow_id] = useState([]);
  const [shipments, setShipments] = useState([]);
  const [movies, setmovies] = useState([]);
  const [dates, setdates] = useState([]);
  const [times, settimes] = useState([]);

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
  }, [location.search, username]);
  const handleOnChange = (e) => {
    const { name, value } = e.target;
    setInputValue({
      ...inputValue,
      [name]: value,
    });
  };

  const handleMovieChange = async (e) => {
    const selectedMovie = e.target.value;
  
    // Fetch data based on the selected movie
    try {
      const response = await axios.get(`http://localhost:8000/my-shows/${username}/${selectedMovie}`);
      const shipments1 = response.data.data;
      setmovies(shipments1);
      console.log(shipments1);
    } catch (error) {
      console.error(error);
      toast.error("Error fetching HALLS", { position: "bottom-left" });
    }
  
    // Update the state based on the selected movie
    setInputValue({
      ...inputValue,
      movie_name: selectedMovie,
    });
  };

  const handleDateChange = async (e) => {
    const selectedMovie = e.target.value;
  
    // Fetch data based on the selected movie
    try {
      const response = await axios.get(`http://localhost:8000/my-shows3/${username}/${inputValue.movie_name}/${selectedMovie}`);
      const shipments1 = response.data.data;
      setdates(shipments1)
      console.log(shipments1);
    } catch (error) {
      console.error(error);
      toast.error("Error fetching Date", { position: "bottom-left" });
    }
  
    // Update the state based on the selected movie
    setInputValue({
      ...inputValue,
      date: selectedMovie,
    });
  };

  const handleTimeChange = async (e) => {
    const selectedMovie = e.target.value;
  
    // Fetch data based on the selected movie
    try {
      const response = await axios.get(`http://localhost:8000/my-shows4/${username}/${inputValue.movie_name}/${inputValue.date}/${selectedMovie}`);
      const shipments1 = response.data.data;
      settimes(shipments1)
      console.log(shipments1);
    } catch (error) {
      console.error(error);
      toast.error("Error fetching Time", { position: "bottom-left" });
    }
  
    // Update the state based on the selected movie
    setInputValue({
      ...inputValue,
      time: selectedMovie,
    });
  };

  const handleHallChange = async (e) => {
    const selectedMovie = e.target.value;
  
    // Fetch data based on the selected movie
    try {
      const response = await axios.get(`http://localhost:8000/my-shows2/${username}/${inputValue.movie_name}/${selectedMovie}/${inputValue.date}/${inputValue.time}`);
      const shipments1 = response.data.data;

      setshow_id(shipments1);
      console.log(shipments1);
    } catch (error) {
      console.error(error);
      toast.error("Error fetching Shows", { position: "bottom-left" });
    }
  
    // Update the state based on the selected movie
    setInputValue({
      ...inputValue,
      hall_name: selectedMovie,
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
        `http://localhost:8000/admin/book-ticket/${show_id[0].id}/${inputValue.user_email}`,
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
    user_email: "",
    });
  };

  return (
    <div className="form_container">
      <h2>Book Ticket</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="movie_name">movie name</label>
          <select
            name="movie_name"
            value={movie_name}
            onChange={handleMovieChange}
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
          <label htmlFor="date">date</label>
          <select
            name="date"
            value={date}
            onChange={handleDateChange}
          >
            <option></option>
            {movies.map((shipment) => (
            <option key={shipment.id} value={shipment.date}>
                {shipment.date}
            </option>
            ))}
          
            </select>
        </div>
        <div>
          <label htmlFor="time">time</label>
          <select
            name="time"
            value={time}
            onChange={handleTimeChange}
          >
            <option></option>
            {dates.map((shipment) => (
            <option key={shipment.id} value={shipment.time}>
                {shipment.time}
            </option>
            ))}
          
            </select>
        </div>
        <div>
          <label htmlFor="hall_name">hall name</label>
          <select
            name="hall_name"
            value={hall_name}
            onChange={handleHallChange}
          >
            <option></option>
            {times.map((shipment) => (
            <option key={shipment.id} value={shipment.hall_name}>
                {shipment.hall_name}
            </option>
            ))}
          
            </select>
        </div>
        <div>
          <label htmlFor="user_email">user email</label>
          <input
                type="text"
                name="user_email"
                value={inputValue.user_email}
                onChange={handleOnChange}
                placeholder="Enter user email"
            />
        </div>
        <button type="submit">Submit</button>
      </form>
      <ToastContainer />
    </div>
  );
};

export default CreateShipment;