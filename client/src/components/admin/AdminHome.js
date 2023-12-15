import { useEffect, useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { useCookies } from "react-cookie";
import axios from "axios";
import { ToastContainer, toast } from "react-toastify";

const SellerHome = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [cookies, removeCookie] = useCookies([]);
  const [username, setUsername] = useState("");

  useEffect(() => {
    // Parse the query parameters to get the username
    const searchParams = new URLSearchParams(location.search);
    const usernameFromParams = searchParams.get("email");
    setUsername(usernameFromParams);
  }, [location.search]);

  const Logout = () => {
    removeCookie("token");
    navigate("/signup");
  };

  const Ship = () => {
    // ... (implement the Ship functionality if needed)
    const usernameQueryParam = `?email=${username}`;
    navigate("/add-movie-hall" + usernameQueryParam);
  };

  const Ship2 = () => {
    // ... (implement the Ship functionality if needed)
    const usernameQueryParam = `?email=${username}`;
    navigate("/my-halls" + usernameQueryParam);
  };
  const Ship3 = () => {
    // ... (implement the Ship functionality if needed)
    const usernameQueryParam = `?email=${username}`;
    navigate("/admin/book-ticket" + usernameQueryParam);
  };
  const Ship4 = () => {
    // ... (implement the Ship functionality if needed)
    const usernameQueryParam = `?email=${username}`;
    navigate("/add-show" + usernameQueryParam);
  };

  return (
    <>
      <div className="home_page">
        <h4>
          {" "}
          Welcome <span>{username}</span>
        </h4>
        <button onClick={Logout}>Logout</button>
        <button onClick={Ship}>Add Movie Hall</button>
        <button onClick={Ship2}>My Halls</button>
        <button onClick={Ship3}>Book Ticket</button>
        <button onClick={Ship4}>Add Show</button>
      </div>
      <ToastContainer />
    </>
  );
};

export default SellerHome;