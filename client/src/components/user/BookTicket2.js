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
    location1: "",
    admin: "",
    movie_name: "",
    hall_name: "",
    date: "",
    time: "",
    user_email: "",
  });
  const { location1, admin, movie_name, hall_name, date, time, user_email } = inputValue;
  const location = useLocation();
  const [cookies, removeCookie] = useCookies([]);
  const [username, setUsername] = useState("");
  const [show_id, setshow_id] = useState([]);
  const [seats_left, setseats_left] = useState([]);
  const [shipments, setShipments] = useState([]);
  const [movies, setmovies] = useState([]);
  const [dates, setdates] = useState([]);
  const [times, settimes] = useState([]);
  const [cinema, setcinema] = useState([]);

  useEffect(() => {
    // Parse the query parameters to get the username
    const searchParams = new URLSearchParams(location.search);
    const usernameFromParams = searchParams.get("email");
    setUsername(usernameFromParams);
  }, [location.search, username]);
  const handleOnChange = (e) => {
    const { name, value } = e.target;
    setInputValue({
      ...inputValue,
      [name]: value,
    });
  };

  const handleLocationChange = async (e) => {
    const selectedMovie = e.target.value;
  
    // Fetch data based on the selected movie
    try {
      const response = await axios.get(`http://localhost:8000/location/${selectedMovie}`);
      const shipments1 = response.data.data;
      const response2 = await axios.get(`http://localhost:8000/${shipments1[0].email}/movie`);
        const shipments2 = response2.data.data;
        setShipments(shipments2);
      setcinema(shipments1);
    } catch (error) {
      console.error(error);
    }
  
    // Update the state based on the selected movie
    setInputValue({
      ...inputValue,
      location1: selectedMovie,
    });
  };

  const handleMovieChange = async (e) => {
    const selectedMovie = e.target.value;
  
    // Fetch data based on the selected movie
    try {
      const response = await axios.get(`http://localhost:8000/my-shows/${cinema[0].email}/${selectedMovie}`);
      const shipments1 = response.data.data;
      setmovies(shipments1);
      console.log(shipments1);
    } catch (error) {
      console.error(error);
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
      const response = await axios.get(`http://localhost:8000/my-shows3/${cinema[0].email}/${inputValue.movie_name}/${selectedMovie}`);
      const shipments1 = response.data.data;
      setdates(shipments1)
      console.log(shipments1);
    } catch (error) {
      console.error(error);
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
      const response = await axios.get(`http://localhost:8000/my-shows4/${cinema[0].email}/${inputValue.movie_name}/${inputValue.date}/${selectedMovie}`);
      const shipments1 = response.data.data;
      settimes(shipments1)
      console.log(shipments1);
    } catch (error) {
      console.error(error);
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
      const response = await axios.get(`http://localhost:8000/my-shows2/${cinema[0].email}/${inputValue.movie_name}/${selectedMovie}/${inputValue.date}/${inputValue.time}`);
      const shipments1 = response.data.data;

      setshow_id(shipments1);

      console.log(shipments1);
    }  catch (error1) {
      // Handle error for the first request
      console.error("Error in Axios request 1:", error1);
    } 
    // Update the state based on the selected movie
    setInputValue({
      ...inputValue,
      hall_name: selectedMovie,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if(show_id[0].seats_left - 1 >= 0){
    try {
      const data1 = await axios.post(
        `http://localhost:8000/admin/book-ticket/${show_id[0].id}/${username}`,
        { withCredentials: true }
      );
      console.log(data1)
      const {data} = data1
      if(data.data != null){
        const usernameQueryParam = `?email=${username}`;
        navigate("/user-home" + usernameQueryParam);
      }
      else if(data.data === null){
        const usernameQueryParam = `?email=${username}`;
        navigate("/user-home" + usernameQueryParam);
      }
      console.log(data);
    } catch (error) {
      console.log(error);
    }

    try {
      const response = await axios.put(
        `http://localhost:8000/admin/book-ticket2/${show_id[0].id}/${show_id[0].seats_left - 1}`);

      const shipments1 = response.data.data;
      setseats_left(shipments1);
      console.log(response)
    }
    
    catch (error2) {
      // Handle error for the second request
      console.error("Error in Axios request 2:", error2);
    }
  }
    else{
      window.alert("no seats left in the show")
    }


    setInputValue({
      ...inputValue,
    location1: "",
    admin: "",
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
          <label htmlFor="location1">location</label>
          <input
                type="text"
                name="location1"
                value={inputValue.location1}
                onChange={handleLocationChange}
                placeholder="Enter location"
            />
        </div>
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
        <button type="submit">Submit</button>
      </form>
      <ToastContainer />
    </div>
  );
};

export default CreateShipment;