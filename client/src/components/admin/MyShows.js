import React, { useState } from "react";
import { Link, useNavigate, useLocation } from "react-router-dom";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";
import { useEffect } from "react";
import { useCookies } from "react-cookie";

const MyShipments = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [username, setUsername] = useState("");
  const [shipments, setShipments] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Parse the query parameters to get the username
    const searchParams = new URLSearchParams(location.search);
    const usernameFromParams = searchParams.get("id");
    setUsername(usernameFromParams);
    const fetchData = async () => {
        try {
            const response = await axios.get(
              `http://localhost:8000/my-shows5/${username}`);
            const shipments1 = response.data.data;
            setShipments(shipments1);
            console.log(shipments1);
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };
    fetchData();
  }, [username]);

  return (
    <div>
      <h1>My Shows</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <table>
          <thead>
            <tr>
            <th>Email</th>
            <th>Movie Name</th>
              <th>Hall Name</th>
              <th>Date</th>
              <th>Time</th>
              <th>Seats Left</th>
            </tr>
          </thead>
          <tbody>
            {shipments.map((shipment) => (
              <tr>
                <td>{shipment.email}</td>
                <td>{shipment.movie_name}</td>
                <td>{shipment.hall_name}</td>
                <td>{shipment.date}</td>
                <td>{shipment.time}</td>
                <td>{shipment.seats_left}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
      <ToastContainer />
    </div>
  );
};

export default MyShipments;