import { Route, Routes } from "react-router-dom";
import { Login, Signup } from "./pages";
import Home from "./pages/Home";
import AdminHome from "./components/admin/AdminHome"
import SuperAdminHome from "./components/super-admin/SuperAdminHome"
import AddMovie from "./components/super-admin/AddMovie"
import AddMovieHall from "./components/admin/AddMovieHall"
import MyHalls from "./components/admin/MyHalls"
import BookTicket from "./components/admin/BookTicket"
import AddShow from "./components/admin/AddShow"
import MyBookings from "./components/admin/MyBookings"
import MyBookings2 from "./components/user/MyBookings2"
import BookTicket2 from "./components/user/BookTicket2"
import MyShows from "./components/admin/MyShows"
import UserHome from "./components/user/UserHome"


function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/admin/book-ticket" element={<BookTicket />} />
        <Route path="/user/book-ticket" element={<BookTicket2 />} />
        <Route path="/my-bookings2" element={<MyBookings2 />} />
        <Route path="/my-halls" element={<MyHalls />} />
        <Route path="/my-bookings" element={<MyBookings />} />
        <Route path="/add-movie-hall" element={<AddMovieHall />} />
        <Route path="/add-movie" element={<AddMovie />} />
        <Route path="/add-show" element={<AddShow />} />
        <Route path="/my-shows" element={<MyShows />} />
        <Route path="/admin-home" element={<AdminHome />} />
        <Route path="/user-home" element={<UserHome />} />
        <Route path="/super-admin-home" element={<SuperAdminHome />} />
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
      </Routes>
    </div>
  );
}

export default App;