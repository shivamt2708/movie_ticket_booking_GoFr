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
    const usernameFromParams = searchParams.get("email");
    setUsername(usernameFromParams);
    const fetchData = async () => {
        try {
            const response = await axios.get(
              `http://localhost:8000/my-halls/${username}`);
            const shipments1 = response.data.data;
            setShipments(shipments1);
            console.log(shipments1);
        } catch (error) {
            console.error(error);
            toast.error("Error fetching halls", { position: "bottom-left" });
        } finally {
            setLoading(false);
        }
    };
    fetchData();
  }, [username]);

  return (
    <div>
      <h1>Your Halls</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Hall Name</th>
              <th>Admin's Email</th>
              <th>Price of one seat</th>
              <th>Total seats</th>
            </tr>
          </thead>
          <tbody>
            {shipments.map((shipment) => (
              <tr>
                <td>{shipment.name}</td>
                <td>{shipment.email}</td>
                <td>{shipment.price}</td>
                <td>{shipment.total_seats}</td>
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