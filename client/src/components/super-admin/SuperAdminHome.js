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
    navigate("/add-movie" + usernameQueryParam);
  };

  return (
    <>
      <div className="home_page">
        <h4>
          {" "}
          Welcome <span>{username}</span>
        </h4>
        <button onClick={Logout}>Logout</button>
        <button onClick={Ship}>Add Movie</button>
      </div>
      <ToastContainer />
    </>
  );
};

export default SellerHome;