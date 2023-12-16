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
              `http://localhost:8000/my-bookings/${username}`);
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
      <h1>My Bookings</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Booking id</th>
              <th>Show id</th>
              <th>User's email</th>
            </tr>
          </thead>
          <tbody>
            {shipments.map((shipment) => (
              <tr>
                <td>{shipment.id}</td>
                <td>
                    <span>
                        <Link to={"/my-shows" + `?id=${shipment.show_id}`}>{shipment.show_id}</Link>
                    </span>
                </td>
                <td>{shipment.user_email}</td>
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