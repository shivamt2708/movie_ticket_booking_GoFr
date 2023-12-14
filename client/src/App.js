import { Route, Routes } from "react-router-dom";
import { Login, Signup } from "./pages";
import Home from "./pages/Home";
import AdminHome from "./components/admin/AdminHome"
import UserHome from "./components/user/UserHome"
import AddMovieHall from "./components/admin/AddMovieHall"
import MyHalls from "./components/admin/MyHalls"

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/my-halls" element={<MyHalls />} />
        <Route path="/add-movie-hall" element={<AddMovieHall />} />
        <Route path="/user-home" element={<UserHome />} />
        <Route path="/admin-home" element={<AdminHome />} />
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
      </Routes>
    </div>
  );
}

export default App;